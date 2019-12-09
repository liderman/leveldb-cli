// Copyright 2015 Osipov Konstantin <k.osipov.msk@gmail.com>. All rights reserved.
// license that can be found in the LICENSE file.

// This file is part of the application source code leveldb-cli
// This software provides a console interface to leveldb.

package commands

import (
	"fmt"
)

// Command deleting records.
// The command to remove records from the database by key.
//
// Returns a string containing information about the result of the operation.
func Delete(key string) string {
	if !isConnected {
		return AppError(ErrDbDoesNotOpen)
	}

	if key == "" {
		return AppError(ErrKeyIsEmpty)
	}

	err := dbh.Delete([]byte(key), nil)
	if err != nil {
		return fmt.Sprintf(
			AppError(ErrUnableToDelete),
			err.Error(),
		)
	}

	return "Success"
}
