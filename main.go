package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sync"
)

const (
	width = 800
	height = 400
	passesPerThread = 50
	concurrent = 4
)

type Image [width*height]Vec3

func main() {
	lookFrom := makeVec3(-2, 2, 1)
	lookAt := makeVec3(0, 0, -1)
	up := makeVec3(0, 1, 0)
	distanceToFocus := subtract(lookFrom, lookAt).length()
	aperture := float32(0.1)
	camera := makeCamera(lookFrom, lookAt, up, 30, float32(width)/float32(height), aperture,distanceToFocus)

	list := []Hittable {

		Sphere{makeVec3(0,0,-1), 0.5, Lambertian{makeVec3(0.1, 0.3, 0.9)}},
		Sphere{makeVec3(0,-100.5,-1), 100, Lambertian{makeVec3(0.8, 0.8, 0.0)}},
		Sphere{makeVec3(1,0,-1), 0.3, Metal{makeVec3(0.83, 0.69, 0.22), 0.05}},
		Sphere{makeVec3(-1,0,-1), 0.4, Dialectric{1.5}},
		Sphere{makeVec3(0.5,-0.2,-0.7), -0.1, Dialectric{1.1}},
	}
	world := World{list}

	var image Image

	ch := make(chan *Image, concurrent)
	wg := sync.WaitGroup{}

	for pass := 0; pass < concurrent; pass++ {
		wg.Add(1)

		go makePass(&camera, &world, ch, &wg)
	}

	wg.Wait()

	close(ch)

	for imagePass := range ch {
		for row := height - 1; row >= 0; row-- {
			for column := 0; column < width; column++ {
				index := row*width + column
				image[index] = add(image[index], imagePass[index])
			}
		}
	}

	printImage(&image, concurrent)
}

func makePass(camera *Camera, world *World, ch chan *Image, wg *sync.WaitGroup) {
	var imagePass Image
	for row := height - 1; row >= 0; row-- {
		for column := 0; column < width; column++ {

			color := makeVec3(0,0,0)
			for i := 0; i < passesPerThread; i++ {
				u := (float32(column) + rand.Float32()) / float32(width)
				v := (float32(row)+ rand.Float32()) / float32(height)

				ray := camera.getRay(u, v)
				color = add(color, ray.color(world, 0))
			}

			color = divideScalar(color, float32(passesPerThread))
			index := row * width + column
			imagePass[index] = color
		}
	}

	ch <- &imagePass
	wg.Done()
}

func printImage(image *Image, numPasses int) {
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
			color = divideScalar(color, float32(numPasses))
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