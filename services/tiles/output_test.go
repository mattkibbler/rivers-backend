package tiles

import (
	"bytes"
	"testing"
)

func TestTileRegionPacketEncoding(t *testing.T) {
	startX := int16(10)
	startY := int16(20)
	endX := int16(12)
	endY := int16(22)
	var packetData [][]TileContent
	for y := startY; y <= endY; y++ {
		var row []TileContent
		for x := startX; x <= endX; x++ {
			row = append(row, TileContent{
				Material: TileMaterial(x + y),
				ZLevel:   uint8(y),
			})
		}
		packetData = append(packetData, row)
	}

	packet := TileRegionPacket{
		Region: TileRegion{
			StartX: startX,
			StartY: startY,
			EndX:   endX,
			EndY:   endY,
		},
		Data: packetData,
	}

	var buf bytes.Buffer
	err := encodeTileRegionPacket(&buf, packet)
	if err != nil {
		t.Error(err)
	}

	expectedOutput := []byte{
		0x0A, 0x00, // startX (10)
		0x14, 0x00, // startY (20)
		0x0C, 0x00, // endX (12)
		0x16, 0x00, // endY (22)
		0x1E, 0x14, // Material for (10,20), ZLevel for (10,20)
		0x1F, 0x14, // Material for (11,20), ZLevel for (11,20)
		0x20, 0x14, // Material for (12,20), ZLevel for (12,20)
		0x1F, 0x15, // Material for (10,21), ZLevel for (10,21)
		0x20, 0x15, // Material for (11,21), ZLevel for (11,21)
		0x21, 0x15, // Material for (12,21), ZLevel for (12,21)
		0x20, 0x16, // Material for (10,22), ZLevel for (10,22)
		0x21, 0x16, // Material for (11,22), ZLevel for (11,22)
		0x22, 0x16, // Material for (12,22), ZLevel for (12,22)
	}

	if !bytes.Equal(buf.Bytes(), expectedOutput) {
		t.Errorf("Encoded output does not match expected output.\nGot: %v\nExpected: %v", buf.Bytes(), expectedOutput)
	}

}
