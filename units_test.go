package main

import (
	"testing"
)


func TestGenerateSoldier(t *testing.T) {

	m := generateUnit(false)

	t.Log(m)

} // TestGenerateSoldier