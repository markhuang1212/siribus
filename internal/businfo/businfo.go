package businfo

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type BusEtaResponse struct {
	Data []struct {
		EtaSeq int    `json:"eta_seq"`
		Eta    string `json:"eta"`
	} `json:"data"`
}

func GetTimeFromNow(route string, busstop string, direction string) ([]time.Duration, error) {

	url := Endpoint + "/transport/kmb/eta/" + busstop + "/" + route + "/" + direction

	log.Println(url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var busEtaResp BusEtaResponse
	json.Unmarshal(body, &busEtaResp)

	var etas []time.Duration
	for _, d := range busEtaResp.Data {
		t, err := time.Parse("2006-01-02T15:04:05Z07:00", d.Eta)
		if err != nil {
			return nil, err
		}
		etas = append(etas, time.Until(t))
	}

	return etas, nil
}
