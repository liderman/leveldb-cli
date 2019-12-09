// Copyright 2015 Osipov Konstantin <k.osipov.msk@gmail.com>. All rights reserved.
// license that can be found in the LICENSE file.

// This file is part of the application source code leveldb-cli
// This software provides a console interface to leveldb.

package commands

import (
	"fmt"
)

// Command close the database.
// It is safe to close the database file leveldb.
//
// Returns a string containing information about the result of the operation.
func Close() string {
	if !isConnected {
		return AppError(ErrDbDoesNotOpen)
	}

	err := dbh.Close()
	if err != nil {
		return fmt.Sprintf(
			AppError(ErrCouldNotCloseDatabase),
			err.Error(),
		)
	}

	isConnected = false
	return "Success"
}
