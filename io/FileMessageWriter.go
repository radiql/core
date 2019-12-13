package io

import (
	"fmt"
	"log"
	"os"
)

// FileMessageWriter writes messages to a specified file
type FileMessageWriter struct {
	// Usually set up on initialisation but can be rolled over into another file path.
	TargetFilePath <-chan string
	
	// A stream of messages to be written.
	Msg            <-chan interface{}
}

// Process writes the incoming message to a file.
func (c *FileMessageWriter) Process() {

	targetFile := TargetFile{}

	for msg := range c.Msg {

		filePath := <-c.TargetFilePath
		f, _ := targetFile.Get(filePath)
		defer f.Close()

		fmt.Fprintf(&f, "%v", msg)
	}

}

/*
// LogFileWriter writes log events to log files
type LogFileWriter struct {
	LogFileSpec <-chan string
	Event       <-chan string

	targetFile TargetFile
}

// KafkaSteamWriter injects events into a Kafka event stream
type KafkaSteamWriter struct {
	KafkaSteamSpec <-chan string
	Event          <-chan string
}

// KafkaStreamReader reads events from a Kafka event stream.
type KafkaStreamReader struct {
	KafkaStreamSpec <-chan string
	Event           chan<- string
}
*/

// TargetFile provides access to a file given a file name.
type TargetFile struct {
	file        os.File
	currentPath string
}

// Get returns a file for appending or creates one.
func (t *TargetFile) Get(filePath string) (os.File, error) {
	// Checks to see if the path has changed.
	if filePath != t.currentPath {
		// close the old file
		if &t.file != nil {
			t.file.Close()
			t.currentPath = ""
		}

		// If the file doesn't exist, create it, or append to the file
		f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		t.file = *f
		t.currentPath = filePath

	}
	return t.file, nil
}
