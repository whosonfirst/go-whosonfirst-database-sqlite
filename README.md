# go-whosonfirst-database-sqlite

Go package providing tools for indexing Who's On First records in SQLite databases.

## Documentation

Documentation is incomplete at this time.

## Under the hood

This package is a very thin wrapper around the [whosonfirst/go-whosonfirst-database](https://github.com/whosonfirst/go-whosonfirst-database) package. That package provides all of the actual functionality for indexing Who's On First records but does NOT load any specific `database/sql` drivers. That happens in this package.

Both the [mattn/go-sqlite3](https://github.com/mattn/go-sqlite3) and the [modernc.org/sqlite](https://pkg.go.dev/modernc.org/sqlite) drivers are made available for tools built with the `-tags mattn` and `-tags modernc` flags, respectively.

Note that because both of the package (and especially the `modernc.org/sqlite` package are very large the `vendor` package is not included by defaut with this package. You will need to `go mod tidy` and `go mod vendor` the dependencies manually.

## Tools

```
$> make cli
go build -tags mattn -mod vendor -ldflags="-s -w" -o bin/wof-sqlite-index cmd/wof-sqlite-index/main.go
```

If you want to specify alternate tags you can do so with the `TAGS=` Makefile variable. For example:

```
$> make cli TAGS=mattn,modernc
```

### wof-sqlite-index

Index one or more Who's On First sources in a SQLite database.

_This tool supersedes the `wof-sqlite-index-features` tool from the [whosonfirst/go-whosonfist-sqlite-features-index](https://github.com/whosonfirst/go-whosonfirst-sqlite-features-index) package._

```
$> ./bin/wof-sqlite-index -h
  -all
    	Index all tables (except the 'search' and 'geometries' tables which you need to specify explicitly)
  -ancestors
    	Index the 'ancestors' tables
  -concordances
    	Index the 'concordances' tables
  -database-uri string
    	A URI in the form of 'sql://{DATABASE_SQL_ENGINE}?dsn={DATABASE_SQL_DSN}'. For example: sql://sqlite3?dsn=test.db
  -geojson
    	Index the 'geojson' table
  -geometries
    	Index the 'geometries' table (requires that libspatialite already be installed)
  -index-alt value
    	Zero or more table names where alt geometry files should be indexed.
  -index-alt-files
    	Index alt geometries. This flag is deprecated, please use -index-alt=TABLE,TABLE,etc. instead. To index alt geometries in all the applicable tables use -index-alt=*
  -index-relations
    	Index the records related to a feature, specifically wof:belongsto, wof:depicts and wof:involves. Alt files for relations are not indexed at this time.
  -index-relations-reader-uri string
    	A valid go-reader.Reader URI from which to read data for a relations candidate.
  -iterator-uri string
    	A valid whosonfirst/go-whosonfirst-iterate/v2 URI. Supported emitter URI schemes are: directory://,featurecollection://,file://,filelist://,geojsonl://,null://,repo:// (default "repo://")
  -names
    	Index the 'names' table
  -optimize
    	Attempt to optimize the database before closing connection (default true)
  -processes int
    	The number of concurrent processes to index data with (default 20)
  -properties
    	Index the 'properties' table
  -rtree
    	Index the 'rtree' table
  -search
    	Index the 'search' table (using SQLite FTS4 full-text indexer)
  -spatial-tables
    	If true then index the necessary tables for use with the whosonfirst/go-whosonfirst-spatial-sqlite package.
  -spelunker
    	Index the 'spelunker' table
  -spelunker-tables
    	If true then index the necessary tables for use with the whosonfirst/go-whosonfirst-spelunker packages
  -spr
    	Index the 'spr' table
  -strict-alt-files
    	Be strict when indexing alt geometries (default true)
  -supersedes
    	Index the 'supersedes' table
  -timings
    	Display timings during and after indexing
  -verbose
    	Enable verbose (debug) logging
```

For example:

```
$> ./bin/wof-sqlite-index \
	-database-uri 'sql://sqlite3?dsn=test.db' \
	-timings -spatial-tables \
	/usr/local/data/whosonfirst-data-admin-us/
	
2024/12/13 10:01:14 INFO Time to index table table=geojson count=98764 time=7.604906428s
2024/12/13 10:01:14 INFO Time to index table table=rtree count=98764 time=14.588749044s
2024/12/13 10:01:14 INFO Time to index table table=properties count=98764 time=6.109436136s
2024/12/13 10:01:14 INFO Time to index table table=spr count=98764 time=29.349679552s
...
2024/12/13 10:04:14 INFO Time to index table table=properties count=410829 time=35.907407165s
2024/12/13 10:04:14 INFO Time to index table table=spr count=410829 time=1m42.546600206s
2024/12/13 10:04:14 INFO Time to index all count=410829 time=4m0.01180175s
2024/12/13 10:05:07 time to index paths (1) 4m52.480048125s
```

## See also

* https://github.com/whosonfirst/go-whosonfirst-database/
* https://github.com/sfomuseum/go-database/
* https://github.com/mattn/go-sqlite3
* https://pkg.go.dev/modernc.org/sqlite