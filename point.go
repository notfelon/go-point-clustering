// Package cluster implements DBScan clustering on (lat, lon) using K-D Tree
package cluster

// Point is longitue, latittude
type Point struct {
	xy   [2]float64
	data []string
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
	return a.xy[0] <= b.xy[0] && a.xy[1] <= b.xy[1]
}

// GreaterEq - a >= b
func (a *Point) GreaterEq(b *Point) bool {
	return a.xy[0] >= b.xy[0] && a.xy[1] >= b.xy[1]
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

		for j := range pt.xy {
			center.xy[j] += pt.xy[j]

			if pt.xy[j] < min.xy[j] {
				min.xy[j] = pt.xy[j]
			}
			if pt.xy[j] > max.xy[j] {
				max.xy[j] = pt.xy[j]
			}
		}
	}

	for j := range center.xy {
		center.xy[j] /= float64(len(c.Points))
	}

	return
}

// Inside checks if (innerMin, innerMax) rectangle is inside (outerMin, outMax) rectangle
func Inside(innerMin, innerMax, outerMin, outerMax *Point) bool {
	return innerMin.GreaterEq(outerMin) && innerMax.LessEq(outerMax)
}
