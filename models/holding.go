package models

type Holding struct {

}

type HoldingResponse struct {
	Holding *Holding `json"holding"`
}

type HoldingsResponse struct {
	Holdings []*Holding `json"holdings"`
}
