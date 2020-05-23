package main

import (
	"testing"
)

func init() {
	initRng()
} // init


func TestGenerateMapPlainSmall(t *testing.T) {

	m := generateMap(MapConfig{
		Height: 100,
		Width:100,
		Terrain: TERRAIN_PLAIN,
	})

	t.Log(m)

} // TestGenerateMapPlainSmall


func TestGenerateMapJungleBig(t *testing.T) {

	m := generateMap(MapConfig{
		Height: 1000,
		Width:1000,
		Terrain: TERRAIN_JUNGLE,
	})

	t.Log(m)

} // TestGenerateMapJungleBig


func TestGenerateMapInvalidWidth(t *testing.T) {

	m := generateMap(MapConfig{
		Height: 100,
		Width:0,
		Terrain: TERRAIN_PLAIN,
	})

	if m != nil {
		t.Error("Should return nil for GameMap")
	}

} // TestGenerateMapInvalidWidth
