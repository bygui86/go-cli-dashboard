# go-cli-dashboard

Golang terminal dashboard based on gizak / termui

## Run

1. Start a Go application exposing `expvar` (debug information like memstats, pprof and others)

	```shell script
	make run-example-app
	```

2. start CLI dashboard

	- based on gizak/termui
	
		```shell script
		make run-termui-based
		```

	- `TODO` based on mum4k/termdash
	
		```shell script
		make run-termdash-based
		```

3. see how the HTTP server behaves using the CLI dashboard 

---

## Links
- https://levelup.gitconnected.com/building-a-terminal-dashboard-in-golang-in-300-lines-of-code-3b9f83f363a8
- https://github.com/gizak/termui
- https://github.com/mum4k/termdash

