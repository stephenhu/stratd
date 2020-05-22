package main

const (
	TERRAIN_PLAIN 					= iota
	TERRAIN_DESERT
	TERRAIN_JUNGLE
	TERRAIN_MOUNTAIN
	TERRAIN_SEA
	TERRAIN_SPACE_STATION
	TERRAIN_STEPPE
	TERRAIN_SWAMP
	TERRAIN_URBAN
	TERRAIN_WASTELAND
)

type TerrainProfile struct {
	Name									string					`json:"name"`
	MaxMountainPct				int							`json:"maxMountainPct"`
	MaxWaterPct						int							`json:"maxWaterPct"`
	MaxForestPct      		int             `json:"maxForestPct"`
	Surface           		int             `json:"surface"`
	MaxMountainAltitude   int             `json:"maxMountainAltitude"`
	MaxHillAltitude       int     				`json:"maxHillAltitude"`
}

var terrainProfiles = map[int] TerrainProfile {
	TERRAIN_PLAIN: {
		Name: "plain",
		MaxMountainPct: 15,
		MaxWaterPct: 15,
		MaxForestPct: 40,
		Surface: SURFACE_GRASS,
		MaxMountainAltitude: 200,
		MaxHillAltitude: 20,
	},
	TERRAIN_DESERT: {
		Name: "desert",
		MaxMountainPct: 5,
		MaxWaterPct: 2,
		MaxForestPct: 1,
		Surface: SURFACE_SAND,
		MaxMountainAltitude: 200,
		MaxHillAltitude: 30,
	},
	TERRAIN_JUNGLE: {
		Name: "jungle",
		MaxMountainPct: 1,
		MaxWaterPct: 15,
		MaxForestPct: 80,
		Surface: SURFACE_MUD,
		MaxMountainAltitude: 200,
		MaxHillAltitude: 2,
	},
	TERRAIN_MOUNTAIN: {
		Name: "mountain",
		MaxMountainPct: 90,
		MaxWaterPct: 1,
		MaxForestPct: 30,
		Surface: SURFACE_ROCK,
		MaxMountainAltitude: 200,
		MaxHillAltitude: 20,
	},
	TERRAIN_SEA: {
		Name: "sea",
		MaxMountainPct: 0,
		MaxWaterPct: 90,
		MaxForestPct: 0,
		Surface: SURFACE_WATER,
		MaxMountainAltitude: 200,
		MaxHillAltitude: 20,
	},
	TERRAIN_STEPPE: {
		Name: "steppe",
		MaxMountainPct: 12,
		MaxWaterPct: 1,
		MaxForestPct: 1,
		Surface: SURFACE_ROCK,
		MaxMountainAltitude: 200,
		MaxHillAltitude: 5,
	},
	TERRAIN_SWAMP: {
		Name: "swamp",
		MaxMountainPct: 0,
		MaxWaterPct: 95,
		MaxForestPct: 70,
		Surface: SURFACE_WATER,
		MaxMountainAltitude: 200,
		MaxHillAltitude: 0,
	},
	TERRAIN_URBAN: {
		Name: "urban",
		MaxMountainPct: 0,
		MaxWaterPct: 20,
		MaxForestPct: 1,
		Surface: SURFACE_CONCRETE,
		MaxMountainAltitude: 200,
		MaxHillAltitude: 20,
	},
	TERRAIN_WASTELAND: {
		Name: "wasteland",
		MaxMountainPct: 10,
		MaxWaterPct: 2,
		MaxForestPct: 0,
		Surface: SURFACE_DIRT,
		MaxMountainAltitude: 200,
		MaxHillAltitude: 20,
	},
}
