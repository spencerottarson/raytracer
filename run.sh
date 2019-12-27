#! /bin/bash

go build
./raytracer
git add image.ppm
git commit -m "add image"
git push -u origin cloud
