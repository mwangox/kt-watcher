package main

import (
	"log"
	"time"
)

type Config struct {
	Database struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Port     string `json:"port"`
	} `json:"database"`
}

func main() {

	config := GetConfig()
	go func() {
		for {
			log.Println("Hello New Config", config)
			time.Sleep(2 * time.Second)
		}

	}()

	ConfigFileWatcher()
	/*watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	defer watcher.Close()
	done := make(chan bool)

	go func() {
		for {
			log.Println("Hello New Config", config)
			time.Sleep(2 * time.Second)
		}

	}()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("Event: ", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("Modified file:", event.Name)
					log.Println("Operation:", event.Op)
					json_f, err := os.Open("C:/Users/mwangokai.VCTZ/developer/go_file_watcher/conf/config.json")
					if err != nil {
						log.Fatal("Failed to open file,", err)
					}
					json_b, _ := ioutil.ReadAll(json_f)
					json.Unmarshal(json_b, &config)
					if err != nil {
						log.Println("Failed to unmarshal", err)
					}
					log.Println(config)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("Error:", err)
			}
		}

	}()

	error := watcher.Add("C:/Users/mwangokai.VCTZ/developer/go_file_watcher/conf/config.json")
	if error != nil {
		log.Fatal(error)
	}
	<-done */

}
