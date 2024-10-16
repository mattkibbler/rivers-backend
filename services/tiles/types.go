package tiles

import (
	"fmt"
	"strconv"
	"strings"
)

type TileMaterial int

const (
	TileMaterialStone TileMaterial = iota
	TileMaterialGrass
	TileMaterialDirt
)

var TileMaterials = []TileMaterial{TileMaterialStone, TileMaterialGrass, TileMaterialDirt}

type TileContent struct {
	Material TileMaterial `json:"material"`
	ZLevel   int          `json:"zLevel"`
}

type TileRegionPacket struct {
	Region TileRegion      `json:"region"`
	Data   [][]TileContent `json:"data"`
}

type TileRegionPacketCollection struct {
	Packets []TileRegionPacket `json:"packets"`
}

// export default interface TileDataPacket {
// 	region: TileRegion;
// 	data: TileData[][];
// }

type RegionQueryParam string

func (r *RegionQueryParam) Parse() (*TileRegion, error) {
	parts := strings.Split(string(*r), ",")
	if numParts := len(parts); numParts != 4 {
		return nil, fmt.Errorf("expected 4 parts in region params, got %v", numParts)
	}
	intParts := [4]int{}
	for i, part := range parts {
		val, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("all parts of query param must be an integer")
		}
		intParts[i] = val
	}
	return &TileRegion{
		StartX: intParts[0],
		StartY: intParts[1],
		EndX:   intParts[2],
		EndY:   intParts[3],
	}, nil
}

type TileRegion struct {
	StartX int `json:"startX"`
	StartY int `json:"startY"`
	EndX   int `json:"endX"`
	EndY   int `json:"endY"`
}

func (r TileRegion) String() string {
	return fmt.Sprintf("StartX: %d, StartY: %d, EndX: %d, EndY: %d", r.StartX, r.StartY, r.EndX, r.EndY)
}
