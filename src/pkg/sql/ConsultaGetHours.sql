SELECT a.disponibles, a.totales, a.tipo, a.nombre FROM
(SELECT hora.*, a.nombre FROM hora INNER JOIN 
	(SELECT * FROM asignatura 
     WHERE id IN 
     	(SELECT perteneceA.idA FROM perteneceA 
         INNER JOIN titulacion ON titulacion.id=perteneceA.idT WHERE titulacion.nombre="uwu")) a ON a.id=hora.idasignatura) a
INNER JOIN
(SELECT * FROM hora WHERE id 
 	IN (SELECT grupodocente.id FROM `grupodocente` 
        INNER JOIN curso ON grupodocente.idcurso=curso.id WHERE curso.numero=1 AND grupodocente.numero=0)) b
ON a.id=b.id;


SELECT b.*, asignatura.nombre FROM (
SELECT a.*, hora.tipo, hora.grupo, hora.semana, hora.diaSemana, hora.idasignatura FROM (
SELECT entradahorario.inicio, entradahorario.fin, entradahorario.idhoras, aula.nombre FROM app_db.entradahorario 
INNER JOIN app_db.aula ON entradahorario.idaula = aula.id
WHERE idgrupo IN (SELECT id FROM app_db.grupodocente 
        WHERE numero = 1 AND idcurso IN (SELECT id FROM app_db.curso 
            WHERE numero = 1 AND idT IN (SELECT id FROM app_db.titulacion
                WHERE nombre = "Ing. Informatica")))) a
INNER JOIN app_db.hora ON a.idhoras = hora.id ) b
INNER JOIN app_db.asignatura ON b.idasignatura = asignatura.id
;