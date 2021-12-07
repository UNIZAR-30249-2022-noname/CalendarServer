#!/bin/sh
docker compose up -d
docker exec -it horarios_db bash -c "mkdir /data"
docker cp src/pkg/sql/Create_Tables.sql horarios_db:/data/
docker cp src/pkg/sql/Drop_tables.sql horarios_db:/data/
docker cp src/pkg/sql/Populate_Tables.sql horarios_db:/data/
docker exec -it horarios_db bash -c "mysql -P 6033 -u user -puser app_db < /data/Drop_tables.sql"
docker exec -it horarios_db bash -c "mysql -P 6033 -u user -puser app_db < /data/Create_Tables.sql"
docker exec -it horarios_db bash -c "mysql -P 6033 -u user -puser app_db < /data/Populate_Tables.sql"