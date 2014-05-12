package adapter

import (
	"fmt"
	"regexp"
)

func Convert(Text string) string {
	replaceHedding6(&Text)
	replaceHedding5(&Text)
	replaceHedding4(&Text)
	replaceHedding3(&Text)
	replaceHedding2(&Text)
	replaceHedding1(&Text)
	fmt.Printf(Text)
	return Text
}

/**
*	Replace <h1> test </h1>
 */
func replaceHedding1(Text *string) {
	reg := regexp.MustCompile(`^\s*#{1}(.*?)$`)
	*Text = reg.ReplaceAllString(*Text, "<h1>$1</h1>")
}

/**
*	Replace <h2> test </h2>
 */
func replaceHedding2(Text *string) {
	reg := regexp.MustCompile(`^\s*#{2}(.*?)$`)
	*Text = reg.ReplaceAllString(*Text, "<h2>$1</h2>")
}

/**
*	Replace <h3> test </h3>
 */
func replaceHedding3(Text *string) {
	reg := regexp.MustCompile(`^\s*#{3}(.*?)$`)
	*Text = reg.ReplaceAllString(*Text, "<h3>$1</h3>")
}

/**
*	Replace <h4> test </h4>
 */
func replaceHedding4(Text *string) {
	reg := regexp.MustCompile(`^\s*#{4}(.*?)$`)
	*Text = reg.ReplaceAllString(*Text, "<h4>$1</h4>")
}

/**
*	Replace <h5> test </h5>
 */
func replaceHedding5(Text *string) {
	reg := regexp.MustCompile(`^\s*#{5}(.*?)$`)
	*Text = reg.ReplaceAllString(*Text, "<h5>$1</h5>")
}

/**
*	Replace <h6> test </h6>
 */
func replaceHedding6(Text *string) {
	reg := regexp.MustCompile(`^\s*#{6}(.*?)$`)
	*Text = reg.ReplaceAllString(*Text, "<h6>$1</h6>")
}
