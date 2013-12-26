package concat

import (
	"fmt"
	"io/ioutil"
)

func Files(files ...string) string {
	content := ""

	for _, file := range files {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}

		content += fmt.Sprintln(string(data))
	}

	return content
}
