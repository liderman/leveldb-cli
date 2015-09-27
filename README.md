# leveldb-cli
Command-line utility for working with levelDB

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

Commands
--------
