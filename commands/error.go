// Copyright 2015 Osipov Konstantin <k.osipov.msk@gmail.com>. All rights reserved.
// license that can be found in the LICENSE file.

// This file is part of the application source code leveldb-cli
// This software provides a console interface to leveldb.

package commands
import "fmt"

const ERR_DB_DOES_NOT_OPEN          = 1001
const ERR_OPENING_DATABASE          = 1002
const ERR_UNABLE_TO_WRITE           = 1003
const ERR_KEY_IS_EMPTY              = 1004
const ERR_UNABLE_TO_DELETE          = 1005
const ERR_COULD_NOT_CLOSE_DATABASE  = 1006

// Error messages list
var errorMessages = map[int]string{
	ERR_DB_DOES_NOT_OPEN:         "Database does not open",
	ERR_OPENING_DATABASE:         "Error opening database `%s`",
	ERR_UNABLE_TO_WRITE:          "Unable to write [`%s`]",
	ERR_KEY_IS_EMPTY:             "Key is exmpty",
	ERR_UNABLE_TO_DELETE:         "Unable to delete [`%s`]",
	ERR_COULD_NOT_CLOSE_DATABASE: "Could not close database [`%s`]",
}

// The wrapper for outputting errors in the application
// Returns the text of the error
func AppError(code int) string {
	msg, ok := errorMessages[code]
	if (ok) {
		return fmt.Sprintf("Error %d: %s!", code, msg)
	}

	return "Unknown error"
}
