package klinutils

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"sync"
)

func Genuuid() ([]byte, error) {
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		return out, err
	}
	return bytes.Split(out, []byte("\n"))[0], nil
}
func captureOutput(f func()) (string, error) {
	reader, writer, err := os.Pipe()
	if err != nil {
		return "", err
	}
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		log.SetOutput(os.Stderr)
	}()
	os.Stdout = writer
	os.Stderr = writer
	log.SetOutput(writer)
	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, reader)
		out <- buf.String()
	}()
	wg.Wait()
	f()
	writer.Close()
	return <-out, nil
}
func Genmac() ([]byte, error) {
	buf := make([]byte, 6)
	_, err := rand.Read(buf)
	if err != nil {
		return []byte(""), err
	}
	// Set the local bit
	buf[0] |= (buf[0] | 2) & 0xfe
	re, err := captureOutput(func() {
		fmt.Printf("%02x:%02x:%02x:%02x:%02x:%02x", buf[0], buf[1], buf[2], buf[3], buf[4], buf[5])
	})
	if err != nil {
		return []byte(""), err
	}
	return []byte(re), nil
}