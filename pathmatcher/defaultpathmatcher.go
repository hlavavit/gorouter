package pathmatcher

import (
	"regexp"
	"strings"
)

// DefaultPathMatcher matches url paths with standard syntax
// * as wildcard, {param:pattern} path variables, ** wildcard for multiple separated paths
// its inspired by https://github.com/spring-projects/spring-framework/blob/master/spring-core/src/main/java/org/springframework/util/AntPathMatcher.java
type DefaultPathMatcher struct {
	pathSeparator        string
	pattern              string
	patternTokens        []string
	patternTokenMatchers map[string]*PathVariableMatcher
	caseSensitive        bool
	urlVariableRegexp    regexp.Regexp
	wildcardChars        []rune
}

//NewDefaultPathMatcher constract url pattern matcher based on url pattern
func NewDefaultPathMatcher(pattern string) *DefaultPathMatcher {

	matcher := &DefaultPathMatcher{
		pathSeparator:        "/",
		pattern:              pattern,
		caseSensitive:        true,
		urlVariableRegexp:    *regexp.MustCompile("\\{[^/]+?\\}"),
		wildcardChars:        []rune{'*', '?', '{'},
		patternTokenMatchers: make(map[string]*PathVariableMatcher),
	}
	matcher.patternTokens = matcher.tokenizePath(pattern)
	for _, patternToken := range matcher.patternTokens {
		//make sure all are cached on startup, to avoid runtime failures in cases when pattern wont compile.
		matcher.getPatternTokenMatcher(patternToken)
	}
	return matcher
}

// Match check whenever path matches pattern
func (matcher *DefaultPathMatcher) Match(path string) (bool, map[string]string) {
	urlVariables := make(map[string]string, 0)

	if strings.HasPrefix(path, matcher.pathSeparator) != strings.HasPrefix(matcher.pattern, matcher.pathSeparator) {
		return false, urlVariables
	}

	if !matcher.isPotentialMatch(path) {
		return false, urlVariables
	}
	pattIdxStart := 0
	pattIdxEnd := len(matcher.patternTokens) - 1

	pathTokens := matcher.tokenizePath(path)
	pathIdxStart := 0
	pathIdxEnd := len(pathTokens) - 1

	for pattIdxStart <= pattIdxEnd && pathIdxStart <= pathIdxEnd {
		pattToken := matcher.patternTokens[pattIdxStart]
		if "**" == (pattToken) {
			break
		}
		match, variables := matcher.getPatternTokenMatcher(pattToken).Match(pathTokens[pathIdxStart])
		if !match {
			return false, urlVariables
		}
		for k, v := range variables {
			urlVariables[k] = v
		}
		pattIdxStart++
		pathIdxStart++
	}
	if pathIdxStart > pathIdxEnd {
		if pattIdxStart > pattIdxEnd {
			return strings.HasSuffix(matcher.pattern, matcher.pathSeparator) == strings.HasSuffix(path, matcher.pathSeparator), urlVariables
		}
		if pattIdxStart == pattIdxEnd && matcher.patternTokens[pattIdxStart] == "*" && strings.HasSuffix(path, matcher.pathSeparator) {
			return true, urlVariables
		}
		// path is exhausted
		for i := pattIdxStart; i <= pattIdxEnd; i++ {
			if matcher.patternTokens[i] != "**" {
				return false, urlVariables
			}
		}
		return true, urlVariables
	}

	if pattIdxStart > pattIdxEnd {
		// String not exhausted, but pattern is. Failure.
		return false, urlVariables
	}

	for pattIdxStart <= pattIdxEnd && pathIdxStart <= pathIdxEnd {
		pattDir := matcher.patternTokens[pattIdxEnd]
		if pattDir == "**" {
			break
		}
		match, variables := matcher.getPatternTokenMatcher(pattDir).Match(pathTokens[pathIdxEnd])
		if !match {
			return false, urlVariables
		}
		for k, v := range variables {
			urlVariables[k] = v
		}
		pattIdxEnd--
		pathIdxEnd--
	}
	if pathIdxStart > pathIdxEnd {
		// String is exhausted
		for i := pattIdxStart; i <= pattIdxEnd; i++ {
			if matcher.patternTokens[i] != "**" {
				return false, urlVariables
			}
		}
		return true, urlVariables
	}

	for pattIdxStart != pattIdxEnd && pathIdxStart <= pathIdxEnd {
		patIdxTmp := -1

		for i := pattIdxStart + 1; i <= pattIdxEnd; i++ {
			if matcher.patternTokens[i] == "**" {
				patIdxTmp = i
				break
			}
		}
		if patIdxTmp == pattIdxStart+1 {
			// '**/**' situation, so skip one
			pattIdxStart++
			continue
		}
		// Find the pattern between padIdxStart & padIdxTmp in str between
		// strIdxStart & strIdxEnd
		patLength := (patIdxTmp - pattIdxStart - 1)
		strLength := (pathIdxEnd - pathIdxStart + 1)
		foundIdx := -1

	strLoop:
		for i := 0; i <= strLength-patLength; i++ {
			for j := 0; j < patLength; j++ {
				subPat := matcher.patternTokens[pattIdxStart+j+1]
				subStr := pathTokens[pathIdxStart+i+j]
				match, variables := matcher.getPatternTokenMatcher(subPat).Match(subStr)
				if !match {
					continue strLoop
				}
				for k, v := range variables {
					urlVariables[k] = v
				}
			}
			foundIdx = pathIdxStart + i
			break
		}
		if foundIdx == -1 {
			return false, urlVariables
		}
		pattIdxStart = patIdxTmp
		pathIdxStart = foundIdx + patLength
	}

	for i := pattIdxStart; i <= pattIdxEnd; i++ {
		if matcher.patternTokens[i] != "**" {
			return false, urlVariables
		}
	}
	return true, urlVariables
}

func (matcher *DefaultPathMatcher) getPatternTokenMatcher(pattern string) (result *PathVariableMatcher) {
	result = matcher.patternTokenMatchers[pattern]
	if result == nil {
		result = NewPathVariableMatcher(pattern)
		matcher.patternTokenMatchers[pattern] = result
	}
	return
}

func (matcher *DefaultPathMatcher) isPotentialMatch(path string) bool {
	var pos int

	for _, patternToken := range matcher.patternTokens {
		skipped := matcher.skipSeparator(path, pos)
		pos += skipped
		skipped = matcher.skipSegment(path, pos, patternToken)
		if skipped < len(patternToken) {
			return skipped > 0 || (len(patternToken) > 0 && matcher.isWildcardChar([]rune(patternToken)[0]))
		}
		pos += skipped
	}

	return true
}

func (matcher *DefaultPathMatcher) skipSeparator(path string, pos int) (skipped int) {
	for strings.HasPrefix(path[pos+skipped:], matcher.pathSeparator) {
		skipped += len(matcher.pathSeparator)
	}
	return
}

func (matcher *DefaultPathMatcher) skipSegment(path string, pos int, segment string) (skipped int) {
	for _, char := range segment {
		if matcher.isWildcardChar(char) {
			return
		}
		currPos := pos + skipped
		if currPos >= len(path) {
			return 0
		}
		if char == []rune(path[currPos:])[0] {
			skipped += len(string(char))
		}
	}
	return
}
func (matcher *DefaultPathMatcher) isWildcardChar(char rune) bool {
	for _, candidate := range matcher.wildcardChars {
		if candidate == char {
			return true
		}
	}
	return false
}

func (matcher *DefaultPathMatcher) tokenizePath(path string) []string {
	tokenCandidates := strings.Split(path, matcher.pathSeparator)
	//remove empty
	tokens := make([]string, 0, len(tokenCandidates))
	for _, token := range tokenCandidates {
		if len(token) > 0 {
			tokens = append(tokens, token)
		}
	}
	return tokens

}
