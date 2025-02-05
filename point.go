// Package cluster implements DBScan clustering on (lat, lon) using K-D Tree
package cluster

// Point is longitue, latittude
type Point struct {
	XY   [2]float64
	Data []string
}

// PointList is a slice of Points
type PointList []Point

// Cluster is a result of DBScan work
type Cluster struct {
	C      int
	Points []int
}

// sqDist returns squared (w/o sqrt & normalization) distance between two points
func (a *Point) sqDist(b *Point) float64 {
	return DistanceSphericalFast(a, b)
}

// LessEq - a <= b
func (a *Point) LessEq(b *Point) bool {
	return a.XY[0] <= b.XY[0] && a.XY[1] <= b.XY[1]
}

// GreaterEq - a >= b
func (a *Point) GreaterEq(b *Point) bool {
	return a.XY[0] >= b.XY[0] && a.XY[1] >= b.XY[1]
}

// CentroidAndBounds calculates center and cluster bounds
func (c *Cluster) CentroidAndBounds(points PointList) (center, min, max Point) {
	if len(c.Points) == 0 {
		panic("empty cluster")
	}

	min = Point{[2]float64{180.0, 90.0}, nil}
	max = Point{[2]float64{180.0, 90.0}, nil}
	max = Point{[2]float64{-180.0, -90.0}, nil}
	min = Point{[2]float64{-180.0, -90.0}, nil}

	for _, i := range c.Points {
		pt := points[i]

		for j := range pt.XY {
			center.XY[j] += pt.XY[j]

			if pt.XY[j] < min.XY[j] {
				min.XY[j] = pt.XY[j]
			}
			if pt.XY[j] > max.XY[j] {
				max.XY[j] = pt.XY[j]
			}
		}
	}

	for j := range center.XY {
		center.XY[j] /= float64(len(c.Points))
	}

	return
}

// Inside checks if (innerMin, innerMax) rectangle is inside (outerMin, outMax) rectangle
func Inside(innerMin, innerMax, outerMin, outerMax *Point) bool {
	return innerMin.GreaterEq(outerMin) && innerMax.LessEq(outerMax)
}
