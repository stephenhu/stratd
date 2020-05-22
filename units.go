package main

import (
	//"log"
	//"math"
	//"math/rand"
)

const (
	HITPOINT_MAX          = 3000
	HITPOINT_MIN          = 100
	AGE_OFFICER_MAX       = 65
	AGE_OFFICER_MIN       = 28
	AGE_SOLDIER_MAX       = 45
	AGE_SOLDIER_MIN       = 16
	EXP_UNIT_MAX       		= 20000
	EXP_UNIT_MIN       		= 100
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
	SIZE_LARGE
	SIZE_XLARGE
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


func generateUnit(rank int) *Unit {

	u := Unit{}

	if rank == RANK_SOLDIER || rank == RANK_SERGEANT || rank == RANK_REGIMENTAL_SERGEANT {
		u.Age 						= attrRange(AGE_SOLDIER_MIN, AGE_SOLDIER_MAX)
	} else {
		u.Age 						= attrRange(AGE_OFFICER_MIN, AGE_OFFICER_MAX)
	}

	u.Rank 						= rank
	u.Experience 			= attrRange(EXP_UNIT_MIN, EXP_UNIT_MAX)
	u.Size 						= attrSelect(SIZE_AVERAGE, SIZE_SMALL, SIZE_LARGE, SIZE_XLARGE)

	u.Intelligence 		= attrStandard(ATTRIBUTE_SIGMA, ATTRIBUTE_MEAN)
	u.Strength 				= attrStandard(ATTRIBUTE_SIGMA, ATTRIBUTE_MEAN)
	u.Vision 					= attrStandard(ATTRIBUTE_SIGMA, ATTRIBUTE_MEAN)
	u.Charisma 				= attrStandard(ATTRIBUTE_SIGMA, ATTRIBUTE_MEAN)
	u.Dexterity 			= attrStandard(ATTRIBUTE_SIGMA, ATTRIBUTE_MEAN)
	u.Endurance 			= attrStandard(ATTRIBUTE_SIGMA, ATTRIBUTE_MEAN)
	u.Eq 							= attrStandard(ATTRIBUTE_SIGMA, ATTRIBUTE_MEAN)
	u.Psychology 			= attrStandard(ATTRIBUTE_SIGMA, ATTRIBUTE_MEAN)
	u.Luck 						= attrStandard(ATTRIBUTE_SIGMA, ATTRIBUTE_MEAN)
	u.Loyalty 				= attrStandard(ATTRIBUTE_SIGMA, ATTRIBUTE_MEAN)
	u.Speed 					= attrStandard(ATTRIBUTE_SIGMA, ATTRIBUTE_MEAN)
	u.Health 					= attrStandard(ATTRIBUTE_SIGMA, ATTRIBUTE_MEAN)
	u.Recovery 				= attrStandard(ATTRIBUTE_SIGMA, ATTRIBUTE_MEAN)
	u.Discipline 			= attrStandard(ATTRIBUTE_SIGMA, ATTRIBUTE_MEAN)
	u.Learning 				= attrStandard(ATTRIBUTE_SIGMA, ATTRIBUTE_MEAN)
	u.Stamina 				= attrStandard(ATTRIBUTE_SIGMA, ATTRIBUTE_MEAN)
	u.Agility 				= attrStandard(ATTRIBUTE_SIGMA, ATTRIBUTE_MEAN)
	u.Wisdom 					= attrStandard(ATTRIBUTE_SIGMA, ATTRIBUTE_MEAN)
	u.Movement 				= attrStandard(ATTRIBUTE_SIGMA, ATTRIBUTE_MEAN)
	u.Morale 					= attrStandard(ATTRIBUTE_SIGMA, ATTRIBUTE_MEAN)

	return &u

} // generateUnit
