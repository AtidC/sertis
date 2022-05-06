package log

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	conf "github.com/spf13/viper"
)

var startProcess time.Time
var logLevel string

type logger struct {
	Timestamp string `json:"@timestamp"`        // The time at which the log message was created (ns) for kibana
	Suffix    string `json:"@suffix"`           // Suffix for kibana
	RequestID string `json:"requestId"`         // The transaction id
	Level     string `json:"level"`             // The log level
	Func      string `json:"func"`              // The function name write log
	Msg       string `json:"msg,omitempty"`     // The log message
	Status    string `json:"status,omitempty"`  // The status log
	Elapsed   int64  `json:"elapsed,omitempty"` // The elapsed time at end of process (ms)
}

func Error(requestID, format string, args ...interface{}) {
	log := logger{}
	pc, _, lineno, ok := runtime.Caller(1)
	if ok {
		str := strings.Split(runtime.FuncForPC(pc).Name(), "/")
		log.Func = fmt.Sprintf("%s:%d", str[len(str)-1], lineno)
	}
	log.Timestamp = time.Now().Format(time.RFC3339Nano)
	log.Suffix = conf.GetString("kibana.suffix")
	log.Level = "ERROR"
	log.RequestID = requestID
	log.Msg = fmt.Sprintf(format, args...)

	enc := json.NewEncoder(os.Stdout)
	enc.SetEscapeHTML(false)
	if conf.GetString("service.env") == "dev" {
		enc.SetIndent("", "    ")
	}
	if err := enc.Encode(log); err != nil {
		fmt.Printf("Failed to encode log, %v\n", err)
	}
}

func Debug(requestID, format string, args ...interface{}) {
	if logLevel == "Debug" {
		log := logger{}
		pc, _, lineno, ok := runtime.Caller(1)
		if ok {
			str := strings.Split(runtime.FuncForPC(pc).Name(), "/")
			log.Func = fmt.Sprintf("%s:%d", str[len(str)-1], lineno)
		}
		log.Timestamp = time.Now().Format(time.RFC3339Nano)
		log.Suffix = conf.GetString("kibana.suffix")
		log.Level = "DEBUG"
		log.RequestID = requestID
		log.Msg = fmt.Sprintf(format, args...)

		enc := json.NewEncoder(os.Stdout)
		enc.SetEscapeHTML(false)
		if conf.GetString("service.env") == "dev" {
			enc.SetIndent("", "    ")
		}
		if err := enc.Encode(log); err != nil {
			fmt.Printf("Failed to encode log, %v\n", err)
		}
	}
}

func Info(requestID, format string, args ...interface{}) {
	log := logger{}
	pc, _, lineno, ok := runtime.Caller(1)
	if ok {
		str := strings.Split(runtime.FuncForPC(pc).Name(), "/")
		log.Func = fmt.Sprintf("%s:%d", str[len(str)-1], lineno)
	}
	log.Timestamp = time.Now().Format(time.RFC3339Nano)
	log.Suffix = conf.GetString("kibana.suffix")
	log.Level = "INFO"
	log.RequestID = requestID

	log.Msg = fmt.Sprintf(format, args...)

	enc := json.NewEncoder(os.Stdout)
	enc.SetEscapeHTML(false)
	if conf.GetString("service.env") == "dev" {
		enc.SetIndent("", "    ")
	}
	if err := enc.Encode(log); err != nil {
		fmt.Printf("Failed to encode log, %v\n", err)
	}
}

func StartAPI(requestID string, c echo.Context) {
	//Info(requestID, "START: %s", path)
	log := logger{}
	pc, _, lineno, ok := runtime.Caller(1)
	if ok {
		str := strings.Split(runtime.FuncForPC(pc).Name(), "/")
		log.Func = fmt.Sprintf("%s:%d", str[len(str)-1], lineno)
	}
	log.Timestamp = time.Now().Format(time.RFC3339Nano)
	log.Suffix = conf.GetString("kibana.suffix")
	log.Level = "INFO"
	log.RequestID = requestID

	log.Msg = fmt.Sprintf("START: %s, %s", c.Request().Method, c.Path())

	enc := json.NewEncoder(os.Stdout)
	enc.SetEscapeHTML(false)
	if conf.GetString("service.env") == "dev" {
		enc.SetIndent("", "    ")
	}
	if err := enc.Encode(log); err != nil {
		fmt.Printf("Failed to encode log, %v\n", err)
	}
}

func End(requestID string, startTime time.Time, err error) {
	log := logger{}
	pc, _, lineno, ok := runtime.Caller(1)
	if ok {
		str := strings.Split(runtime.FuncForPC(pc).Name(), "/")
		log.Func = fmt.Sprintf("%s:%d", str[len(str)-1], lineno)
	}
	log.Timestamp = time.Now().Format(time.RFC3339Nano)
	log.Suffix = conf.GetString("kibana.suffix")

	elapsedTime := ElapsedTime(startTime)
	if err != nil {
		log.Level = "ERROR"
		log.RequestID = requestID
		log.Msg = fmt.Sprintf("done: %d ms., Error: %s ", elapsedTime, err.Error())
	} else {
		log.Level = "INFO"
		log.RequestID = requestID
		log.Msg = fmt.Sprintf("done: %d ms.", elapsedTime)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetEscapeHTML(false)
	if conf.GetString("service.env") == "dev" {
		enc.SetIndent("", "    ")
	}
	if err := enc.Encode(log); err != nil {
		fmt.Printf("Failed to encode log, %v\n", err)
	}
}

func EndAPI(requestID string, startTime time.Time, statusCode int, response interface{}) {
	log := logger{}
	pc, _, lineno, ok := runtime.Caller(1)
	if ok {
		str := strings.Split(runtime.FuncForPC(pc).Name(), "/")
		log.Func = fmt.Sprintf("%s:%d", str[len(str)-1], lineno)
	}
	log.Timestamp = time.Now().Format(time.RFC3339Nano)
	log.Suffix = conf.GetString("kibana.suffix")
	log.Level = "INFO"
	log.RequestID = requestID

	elapsedTime := ElapsedTime(startTime)
	if statusCode == 200 {
		resByte, err := json.Marshal(response)
		if err != nil {
			log.Msg = fmt.Sprintf("END: %d ms., Status: %d, Response: %s ",
				elapsedTime, statusCode, response)
		}
		log.Msg = fmt.Sprintf("END: %d ms., Status: %d, Response: %s ",
			elapsedTime, statusCode, string(resByte))
	} else {
		log.Msg = fmt.Sprintf("END: %d ms., Status: %d, Response: %s ",
			elapsedTime, statusCode, response)
	}
	log.Elapsed = elapsedTime

	enc := json.NewEncoder(os.Stdout)
	enc.SetEscapeHTML(false)
	if conf.GetString("service.env") == "dev" {
		enc.SetIndent("", "    ")
	}
	if err := enc.Encode(log); err != nil {
		fmt.Printf("Failed to encode log, %v\n", err)
	}
}

func ToJson(response interface{}) interface{} {
	resByte, err := json.Marshal(response)
	if err != nil {
		return response
	}
	return string(resByte)
}

func SetLogLevel(level string) {
	logLevel = level
}

func getElapsedTime() int64 {
	return time.Since(startProcess).Nanoseconds() / int64(time.Millisecond)
}

func ElapsedTime(startTime time.Time) int64 {
	return time.Since(startTime).Nanoseconds() / int64(time.Millisecond)
}
