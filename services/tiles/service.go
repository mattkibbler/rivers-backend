package tiles

import (
	"database/sql"
	"net/http"

	"github.com/mattkibbler/rivers-backend/api"
	"github.com/mattkibbler/rivers-backend/output"
)

type Service struct {
	store *Store
}

func NewService(db *sql.DB) *Service {
	store := NewStore(db)
	return &Service{
		store: store,
	}
}

func (s *Service) RegisterRoutes(server *api.ApiServer) {
	server.Get("/api/v1/tiles/regions", s.handleGetRegions)
}

func (s *Service) handleGetRegions(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow any origin

	q := r.URL.Query()
	qr := q["regions[]"]
	var regions []TileRegion
	for _, regStr := range qr {
		regQueryParam := RegionQueryParam(regStr)
		parsedRegion, err := regQueryParam.Parse()
		if err != nil {
			output.WriteError(w, http.StatusBadRequest, err)
			return
		}
		regions = append(regions, *parsedRegion)
	}

	packets := []TileRegionPacket{}
	for _, region := range regions {
		packet := s.store.GetPacket(region)
		packets = append(packets, packet)
	}

	output.WriteJSON(w, http.StatusOK, TileRegionPacketCollection{
		Packets: packets,
	})

}
