#!/bin/bash
echo "¿Seguro que quieres dropear las tablas?"

echo -n "y/n: "
read -r respuesta

y="y"

if [ $respuesta = $y ]; then
    echo "DROPEANDO..."
    mysql -h 127.0.0.1 -P 6033 -u user -puser app_db < Drop_tables.sql        
    exit 1
fi

echo "No hay dropeo (su respuesta: $respuesta)"

