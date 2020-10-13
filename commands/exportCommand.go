// license that can be found in the LICENSE file.

// This file is part of the application source code leveldb-cli
// This software provides a console interface to leveldb.

package commands

import (
	"io/ioutil"
)

// The command exports a value for the selected key
// to file with specified filename.
//
// Returns a string containing information about the result of the operation.
func Export(key, fileName string) string {
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

	fileErr := ioutil.WriteFile(fileName, value, 0644)
	if fileErr != nil {
		return AppError(FileWriteErr)
	}

	return "Success"
}
