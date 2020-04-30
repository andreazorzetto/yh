package main

import (
	"testing"
)

func TestIsKeyValue(t *testing.T) {
	l := line{
		line: "serviceAccount: hashicorp-consul-server",
	}

	if l.isKeyValue() != true {
		t.Errorf("Expected True but got %v", l.isKeyValue())
	}
}

func TestGetKeyValue(t *testing.T) {
	l := line{
		line: "foo: bar",
	}

	l.getKeyValue()

	if l.key != "foo" {
		t.Errorf("Expected 'foo' but got %v", l.key)
	}
	if l.value != "bar" {
		t.Errorf("Expected 'bar' but got %v", l.value)
	}
}

func TestIsNumberOrIP(t *testing.T) {
	n := "657890"
	f := "123123.12312"
	ip := "127.0.0.1"
	w := "test not a number"

	if isNumberOrIP(n) != true {
		t.Errorf("Expected true but got %v", isNumberOrIP(n))
	}
	if isNumberOrIP(f) != true {
		t.Errorf("Expected true but got %v", isNumberOrIP(f))

	}
	if isNumberOrIP(ip) != true {
		t.Errorf("Expected true but got %v", isNumberOrIP(ip))
	}
	if isNumberOrIP(w) != false {
		t.Errorf("Expected false but got %v", isNumberOrIP(w))
	}

}

func TestIsComment(t *testing.T) {
	c1 := "\t#this is a comment"
	c2 := "#this is another comment"
	nc := "- item"

	if isComment(c1) != true {
		t.Errorf("Expected true but got %v", isComment(c1))
	}

	if isComment(c2) != true {
		t.Errorf("Expected true but got %v", isComment(c2))
	}

	if isComment(nc) != false {
		t.Errorf("Expected false but got %v", isComment(nc))
	}

}

func TestIsBoolean(t *testing.T) {
	b1 := "true"
	b2 := "false"
	b3 := "False"
	w := "asdasd"

	if isBoolean(b1) != true {
		t.Errorf("Expected true but got %v", isComment(b1))
	}

	if isBoolean(b2) != true {
		t.Errorf("Expected true but got %v", isComment(b2))
	}

	if isBoolean(b3) != true {
		t.Errorf("Expected true but got %v", isComment(b3))
	}

	if isBoolean(w) != false {
		t.Errorf("Expected false but got %v", isComment(w))
	}

}
