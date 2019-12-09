// Copyright 2015 Osipov Konstantin <k.osipov.msk@gmail.com>. All rights reserved.
// license that can be found in the LICENSE file.

// This file is part of the application source code leveldb-cli
// This software provides a console interface to leveldb.

package commands

import (
	"github.com/liderman/leveldb-cli/cliutil"
)

// The command get a value.
// It gets the value for the selected key.
//
// Returns a string containing information about the result of the operation.
func Get(key, format string) string {
	if !isConnected {
		return AppError(ErrDbDoesNotOpen)
	}

	if key == "" {
		return AppError(ErrKeyIsEmpty)
	}

	value, err := dbh.Get([]byte(key), nil)
	if err != nil {
		return AppError(ErrKeyNotFound)
	}

	return cliutil.ToString(format, value)
}
