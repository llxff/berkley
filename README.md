# BerkeleyDB excexerciseersice

A command-line utility for working with BerkeleyDB databases, built during the [BerkeleyDB B-Tree tutorial series](https://transactional.blog/building-berkeleydb).

## Basic Example

Generate random key/value pairs that can be loaded into a BerkeleyDB database using `db_load`:

```bash
./berkley generate -n 20 | db_load -T -t btree testdata.bdb
```

This command:
1. Generates 20 random key/value pairs
2. Pipes the output to `db_load`
3. Creates a BerkeleyDB btree database file named `testdata.bdb`

### Options

- `-n <number>`: Specify the number of key/value pairs to generate

## Requirements

- BerkeleyDB utilities (`db_load`) must be installed on your system
- This tool must be in your system PATH
