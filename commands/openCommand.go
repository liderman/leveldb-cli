// Copyright 2015 Osipov Konstantin <k.osipov.msk@gmail.com>. All rights reserved.
// license that can be found in the LICENSE file.

// This file is part of the application source code leveldb-cli
// This software provides a console interface to leveldb.

package commands

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
)

// Database connection
var dbh *leveldb.DB

// DB connection flag
var isConnected bool

// The command to open the database.
// If the database does not exist, it will create a new database.
//
// Returns a string containing information about the result of the operation.
func Open(file string) string {
	var err error
	dbh, err = leveldb.OpenFile(file, nil)
	if err != nil {
		return fmt.Sprintf(AppError(ErrOpeningDatabase), file)
	}

	isConnected = true
	return "Success"
}
