package main

type Metal struct {
	albedo Vec3
	fuzz float32
}

func reflect(vIn Vec3, normal *Vec3) Vec3 {
	return subtract(vIn, multiplyByValue(*normal, 2*dot(vIn, *normal)))
}

func (metal Metal) scatter(ray Ray, record *HitRecord) (bool, Vec3, Ray) {
	reflected := reflect(ray.direction().makeUnitVector(), &record.normal)
	scattered := makeRay(record.p, add(reflected, multiplyByValue(randomInUnitSphere(), metal.fuzz)))
	if dot(scattered.direction(), record.normal) > 0 {
		return true, metal.albedo, scattered
	}

	return false, metal.albedo, scattered
}
