package generator

import (
	"github.com/RH12503/Triangula/normgeom"
	"math/rand" 
	"time" 
)

// A RandomGenerator generates a point group filled with random points.
type RandomGenerator struct {
}

// Generate returns a set of randomly distributed points.
func (r RandomGenerator) Generate(n int) normgeom.NormPointGroup { 
	return randomPoints(n)
}

// randomPoints returns a point group with a specified number of points.
func randomPoints(n int) normgeom.NormPointGroup {
	points := normgeom.NormPointGroup{} 
	
	rand.Seed(time.Now().UnixNano()) // Initialize random seed 
	scale:=120
	for i := 0; i < n; i++ {
		randomInt1 := rand.Intn(scale+1) 
		randomFloat1 := float64(randomInt1) 
		randomInt2 := rand.Intn(scale+1) 
		randomFloat2 := float64(randomInt2)  

		points = append(points, normgeom.NormPoint{X: randomFloat1*1/float64(scale), Y: randomFloat2*1/float64(scale)})
	}

	return points
}
