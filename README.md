LevelDB-CLI: a simple utility for debugging LevelDB
===========
Command-line utility for working with levelDB.
This utility is useful for debugging applications using the database LevelDB


![Demo GIF](https://raw.githubusercontent.com/liderman/leveldb-cli/master/docs/live-demo.gif)

Installation and build
----------------------

```
go get github.com/liderman/leveldb-cli
go install
```

Requirements
------------
 * `go1.5` or newer.

Usage
-----

```
# ./leveldb-cli
```

```
» open testdb
Database not exist! Create new database.
Success
testdb» set key100 value100
Success
testdb» set key200 value200
Success
testdb» set key300 value300
Success
testdb» set "key \"123" value
Success
testdb» show prefix key
Key	      | Value
key100	  | value100
key200	  | value200
key300	  | value300
key \"123 | value

testdb» show range key2 key3
Key	| Value
key200	| value200

testdb» close
Success
» exit
```

Commands
--------

### open
> open `DATABASE_NAME`

Opens database.
If the database does not exist, it is created.
You can use this method to create a new database.
 * `DATABASE_NAME` - The database name or path

### close
> close

It closes a previously opened database.

### set
> set `KEY` `VALUE`

Set the value of for a key.
 * `KEY` - The key
 * `VALUE` - The value

### delete
> delete `KEY`

Delete the record by key.
 * `KEY` - The key

### get
> get `KEY` [`FORMAT`]

Display value by key.
 * `KEY` - The key
 * `FORMAT` - Data Display Format (Optional)

### export
> export `KEY` `FILENAME`

Write value by key to file with specified filename or path.
 * `KEY` - The key
 * `FILENAME` - File name of the output file

### show
> show prefix `KEY_PREFIX` [`FORMAT`]

Displays all values the keys that begin with the prefix.
 * `KEY_PREFIX` - The prefix to list of keys
 * `FORMAT` - Data Display Format (Optional)

> show range `START` `LIMIT` [`FORMAT`]

Displays all values, the keys of which are in the range between "START" and "LIMIT".
 * `START` - The key or key prefix indicating the beginning of the range
 * `LIMIT` - The key or key prefix indicating the end of the range
 * `FORMAT` - Data Display Format (Optional)

#### The list of formats available to display
 * `raw` - Raw data without processing (default)
 * `bson` - Attempts to convert the data to be displayed from `bson` to `json`
 * `geohash` - Attempts to convert the data format of the `geohash` in the coordinates of the center (lat, lng)
 * `int64` - Attempts to display the data as an integer 64-bit number
 * `float64` - Attempts to display the data as a 64-bit number c with a floating point
 
### help
> help

Displays short usage software

### version
> version

Displays the current version of software and operating systems on which it runs

LICENSE
-------
Project distributes with standard MIT license
