package tiles

import (
	"github.com/aquilax/go-perlin"
)

var p *perlin.Perlin

func init() {
	p = newPerlinNoiseGenerator()
}

func newPerlinNoiseGenerator() *perlin.Perlin {
	// Set up the Perlin noise generator
	alpha := 5.0      // The 'persistence' value
	beta := 2.0       // The 'frequency' value
	n := int32(3)     // Number of octaves (noise layers)
	seed := int64(42) // Seed for randomness

	return perlin.NewPerlin(alpha, beta, n, seed)
}

// Convert Perlin noise value (-1 to 1) to a height (0 to 256) as an integer
func perlinNoiseToHeight(noise float64) int {
	return int((noise + 1) / 2 * 256)
}

func GenerateTile(x int, y int) TileContent {
	// Add a scale factor to vary the input to the Perlin noise function
	scale := 0.06 // Adjust this value to control the "zoom level" of the noise
	// Pick a random index from the slice
	//material := TileMaterials[rand.Intn(len(TileMaterials))]
	material := TileMaterialStone
	noiseValue := p.Noise2D(float64(x)*scale, float64(y)*scale)

	return TileContent{
		Material: material,
		ZLevel:   perlinNoiseToHeight(noiseValue),
	}
}
