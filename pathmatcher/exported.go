package pathmatcher

// Matcher interface with single method that matches url string returns true if matches and map of extraced variables
type Matcher interface {
	//Match path agains underlying pattern and extract path variables
	Match(string) (bool, map[string]string)
}

// NewDefaultMatcher creates new instance of DefaultMatcher based on pattern
// DefaultMatcher matches url paths with support for following wildcards and path variables
//
// * /test matches exact urls /test and /test/
//
// * /test/* matches /test/ and /test/something but not /test/something/other
//
// * /test/** matches any path starting with /test/ meaning /test/something/other
//
// * /t?st matches /test and /tast
//
// * /test/{param} extracts path variable and names it param for /test/value1 matches and returns param: value1
func NewDefaultMatcher(pattern string) Matcher {
	return newDefaultPathMatcher(pattern)
}

//Match path agains underlying pattern and extract path variables
func (matcher defaultPathMatcher) Match(path string) (bool, map[string]string) {
	return matcher.match(path)
}
