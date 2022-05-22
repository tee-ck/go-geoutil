package geoutil

import (
	"fmt"
	"runtime"
	"testing"
)

func BenchmarkGetDistance(b *testing.B) {
	p1 := NewPoint(3.3000716307302, 101.57032339298446)
	p2 := NewPoint(3.3051875248786815, 101.5731201632972)

	for i := 0; i < b.N; i++ {
		GetDistance(p1, p2)
	}
}

func BenchmarkFastGetDistance(b *testing.B) {
	p1 := NewPoint(3.3000716307302, 101.57032339298446)
	p2 := NewPoint(3.3051875248786815, 101.5731201632972)

	for i := 0; i < b.N; i++ {
		FastGetDistance(p1, p2)
	}
}

func BenchmarkGetDistanceHaversine(b *testing.B) {
	p1 := NewPoint(3.3000716307302, 101.57032339298446)
	p2 := NewPoint(3.3051875248786815, 101.5731201632972)

	for i := 0; i < b.N; i++ {
		GetDistanceHaversine(p1, p2)
	}
}

func BenchmarkGetBoundary(b *testing.B) {
	p := NewPoint(3.3000716307302, 101.57032339298446)
	distance := 10.0 * Meter

	for i := 0; i < b.N; i++ {
		GetBoundary(p, distance)
	}
}

func BenchmarkGeoHashEncode8(b *testing.B) {
	p := NewPoint(3.300071631, 101.570323393)

	for i := 0; i < b.N; i++ {
		GeoHashEncode(p, 8)
	}
}

func BenchmarkGeoHashDecode8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GeoHashDecode("w284z2c2")
	}
}

func BenchmarkGeoHashEncode12(b *testing.B) {
	p := NewPoint(3.300071631, 101.570323393)

	for i := 0; i < b.N; i++ {
		GeoHashEncode(p, 12)
	}
}

func BenchmarkGeoHashDecode12(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GeoHashDecode("w284z2c221fq")
	}
}

func BenchmarkGeoHashEncode22(b *testing.B) {
	p := NewPoint(3.300071631, 101.570323393)

	for i := 0; i < b.N; i++ {
		GeoHashEncode(p, 22)
	}
}

func BenchmarkGeoHashDecode22(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GeoHashDecode("w284z2c221fqjf5yt97q9y")
	}
}

func TestDistanceMeasure(t *testing.T) {
	start := NewPoint(1.4651887210464056, 103.76371843105089)
	paths := []*Point{
		NewPoint(1.4664832398638055, 103.7622076877256),
		NewPoint(1.4721961443069989, 103.75729928463724),
		NewPoint(1.4808530942869949, 103.75039246029152),
		NewPoint(1.5144551586540216, 103.72812948115704),
		NewPoint(1.5993485177062203, 103.66178554408644),
		NewPoint(1.7950913251194864, 103.50946273972838),
		NewPoint(2.084919122258202, 103.26650368581207),
		NewPoint(2.465410754148679, 102.95076322420313),
		NewPoint(2.8812187025852984, 102.572948714497),
		NewPoint(3.5422561348163066, 102.10772445379378),
		NewPoint(4.381767203988625, 101.51148358056182),
		NewPoint(5.393181816403189, 100.8510902072363),
		NewPoint(6.5742590500672184, 100.2186383364558),
	}

	for _, p := range paths {
		fmt.Println(GetDistance(start, p))
		fmt.Println(FastGetDistance(start, p))
		fmt.Println()
	}
}

func TestGetDistanceHaversine(t *testing.T) {
	p1 := NewPoint(3.096249444382203, 101.53713780926184)
	p2 := NewPoint(1.3361333233860981, 103.83548216216576)

	fmt.Println(GetDistanceHaversine(p1, p2) / 1000)
}

func TestPoint_Distance(t *testing.T) {
	p1 := NewPoint(3.096249444382203, 101.53713780926184)
	p2 := NewPoint(1.3361333233860981, 103.83548216216576)

	fmt.Println(p1.DistanceTo(p2) / 1000)
}

func TestPoint_BoundaryOf(t *testing.T) {
	p := NewPoint(3.3000716307302, 101.57032339298446)
	distance := 10.0 * Meter

	rect := p.BoundaryOf(distance)
	fmt.Println(rect)
	fmt.Println(p.DistanceTo(NewPoint(rect.Min.Lat, rect.Min.Lng)))
	fmt.Println(p.DistanceTo(NewPoint(rect.Max.Lat, rect.Max.Lng)))
	fmt.Println(p.DistanceTo(NewPoint(rect.Min.Lat, rect.Max.Lng)))
	fmt.Println(p.DistanceTo(NewPoint(rect.Max.Lat, rect.Min.Lng)))
}

func TestGeoHashEncode(t *testing.T) {
	//p := NewPoint(3.3000716307302, 101.57032339298446)
	p := NewPoint(39.9257460000, 116.5998310000)

	for i := 1; i < 23; i++ {
		fmt.Println(p, ":", string(GeoHashEncode(p, i)))
	}
	fmt.Println(int64(90e12))
}

func TestGeoHashDecode(t *testing.T) {
	geohash := "w284z2c221fq"

	fmt.Println(string(geohash), ":", GeoHashDecode(geohash))
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
