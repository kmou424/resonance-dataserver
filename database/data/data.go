package data

import (
	"embed"
	"github.com/charmbracelet/log"
	"regexp"
	"strings"
)

//go:embed files
var data embed.FS

func ReadGoodsCitiesMapper() []byte {
	content, err := data.ReadFile("files/goods_cities_mapper.json")
	if err != nil {
		log.Fatal("can't read mapper data", "error", err)
	}
	return removeJsonComment(string(content))
}

func removeJsonComment(src string) []byte {
	var sb strings.Builder
	commentStartReg := regexp.MustCompile(`-.*->`)
	commentEndReg := regexp.MustCompile(`<-.*-`)
	lines := strings.Split(src, "\n")
	for _, line := range lines {
		commentStart := commentStartReg.FindString(line)
		commentEnd := commentEndReg.FindString(line)
		// filter comments
		if commentStart != "" || commentEnd != "" {
			continue
		}
		sb.WriteString(line)
		sb.WriteRune('\n')
	}
	res := sb.String()

	removeRedundantComma := func() {
		reg := regexp.MustCompile(`}.*\n]`)
		src := reg.FindString(res)
		dst := strings.Replace(src, ",", "", 1)
		res = strings.Replace(res, src, dst, 1)
	}
	removeRedundantComma()

	return []byte(res)
}
