package geoutil

import (
	"fmt"
	"testing"
)

func BenchmarkGetDistance(b *testing.B) {
	p1 := &Point{3.3000716307302, 101.57032339298446}
	p2 := &Point{3.3051875248786815, 101.5731201632972}

	for i := 0; i < b.N; i++ {
		GetDistance(p1, p2)
	}
}

func BenchmarkGetDistanceHaversine(b *testing.B) {
	p1 := &Point{3.3000716307302, 101.57032339298446}
	p2 := &Point{3.3051875248786815, 101.5731201632972}

	for i := 0; i < b.N; i++ {
		GetDistanceHaversine(p1, p2)
	}
}

func BenchmarkGetBoundary(b *testing.B) {
	p := &Point{3.3000716307302, 101.57032339298446}
	meters := 10.0

	for i := 0; i < b.N; i++ {
		GetBoundary(p, meters)
	}
}

func BenchmarkGeoHashEncode(b *testing.B) {
	p := &Point{3.300071631, 101.570323393}

	for i := 0; i < b.N; i++ {
		GeoHashEncode(p, 12)
	}
}

func BenchmarkGeoHashDecode(b *testing.B) {
	geohash := []byte("w284z2c221fq")

	for i := 0; i < b.N; i++ {
		GeoHashDecode(geohash)
	}
}

func TestPoint_Distance(t *testing.T) {
	p1 := &Point{3.096249444382203, 101.53713780926184}
	p2 := &Point{1.3361333233860981, 103.83548216216576}

	fmt.Println(p1.DistanceTo(p2) / 1000)
}

func TestGetDistanceHaversine(t *testing.T) {
	p1 := &Point{3.096249444382203, 101.53713780926184}
	p2 := &Point{1.3361333233860981, 103.83548216216576}

	fmt.Println(GetDistanceHaversine(p1, p2) / 1000)
}

func TestPoint_BoundaryOf(t *testing.T) {
	p := &Point{3.3000716307302, 101.57032339298446}
	meters := 10.0

	rect := p.BoundaryOf(meters)
	fmt.Println(rect)
	fmt.Println(p.DistanceTo(&Point{rect.Min.Lat, rect.Min.Lng}))
	fmt.Println(p.DistanceTo(&Point{rect.Max.Lat, rect.Max.Lng}))
	fmt.Println(p.DistanceTo(&Point{rect.Min.Lat, rect.Max.Lng}))
	fmt.Println(p.DistanceTo(&Point{rect.Max.Lat, rect.Min.Lng}))
}

func TestGeoHashEncode(t *testing.T) {
	p := &Point{3.3000716307302, 101.57032339298446}

	fmt.Println(p, ":", string(GeoHashEncode(p, 12)))
}

func TestGeoHashDecode(t *testing.T) {
	geohash := []byte("w284z2c221fq")

	fmt.Println(string(geohash), ":", GeoHashDecode(geohash))
}
