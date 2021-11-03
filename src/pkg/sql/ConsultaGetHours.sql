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
