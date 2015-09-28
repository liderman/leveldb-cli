# leveldb-cli
Command-line utility for working with levelDB

## Installation and build

```
go get github.com/liderman/leveldb-cli
go install
```

## Requirements
 * `go1.5` or newer.


## Usage

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
testdb» show prefix key
Key	| Value
key100	| value100
key200	| value200
key300	| value300

testdb» show range key2 key3
Key	| Value
key200	| value200

testdb» close
Success
» exit
```

## Commands

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

### show
> show prefix `KEY_PREFIX`

Displays all values the keys that begin with the prefix.
 * `KEY_PREFIX` - The prefix to list of keys

> show range `START` `LIMIT`

Displays all values, the keys of which are in the range between "START" and "LIMIT".
 * `START` - The key or key prefix indicating the beginning of the range
 * `LIMIT` - The key or key prefix indicating the end of the range
 
### help
> help

Displays short usage software

### version
> version

Displays the current version of software and operating systems on which it runs
