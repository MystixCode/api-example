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
		logger = log.New(os.Stdout, "LOG ", log.Lshortfile)
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
}

// Debug is logged here
func Debug(args ...interface{}) {
	if viper.GetBool("logger.debug") {
		col := color.New(color.FgHiWhite, color.BgBlue, color.Bold).SprintfFunc()
		logger.SetPrefix(col("DEBUG\t"))
		logger.Println(fmt.Sprint(args...))
	}
}

// Info is logged here
func Info(args ...interface{}) {
	if viper.GetBool("logger.debug") {
		col := color.New(color.FgHiWhite, color.BgGreen, color.Bold).SprintfFunc()
		logger.SetPrefix(col("INFO\t"))
	} else {
		logger.SetPrefix("INFO\t")
	}
	logger.Println(fmt.Sprint(args...))
}

// Warning is logged here
func Warning(args ...interface{}) {
	if viper.GetBool("logger.debug") {
		col := color.New(color.FgHiWhite, color.BgYellow, color.Bold).SprintfFunc()
		logger.SetPrefix(col("WARN\t"))
	} else {
		logger.SetPrefix("WARN\t")
	}
	logger.Println(fmt.Sprint(args...))
}

// Error (normal error) is logged here
func Error(args ...interface{}) {
	if viper.GetBool("logger.debug") {
		col := color.New(color.FgHiWhite, color.BgHiRed, color.Bold).SprintfFunc()
		logger.SetPrefix(col("ERROR\t"))
	} else {
		logger.SetPrefix("ERROR\t")
	}
	logger.Println(fmt.Sprint(args...))
}

// Fatal error is logged here
func Fatal(args ...interface{}) {
	if viper.GetBool("logger.debug") {
		col := color.New(color.FgHiWhite, color.BgRed, color.Bold).SprintfFunc()
		logger.SetPrefix(col("FATAL\t"))
	} else {
		logger.SetPrefix("FATAL\t")
	}
	logger.Fatal(fmt.Sprint(args...))
}

// Panic error is logged here
func Panic(args ...interface{}) {
	if viper.GetBool("logger.debug") {
		col := color.New(color.FgHiWhite, color.BgHiMagenta, color.Bold).SprintfFunc()
		logger.SetPrefix(col("PANIC\t"))
	} else {
		logger.SetPrefix("PANIC\t")
	}
	logger.Panic(fmt.Sprint(args...))
}
