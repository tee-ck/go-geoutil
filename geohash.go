package geoutil

// b32chars is the base32 character set used by geohash.
var b32chars = []byte("0123456789bcdefghjkmnpqrstuvwxyz")

// b32maps maps each rune in b32chars to its index.
var b32maps = map[byte]int{
	'0': 0x00, '1': 0x01, '2': 0x02, '3': 0x03, '4': 0x04, '5': 0x05, '6': 0x06, '7': 0x07,
	'8': 0x08, '9': 0x09, 'b': 0x0a, 'c': 0x0b, 'd': 0x0c, 'e': 0x0d, 'f': 0x0e, 'g': 0x0f,
	'h': 0x10, 'j': 0x11, 'k': 0x12, 'm': 0x13, 'n': 0x14, 'p': 0x15, 'q': 0x16, 'r': 0x17,
	's': 0x18, 't': 0x19, 'u': 0x1a, 'v': 0x1b, 'w': 0x1c, 'x': 0x1d, 'y': 0x1e, 'z': 0x1f,
}

// GeoHashDecode decodes a geohash bytearray into a point.
func GeoHashDecode(geohash []byte) (p *Point) {
	// from: https://github.com/vinsci/geohash
	p = new(Point)
	minLat, maxLat, minLng, maxLng := -90.0, 90.0, -180.0, 180.0
	even := true

	masks := []int{16, 8, 4, 2, 1}
	for _, c := range geohash {
		cd := b32maps[c]
		for _, m := range masks {
			if even {
				if (cd & m) == 0 {
					maxLng = (minLng + maxLng) / 2
				} else {
					minLng = (minLng + maxLng) / 2
				}
			} else {
				if (cd & m) == 0 {
					maxLat = (minLat + maxLat) / 2
				} else {
					minLat = (minLat + maxLat) / 2
				}
			}

			even = !even
		}
	}

	p.Lat = (minLat + maxLat) / 2
	p.Lng = (minLng + maxLng) / 2

	return p
}

// GeoHashEncode encodes a point into a geohash bytearray.
func GeoHashEncode(p *Point, precision int) []byte {
	minLat, maxLat, minLng, maxLng := -90.0, 90.0, -180.0, 180.0
	if precision < 1 || precision > 22 {
		precision = 12
	}
	geohash := make([]byte, precision)

	midLat, midLng := (minLat+maxLat)/2, (minLng+maxLng)/2
	for idx, bit := 0, 1; idx < precision; midLat, midLng = (minLat+maxLat)/2, (minLng+maxLng)/2 {
		bit <<= 1
		if p.Lng > midLng {
			bit |= 0b1
			minLng = midLng

		} else {
			maxLng = midLng
		}

		if bit > 0b11111 {
			geohash[idx] = b32chars[bit&0b11111]
			idx, bit = idx+1, 1

			if idx >= precision {
				break
			}
		}

		bit <<= 1
		if p.Lat > midLat {
			bit |= 0b1
			minLat = midLat

		} else {
			maxLat = midLat
		}

		if bit > 0b11111 {
			geohash[idx] = b32chars[bit&0b11111]
			idx, bit = idx+1, 1

			if idx >= precision {
				break
			}
		}
	}

	return geohash
}
