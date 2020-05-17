package main

import (
	"testing"
)


func TestGenerateMapPlainSmall(t *testing.T) {

	m := generateMap(MapConfig{
		Height: 100,
		Width:100,
		Terrain: TERRAIN_PLAIN,
	})

	t.Log(m)

} // TestGenerateMapPlainSmall
