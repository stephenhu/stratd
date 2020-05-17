package main

import (
	"fmt"
	//"log"
	"math/rand"
)

const (
	TERRAIN_PLAIN 					= iota
	TERRAIN_DESERT
	TERRAIN_JUNGLE
	TERRAIN_MOUNTAIN
	TERRAIN_STEPPE
	TERRAIN_SWAMP
	TERRAIN_URBAN
	TERRAIN_WASTELAND
)

const (
	SURFACE_DIRT 					= iota
	SURFACE_GRAVEL
	SURFACE_ROCK
	SURFACE_CONCRETE
	SURFACE_SAND
	SURFACE_DEEP_WATER
	SURFACE_MEDIUM_WATER
	SURFACE_WATER
)

const (
	MAX_WIDTH							= 8192
	MAX_HEIGHT						= 8192
	MAX_WATER_PCT         = 100
	MAX_MOUNTAIN_PCT      = 100
	MIN_WIDTH							= 64
	MIN_HEIGHT						= 32
	MIN_WATER_PCT         = 0
	MIN_MOUNTAIN_PCT      = 0
)

type Plot struct {
	Altitude        int         `json:"altitude"`
	Forest          int         `json:"forest"`
	Surface         int        	`json:"surface"`
	Snow            bool        `json:"snow"`
	Ice             bool        `json:"ice"`
	Rain            bool        `json:"rain"`
	Wind            bool        `json:"wind`
	UnitID          int         `json:"unitId"`
}

type GameMap struct {
	Plots        		[][]Plot		`json:"plots"`
	Terrain         int         `json:"terrain"`
	Weather         int         `json:"weather"`
}


func generateMap(conf MapConfig) *GameMap {

	if conf.Width > MAX_WIDTH || conf.Width < MIN_WIDTH {

		logf(ERROR, fmt.Sprintf("Width: %d, must be between %d - %d", conf.Width,
			MIN_WIDTH, MAX_WIDTH))
		return nil

	} else if conf.Height > MAX_HEIGHT || conf.Height < MIN_HEIGHT {

		logf(ERROR, fmt.Sprintf("Length: %d, must be between %d - %d", conf.Height,
			MIN_HEIGHT, MAX_HEIGHT))
		return nil

	} else if conf.Terrain > TERRAIN_WASTELAND || conf.Terrain < TERRAIN_PLAIN {

		logf(ERROR, fmt.Sprintf("Terrain: %d, must be between %d - %d", conf.Terrain,
			TERRAIN_PLAIN, TERRAIN_WASTELAND))
		return nil

	}

	gm := GameMap{
		Terrain: conf.Terrain,
		Plots: make([][]Plot, conf.Width),
	}

	totalWater 	:= 0
	//totalMtn  	:= 0

	for x := 0; x < conf.Width; x++ {

		gm.Plots[x] = make([]Plot, conf.Height)

		for y := 0; y < conf.Height; y++ {

			v := rand.Intn(SURFACE_WATER)

			if v == SURFACE_WATER || v == SURFACE_DEEP_WATER ||
			  v == SURFACE_MEDIUM_WATER {

				totalWater++

				if totalWater > conf.MaxWaterPct {
					v = SURFACE_DIRT
				}

			}

			gm.Plots[x][y] = Plot{
				Surface: v,
			}

		}

	}

	return &gm

} // generateMap
