package mapperdata

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
	var sb strings.Builder
	commentReg := regexp.MustCompile(`//.*`)
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		result := commentReg.FindString(line)
		// filter comments
		if result != "" {
			continue
		}
		sb.WriteString(line)
		sb.WriteRune('\n')
	}
	return []byte(sb.String())
}
