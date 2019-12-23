package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	width := 200
	height := 100

	file, err := os.Create("image.ppm")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	fmt.Fprintln(writer, "P3")
	fmt.Fprintf(writer, "%d %d\n", width, height)
	fmt.Fprintln(writer, "255")

	for row := height - 1; row >= 0; row-- {
		for column := 0; column < width; column++ {
			vec3 := makeVec3(float32(column) / float32(width), float32(row) / float32(height), float32(0.2))

			rInt := int16(255.99*vec3.r())
			gInt := int16(255.99*vec3.g())
			bInt := int16(255.99*vec3.b())
			fmt.Fprintf(writer, "%d %d %d\n", rInt, gInt, bInt)
		}
	}

	writer.Flush()
}
