package pathmatcher

import (
	"fmt"
	"testing"

	log "github.com/sirupsen/logrus"
)

//https://github.com/spring-projects/spring-framework/blob/master/spring-core/src/test/java/org/springframework/util/AntPathMatcherTests.java
func TestPathVariableMatcher(t *testing.T) {
	log.SetLevel(log.TraceLevel)
	//basic
	checkPathVariable(t, "{hotel}", "1", true, map[string]string{"hotel": "1"})
	checkPathVariable(t, "h?tel", "hotel", true, map[string]string{})
	checkPathVariable(t, "hotel", "hotel", true, map[string]string{})
	checkPathVariable(t, "hotel", "hell", false, map[string]string{})
	checkPathVariable(t, "/*/hotels/*/{hotel}", "/foo/hotels/bar/1", true, map[string]string{"hotel": "1"})
	checkPathVariable(t, "/{page}.html", "/42.html", true, map[string]string{"page": "42"})
	checkPathVariable(t, "/{page}.*", "/42.html", true, map[string]string{"page": "42"})
	checkPathVariable(t, "/A-{B}-C", "/A-b-C", true, map[string]string{"B": "b"})
	checkPathVariable(t, "/{name}.{extension}", "/test.html", true, map[string]string{"name": "test", "extension": "html"})
	//regexp
	checkPathVariable(t, "{symbolicName:[\\w\\.]+}-{version:[\\w\\.]+}.jar", "com.example-1.0.0.jar", true, map[string]string{"symbolicName": "com.example", "version": "1.0.0"})
	checkPathVariable(t, "{symbolicName:[\\w\\.]+}-sources-{version:[\\w\\.]+}.jar", "com.example-sources-1.0.0.jar", true, map[string]string{"symbolicName": "com.example", "version": "1.0.0"})

	//regexp quialifiers
	checkPathVariable(t, "{symbolicName:[\\p{L}\\.]+}-sources-{version:[\\p{N}\\.]+}.jar", "com.example-sources-1.0.0.jar", true, map[string]string{"symbolicName": "com.example", "version": "1.0.0"})
	checkPathVariable(t, "{symbolicName:[\\w\\.]+}-sources-{version:[\\d\\.]+}-{year:\\d{4}}{month:\\d{2}}{day:\\d{2}}.jar", "com.example-sources-1.0.0-20100220.jar", true, map[string]string{"symbolicName": "com.example", "version": "1.0.0", "year": "2010", "month": "02", "day": "20"})
	checkPathVariable(t, "{symbolicName:[\\p{L}\\.]+}-sources-{version:[\\p{N}\\.\\{\\}]+}.jar", "com.example-sources-1.0.0.{12}.jar", true, map[string]string{"symbolicName": "com.example", "version": "1.0.0.{12}"})
	//checkPathVariable(t, "", "", true, map[string]string{"": ""})

}

func checkPathVariable(t *testing.T, pattern string, value string, expected bool, extracted map[string]string) {
	match, variables := NewPathVariableMatcher(pattern).Match(value)
	if match != expected {
		t.Fatal(fmt.Sprintf("For pattern=%v and value=%v expected match to be %v", pattern, value, expected))
	}
	for key, varValue := range extracted {
		if variables[key] != varValue {
			t.Fatal(fmt.Sprintf("For pattern=%v and value=%v expected variable %v to be %v", pattern, value, key, varValue))
		}
	}
}
