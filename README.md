# CalendarServer
Server for the project of the new calendar and schedule of Universidad de Zaragoza

# Docker

Para contruir la imegen desde la carpeta root ejecutar
```bash
docker build -t arejula27/calendarunizar:0.0.1 -f docker/Dockerfile .
```
Si ya esta  construirla o si se quiere descargar de internet usa
```bash
docker run -d -p 8080:8080 arejula27/calendarunizar:0.0.1 
```
Para probar su funcionamiento hacer `localhost:8080/ping`
