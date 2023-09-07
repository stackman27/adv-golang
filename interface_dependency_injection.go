package main

import "fmt"

// Dependency Injection
// Learn more https://www.youtube.com/watch?v=UX4XjxWcDB4

type SafetyPlacer interface {
	PlaceSafeties()
}

type RockClimber struct {
	kind int
	rocksClimbed int
	// rockClimber still depends on safetyPlacer but it doesnot depend on the implementation of safetyPlacer, it implements on a behavior of the safetyPlacer
	sp SafetyPlacer
}

// this is dependency injection constructor
func newRockClimber(sp SafetyPlacer) *RockClimber {
	return &RockClimber{
		sp: sp,
	}
}

type IceSafetyPlacer struct {}

func (sp IceSafetyPlacer) PlaceSafeties() {
	fmt.Println("placing my ICE safeties")
}

type NOOPSafetyPlacer struct {}
// this is just a placeholder
func (sp NOOPSafetyPlacer) PlaceSafeties() {
	fmt.Println("placing my NO-OP safeties")
}

func (rc *RockClimber) climbRock() {
	rc.rocksClimbed ++
	if rc.rocksClimbed == 10 {
		rc.sp.PlaceSafeties()
	}
}

func DepInjMain() {

	rc := newRockClimber(IceSafetyPlacer{})

	for i:=0; i < 11; i++ {
		rc.climbRock()
	}

}