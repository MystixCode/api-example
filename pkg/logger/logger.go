package logger

import (
	"github.com/fatih/color"
	"github.com/spf13/viper"

	"fmt"
	"log"
	"os"
)

var (
	logger *log.Logger
)

func Init() {
	if viper.GetBool("logger.debug") {
		logger = log.New(os.Stdout, "LOG ", log.Ldate|log.Ltime)
	} else {
		if _, err := os.Stat(viper.GetString("logger.logfile_path")); os.IsNotExist(err) {
			err = os.Mkdir(viper.GetString("logger.logfile_path"), 0755)
			if err != nil {
				logger.Fatal(err)
			}
		}
		lf := viper.GetString("logger.logfile_path") + viper.GetString("logger.logfile_name")
		f, err := os.OpenFile(lf, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		logger = log.New(f, "LOG ", log.Ldate|log.Ltime)
	}

	if viper.GetBool("logger.debug") {
		logger.SetFlags(log.Lshortfile)
	}
}

// Debug are logged here
func Debug(args ...interface{}) {
	if viper.GetBool("logger.debug") {
		col := color.New(color.FgHiBlack, color.BgBlue, color.Bold).SprintfFunc()
		logger.SetPrefix(col("DEBUG\t"))
		logger.Println(fmt.Sprint(args...))
	}
}

// Info  are logged here
func Info(args ...interface{}) {
	if viper.GetBool("logger.debug") {
		col := color.New(color.FgHiBlack, color.BgGreen, color.Bold).SprintfFunc()
		logger.SetPrefix(col("INFO\t"))
	} else {
		logger.SetPrefix("INFO\t")
	}
	logger.Println(fmt.Sprint(args...))
}

// Warning  are logged here
func Warning(args ...interface{}) {
	if viper.GetBool("logger.debug") {
		col := color.New(color.FgHiBlack, color.BgYellow, color.Bold).SprintfFunc()
		logger.SetPrefix(col("WARN\t"))
	} else {
		logger.SetPrefix("WARN\t")
	}
	logger.Println(fmt.Sprint(args...))
}

// Error (normal errors) are logged here
func Error(args ...interface{}) {
	if viper.GetBool("logger.debug") {
		col := color.New(color.FgHiBlack, color.BgHiRed, color.Bold).SprintfFunc()
		logger.SetPrefix(col("ERROR\t"))
	} else {
		logger.SetPrefix("ERROR\t")
	}
	logger.Println(fmt.Sprint(args...))
}

// Fatal errors are logged here
func Fatal(args ...interface{}) {
	if viper.GetBool("logger.debug") {
		col := color.New(color.FgHiBlack, color.BgRed, color.Bold).SprintfFunc()
		logger.SetPrefix(col("FATAL\t"))
	} else {
		logger.SetPrefix("FATAL\t")
	}
	logger.Fatal(fmt.Sprint(args...))
}

// Panic errors are logged here
func Panic(args ...interface{}) {
	if viper.GetBool("logger.debug") {
		col := color.New(color.FgHiBlack, color.BgHiMagenta, color.Bold).SprintfFunc()
		logger.SetPrefix(col("PANIC\t"))
	} else {
		logger.SetPrefix("PANIC\t")
	}
	logger.Panic(fmt.Sprint(args...))
}
