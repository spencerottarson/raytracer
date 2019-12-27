package main

import "math/rand"

type World struct {
	list []Hittable
}

func (list World) hit(ray *Ray, tMin float32, tMax float32) (bool, HitRecord) {
	hitAnything := false
	closestSoFar := tMax
	var record HitRecord

	for _, item := range list.list {
		hit, tempRecord := item.hit(ray, tMin, closestSoFar)
		if hit {
			hitAnything = true
			closestSoFar = tempRecord.t
			record = tempRecord
		}
	}

	return hitAnything, record
}

func randomWorld() *World {
	var list []Hittable

	planet := Sphere{makeVec3(0, -1000, 0), 1000, Lambertian{makeVec3(0.5, 0.5, 0.5)}}
	orb1 := Sphere{makeVec3(0, 1, 0), 1, Dialectric{1.5}}
	orb2 := Sphere{makeVec3(-4, 1, 0), 1, Lambertian{makeVec3(0.4, 0.2, 0.1)}}
	orb3 := Sphere{makeVec3(4, 1, 0), 1, Metal{makeVec3(0.83, 0.69, 0.22), 0.05}}
	list = append(list, planet)
	list = append(list, orb1)
	list = append(list, orb2)
	list = append(list, orb3)

	//new sphere(vec3(0, 1, 0), 1.0, new dielectric(1.5));
	//list[i++] = new sphere(vec3(-4, 1, 0), 1.0, new lambertian(vec3(0.4, 0.2, 0.1)));
	//list[i++] = new sphere(vec3(4, 1, 0), 1.0, new metal(vec3(0.7, 0.6, 0.5), 0.0));

	for a := float32(-11); a < float32(11); a++ {
		for b := float32(-11); b < float32(11); b++ {
			chooseMaterial := rand.Float32()
			size := rand.Float32() * 0.2 + 0.1
			center := makeVec3(a + rand.Float32() * 0.85, size, b + rand.Float32() * 0.85)
			if subtract(center, makeVec3(4, 0.2, 0)).length() > 0.9 {
				if chooseMaterial < 0.6 {
					sphere := Sphere{
						center,
						size,
						Lambertian{
						makeVec3(rand.Float32() * rand.Float32(),
							rand.Float32()*rand.Float32(),
							rand.Float32() * rand.Float32()),
						},
					}
					list = append(list, sphere)
				} else if chooseMaterial < 0.8 {
					sphere := Sphere{
						center,
						size,
						Metal{
							makeVec3(0.5 * (1+rand.Float32()),
								0.5 * (1+rand.Float32()),
								0.5 * (1+rand.Float32())),
								0.5 * rand.Float32(),
						},
					}
					list = append(list, sphere)
				} else if chooseMaterial < 0.9 {
					sphere := Sphere{
						center,
						size,
						Dialectric{
							1.5,
						},
					}
					list = append(list, sphere)
				} else {
					sphere := Sphere{
						center,
						-size,
						Dialectric{
							1+(0.5*rand.Float32()*rand.Float32()),
						},
					}
					list = append(list, sphere)
				}
			}
		}
	}

	return &World{list}
}