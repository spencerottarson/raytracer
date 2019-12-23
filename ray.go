package main

type Ray struct {
	A, B Vec3
}

func makeRay(a, b Vec3) Ray { return Ray{a, b} }
func (r Ray) origin() Vec3 { return r.A }
func (r Ray) direction() Vec3 { return r.B }
func (r Ray) pointAtParameter(t float32) Vec3 { return r.A.add(r.B.multiplyByValue(t)) }

func (r Ray) color() Vec3 {
	if r.hitSphere(makeVec3(0, 0, 1), 0.5) {
		return makeVec3(1,0,0)
	}

	unit_direction := r.direction().makeUnitVector()
	t := 0.5 * (unit_direction.y() + 1.0)
	return makeVec3(1.0, 1.0, 1.0).multiplyByValue(1.0-t).add(makeVec3(0.5, 0.7, 1.0).multiplyByValue(t))
}

func (r Ray) hitSphere(center Vec3, radius float32) bool {
	oc := r.origin().add(center.negative())
	a := r.direction().dot(r.direction())
	b := 2.0 * oc.dot(r.direction())
	c := oc.dot(oc) - radius * radius
	discriminant := b * b - 4 * a * c

	return discriminant > 0
}