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
			r := float32(column) / float32(width)
			g := float32(row) / float32(height)
			b := float32(0.2)

			rInt := int16(255.99*r)
			gInt := int16(255.99*g)
			bInt := int16(255.99*b)
			fmt.Fprintf(writer, "%d %d %d\n", rInt, gInt, bInt)
		}
	}

	writer.Flush()
}
