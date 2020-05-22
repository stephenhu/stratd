package main

import (
	"math/rand"
	"time"
)

const (
	ATTRIBUTE_MAX					= 1000
	ATTRIBUTE_DEFAULT     = 100
	HITPOINT_MAX          = 3000
	AGE_OFFICER_MAX       = 65
	AGE_SOLDIER_MAX       = 45
)

const (
	RANK_SOLDIER = iota
	RANK_CORPORAL
	RANK_SERGEANT
	RANK_REGIMENTAL_SERGEANT
	RANK_LIEUTENANT
	RANK_CAPTAIN
	RANK_MAJOR
	RANK_LIEUTENANT_COLONEL
	RANK_COLONEL
	RANK_BRIGADIER_GENERAL
	RANK_MAJOR_GENERAL
	RANK_LIEUTENANT_GENERAL
	RANK_GENERAL
	RANK_FIELD_MARSHAL
)

const (
	SIZE_AVERAGE = iota
	SIZE_SMALL
	SIZE_BIG
	SIZE_HUGE
)

type Unit struct {
	Age							int 					`json:"age"`
	Rank						int 					`json:"rank"`
	Experience			int 					`json:"experience"`
	Size						int 					`json:"size"`
	Hitpoints				int 					`json:"hitpoints"`
	Intelligence		int 					`json:"intelligence"`
	Strength				int 					`json:"strength"`
	Vision					int 					`json:"vision"`
	Charisma				int 					`json:"charisma"`
	Dexterity				int 					`json:"dexterity"`
	Endurance				int 					`json:"endurance"`
	Eq							int 					`json:"eq"`
	Psychology			int 					`json:"psychology"`
	Luck						int 					`json:"luck"`
	Loyalty					int 					`json:"loyalty"`
	Morale					int 					`json:"morale"`
	Speed						int 					`json:"speed"`
	Health					int 					`json:"health"`
	Recovery				int 					`json:"recovery"`
	Discipline			int 					`json:"discipline"`
	Learning				int 					`json:"learning"`
	Stamina					int 					`json:"stamina"`
	Agility					int 					`json:"agility"`
	Wisdom					int 					`json:"wisdom"`
	Movement				int 					`json:"movement"`
}


func generateUnit(o bool) *Unit {

	u := Unit{}

	seed := rand.NewSource(time.Now().UnixNano())

	r := rand.New(seed)

	if o {
		u.Age 						= r.Intn(AGE_SOLDIER_MAX)
	} else {
		u.Age 						= r.Intn(AGE_OFFICER_MAX)
	}

	if o {
		u.Rank 						= RANK_SOLDIER
	} else {
		u.Rank 						= RANK_LIEUTENANT
	}

	u.Size 						= r.Intn(SIZE_HUGE)
	u.Intelligence 		= r.Intn(ATTRIBUTE_MAX)
	u.Strength 				= r.Intn(ATTRIBUTE_MAX)
	u.Vision 					= r.Intn(ATTRIBUTE_MAX)
	u.Charisma 				= r.Intn(ATTRIBUTE_MAX)
	u.Dexterity 			= r.Intn(ATTRIBUTE_MAX)
	u.Endurance 			= r.Intn(ATTRIBUTE_MAX)
	u.Eq 							= r.Intn(ATTRIBUTE_MAX)
	u.Psychology 			= r.Intn(ATTRIBUTE_MAX)
	u.Luck 						= r.Intn(ATTRIBUTE_MAX)
	u.Loyalty 				= r.Intn(ATTRIBUTE_MAX)
	u.Morale 					= ATTRIBUTE_DEFAULT
	u.Speed 					= r.Intn(ATTRIBUTE_MAX)
	u.Health 					= r.Intn(AGE_SOLDIER_MAX)
	u.Recovery 				= r.Intn(AGE_SOLDIER_MAX)
	u.Discipline 			= r.Intn(AGE_SOLDIER_MAX)
	u.Learning 				= r.Intn(AGE_SOLDIER_MAX)
	u.Stamina 				= r.Intn(AGE_SOLDIER_MAX)
	u.Agility 				= r.Intn(AGE_SOLDIER_MAX)
	u.Wisdom 					= r.Intn(AGE_SOLDIER_MAX)
	u.Movement 				= r.Intn(AGE_SOLDIER_MAX)

	return &u

} // generateUnit
