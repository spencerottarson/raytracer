package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	width := 200
	height := 100
	numPasses := 100

	file, err := os.Create("image.ppm")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	fmt.Fprintln(writer, "P3")
	fmt.Fprintf(writer, "%d %d\n", width, height)
	fmt.Fprintln(writer, "255")

	camera := makeCamera()

	list := []Hittable {
		Sphere{makeVec3(0,0,-1), 0.5},
		Sphere{makeVec3(0,-100.5,-1), 100},
	}
	world := HittableList{list}

	for row := height - 1; row >= 0; row-- {
		for column := 0; column < width; column++ {

			color := makeVec3(0,0,0)
			for i := 0; i < numPasses; i++ {
				u := (float32(column) + rand.Float32()) / float32(width)
				v := (float32(row)+ rand.Float32()) / float32(height)

				ray := camera.getRay(u, v)
				color = add(color, ray.color(world))
			}

			color = divideByValue(color, float32(numPasses))
			color = makeVec3(sqrt(color.r()), sqrt(color.g()), sqrt(color.b()))

			rInt := int16(255.99*color.r())
			gInt := int16(255.99*color.g())
			bInt := int16(255.99*color.b())
			fmt.Fprintf(writer, "%d %d %d\n", rInt, gInt, bInt)
		}
	}

	writer.Flush()
}
