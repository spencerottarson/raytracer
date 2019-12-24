package main

import "math"

type Ray struct {
	A, B Vec3
}

func makeRay(a, b Vec3) Ray { return Ray{a, b} }
func (r Ray) origin() Vec3 { return r.A }
func (r Ray) direction() Vec3 { return r.B }
func (r Ray) pointAtParameter(t float32) Vec3 { return add(r.A, multiplyScalar(r.B, t)) }

func (r Ray) color(world Hittable, depth int) Vec3 {
	hit, tempRecord := world.hit(&r, 0.001, math.MaxFloat32)
	if hit {
		if depth < 50 {
			shouldReflect, attenuation, scattered := tempRecord.matPtr.scatter(r, &tempRecord)
			if shouldReflect {
				return multiply(scattered.color(world, depth+2), attenuation)
			}
		}
		return makeVec3(0,0,0)
	} else {
		unitDirection := r.direction().makeUnitVector()
		t := 0.5 * (unitDirection.y() + 1.0)
		return add(multiplyScalar(makeVec3(1.0, 1.0, 1.0), 1.0-t), multiplyScalar(makeVec3(0.5, 0.7, 1.0), t))
	}
}
