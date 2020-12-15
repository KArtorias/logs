package logs

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	wg         sync.WaitGroup
	logger     *Logger
	fileFormat = "./%slog_message.%s.log"
	defLogKey  = "LOG_ID"
)

type SplitType int

const (
	DAY    SplitType = 1
	HOUR   SplitType = 2
	MINUTE SplitType = 3
)

var (
	DATA_FORMAT = map[SplitType]string{
		DAY:    "2006-01-02",
		HOUR:   "2006-01-02 15",
		MINUTE: "2006-01-02 15-04",
	}
)

type Logger struct {
	mutex      sync.RWMutex
	logFile    *os.File
	suffix     string
	formatType SplitType
	dir        string
	logKey     string
}

func init() {
	timeStr := time.Now().Format(DATA_FORMAT[DAY])
	suffix := strings.Join(strings.Split(timeStr, " "), "_")
	logger = &Logger{
		suffix:     suffix,
		formatType: DAY,
		logKey:     defLogKey,
	}
}

func SetLogSplitType(splitType SplitType) {
	logger.formatType = splitType
	timeStr := time.Now().Format(DATA_FORMAT[logger.formatType])
	logger.suffix = strings.Join(strings.Split(timeStr, " "), "_")
}

func SetLogKey(logKey string) {
	logger.logKey = logKey
}

func SetDir(dir string) {
	if dir == "/" || dir == "" {
		return
	}
	if strings.HasPrefix(dir, "/") {
		dir = dir[1:]
	}
	if !strings.HasSuffix(dir, "/") {
		dir = dir + "/"
	}
	logger.dir = dir
}

func StartLog() {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()
	if logger.dir != "" {
		os.MkdirAll(logger.dir, os.ModePerm)
	}
	file := fmt.Sprintf(fileFormat, logger.dir, logger.suffix)
	var err error
	logger.logFile, err = os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logger.logFile)
	log.SetOutput(mw)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func Stop() {
	wg.Wait()
	logger.logFile.Close()
}

func flashLogFile() {
	timeStr := time.Now().Format(DATA_FORMAT[logger.formatType])
	suffix := strings.Join(strings.Split(timeStr, " "), "_")
	if suffix == logger.suffix {
		return
	}

	logger.mutex.Lock()
	defer logger.mutex.Unlock()
	wg.Wait()
	logger.logFile.Close()

	logger.suffix = suffix
	file := fmt.Sprintf(fileFormat, logger.dir, logger.suffix)
	var err error
	logger.logFile, err = os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logger.logFile)
	log.SetOutput(mw)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
