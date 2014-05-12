package adapter

import (
	"testing"
)

func TestConvert(t *testing.T) {
	text := `#test
##test2
####test4
###test3`
	//text = `     #test`
	result := "<h1>test</h1><h2>test</h2><h4>test</h4><h3>test</h3>"
	if Convert(text) != result {
		t.Errorf("expected"+result+", got", text)
	}
}

func TestReplaceHedding1(t *testing.T) {

	text := `#test
			##test2
####test4
###test3`
	replaceHedding1(&text)

	if text != "<h1>test</h1>##test\n####test\n###test" {
		t.Errorf("expected <h1>test</h1>, got", text)
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
