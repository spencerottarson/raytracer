package main

import "math"

type Vec3 struct {
	e [3]float32
}

func makeVec3(e0 float32, e1 float32, e2 float32) Vec3 {
	return Vec3{[3]float32{e0, e1, e2}}
}

func (v Vec3) x() float32 { return v.e[0] }
func (v Vec3) y() float32 { return v.e[1] }
func (v Vec3) z() float32 { return v.e[2] }

func (v Vec3) r() float32 { return v.e[0] }
func (v Vec3) g() float32 { return v.e[1] }
func (v Vec3) b() float32 { return v.e[2] }

func (v Vec3) negative() Vec3 { return makeVec3(-v.e[0], -v.e[1], -v.e[2]) }

func (v Vec3) get(i int) float32 { return v.e[i] }

func (v Vec3) add(v2 Vec3) Vec3    { return makeVec3(v.e[0] + v2.e[0], v.e[1] + v2.e[1], v.e[2] + v2.e[2]) }
func (v Vec3) subtract(v2 Vec3) Vec3 { return makeVec3(v.e[0] - v2.e[0], v.e[1] - v2.e[1], v.e[2] - v2.e[2]) }
func (v Vec3) multiplyBy(v2 Vec3) Vec3 { return makeVec3(v.e[0] * v2.e[0], v.e[1] * v2.e[1], v.e[2] * v2.e[2]) }
func (v Vec3) multiplyByValue(val float32) Vec3 { return makeVec3(v.e[0] * val, v.e[1] * val, v.e[2] * val) }
func (v Vec3) divideBy(v2 Vec3) Vec3 { return makeVec3(v.e[0] / v2.e[0], v.e[1] / v2.e[1], v.e[2] / v2.e[2]) }
func (v Vec3) divideByValue(val float32) Vec3 { return makeVec3(v.e[0] / val, v.e[1] / val, v.e[2] / val) }
func (v Vec3) dot(v2 Vec3) float32 { return v.e[0] * v2.e[0] + v.e[1] * v2.e[1] + v.e[2] * v2.e[2] }
func (v Vec3) cross(v2 Vec3) Vec3 {
	return makeVec3(
		v.e[1] * v2.e[2] - v.e[2] * v2.e[1],
		v.e[2] * v2.e[0] - v.e[0] * v2.e[2],
		v.e[0] * v2.e[1] - v.e[1] * v2.e[0],
	)
}
func (v Vec3) squaredLength() float32 { return v.e[0] * v.e[0] + v.e[1] * v.e[1] + v.e[2] * v.e[2] }
func (v Vec3) length() float32 { return float32(math.Sqrt(float64(v.squaredLength()))) }
func (v Vec3) makeUnitVector() Vec3 { return v.divideByValue(v.length()) }