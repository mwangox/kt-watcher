package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"

	_ "flag"

	"github.com/fsnotify/fsnotify"
)

var config Config

var filename string

func init() {
	config = Config{}
	flag.StringVar(&filename, "config.file", "./conf/config.json", "Configuration file")
	flag.Parse()
}

func GetConfig() *Config {
	json_f, err := os.Open(filename)
	if err != nil {
		log.Fatal("Failed to open file,", err)
	}
	json_b, _ := ioutil.ReadAll(json_f)

	json.Unmarshal(json_b, &config)
	if err != nil {
		log.Fatal("Failed to unmarshal", err)
	}
	defer json_f.Close()

	return &config
}

func ConfigFileWatcher() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	defer watcher.Close()
	done := make(chan bool)

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("Config file changed:", event.Name)

					json_f, err := os.Open(filename)
					if err != nil {
						log.Fatal("Failed to open file,", err)
					}
					json_b, _ := ioutil.ReadAll(json_f)
					json.Unmarshal(json_b, &config)
					if err != nil {
						log.Fatal("Failed to unmarshal", err)
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("Error:", err)
			}
		}

	}()

	error := watcher.Add(filename)
	if error != nil {
		log.Fatal(error)
	}
	<-done
}
