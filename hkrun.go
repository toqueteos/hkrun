package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

var logger = log.New(os.Stdout, "[hkrun] ", 0)

func main() {
	// Read config file
	bytes, err := ioutil.ReadFile("hk.env")
	if err != nil {
		logger.Fatal(err)
	}

	// Remove empty lines and/or extra blank characters.
	body := strings.TrimSpace(string(bytes))

	vars := strings.Split(body, "\n")
	if len(vars) != 0 {
		logger.Println("Setting environment...")
	}

	for _, line := range vars {
		e := strings.SplitN(line, "=", 2)
		fmt.Printf("\t%s = %q\n", e[0], e[1])
		os.Setenv(e[0], e[1])
	}

	cmd := exec.Command("gin")
	out, outErr := cmd.StdoutPipe()
	if outErr != nil {
		logger.Fatal(outErr)
	}
	cmd.Start()
	go io.Copy(os.Stdout, out)
	cmd.Wait()
	out.Close()
}
