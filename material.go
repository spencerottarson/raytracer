package main

import "math"

type Material interface {
	scatter(ray Ray, record *HitRecord) (bool, Vec3, Ray)
}

func reflect(vIn Vec3, normal *Vec3) Vec3 {
	return subtract(vIn, multiplyScalar(*normal, 2*dot(vIn, *normal)))
}

func refract(vIn Vec3, normal *Vec3, nIOverNT float32) (bool, Vec3) {
	uv := vIn.makeUnitVector()

	dt := dot(uv, *normal)

	discriminant := 1.0 - nIOverNT*nIOverNT*(1-dt*dt)

	if discriminant > 0 {
		refracted := subtract(multiplyScalar(subtract(uv, multiplyScalar(*normal, dt)), nIOverNT), multiplyScalar(*normal, sqrt(discriminant)))
		return true, refracted
	} else {
		return false, makeVec3(0,0,0)
	}
}

func schlick(cosine, indexOfRefraction float32) float32 {
	r0 := (1-indexOfRefraction) / (1+indexOfRefraction)
	r0 = r0 * r0
	return r0 + (1-r0)*float32(math.Pow(1 - float64(cosine),5))
}