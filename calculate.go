package geoutil

import (
	"math"
)

const (
	EarthRadius = Distance(6370996.81) // meters
)

// radians converts degrees to radians.
func radians(degrees float64) float64 {
	return degrees * (math.Pi / 180.0)
}

// GetBoundary returns the boundary of a circle.
func GetBoundary(point *Point, meters float64) (rect *Boundary) {
	return &Boundary{
		Min: Point{
			Lat: point.Lat - (meters / 111_000.0),
			Lng: point.Lng - (meters / 111_000.0 / math.Cos(radians(point.Lat))),
		},
		Max: Point{
			Lat: point.Lat + (meters / 111_000.0),
			Lng: point.Lng + (meters / 111_000.0 / math.Cos(radians(point.Lat))),
		},
	}
}

// GetDistance returns the distance between two points in meters.
func GetDistance(p1, p2 *Point) Distance {
	rLat1 := radians(p1.Lat)
	rLat2 := radians(p2.Lat)
	dLng := radians(p2.Lng - p1.Lng)

	return Distance(EarthRadius.Meters() * math.Acos(math.Cos(rLat1)*math.Cos(rLat2)*math.Cos(dLng)+math.Sin(rLat1)*math.Sin(rLat2)))
}

// IsValidPoint returns true if the point is valid.
func IsValidPoint(point *Point) bool {
	return point.Lat >= -90 && point.Lat <= 90 && point.Lng >= -180 && point.Lng <= 180
}

// haversine returns the haversine of the angle.
func haversine(theta float64) float64 {
	return (1 - math.Cos(theta)) / 2
}

// GetDistanceHaversine returns the distance between two points in meters.
func GetDistanceHaversine(point1, point2 *Point) Distance {
	rLat1 := radians(point1.Lat)
	rLat2 := radians(point2.Lat)

	dLat := rLat2 - rLat1
	dLng := radians(point2.Lng) - radians(point1.Lng)

	_a := haversine(dLat) + (math.Cos(rLat1) * math.Cos(rLat2) * haversine(dLng))
	_c := 2 * math.Atan2(math.Sqrt(_a), math.Sqrt(1-_a))

	return Distance(EarthRadius.Meters() * _c)
}
