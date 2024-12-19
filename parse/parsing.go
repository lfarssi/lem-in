package parsing

import (
	"strings"
)

func (g *Farm) ParseFile(file string) string {
	var res string
	arr := strings.Split(file, "\n")
	for _, line := range arr {
		if strings.HasPrefix(line, "#") {
			res += "comment"
		} else if strings.Contains(line, "##start") {
			res += "start"
		} else if strings.Contains(line, "##end") {
			res += "end"
		}
	}

	return file
}
