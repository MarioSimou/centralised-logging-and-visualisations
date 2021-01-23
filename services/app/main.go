package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

func main(){
	var portPtr = flag.String("port", "8080", "The port the app listens to")
	var appPathPtr = flag.String("appPath", "/go/src/app", "The app location")

	flag.Parse()

	var logger = logrus.Logger{
		Formatter: &logrus.JSONFormatter{},
		Out: os.Stderr,
		Level: logrus.TraceLevel,
		Hooks: make(logrus.LevelHooks),
	}

	var fileSystemHook = lfshook.NewHook(
		lfshook.PathMap{
			logrus.InfoLevel: strings.Join([]string{*appPathPtr, "app.log"}, "/"),
			logrus.ErrorLevel: strings.Join([]string{*appPathPtr, "app.log"}, "/"),
		},
		&logrus.JSONFormatter{},
	)
	logger.Hooks.Add(fileSystemHook)

	var address = fmt.Sprintf(":%s", *portPtr)
	
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
		var fields = map[string]interface{}{
			"userAgent":  r.Header.Get("User-Agent"),
			"referer": r.Header.Get("Referer"),
			"path": r.RequestURI,
			"host": r.Host,
			"requestID": r.Header.Get("X-Request-ID"),
			"proto": r.Proto,
			"method": r.Method,
			"contentLength": r.ContentLength,
		}

		logger.WithFields(fields).Infof("new request for /hello")
		fmt.Fprintln(w, "hello world")
	})

	logger.Infof("The app listens on address '%s'", address)
	logger.Fatalln(http.ListenAndServe(address, nil))	
}