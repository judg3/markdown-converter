package adapter

import (
	"testing"
)

func TestConvert(t *testing.T) {
	text := `#test
  ##test2
####test4
###test3
test *bold* **italic**`
	//text = `     #test`
	expectedResult := "<h1>test</h1>\n<h2>test2</h2>\n<h4>test4</h4>\n<h3>test3</h3>\ntest <strong>bold</strong> <em>italic</em>"
	result := Convert(text)

	if expectedResult != result {
		t.Errorf("expected"+expectedResult+", got: ", result)
	}
}

func TestReplaceHedding1(t *testing.T) {

	text := "#test\n##test2\n####test4\n###test3"
	replaceHedding1(&text)

	expectedResult := "<h1>test</h1>\n##test2\n####test4\n###test3"

	if text != expectedResult {
		t.Errorf("expected "+expectedResult+", got: ", text)
	}
}

func TestReplaceHedding2(t *testing.T) {

	text := "##test"
	replaceHedding2(&text)

	if text != "<h2>test</h2>" {
		t.Errorf("expected <h2>test</h2>, got", text)
	}
}

func TestReplaceHedding3(t *testing.T) {

	text := "###test"
	replaceHedding3(&text)

	if text != "<h3>test</h3>" {
		t.Errorf("expected <h3>test</h3>, got", text)
	}
}

func TestReplaceHedding4(t *testing.T) {

	text := "####test"
	replaceHedding4(&text)

	if text != "<h4>test</h4>" {
		t.Errorf("expected <h4>test</h4>, got", text)
	}
}

func TestReplaceHedding5(t *testing.T) {

	text := "#####test"
	replaceHedding5(&text)

	if text != "<h5>test</h5>" {
		t.Errorf("expected <h5>test</h5>, got", text)
	}
}

func TestReplaceHedding6(t *testing.T) {

	text := "######test"
	replaceHedding6(&text)

	if text != "<h6>test</h6>" {
		t.Errorf("expected <h6>test</h6>, got", text)
	}
}

func TestReplaceBold(t *testing.T) {
	text := "blablabla *bold*"
	replaceBold(&text)

	if text != "blablabla <strong>bold</strong>" {
		t.Errorf("expeced: blablabla <strong>bold</strong>, got: ", text)
	}
}

func TestReplaceItalic(t *testing.T) {
	text := "blablabla **bold**"
	replaceItalic(&text)

	if text != "blablabla <em>bold</em>" {
		t.Errorf("expeced: blablabla <em>bold</em>, got: ", text)
	}
}

func TestReplaceUl(t *testing.T) {
	text := `
	*test line 1
	*test line 2
	*test line 3
	test
	*test line 5
	*test line 6
	`

	expextedText := `<ul><li>test line 1</li>
<li>test line 2</li>
<li>test line 3</li></ul>
	test
<ul><li>test line 5</li>
<li>test line 6</li></ul>
	`
	replaceUl(&text)

	if text != expextedText {
		t.Errorf("expeced: "+expextedText+"got: ", text)
	}
}
