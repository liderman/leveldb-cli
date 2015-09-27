// Copyright 2015 Osipov Konstantin <k.osipov.msk@gmail.com>. All rights reserved.
// license that can be found in the LICENSE file.

// This file is part of the application source code leveldb-cli
// This software provides a console interface to leveldb.

package cliutil

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/TomiHiltunen/geohash-golang"
	"fmt"
	"encoding/binary"
	"math"
)

func ToString(format string, value []byte) string {
	switch format {
	case "bson":
		return bsonToString(value)
	case "geohash":
		return geohashToString(value)
	case "int64":
		return int64ToString(value)
	case "float64":
		return float64ToString(value)
	case "raw":
	default:
	}

	return string(value)
}

func bsonToString(value []byte) string {
	var dst interface{}
	err := bson.Unmarshal(value, &dst)
	if err != nil {
		return "Error converting!"
	}

	return fmt.Sprintf("%#v", dst)
}

func geohashToString(value []byte) string {
	position := geohash.Decode(string(value))

	return fmt.Sprintf("lat: %f lng: %f", position.Center().Lat(), position.Center().Lng())
}


func int64ToString(value []byte) string {
	return fmt.Sprintf("%d", binary.BigEndian.Uint64(value))
}

func float64ToString(value []byte) string {
	return fmt.Sprintf("%f", math.Float64frombits(
		binary.LittleEndian.Uint64(value),
	))
}