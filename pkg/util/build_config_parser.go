package util

import (
	"io/ioutil"
	"regexp"
)

const (
	DEPENDENCY = "dependency"
)

func ParseConfigComments(languageLineComment string, functionLocation string) (map[string][]string, error) {
	fileContentBytes, err := ioutil.ReadFile(functionLocation)
	if err != nil {
		return nil, err
	}
	fileContent := string(fileContentBytes)

	result := make(map[string][]string, 0)

	// Group 1 contains key, Group 2 contains value
	commentRegex := regexp.MustCompile(`(?m)^[[:space:]]*\Q` + languageLineComment + `\E[[:space:]]*kfn:([a-zA-Z\-]+)[[:space:]]+(.*)$`)
	submatches := commentRegex.FindAllStringSubmatch(fileContent, -1)

	for _, submatch := range submatches {
		key := submatch[1]
		value := submatch[2]
		entries, ok := result[key]
		if ok {
			result[key] = append(entries, value)
		} else {
			result[key] = make([]string, 1)
			result[key][0] = value
		}
	}

	return result, nil
}
