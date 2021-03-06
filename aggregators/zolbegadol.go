package aggregators

// An aggregator for the Zol-Begadol chain.

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"regexp"
	"path/filepath"
	"log"
	"os"
)

// Home page of the Zol-Begadol price site.
const zolbegadolHome = "http://zolvebegadol.com/"

// An aggregator for the Zol-Begadol chain.
type zolbegadolAggregator struct{}

// Returns a new Zol-Begadol aggregator.
func Zolbegadol() Aggregator {
	return &zolbegadolAggregator{}
}

func (a *zolbegadolAggregator) Aggregate(dir string) error {
	// Create output directory.
	err := os.MkdirAll(dir, 0700)
	if err != nil {
		return fmt.Errorf("Failed to make dir: %v", err)
	}
	
	// Start downloader threads.
	files, filesErr := a.getFilesChannel()
	done := make(chan error, numberOfThreads)
	
	for i := 0; i < numberOfThreads; i++ {
		go func() {
			for df := range files {
				to := filepath.Join(dir, df.file)
				_, err := downloadIfNotExists(zolbegadolHome + df.dir + df.file,
						to, nil)
				if err != nil {
					done <- err
					return
				}
			}
			
			done <- nil
		}()
	}
	
	// Wait for threads to finish (including pusher thread).
	for i := 0; i < numberOfThreads; i++ {
		e := <- done
		if e != nil {
			err = e
		}
	}
	
	// Drain file channel.
	for range files {}
	
	// Check for errors in file getter.
	e := <-filesErr
	if e != nil {
		err = e
	}

	return err
}

// Returns paths of subdirectories of the price page.
func (a *zolbegadolAggregator) getDirectories() ([]string, error) {
	// Get home page.
	res, err := http.Get(zolbegadolHome)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Failed to get page (status %s).", res.Status)
	}
	
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	
	// Parse directory names.
	re := regexp.MustCompile("<a href=\"(201\\d{5}/)\"")
	match := re.FindAllSubmatch(body, -1)
	
	if len(match) == 0 {
		return nil, fmt.Errorf("Found no directories.")
	}
	
	dirs := make([]string, len(match))
	for i := range match {
		dirs[i] = string(match[i][1]) + "/gz/"
	}
	
	return dirs, nil
}

// Returns paths of files in a subdirectory. The paths are ready for download.
// dir should be as returned from getDirectories.
func (a *zolbegadolAggregator) getFiles(dir string) ([]string, error) {
	// Get home page.
	res, err := http.Get(zolbegadolHome + dir)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Failed to get page (status %s).", res.Status)
	}
	
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// Parse file names.
	re := regexp.MustCompile("<a href=\"([^\"]*\\.gz)\"")
	match := re.FindAllSubmatch(body, -1)
	
	if len(match) == 0 {
		return nil, fmt.Errorf("Found no files in directory '%s'.", dir)
	}
	
	files := make([]string, len(match))
	for i := range match {
		files[i] = string(match[i][1])
	}
	
	return files, nil
}

// Returns a channel through which file names for download will be returned,
// and a channel for error reporting. The error reporting will only report one
// error, which will be nil if everything went ok.
//
// This function was created because going over all directories in a single
// thread takes too long.
func (a *zolbegadolAggregator) getFilesChannel() (files chan *dirFile,
		done chan error) {
	// Initialize channels.
	files = make(chan *dirFile, numberOfThreads)
	done = make(chan error, 1)
		
	// Get files for download.
	dirs, err := a.getDirectories()
	if err != nil {
		done <- err
		close(files)
		return
	}
	if len(dirs) == 0 {
		done <- fmt.Errorf("Found 0 directories.")
		close(files)
		return
	}
	log.Printf("Found %d directories.", len(dirs))
	
	// Create pusher threads.
	dirChan := make(chan string, numberOfThreads)
	pushDones := make(chan error, numberOfThreads)
	
	for i := 0; i < numberOfThreads; i++ {
		go func() {
			for dir := range dirChan {
				// Download file list.
				fileList, err := a.getFiles(dir)
				if err != nil {
					pushDones <- err
					return
				}
				if len(fileList) == 0 {
					pushDones <- fmt.Errorf("Found 0 files in directory %s.",
							dir)
					return
				}
				
				// Push files into channel.
				for _, file := range fileList {
					files <- &dirFile{dir, file}
				}
			}
			
			pushDones <- nil
		}()
	}
	
	// Dir pusher thread.
	go func() {
		for _, dir := range dirs {
			dirChan <- dir
		}
		close(dirChan)
	}()
	
	// Waiter thread.
	go func() {
		// Wait for pusher threads.
		var err error
		for i := 0; i < numberOfThreads; i++ {
			e := <-pushDones
			if e != nil {
				err = e
			}
		}
		done <- err
		close(done)
		
		// Drain dir pusher.
		for range dirChan {}
		
		// We're done!
		close(files)
	}()
	
	return
}
