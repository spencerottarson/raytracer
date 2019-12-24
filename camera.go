package main

import "math"

type Camera struct {
	lowerLeftCorner, horizontal, vertical, origin, w, u, v Vec3
	lensRadius float32
}

func makeCamera(lookFrom Vec3, lookAt Vec3, up Vec3, fieldOfView float32, aspectRatio float32, aperture float32, focusDistance float32) Camera {
	lensRadius := aperture / 2.0
	theta := fieldOfView * math.Pi/180.0
	halfHeight := float32(math.Tan(float64(theta / 2.0)))
	halfWidth := aspectRatio * halfHeight

	w := subtract(lookFrom, lookAt).makeUnitVector()
	u := cross(up, w).makeUnitVector()
	v := cross(w, u)

	return Camera{
		lowerLeftCorner: subtract(subtract(subtract(lookFrom, multiplyScalar(u, halfWidth*focusDistance)), multiplyScalar(v, halfHeight*focusDistance)), multiplyScalar(w, focusDistance)),
		horizontal: multiplyScalar(u, 2 * halfWidth*focusDistance),
		vertical: multiplyScalar(v, 2 * halfHeight*focusDistance),
		origin: lookFrom,
		w: w,
		u: u,
		v: v,
		lensRadius: lensRadius,
	}
}

func (camera Camera) getRay(s, t float32) Ray {
	rd := multiplyScalar(randomInUnitDisk(), camera.lensRadius)
	offset := add(multiplyScalar(camera.u, rd.x()), multiplyScalar(camera.v, rd.y()))
	return makeRay(add(camera.origin, offset),
		subtract(subtract(add(add(camera.lowerLeftCorner,
			multiplyScalar(camera.horizontal, s)),
			multiplyScalar(camera.vertical, t)), camera.origin), offset))
}
