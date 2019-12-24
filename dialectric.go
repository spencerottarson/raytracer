package main

import "math/rand"

type Dialectric struct {
	indexOfRefraction float32
}

func (d Dialectric) scatter(ray Ray, record *HitRecord) (bool, Vec3, Ray) {
	reflected := reflect(ray.direction(), &record.normal)
	attenuation := makeVec3(1, 1, 1)

	var outwardNormal Vec3
	var nIOverNT float32
	var reflectProb float32
	var cosine float32

	if dot(ray.direction(), record.normal) > 0 {
		outwardNormal = record.normal.negative()
		nIOverNT = d.indexOfRefraction
		cosine = d.indexOfRefraction * dot(ray.direction(), record.normal) / ray.direction().length()
	} else {
		outwardNormal = record.normal
		nIOverNT = 1.0 / d.indexOfRefraction
		cosine = -dot(ray.direction(), record.normal) / ray.direction().length()
	}

	shouldRefract, refracted := refract(ray.direction(), &outwardNormal, nIOverNT)

	var scattered Ray
	if shouldRefract {
		reflectProb = schlick(cosine, d.indexOfRefraction)
	} else {
		reflectProb = 1.0
	}

	if rand.Float32() < reflectProb {
		scattered = makeRay(record.p, reflected)
	} else {
		scattered = makeRay(record.p, refracted)
	}

	return true, attenuation, scattered
}