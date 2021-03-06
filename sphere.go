package main

type Sphere struct {
	center Vec3
	radius float32
	material Material
}

func (sphere Sphere) hit(ray *Ray, tMin float32, tMax float32) (bool, HitRecord) {
	oc := subtract(ray.origin(), sphere.center)
	a := dot(ray.direction(), ray.direction())
	b := dot(oc, ray.direction())
	c := dot(oc, oc) - sphere.radius * sphere.radius
	discriminant := b * b - a * c

	if discriminant > 0 {
		temp := (-b - sqrt(discriminant)) / a
		if temp < tMax && temp > tMin {
			point := ray.pointAtParameter(temp)
			normal := divideScalar(subtract(point, sphere.center), sphere.radius)
			return true, HitRecord{temp, point, normal, sphere.material}
		}
		temp = (-b + sqrt(discriminant)) / a
		if temp < tMax && temp > tMin {
			point := ray.pointAtParameter(temp)
			normal := divideScalar(subtract(point, sphere.center), sphere.radius)
			return true, HitRecord{temp, point, normal, sphere.material}
		}
	}
	return false, HitRecord{}
}
