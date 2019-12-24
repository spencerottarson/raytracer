package main

type Camera struct {
	lowerLeftCorner, horizontal, vertical, origin Vec3
}

func makeCamera() Camera {
	return Camera{
		lowerLeftCorner: makeVec3(-2.0, -1.0, -1.0),
		horizontal:      makeVec3(4.0, 0.0, 0.0),
		vertical:        makeVec3(0.0, 2.0, 0.0),
		origin:          makeVec3(0.0, 0.0, 0.0),
	}
}

func (camera Camera) getRay(u, v float32) Ray {
	return makeRay(camera.origin,
		add(add(camera.lowerLeftCorner,
			multiplyByValue(camera.horizontal, u)),
			multiplyByValue(camera.vertical, v)))
}
