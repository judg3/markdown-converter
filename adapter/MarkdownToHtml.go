package adapter

import (
	"regexp"
)

func Convert(Text string) string {
	replaceHedding6(&Text)
	replaceHedding5(&Text)
	replaceHedding4(&Text)
	replaceHedding3(&Text)
	replaceHedding2(&Text)
	replaceHedding1(&Text)
	replaceItalic(&Text)
	replaceBold(&Text)

	return Text
}

/**
*	Replace "* test *"  => "<h1>test</h1>"
 */
func replaceHedding1(Text *string) {
	reg := regexp.MustCompile(`(?m)^\s*#([^#].*?)$`)
	*Text = reg.ReplaceAllString(*Text, "<h1>$1</h1>")
}

/**
*	Replace "**test"  =>  "<h2>test</h2>"
 */
func replaceHedding2(Text *string) {
	reg := regexp.MustCompile(`(?m)^\s*#{2}([^#].*?)$`)
	*Text = reg.ReplaceAllString(*Text, "<h2>$1</h2>")
}

/**
*	Replace "***test" => "<h3>test</h3>"
 */
func replaceHedding3(Text *string) {
	reg := regexp.MustCompile(`(?m)^\s*#{3}([^#].*?)$`)
	*Text = reg.ReplaceAllString(*Text, "<h3>$1</h3>")
}

/**
*	Replace "****test" => "<h4>test</h4>"
 */
func replaceHedding4(Text *string) {
	reg := regexp.MustCompile(`(?m)^\s*#{4}([^#].*?)$`)
	*Text = reg.ReplaceAllString(*Text, "<h4>$1</h4>")
}

/**
*	Replace "*****test" => "<h5>test</h5>"
 */
func replaceHedding5(Text *string) {
	reg := regexp.MustCompile(`(?m)^\s*#{5}([^#].*?)$`)
	*Text = reg.ReplaceAllString(*Text, "<h5>$1</h5>")
}

/**
*	Replace "******test" => "<h6>test</h6>"
 */
func replaceHedding6(Text *string) {
	reg := regexp.MustCompile(`(?m)^\s*#{6}([^#].*?)$`)
	*Text = reg.ReplaceAllString(*Text, "<h6>$1</h6>")
}

/**
*	Replace "*BOLD TEXT*" => "<strong>BOLD TEXT</strong>"
 */
func replaceBold(Text *string) {
	reg := regexp.MustCompile(`\*([^\*].*?[^\*])\*`)
	*Text = reg.ReplaceAllString(*Text, "<strong>$1</strong>")
}

/**
*	Replace "*Italic TEXT*" => "<em>Italic TEXT</em>"
 */
func replaceItalic(Text *string) {
	reg := regexp.MustCompile(`\*\*([^\*].*?[^\*])\*\*`)
	*Text = reg.ReplaceAllString(*Text, "<em>$1</em>")
}

/**
*	Replace "*text1 \n *text2" => "<ul><li>text1</li> \n <li>text2<\li><\ul>"
 */
func replaceUl(Text *string) {
	reg := regexp.MustCompile(`(?m)^\s*\*(.*?)$`)
	*Text = reg.ReplaceAllString(*Text, "<ul><li>$1</li></ul>")
	reg = regexp.MustCompile(`(?s)</ul>\s*<ul>`)
	*Text = reg.ReplaceAllString(*Text, "\n")
}
