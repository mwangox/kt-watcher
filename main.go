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
        //Start watching
	ConfigFileWatcher()

}
