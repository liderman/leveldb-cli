// Copyright 2015 Osipov Konstantin <k.osipov.msk@gmail.com>. All rights reserved.
// license that can be found in the LICENSE file.

// This file is part of the application source code leveldb-cli
// This software provides a console interface to leveldb.

package commands

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/liderman/leveldb-cli/cliutil"
	"github.com/syndtr/goleveldb/leveldb/iterator"
	"github.com/syndtr/goleveldb/leveldb/util"
	"text/tabwriter"
)

// It shows the contents of the database prefix filtering.
// Use the field `format` for specifying the display format of data.
// The list of possible values of format options: raw (default), geohash, bson, int64, float64
//
// Returns a string containing information about the result of the operation.
func ShowByPrefix(prefix, format string) string {
	if !isConnected {
		return AppError(ErrDbDoesNotOpen)
	}

	return showByIterator(
		dbh.NewIterator(util.BytesPrefix([]byte(prefix)), nil),
		format,
		0,
	)
}

// It shows the contents of the database range filtering.
// Use the field `format` for specifying the display format of data.
// The list of possible values of format options: raw (default), geohash, bson, int64, float64
//
// Returns a string containing information about the result of the operation.
func ShowByRange(start, limit, format string) string {
	if !isConnected {
		return AppError(ErrDbDoesNotOpen)
	}

	return showByIterator(
		dbh.NewIterator(&util.Range{Start: []byte(start), Limit: []byte(limit)}, nil),
		format,
		0,
	)
}

// ShowLimit It shows all content of the database limited by the limit.
// Use the field `format` for specifying the display format of data.
// The list of possible values of format options: raw (default), geohash, bson, int64, float64
//
// Returns a string containing information about the result of the operation.
func ShowLimit(limit int, format string) string {
	if !isConnected {
		return AppError(ErrDbDoesNotOpen)
	}

	return showByIterator(
		dbh.NewIterator(nil, nil),
		format,
		limit,
	)
}

// Show by iterator
//
// Returns a string containing information about the result of the operation.
func showByIterator(iter iterator.Iterator, format string, limit int) string {
	if iter.Error() != nil {
		return "Empty result!"
	}

	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	w := new(tabwriter.Writer)

	w.Init(writer, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, "Key\t| Value")

	count := 0
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()

		fmt.Fprintf(w, "%s\t| %s\n", string(key), cliutil.ToString(format, value))

		count++
		if limit != 0 && count >= limit {
			break
		}
	}

	w.Flush()

	iter.Release()
	err := iter.Error()
	if err != nil {
		return "Error iterator!"
	}

	writer.Flush()
	return string(b.Bytes())
}
