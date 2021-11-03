#!/bin/bash
echo "El pwd es user"
mysql -h 127.0.0.1 -P 6033 -u user -p app_db < Create_Tables.sql