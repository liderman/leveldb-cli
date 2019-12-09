// Copyright 2015 Osipov Konstantin <k.osipov.msk@gmail.com>. All rights reserved.
// license that can be found in the LICENSE file.

// This file is part of the application source code leveldb-cli
// This software provides a console interface to leveldb.

package commands

import (
	"fmt"
)

// The command set a value.
// It sets the value for the selected key.
//
// Returns a string containing information about the result of the operation.
func Set(key, value string) string {
	if !isConnected {
		return AppError(ErrDbDoesNotOpen)
	}

	if key == "" {
		return AppError(ErrKeyIsEmpty)
	}

	err := dbh.Put([]byte(key), []byte(value), nil)
	if err != nil {
		return fmt.Sprintf(
			AppError(ErrUnableToWrite),
			err.Error(),
		)
	}

	return "Success"
}
