--Tenemos de Entry el inicio y fin, ultModificacion lo sacaremos con golang
-- idaula lo sacaremos de room e idhoras e idgrupo lo sacaremos de Avaiablehours

--Sacar que hours es -> idhoras idgrupo
SELECT id, idgrupo 
    FROM app_db.hora 
        WHERE hora.tipo=0 AND hora.disponibles=20 AND hora.totales=21 AND hora.idasignatura 
            IN (SELECT id FROM app_db.asignatura WHERE asignatura.nombre='si')

--Sacar id aula
SELECT id FROM app_db.aula WHERE aula.nombre='1'

--El insert
INSERT INTO app_db.entradahorario (inicio, fin, idhoras, idaula, idgrupo, ultModificacion)
VALUES (1, 2, 1, 1, 1, '2008-7-04');

--Actualizar horas disponibles