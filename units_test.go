package main

import (
	"testing"
)


func TestGenerateColonel(t *testing.T) {

	initRng()

	m := generateUnit(RANK_COLONEL)

	t.Logf("%+v", m)

} // TestGenerateColonel


func TestGenerateSoldier(t *testing.T) {

	initRng()

	m := generateUnit(RANK_SOLDIER)

	t.Logf("%+v", m)

} // TestGenerateSoldier
