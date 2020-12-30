package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"ranlay.com/sql/other/pprof/data"
)

func main() {
	go func() {
		for {
			log.Println(data.Add("https://www.baidu.com"))
		}
	}()

	http.ListenAndServe("0.0.0.0:6060", nil)
}
