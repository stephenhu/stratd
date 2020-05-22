package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

const (

	SURFACE_DIRT 					= iota
	SURFACE_MUD
	SURFACE_GRASS
	SURFACE_HIGH_GRASS
	SURFACE_DRY_GRASS
	SURFACE_GRAVEL
	SURFACE_ROCK
	SURFACE_CONCRETE
	SURFACE_SAND
	SURFACE_DEEP_WATER
	SURFACE_WATER
)

const (
	MAX_WIDTH							= 8192
	MAX_HEIGHT						= 8192
	MAX_ALTITUDE_M     		= 5000
	MAX_WATER_PCT         = 100
	MAX_MOUNTAIN_PCT      = 100
	MIN_WIDTH							= 64
	MIN_HEIGHT						= 32
	MIN_ALTITUDE_M     		= -200
	MIN_WATER_PCT         = 0
	MIN_MOUNTAIN_PCT      = 0
	HILL_ALTITUDE_M       = 50
)

type Plot struct {
	Altitude        int         `json:"altitude"`
	Forest          int         `json:"forest"`
	Surface         int        	`json:"surface"`
	UnitID          int         `json:"unitId"`
}

type GameMap struct {
	Plots        		[][]Plot		`json:"plots"`
	Terrain         int         `json:"terrain"`
	//Weather         int         `json:"weather"`
}


func buildTerrain(gm GameMap, conf MapConfig) {

	t := terrainProfiles[conf.Terrain]

	totalWater 		:= 0
	totalMtn   		:= 0
	//totalForest		:= 0
	totalSpaces 	:= conf.Width * conf.Height

	seed := rand.NewSource(time.Now().UnixNano())

	r := rand.New(seed)

	for x := 0; x < conf.Width; x++ {

		gm.Plots[x] = make([]Plot, conf.Height)

		for y := 0; y < conf.Height; y++ {

			v := r.Intn(SURFACE_WATER)
			a := 0

			if v == SURFACE_WATER || v == SURFACE_DEEP_WATER {

				if int(totalWater/totalSpaces*100)  > t.MaxWaterPct {
					v = SURFACE_GRASS
				} else {
					totalWater++
				}

			} else {

				if conf.Terrain == TERRAIN_MOUNTAIN {
					a = r.Intn(t.MaxMountainAltitude)
					totalMtn++
				} else {
					a = r.Intn(t.MaxHillAltitude)
				}

				if totalMtn > t.MaxMountainPct {
					a = r.Intn(t.MaxHillAltitude)
				}

			}

			gm.Plots[x][y] = Plot{
				Surface: v,
				Altitude: a,
			}

		}

	}

	log.Println(totalWater)
	log.Println(totalMtn)
	log.Println(totalSpaces)

} // buildTerrain


func generateMap(conf MapConfig) *GameMap {

	if conf.Width > MAX_WIDTH || conf.Width < MIN_WIDTH {

		logf(ERROR, "generateMap", fmt.Sprintf("Width: %d, must be between %d - %d", conf.Width,
			MIN_WIDTH, MAX_WIDTH))
		return nil

	} else if conf.Height > MAX_HEIGHT || conf.Height < MIN_HEIGHT {

		logf(ERROR, "generateMap", fmt.Sprintf("Length: %d, must be between %d - %d", conf.Height,
			MIN_HEIGHT, MAX_HEIGHT))
		return nil

	} else if conf.Terrain > TERRAIN_WASTELAND || conf.Terrain < TERRAIN_PLAIN {

		logf(ERROR, "generateMap", fmt.Sprintf("Terrain: %d, must be between %d - %d", conf.Terrain,
			TERRAIN_PLAIN, TERRAIN_WASTELAND))
		return nil

	}

	gm := GameMap{
		Terrain: conf.Terrain,
		Plots: make([][]Plot, conf.Width),
	}

	buildTerrain(gm, conf)

	return &gm

} // generateMap
