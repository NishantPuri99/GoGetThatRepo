package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

const RESPONSE_BODY_READ_ERROR = "error reading response body: %v"
const ATTEMPT_FAIL_MSG = "Attempt %d failed: %v\n"
const FILE_NOT_FOUND_MSG = "file not found: %s"
const ATTEMPT_FAIL_MSG_STATUS = "Attempt %d failed with status %d\n"
const RETRYING_MSG = "Retrying..."
const RETRY_FAIL_MSG = "error fetching file after 3 retries: %v"

func GetRawGitHubURL(repo string, filename string) string {
	return fmt.Sprintf("https://raw.githubusercontent.com/%s/main/%s", repo, filename)
}

func FetchFileContent(url string) (string, error) {
	var err error
	for attempt := 1; attempt <= 3; attempt++ {
		resp, err := http.Get(url)
		if err == nil && resp.StatusCode == http.StatusOK {
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return "", fmt.Errorf(RESPONSE_BODY_READ_ERROR, err)
			}
			return string(body), nil
		}

		if err != nil {
			fmt.Printf(ATTEMPT_FAIL_MSG, attempt, err)
		} else {
			if resp.StatusCode == http.StatusNotFound {
				return "", fmt.Errorf(FILE_NOT_FOUND_MSG, url)
			}
			log.Printf(ATTEMPT_FAIL_MSG_STATUS, attempt, resp.StatusCode)
		}
	}
	return "", fmt.Errorf(RETRY_FAIL_MSG, err)
}

func FetchMultipleFiles(urls []string) (map[string]string, []error) {
	var wg sync.WaitGroup
	results := make(map[string]string)
	mu := sync.Mutex{}
	errChan := make(chan error, len(urls))

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			content, err := FetchFileContent(url)
			if err != nil {
				errChan <- err
				return
			}
			mu.Lock()
			defer mu.Unlock()
			results[url] = content
		}(url)
	}

	wg.Wait()
	close(errChan)

	var errors []error
	for err := range errChan {
		errors = append(errors, err)
	}
	if len(errors) > 0 {
		return nil, errors
	}

	return results, nil
}
