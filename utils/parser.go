package utils

import (
	"errors"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)
var r = []*regexp.Regexp{
	regexp.MustCompile(`(.*) - (\d{1,4}(?:\.\d{1,2})?)(?:v\d{1,2})?(?: )?(?:END)?(.*)`),
	regexp.MustCompile(`(.*)[\[\ E](\d{1,4}|\d{1,4}\.\d{1,2})(?:v\d{1,2})?(?: )?(?:END)?[\]\ ](.*)`),
	regexp.MustCompile(`(.*)\[(?:第)?(\d*\.*\d*)[话集話](?:END)?\](.*)`),
	regexp.MustCompile(`(.*)第(\d*\.*\d*)[话話集](?:END)?(.*)`),
	regexp.MustCompile(`(.*)(?:S\d{2})?EP?(\d+)(.*)`),
}

var re = regexp.MustCompile(`([Ss]|Season )\d{1,3}`)
var reSeason = regexp.MustCompile(`([Ss]|Season )(\d{1,3})`)
// var filter = regexp.MustCompile(`[\[\]]`)
type Information struct{
	Title string
	Season int
	Episode int
	Extension string
	Language string
}

func getSeasonAndTitle(seasonAndTitle string) (string, int) {
	title := re.ReplaceAllString(seasonAndTitle, "")
	// title = filter.ReplaceAllString(title, "")
	title = strings.TrimSpace(title)
	match := reSeason.FindStringSubmatch(seasonAndTitle)
	var season int
	if len(match) > 2 {
		season, _ = strconv.Atoi(match[2])
	} else {
		season = 1
	}

	return title, season
}

func getGroup(groupAndTitle string) (string, string) {
	re := regexp.MustCompile(`[\[\]()【】（）]`)
	parts := re.Split(groupAndTitle, -1)
	var filteredParts []string
	for _, part := range parts {
		if part != "" {
			filteredParts = append(filteredParts, part)
		}
	}

	if len(filteredParts) > 1 {
		if matched, _ := regexp.MatchString(`^\d+$`, filteredParts[1]); matched {
			return "", groupAndTitle
		}
		return filteredParts[0], filteredParts[1]
	} else {
		return "", filteredParts[0]
	}
}

func containsSubstrings(str string, substrings []string) bool {
	str = strings.ToLower(str)
	for _, substring := range substrings {
			if strings.Contains(str, substring) {
					return true
			}
	}
	return false
}

func getLanguage(name string)string{
	tc := []string{"tc", "cht", "繁", "zh-tw"}
	sc := []string{"sc", "chs", "简", "zh"}
	if containsSubstrings(name,tc){
		return ".tc"
	}
	if containsSubstrings(name,sc){
		return ".sc"
	}
	return ""
}

func Parse(title string)(*Information,error){
	var info Information
	
	for  _,i := range r{
		match := i.FindStringSubmatch(title)
		if match != nil {
			_,t:=getGroup(match[1])
			info.Title,info.Season = getSeasonAndTitle(t)
			info.Episode,_ = strconv.Atoi(match[2])
			info.Extension = filepath.Ext(title)
			if info.Extension == ".ass"||info.Extension == ".ssa"||info.Extension ==".srt"{
				info.Language = getLanguage(match[3])
			}
			return &info,nil
	} 
}
return &info,errors.New("parse error")
}