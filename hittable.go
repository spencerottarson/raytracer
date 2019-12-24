package main

type HitRecord struct {
	t float32
	p Vec3
	normal Vec3
}

type Hittable interface {
	hit(ray *Ray, tMin float32, tMax float32) (bool, HitRecord)
}
