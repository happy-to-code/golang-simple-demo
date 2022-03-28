package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var allRelations = []Relation{
		{
			Uid:     "7",
			Address: "7",
			Parent:  "8",
			Level:   0,
		},
		{
			Uid:     "3",
			Address: "4",
			Parent:  "7",
			Level:   1,
		},
		{
			Uid:     "2",
			Address: "5",
			Parent:  "4",
			Level:   2,
		},
		{
			Uid:     "5",
			Address: "3",
			Parent:  "4",
			Level:   2,
		},
		{
			Uid:     "1",
			Address: "1",
			Parent:  "4",
			Level:   2,
		},
		{
			Uid:     "6",
			Address: "6",
			Parent:  "5",
			Level:   3,
		},
		{
			Uid:     "4",
			Address: "2",
			Parent:  "3",
			Level:   3,
		},
	}

	fmt.Println(allRelations)

	level2Relations := make(map[int][]Relation)
	for _, rela := range allRelations {
		level := rela.Level
		relation, _ := level2Relations[level]
		relation = append(relation, rela)
		level2Relations[level] = relation
	}
	for k, v := range level2Relations {
		fmt.Println(k, "==>", v)
	}
	fmt.Println("---------------------")
	var treeRelations *[]TreeRelation
	for _, rela := range allRelations {
		treeRelation := TreeRelation{
			Uid:     rela.Uid,
			Address: rela.Address,
			Level:   rela.Level,
		}
		// 	判断treeRelations 列表中是否存在此对象
		exist, relations := isExist(*treeRelations, treeRelation)
		if exist {
			relations = append(relations, rela)
			//	替换
			for _, tr := range *treeRelations {
				if tr.Level == rela.Level {
					tr.Relations = relations
				}
			}
		} else {
			treeRelation.Relations = []Relation{rela}

			*treeRelations = append(*treeRelations, treeRelation)
		}

	}
	marshal, _ := json.Marshal(treeRelations)
	fmt.Println(string(marshal))
	fmt.Println("------------------------------------------------------------")
	fmt.Printf("%+v\n", treeRelations)

}

func isExist(treeRelations []TreeRelation, targetTreeRelation TreeRelation) (bool, []Relation) {
	for _, tr := range treeRelations {
		// 存在
		if tr.Level == targetTreeRelation.Level {
			return true, tr.Relations
		}
	}
	return false, nil
}
