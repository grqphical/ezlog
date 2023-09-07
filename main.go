// Simple logging library that is easily embedable into larger applications.
package ezlog

import (
	"fmt"
	"os"
    "time"
)

// Colours used by the logger.
const (
        // Normal
        InfoColour    = "\033[1;39m%s\033[0m\n"
        // Yellow/gold
        WarningColour = "\033[1;33m%s\033[0m\n"
        // Red
        ErrorColour   = "\033[1;31m%s\033[0m\n"
        // Cyan
        DebugColour   = "\033[1;36m%s\033[0m\n"
        // Dark Red
        CriticalColour = "\033[38;2;85;0;0m%s\033[0m\n"

)

// Internal method to print logs nicely.
func PrintFormatted(message string, severity string, colour string, date time.Time) {
    fmt.Printf(colour, fmt.Sprintf("[%s] %s - %s", severity, date.Format(time.RFC822), message))
}

// Logs a message at an info level
func LogInfo(message string) {
    currentTime := time.Now()
    PrintFormatted(message, "INFO", InfoColour, currentTime) 
}

// Logs a message as a warning
func LogWarning(message string) {
    currentTime := time.Now()
    PrintFormatted(message, "WARNING", WarningColour, currentTime)
}

// Logs a message as an error
func LogError(message string) {
    currentTime := time.Now()
    PrintFormatted(message, "ERROR", ErrorColour, currentTime)
}

// Logs a message as debug
func LogDebug(message string) {
    currentTime := time.Now()
    PrintFormatted(message, "DEBUG", DebugColour, currentTime)
}

// Logs a message as an error and crashes the app
func LogCritical(message string) {
    currentTime := time.Now()
    PrintFormatted(message, "CRITICAL", CriticalColour, currentTime)       
    os.Exit(1)
}

