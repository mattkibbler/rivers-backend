package tiles

import (
	"encoding/binary"
	"io"
)

func encodeTileRegionPackets(buf io.Writer, pCol TileRegionPacketCollection) error {
	for _, packet := range pCol.Packets {
		if err := encodeTileRegionPacket(buf, packet); err != nil {
			return err
		}
	}
	return nil
}

func encodeTileRegionPacket(buf io.Writer, packet TileRegionPacket) error {
	if err := binary.Write(buf, binary.LittleEndian, packet.Region.StartX); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.LittleEndian, packet.Region.StartY); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.LittleEndian, packet.Region.EndX); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.LittleEndian, packet.Region.EndY); err != nil {
		return err
	}

	for _, y := range packet.Data {
		for _, x := range y {
			if err := binary.Write(buf, binary.LittleEndian, x.Material); err != nil {
				return err
			}
			if err := binary.Write(buf, binary.LittleEndian, x.ZLevel); err != nil {
				return err
			}
		}
	}

	return nil
}
