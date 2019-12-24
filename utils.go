package main

import (
	"math"
	"math/rand"
)

func sqrt(value float32) float32 {
	return float32(math.Sqrt(float64(value)))
}

func randomInUnitSphere() Vec3 {
	for {
		vector := subtract(multiplyScalar(makeVec3(rand.Float32(), rand.Float32(), rand.Float32()), 2.0), makeVec3(1,1,1))
		if vector.squaredLength() < 1.0 {
			return vector
		}
	}
}