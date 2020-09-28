# ABS Challenge

To run

`git clone https://github.com/sunejax/abs.git`

cd into the newly created directory

`cd abs`

docker-compose should be able to handle the rest from here

`docker-compose up -d` 

Race condition is solved using an UDF on aerospike, to get that working register the udf with the running instance of aerospike (this could be added to the docker-compose file)

in aql use
`register module "allocate.lua"` 

