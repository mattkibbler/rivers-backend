package tiles

import "math/rand"

func GenerateTile(x int, y int) TileContent {
	// Pick a random index from the slice
	material := TileMaterials[rand.Intn(len(TileMaterials))]

	return TileContent{
		Material: material,
		ZLevel:   1,
	}
}
