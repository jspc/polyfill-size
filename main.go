package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	url string

	features    = flag.String("features", "default,fetch,includes,HTMLPictureElement,Array.prototype.entries,Object.assign", "'Features' string to request a polyfil for/from")
	agents      = flag.String("agents", "./agents.txt", "Path to file containing useragents to test")
	concurrency = flag.Int("concurrency", 50, "Number of downloads to run at once")
	minify      = flag.Bool("min", true, "Download minified version")

	httpAgent = &http.Client{}
)

func main() {
	flag.Parse()
	url = func() string {
		if *minify {
			return fmt.Sprintf("https://cdn.polyfill.io/v2/polyfill.min.js?features=%s", *features)
		} else {
			return fmt.Sprintf("https://cdn.polyfill.io/v2/polyfill.js?features=%s", *features)
		}
	}()

	log.Print("Starting")

	c := make(chan int)
	lines := 0
	count := 0
	running := 0

	file, err := os.Open(*agents)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = lines + 1

		l := scanner.Text()
		go func(l string) {
			for running >= *concurrency {
				time.Sleep(10 * time.Millisecond)
			}
			running = running + 1

			err := grab(l)
			running = running - 1

			if err != nil {
				log.Fatal(err)
			}

			c <- 1
		}(l)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _ = range c {
		count = count + 1
		if count == lines {
			break
		}
	}

	log.Print("Completed")
}

func grab(s string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", s)

	resp, err := httpAgent.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename(s), body, 0644)
}

func filename(s string) string {
	s = strings.Replace(s, " ", "_", -1)
	s = strings.Replace(s, "(", "_", -1)
	s = strings.Replace(s, ")", "_", -1)
	s = strings.Replace(s, ";", "_", -1)
	s = strings.Replace(s, "/", "", -1)

	return fmt.Sprintf("output/%s", s)
}
