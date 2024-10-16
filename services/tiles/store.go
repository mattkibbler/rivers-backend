package tiles

import (
	"database/sql"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetPacket(region TileRegion) TileRegionPacket {
	var result [][]TileContent
	for y := region.StartY; y <= region.EndY; y++ {
		var row []TileContent
		for x := region.StartX; x <= region.EndX; x++ {
			row = append(row, GenerateTile(x, y))
		}
		result = append(result, row)
	}
	return TileRegionPacket{
		Region: region,
		Data:   result,
	}

	// 		const result = [];
	// 		for (let y = tileRegions[i].startY; y <= tileRegions[i].endY; y++) {
	// 			const row = [];
	// 			for (let x = tileRegions[i].startX; x <= tileRegions[i].endX; x++) {
	// 				row.push(this.getTile(x, y));
	// 			}
	// 			result.push(row);
	// 		}
	// 		packets.push({
	// 			data: result,
	// 			region: tileRegions[i],
	// 		});
}
