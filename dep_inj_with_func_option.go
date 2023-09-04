package main

import "fmt"

// Note: Need to put "1" because can't have 2 same functions in same packages
// Dependency Injection
// Learn more https://www.youtube.com/watch?v=UX4XjxWcDB4

type RockClimberFunc1 func (*RockClimber1) 

type SafetyPlacer1 interface {
	PlaceSafeties() 
}

type RockClimber1 struct {
	kind int 
	rocksClimbed int 
	// rockClimber still depends on safetyPlacer but it doesnot depend on the implementation of safetyPlacer, it implements on a behavior of the safetyPlacer
	sp SafetyPlacer1  
}

func defaultRockClimberOts1() RockClimber1{
	return RockClimber1{
		kind: 0, 
		rocksClimbed: 0,
		sp: NOOPSafetyPlacer1{},
	}
}

func withSafety1(safety SafetyPlacer1) RockClimberFunc1{
	return func(opts *RockClimber1) {
		opts.sp = safety
	}
} 


// this is dependency injection constructor 
func newRockClimber1(sp ...RockClimberFunc1) *RockClimber1 {
	o := defaultRockClimberOts1()

	for _, fn := range sp {
		fn(&o)
	}

	return &RockClimber1{
		sp: o.sp,
	}
}


type IceSafetyPlacer1 struct {}
func (sp IceSafetyPlacer1) PlaceSafeties() {
	fmt.Println("placing my ICE safeties")
}

type NOOPSafetyPlacer1 struct {}
// this is just a placeholder
func (sp NOOPSafetyPlacer1) PlaceSafeties() {
	fmt.Println("placing my NO-OP safeties")
}


func (rc *RockClimber1) climbRock() {
	rc.rocksClimbed ++ 
	if rc.rocksClimbed == 10 {
		rc.sp.PlaceSafeties()
	}
}

func FuncMain1() {
	// rc := newRockClimber() // default no-op safety
	rc := newRockClimber1(withSafety1(IceSafetyPlacer1{}))

	for i:=0; i < 11; i++ {
		rc.climbRock()
	}
 
}