package main

import "math"

type Ray struct {
	A, B Vec3
}

func makeRay(a, b Vec3) Ray { return Ray{a, b} }
func (r Ray) origin() Vec3 { return r.A }
func (r Ray) direction() Vec3 { return r.B }
func (r Ray) pointAtParameter(t float32) Vec3 { return add(r.A, multiplyByValue(r.B, t)) }

func (r Ray) color(world Hittable) Vec3 {
	hit, tempRecord := world.hit(&r, 0.0, math.MaxFloat32)
	if hit {
		return multiplyByValue(makeVec3(tempRecord.normal.x()+1, tempRecord.normal.y()+1, tempRecord.normal.z()+1), 0.5)
	} else {
		unitDirection := r.direction().makeUnitVector()
		t := 0.5 * (unitDirection.y() + 1.0)
		return add(multiplyByValue(makeVec3(1.0, 1.0, 1.0), 1.0-t), multiplyByValue(makeVec3(0.5, 0.7, 1.0), t))
	}
}
