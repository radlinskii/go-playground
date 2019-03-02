package file

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// SetEnv sets env variables specified in the .env file in cwd.
func SetEnv() {
	data, err := ioutil.ReadFile(".env")
	if err != nil {
		log.Fatalln(err)
	}

	content := string(data)
	lines := strings.Split(content, "\n")

	for _, line := range lines {
		if line != "" {
			varinfo := strings.SplitN(line, "=", 2)
			if len(varinfo) != 2 {
				log.Fatalln("Error parsing .env file")
			}
			err := os.Setenv(varinfo[0], varinfo[1])
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}
