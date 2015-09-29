// Copyright 2015 Osipov Konstantin <k.osipov.msk@gmail.com>. All rights reserved.
// license that can be found in the LICENSE file.

// This file is part of the application source code leveldb-cli
// This software provides a console interface to leveldb.

package commands

import (
	"github.com/syndtr/goleveldb/leveldb/util"
	"github.com/syndtr/goleveldb/leveldb/iterator"
	"fmt"
	"text/tabwriter"
	"github.com/liderman/leveldb-cli/cliutil"
	"bytes"
	"bufio"
)

// It shows the contents of the database prefix filtering.
// Use the field `format` for specifying the display format of data.
// The list of possible values of format options: raw (default), geohash, bson, int64, float64
//
// Returns a string containing information about the result of the operation.
func ShowByPrefix(prefix, format string) string {
	if (!isConnected) {
		return AppError(ERR_DB_DOES_NOT_OPEN)
	}

	return showByIterator(
		dbh.NewIterator(util.BytesPrefix([]byte(prefix)), nil),
		format,
	)
}

// It shows the contents of the database range filtering.
// Use the field `format` for specifying the display format of data.
// The list of possible values of format options: raw (default), geohash, bson, int64, float64
//
// Returns a string containing information about the result of the operation.
func ShowByRange(start, limit, format string) string {
	if (!isConnected) {
		return AppError(ERR_DB_DOES_NOT_OPEN)
	}

	return showByIterator(
		dbh.NewIterator(&util.Range{Start: []byte(start), Limit: []byte(limit)}, nil),
		format,
	)
}

// Show by iterator
//
// Returns a string containing information about the result of the operation.
func showByIterator(iter iterator.Iterator, format string) string {
	if iter.Error() != nil {
		return "Empty result!"
	}

	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	w := new(tabwriter.Writer)

	w.Init(writer, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, "Key\t| Value")
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()

		fmt.Fprintf(w, "%s\t| %s\n", string(key), cliutil.ToString(format, value))
	}

	w.Flush()

	iter.Release()
	err := iter.Error()
	if (err != nil) {
		return "Error iterator!"
	}

	writer.Flush()
	return string(b.Bytes())
}
