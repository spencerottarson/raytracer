package main

type Material interface {
	scatter(ray Ray, record *HitRecord) (bool, Vec3, Ray)
}
