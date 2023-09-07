package ezlog

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Represents a logger object
type Logger struct {
    // Which levels to log (1 = DEBUG, 2 = INFO, 3 = WARNING, 4 = ERROR, 5 = CRITICAL)
    MinimumSeverity int
    // Where to output logs to. Set it as "" for no file
    OutputFile string
    // Name of logger
    Name string
    // How to format the date in the logs. I recommend using the formats in the time module
    TimeFormat string
    // Should the logger append to the log file or truncate it
    AppendFile bool
    // Should the logger print the logs to the console
    WriteToStdout bool
}

// Constructor for the logger 
func NewLogger(minimumSeverity int, outputFile string, name string, timeFormat string, appendFile bool, writeToStdout bool) Logger {
    return Logger {minimumSeverity, outputFile, name, timeFormat, appendFile, writeToStdout}
}

// Internal method for printing to the console and writing to a file
func (l Logger) PrintFormatted(message string, severity string, colour string, date time.Time) {
    if l.WriteToStdout {
        fmt.Printf(colour, fmt.Sprintf("[%s] (%s) %s - %s", severity, l.Name, date.Format(l.TimeFormat), message))
    }
    if l.OutputFile != "" {
        var mode int
        if l.AppendFile {
            mode = os.O_APPEND
        } else {
            mode = os.O_TRUNC
        }

        file, err := os.OpenFile(l.OutputFile, mode|os.O_CREATE|os.O_WRONLY, 0644)
        if err != nil {
            log.Fatal(err)
        }

        fmt.Fprintf(file, "[%s] (%s) %s - %s\n", severity, l.Name, date.Format(l.TimeFormat), message) 
        if err != nil {
            log.Fatal(err)
        }
    } else if !l.WriteToStdout {
        LogWarning(fmt.Sprintf("No output set for logger '%s'", l.Name))
    }
}

func (l Logger) LogInfo(message string) {
    if l.MinimumSeverity > 2 { return }
    currentTime := time.Now()
    l.PrintFormatted(message, "INFO", InfoColour, currentTime) 
}


func (l Logger) LogWarning(message string) {
    if l.MinimumSeverity > 3 { return }
    currentTime := time.Now()
    l.PrintFormatted(message, "WARNING", WarningColour, currentTime)
}


func (l Logger) LogError(message string) {
    if l.MinimumSeverity > 4 { return }
    currentTime := time.Now()
    l.PrintFormatted(message, "ERROR", ErrorColour, currentTime)
}

func(l Logger)  LogDebug(message string) {
    if l.MinimumSeverity > 5 { return }
    currentTime := time.Now()
    l.PrintFormatted(message, "DEBUG", DebugColour, currentTime)
}


func (l Logger) LogCritical(message string) {
    if l.MinimumSeverity > 1 { return }
    currentTime := time.Now()
    l.PrintFormatted(message, "CRITICAL", CriticalColour, currentTime)       
    os.Exit(1)
}
