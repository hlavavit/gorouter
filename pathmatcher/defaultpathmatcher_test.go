package pathmatcher

import (
	"fmt"
	"testing"

	log "github.com/sirupsen/logrus"
)

//https://github.com/spring-projects/spring-framework/blob/master/spring-core/src/test/java/org/springframework/util/AntPathMatcherTests.java
func TestDefaultPathMatcher(t *testing.T) {
	log.SetLevel(log.InfoLevel)
	checkPathMatcher(t, "test", "test", true, map[string]string{})
	checkPathMatcher(t, "/test", "/test", true, map[string]string{})
	checkPathMatcher(t, "https://example.org", "https://example.org", true, map[string]string{})
	checkPathMatcher(t, "/test.jpg", "test.jpg", false, map[string]string{})
	checkPathMatcher(t, "test", "/test", false, map[string]string{})
	checkPathMatcher(t, "/test", "test", false, map[string]string{})
	//? wildcard
	checkPathMatcher(t, "t?st", "test", true, map[string]string{})
	checkPathMatcher(t, "??st", "test", true, map[string]string{})
	checkPathMatcher(t, "tes?", "test", true, map[string]string{})
	checkPathMatcher(t, "te??", "test", true, map[string]string{})
	checkPathMatcher(t, "?es?", "test", true, map[string]string{})
	checkPathMatcher(t, "tes?", "tes", false, map[string]string{})
	checkPathMatcher(t, "tes?", "testt", false, map[string]string{})
	checkPathMatcher(t, "tes?", "tsst", false, map[string]string{})
	// * wildcard
	checkPathMatcher(t, "*", "test", true, map[string]string{})
	checkPathMatcher(t, "test*", "test", true, map[string]string{})
	checkPathMatcher(t, "test*", "testTest", true, map[string]string{})
	checkPathMatcher(t, "test/*", "test/Test", true, map[string]string{})
	checkPathMatcher(t, "test/*", "test/t", true, map[string]string{})
	checkPathMatcher(t, "test/*", "test/", true, map[string]string{})
	checkPathMatcher(t, "*test*", "AnothertestTest", true, map[string]string{})
	checkPathMatcher(t, "*test", "Anothertest", true, map[string]string{})
	checkPathMatcher(t, "*.*", "test.", true, map[string]string{})
	checkPathMatcher(t, "*.*", "test.test", true, map[string]string{})
	checkPathMatcher(t, "*.*", "test.test.test", true, map[string]string{})
	checkPathMatcher(t, "test*aaa", "testblaaaa", true, map[string]string{})
	checkPathMatcher(t, "test*", "tst", false, map[string]string{})
	checkPathMatcher(t, "test*", "tsttest", false, map[string]string{})
	checkPathMatcher(t, "test*", "test/", false, map[string]string{})
	checkPathMatcher(t, "test*", "test/t", false, map[string]string{})
	checkPathMatcher(t, "test/*", "test", false, map[string]string{})
	checkPathMatcher(t, "*test*", "tsttst", false, map[string]string{})
	checkPathMatcher(t, "*test", "tsttst", false, map[string]string{})
	checkPathMatcher(t, "*.*", "tsttst", false, map[string]string{})
	checkPathMatcher(t, "test*aaa", "test", false, map[string]string{})
	checkPathMatcher(t, "test*aaa", "testblaaab", false, map[string]string{})

	// test matching with ?'s and /'s
	checkPathMatcher(t, "/?", "/a", true, map[string]string{})
	checkPathMatcher(t, "/?/a", "/a/a", true, map[string]string{})
	checkPathMatcher(t, "/a/?", "/a/b", true, map[string]string{})
	checkPathMatcher(t, "/??/a", "/aa/a", true, map[string]string{})
	checkPathMatcher(t, "/a/??", "/a/bb", true, map[string]string{})
	checkPathMatcher(t, "/?", "/a", true, map[string]string{})

	// test matching with **'s
	checkPathMatcher(t, "/**", "/testing/testing", true, map[string]string{})
	checkPathMatcher(t, "/*/**", "/testing/testing", true, map[string]string{})
	checkPathMatcher(t, "/**/*", "/testing/testing", true, map[string]string{})
	checkPathMatcher(t, "/bla/**/bla", "/bla/testing/testing/bla", true, map[string]string{})
	checkPathMatcher(t, "/bla/**/bla", "/bla/testing/testing/bla/bla", true, map[string]string{})
	checkPathMatcher(t, "/**/test", "/bla/bla/test", true, map[string]string{})
	checkPathMatcher(t, "/bla/**/**/bla", "/bla/bla/bla/bla/bla/bla", true, map[string]string{})
	checkPathMatcher(t, "/bla*bla/test", "/blaXXXbla/test", true, map[string]string{})
	checkPathMatcher(t, "/*bla/test", "/XXXbla/test", true, map[string]string{})
	checkPathMatcher(t, "/bla*bla/test", "/blaXXXbl/test", false, map[string]string{})
	checkPathMatcher(t, "/*bla/test", "XXXblab/test", false, map[string]string{})
	checkPathMatcher(t, "/*bla/test", "XXXbl/test", false, map[string]string{})

	checkPathMatcher(t, "/????", "/bala/bla", false, map[string]string{})
	checkPathMatcher(t, "/**/*bla", "/bla/bla/bla/bbb", false, map[string]string{})

	checkPathMatcher(t, "/*bla*/**/bla/**", "/XXXblaXXXX/testing/testing/bla/testing/testing/", true, map[string]string{})
	checkPathMatcher(t, "/*bla*/**/bla/*", "/XXXblaXXXX/testing/testing/bla/testing", true, map[string]string{})
	checkPathMatcher(t, "/*bla*/**/bla/**", "/XXXblaXXXX/testing/testing/bla/testing/testing", true, map[string]string{})
	checkPathMatcher(t, "/*bla*/**/bla/**", "/XXXblaXXXX/testing/testing/bla/testing/testing.jpg", true, map[string]string{})

	checkPathMatcher(t, "*bla*/**/bla/**", "XXXblaXXXX/testing/testing/bla/testing/testing/", true, map[string]string{})
	checkPathMatcher(t, "*bla*/**/bla/*", "XXXblaXXXX/testing/testing/bla/testing", true, map[string]string{})
	checkPathMatcher(t, "*bla*/**/bla/**", "XXXblaXXXX/testing/testing/bla/testing/testing", true, map[string]string{})
	checkPathMatcher(t, "*bla*/**/bla/*", "XXXblaXXXX/testing/testing/bla/testing/testing", false, map[string]string{})

	checkPathMatcher(t, "/x/x/**/bla", "/x/x/x/", false, map[string]string{})

	checkPathMatcher(t, "/foo/bar/**", "/foo/bar", true, map[string]string{})

	checkPathMatcher(t, "", "", true, map[string]string{})

	checkPathMatcher(t, "/{bla}.*", "/testing.html", true, map[string]string{"bla": "testing"})

	checkPathMatcher(t, "/test", "", false, map[string]string{})
	checkPathMatcher(t, "/", "", false, map[string]string{})

	checkPathMatcher(t, "{hotel}", "1", true, map[string]string{"hotel": "1"})
	checkPathMatcher(t, "h?tel", "hotel", true, map[string]string{})
	checkPathMatcher(t, "hotel", "hotel", true, map[string]string{})
	checkPathMatcher(t, "hotel", "hell", false, map[string]string{})
	checkPathMatcher(t, "/*/hotels/*/{hotel}", "/foo/hotels/bar/1", true, map[string]string{"hotel": "1"})
	checkPathMatcher(t, "/{page}.html", "/42.html", true, map[string]string{"page": "42"})
	checkPathMatcher(t, "/{page}.*", "/42.html", true, map[string]string{"page": "42"})
	checkPathMatcher(t, "/A-{B}-C", "/A-b-C", true, map[string]string{"B": "b"})
	checkPathMatcher(t, "/{name}.{extension}", "/test.html", true, map[string]string{"name": "test", "extension": "html"})
	//regexp
	checkPathMatcher(t, "{symbolicName:[\\w\\.]+}-{version:[\\w\\.]+}.jar", "com.example-1.0.0.jar", true, map[string]string{"symbolicName": "com.example", "version": "1.0.0"})
	checkPathMatcher(t, "{symbolicName:[\\w\\.]+}-sources-{version:[\\w\\.]+}.jar", "com.example-sources-1.0.0.jar", true, map[string]string{"symbolicName": "com.example", "version": "1.0.0"})

	//regexp quialifiers
	checkPathMatcher(t, "{symbolicName:[\\p{L}\\.]+}-sources-{version:[\\p{N}\\.]+}.jar", "com.example-sources-1.0.0.jar", true, map[string]string{"symbolicName": "com.example", "version": "1.0.0"})
	checkPathMatcher(t, "{symbolicName:[\\w\\.]+}-sources-{version:[\\d\\.]+}-{year:\\d{4}}{month:\\d{2}}{day:\\d{2}}.jar", "com.example-sources-1.0.0-20100220.jar", true, map[string]string{"symbolicName": "com.example", "version": "1.0.0", "year": "2010", "month": "02", "day": "20"})
	checkPathMatcher(t, "{symbolicName:[\\p{L}\\.]+}-sources-{version:[\\p{N}\\.\\{\\}]+}.jar", "com.example-sources-1.0.0.{12}.jar", true, map[string]string{"symbolicName": "com.example", "version": "1.0.0.{12}"})

}

func checkPathMatcher(t *testing.T, pattern string, value string, expected bool, extracted map[string]string) {
	match, variables := NewDefaultPathMatcher(pattern).Match(value)
	if match != expected {
		t.Fatal(fmt.Sprintf("For pattern=%v and value=%v expected match to be %v", pattern, value, expected))
	}
	for key, varValue := range extracted {
		if variables[key] != varValue {
			t.Fatal(fmt.Sprintf("For pattern=%v and value=%v expected variable %v to be %v", pattern, value, key, varValue))
		}
	}
}
