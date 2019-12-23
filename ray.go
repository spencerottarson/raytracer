package main

type Ray struct {
	A, B Vec3
}

func makeRay(a, b Vec3) Ray { return Ray{a, b} }
func (r Ray) origin() Vec3 { return r.A }
func (r Ray) direction() Vec3 { return r.B }
func (r Ray) pointAtParameter(t float32) Vec3 { return r.A.add(r.B.multiplyByValue(t)) }

func (r Ray) backgroundColor() Vec3 {
	unit_direction := r.direction().makeUnitVector()
	t := 0.5 * (unit_direction.y() + 1.0)
	return makeVec3(1.0, 1.0, 1.0).multiplyByValue(1.0-t).add(makeVec3(0.5, 0.7, 1.0).multiplyByValue(t))
}