package storage

import (
	"testing"
)

func TestSerializeProps(t *testing.T) {
	if serializeProps("", []string{"a", "b", "c"}) != "a,b,c" {
		t.Error("error joining")
	}
}

func TestSerializePropsWithPrefix(t *testing.T) {
	if serializeProps(":", []string{"a", "b", "c"}) != ":a,:b,:c" {
		t.Error("prefixing error")
	}
}
