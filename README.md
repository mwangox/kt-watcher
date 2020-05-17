# go_file_watcher

Go file watcher is used to monitor config file (config.json) and reload the changes dynamically. 
No need of application restart.

#Deployment
You can build it or just run the following
go run .\config.go .\main.go --config.file="./conf/config.json"

Try to change (add,remove, modify) entries in config.json and observe how output changes.
