package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

const (
	width = 600
	height = 300
	numPasses = 100
)

func main() {
	camera := makeCamera()

	list := []Hittable {
		Sphere{makeVec3(0,0,-1), 0.5, Lambertian{makeVec3(0.1, 0.3, 0.9)}},
		Sphere{makeVec3(0,-100.5,-1), 100, Lambertian{makeVec3(0.8, 0.8, 0.0)}},
		Sphere{makeVec3(1,0,-1), 0.3, Metal{makeVec3(0.83, 0.69, 0.22), 0.05}},
		Sphere{makeVec3(-1,0,-1), 0.4, Dialectric{1.5}},
		Sphere{makeVec3(0.5,-0.2,-0.7), -0.1, Dialectric{1.1}},
	}
	world := HittableList{list}

	var image [height * width]Vec3

	for pass := 0; pass < numPasses; pass++ {
		fmt.Println(pass)
		var imagePass [height * width]Vec3
		for row := height - 1; row >= 0; row-- {
			for column := 0; column < width; column++ {
				u := (float32(column) + rand.Float32()) / float32(width)
				v := (float32(row)+ rand.Float32()) / float32(height)

				ray := camera.getRay(u, v)
				color := ray.color(world, 0)
				index := row * width + column
				imagePass[index] = color
			}
		}

		for row := height - 1; row >= 0; row-- {
			for column := 0; column < width; column++ {
				index := row * width + column
				image[index] = add(image[index], imagePass[index])
			}
		}
	}

	printImage(&image, numPasses)
}

func printImage(image *[height * width]Vec3, numPasses int) {
	file, err := os.Create("image.ppm")
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(file)

	fmt.Fprintln(writer, "P3")
	fmt.Fprintf(writer, "%d %d\n", width, height)
	fmt.Fprintln(writer, "255")

	for row := height - 1; row >= 0; row-- {
		//fmt.Println(row)
		for column := 0; column < width; column++ {
			index := row * width + column
			color := image[index]
			color = divideByValue(color, float32(numPasses))
			color = makeVec3(sqrt(color.r()), sqrt(color.g()), sqrt(color.b()))

			rInt := int16(255.99*color.r())
			gInt := int16(255.99*color.g())
			bInt := int16(255.99*color.b())
			fmt.Fprintf(writer, "%d %d %d\n", rInt, gInt, bInt)
		}
	}

	writer.Flush()

	file.Close()
}