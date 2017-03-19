package models

import (
	"gopkg.in/mgo.v2/bson"
)

type ShigeDoc struct {
	Title string
	Author string
	Period string
	Doc string
	Labels string
}


func ShigeData() ShigeDoc {
	c := MgoDb("shige").C("shigedoc")
	reult := ShigeDoc{}

	c.Find(bson.M{}).One(&reult)
	return reult
}

func ShigeDataByQ(q bson.M) []ShigeDoc {
	c := MgoDb("shige").C("shigedoc")
	reult := []ShigeDoc{}
	c.Find(bson.M{}).All(&reult)
	return reult
}