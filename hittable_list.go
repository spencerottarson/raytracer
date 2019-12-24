package main

type HittableList struct {
	list []Hittable
}

func (list HittableList) hit(ray *Ray, tMin float32, tMax float32) (bool, HitRecord) {
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