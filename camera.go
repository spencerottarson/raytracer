package main

import "math"

type Camera struct {
	lowerLeftCorner, horizontal, vertical, origin Vec3
}

func makeCamera(lookFrom Vec3, lookAt Vec3, up Vec3, fieldOfView float32, aspectRatio float32) Camera {
	theta := fieldOfView * math.Pi/180.0
	halfHeight := float32(math.Tan(float64(theta / 2.0)))
	halfWidth := aspectRatio * halfHeight

	w := subtract(lookFrom, lookAt).makeUnitVector()
	u := cross(up, w).makeUnitVector()
	v := cross(w, u)

	return Camera{
		lowerLeftCorner: subtract(subtract(subtract(lookFrom, multiplyScalar(u, halfWidth)), multiplyScalar(v, halfHeight)), w),
		horizontal:      multiplyScalar(u, 2 * halfWidth),
		vertical:        multiplyScalar(v, 2 * halfHeight),
		origin:          lookFrom,
	}
}

func (camera Camera) getRay(s, t float32) Ray {
	return makeRay(camera.origin,
		subtract(add(add(camera.lowerLeftCorner,
			multiplyScalar(camera.horizontal, s)),
			multiplyScalar(camera.vertical, t)), camera.origin))
}
