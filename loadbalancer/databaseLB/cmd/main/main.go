package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"

	"github.com/pranay999000/databaseLB/models"
)

var cfg models.Config

func isAlive(url *url.URL) bool {
	conn, err := net.DialTimeout("tcp", url.Host, time.Minute * 1)

	if err != nil {
		log.Printf("Unreachable to %v, error:%v", url.Host, err.Error())
		return false
	}

	defer conn.Close()
	return true
}

func healthCheck() {
	t := time.NewTicker(time.Minute * 1)

	for {
		select {
		case <- t.C:
			for _, database := range cfg.Databases {
				pingUrl, err := url.Parse(database.URL)

				if err != nil {
					log.Fatalln(err.Error())
				}

				isAlive := isAlive(pingUrl)
				database.SetDead(!isAlive)
				msg := "ok"

				if !isAlive {
					msg = "dead"
				}

				log.Printf("%v checked %v by healthcheck", database.URL, msg)
			}
		}
	}
}

var mutex sync.Mutex
var id int = 0

func lbHandler(w http.ResponseWriter, r *http.Request) {
	maxLen := len(cfg.Databases)

	mutex.Lock()

	currentDatabase := &cfg.Databases[id % maxLen]
	if currentDatabase.GetIsDead() {
		id += 1
	}

	targetURL, err := url.Parse(cfg.Databases[id % maxLen].URL)
	
	if err != nil {
		log.Fatalln(err.Error())
	}

	id += 1
	mutex.Unlock()

	reverseProxy := httputil.NewSingleHostReverseProxy(targetURL)
	reverseProxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, e error) {
		log.Printf("%v is dead.", targetURL)
		currentDatabase.SetDead(true)
		lbHandler(w, r)
	}
	reverseProxy.ServeHTTP(w, r)
}

func main() {
	data, err := ioutil.ReadFile("./config.json")

	if err != nil {
		log.Fatalln(err.Error())
	}

	json.Unmarshal(data, &cfg)

	go healthCheck()

	s := http.Server {
		Addr: ":" + cfg.Proxy.Port,
		Handler: http.HandlerFunc(lbHandler),
	}

	if err = s.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}