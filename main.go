package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	width := 1200
	height := 600

	file, err := os.Create("image.ppm")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	fmt.Fprintln(writer, "P3")
	fmt.Fprintf(writer, "%d %d\n", width, height)
	fmt.Fprintln(writer, "255")

	lowerLeftCorner := makeVec3(-2.0, -1.0, -1.0)
	horizontal := makeVec3(4.0, 0.0, 0.0)
	vertical := makeVec3(0.0, 2.0, 0.0)
	origin := makeVec3(0.0, 0.0, 0.0)

	list := []Hittable {
		Sphere{makeVec3(0,0,-1), 0.5},
		Sphere{makeVec3(0,-100.5,-1), 100},
	}
	world := HittableList{list}

	for row := height - 1; row >= 0; row-- {
		for column := 0; column < width; column++ {
			u := float32(column) / float32(width)
			v := float32(row) / float32(height)

			ray := makeRay(origin, add(add(lowerLeftCorner, multiplyByValue(horizontal, u)), multiplyByValue(vertical, v)))
			color := ray.color(world)

			rInt := int16(255.99*color.r())
			gInt := int16(255.99*color.g())
			bInt := int16(255.99*color.b())
			fmt.Fprintf(writer, "%d %d %d\n", rInt, gInt, bInt)
		}
	}

	writer.Flush()
}
