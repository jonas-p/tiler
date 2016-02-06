package main

type Point struct {
	X, Y float64
}

type BoundingBox struct {
	UpperLeft, LowerLeft, UpperRight, LowerRight Point
}

type ImageRepresentation struct {
	BoundingBox BoundingBox
	EpsgCode    int
	Proj4       string
}
