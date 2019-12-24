package main

type Lambertian struct {
	albedo Vec3
}

func (l Lambertian) scatter(ray Ray, record *HitRecord) (bool, Vec3, Ray) {
	target := add(add(record.p, record.normal), randomInUnitSphere())
	scattered := makeRay(record.p, subtract(target, record.p))

	return true, l.albedo, scattered
}