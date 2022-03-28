package main

type Relation struct {
	Uid     string `json:"uid"`
	Address string `json:"address"`
	Parent  string `json:"parent"`
	Level   int    `json:"level"`
}

type ReturnRelation struct {
	Relation     `json:"relation"`
	RelationList []Relation `json:"relation_list"`
}

type TreeRelation struct {
	Uid       string     `json:"uid"`
	Address   string     `json:"address"`
	Level     int        `json:"level"`
	Relations []Relation `json:"relations"`
}
