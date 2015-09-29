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
	if (!isConnected) {
		return AppError(ERR_DB_DOES_NOT_OPEN)
	}

	if (key == "") {
		return AppError(ERR_KEY_IS_EMPTY)
	}

	err := dbh.Delete([]byte(key), nil)
	if err != nil {
		return fmt.Sprintf(
			AppError(ERR_UNABLE_TO_DELETE),
			err.Error(),
		)
	}

	return "Success"
}
