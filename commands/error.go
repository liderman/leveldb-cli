// Copyright 2015 Osipov Konstantin <k.osipov.msk@gmail.com>. All rights reserved.
// license that can be found in the LICENSE file.

// This file is part of the application source code leveldb-cli
// This software provides a console interface to leveldb.

package commands

import "fmt"

const ErrDbDoesNotOpen = 1001
const ErrOpeningDatabase = 1002
const ErrUnableToWrite = 1003
const ErrKeyIsEmpty = 1004
const ErrUnableToDelete = 1005
const ErrCouldNotCloseDatabase = 1006
const ErrKeyNotFound = 1007

// Error messages list
var errorMessages = map[int]string{
	ErrDbDoesNotOpen:         "Database does not open",
	ErrOpeningDatabase:       "Error opening database `%s`",
	ErrUnableToWrite:         "Unable to write [`%s`]",
	ErrKeyIsEmpty:            "Key is exmpty",
	ErrUnableToDelete:        "Unable to delete [`%s`]",
	ErrCouldNotCloseDatabase: "Could not close database [`%s`]",
	ErrKeyNotFound:           "Key not found",
}

// The wrapper for outputting errors in the application
// Returns the text of the error
func AppError(code int) string {
	msg, ok := errorMessages[code]
	if ok {
		return fmt.Sprintf("Error %d: %s!", code, msg)
	}

	return "Unknown error"
}
