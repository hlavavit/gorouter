package pathmatcher

import (
	"fmt"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
)

var pathVariableGlobMatcher = regexp.MustCompile("\\?|\\*|\\{((?:\\{[^/]+?\\}|[^/{}]|\\\\[{}])+?)\\}")

const (
	defaultPathVariablePattern = "(.*)"
)

// pathVariableMatcher matches and extracts patterns from url tokens
//
// Ispired by AntPathStringMatcher class https://github.com/spring-projects/spring-framework/blob/master/spring-core/src/main/java/org/springframework/util/AntPathMatcher.java
type pathVariableMatcher struct {
	variableNames []string
	pattern       *regexp.Regexp
}

// NewpathVariableMatcher creates new pathVariableMatcher. provided simple pattern /*-t?st/*-{val:[a-z]}/{val2}.ext constucts regexp patern.
func newPathVariableMatcher(pattern string) *pathVariableMatcher {
	matches := pathVariableGlobMatcher.FindAllStringSubmatchIndex(pattern, -1)
	end := 0
	patternBuilder := ""
	variableNames := make([]string, 0)
	for _, matchIndex := range matches {
		patternBuilder += regexp.QuoteMeta(pattern[end:matchIndex[0]])
		match := pattern[matchIndex[0]:matchIndex[1]]
		if "?" == match {
			patternBuilder += "."
		} else if "*" == match {
			patternBuilder += ".*"
		} else if strings.HasPrefix(match, "{") && strings.HasSuffix(match, "}") {
			group := pattern[matchIndex[2]:matchIndex[3]] //2 and 3 are where is stored first group match
			colonIdx := strings.Index(group, ":")

			if colonIdx < 0 {
				patternBuilder += defaultPathVariablePattern
				variableNames = append(variableNames, group)
			} else {
				variableNames = append(variableNames, group[:colonIdx])
				patternBuilder += "(" + group[colonIdx+1:] + ")"
			}
		}
		end = matchIndex[1]
	}
	patternBuilder += regexp.QuoteMeta(pattern[end:])
	patternBuilder = "^" + patternBuilder + "$"
	log.Tracef("pathVariableMatcher created from pattern=%v, with extracted variables=%v and pattern=%v", pattern, variableNames, patternBuilder)
	matcher := pathVariableMatcher{
		variableNames: variableNames,
		pattern:       regexp.MustCompile(patternBuilder),
	}
	return &matcher
}

//Match evaluates whenever matcher matches provided string and extracts variables from said string
func (m pathVariableMatcher) match(str string) (match bool, variables map[string]string) {
	variables = make(map[string]string)
	matches := m.pattern.FindStringSubmatchIndex(str)
	expectedMatches := 1 + len(m.variableNames)
	expectedMatches *= 2 // each mach has start and end index

	if len(matches) != expectedMatches {
		log.Tracef("pathVariableMatcher failed to match string=%v to pattern %v", str, *m.pattern)
		return
	}
	groups := matches[2:]
	for i, value := range m.variableNames {
		start, end := i*2, i*2+1
		variables[value] = str[groups[start]:groups[end]]
	}

	return true, variables
}

func (m pathVariableMatcher) String() (text string) {
	text = fmt.Sprint("variableNames:", m.variableNames, ", pattern:", *m.pattern)
	return
}
