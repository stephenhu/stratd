package main

import (
	"testing"
)


func TestAttrStandard(t *testing.T) {

	s900 	:= 0
	s800 	:= 0
	s700 	:= 0
	s600 	:= 0
	s500 	:= 0
	s400 	:= 0
	s300 	:= 0
	s200 	:= 0
	s100 	:= 0
	s000 	:= 0

	for i := 0; i < 1000000; i++ {

		mu := attrStandard(200.0, 375.0)

		if mu > 900 && mu <= 1000 {
			s900++
		} else if mu > 800 && mu <= 900 {
			s800++
		} else if mu > 700 && mu <= 800 {
			s700++
		} else if mu > 600 && mu <= 700 {
			s600++
		} else if mu > 500 && mu <= 600 {
			s500++
		} else if mu > 400 && mu <= 500 {
			s400++
		} else if mu > 300 && mu <= 400 {
			s300++
		} else if mu > 200 && mu <= 300 {
			s200++
		} else if mu > 100 && mu <= 200 {
			s100++
		} else if mu > 0 && mu <= 100 {
			s000++
		}

	}

	t.Logf("900-1000: %d (%.1f%%)", s900, float32(s900)/float32(1000000)*100.0)
	t.Logf("800-900: %d (%.1f%%)", s800, float32(s800)/float32(1000000)*100.0)
	t.Logf("700-800: %d (%.1f%%)", s700, float32(s700)/float32(1000000)*100.0)
	t.Logf("600-700: %d (%.1f%%)", s600, float32(s600)/float32(1000000)*100.0)
	t.Logf("500-600: %d (%.1f%%)", s500, float32(s500)/float32(1000000)*100.0)
	t.Logf("400-500: %d (%.1f%%)", s400, float32(s400)/float32(1000000)*100.0)
	t.Logf("300-400: %d (%.1f%%)", s300, float32(s300)/float32(1000000)*100.0)
	t.Logf("200-300: %d (%.1f%%)", s200, float32(s200)/float32(1000000)*100.0)
	t.Logf("100-200: %d (%.1f%%)", s100, float32(s100)/float32(1000000)*100.0)
	t.Logf("0-100: %d (%.1f%%)", s000, float32(s000)/float32(1000000)*100.0)

} // TestAttrStandard
