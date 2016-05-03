package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

var (
	LoggerTrace   *log.Logger
	LoggerInfo    *log.Logger
	LoggerWarning *log.Logger
	LoggerError   *log.Logger
)

func InitLogger(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	LoggerTrace = log.New(traceHandle, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
	LoggerInfo = log.New(infoHandle, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	LoggerWarning = log.New(warningHandle, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	LoggerError = log.New(errorHandle, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		//LoggerInfo.Printf("%s\t%s\t%s\t%s", r.Method, r.RequestURI, name, time.Since(start))
		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start))
	})
}
