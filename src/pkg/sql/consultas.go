package consultas

const SelectAvaiableHours = "SELECT a.disponibles, a.totales, a.tipo, a.nombre FROM " +
	"(SELECT hora.*, a.nombre FROM hora INNER JOIN " +
	"(SELECT asignatura.id, asignatura.nombre FROM asignatura " +
	"INNER JOIN titulacion ON asignatura.idT=titulacion.id WHERE titulacion.nombre=?) a " +
	"ON a.id=hora.idasignatura) a " +
	"INNER JOIN " +
	"(SELECT * FROM hora WHERE idgrupo " +
	"IN (SELECT grupodocente.id FROM `grupodocente` " +
	"INNER JOIN curso ON grupodocente.idcurso=curso.id WHERE curso.numero=? AND grupodocente.numero=?)) b " +
	"ON a.id=b.id ORDER BY a.nombre"

const SelectIdHoraGroup = "SELECT id, idgrupo " +
	"FROM app_db.hora " +
	"WHERE hora.tipo=? AND hora.grupo = ? AND hora.semana = ? AND hora.idasignatura " +
	"IN (SELECT id FROM app_db.asignatura WHERE asignatura.nombre=?)"

const SelectIdAula = "SELECT id FROM app_db.aula WHERE aula.nombre=?"

const InsertEntradaHorario = "INSERT INTO app_db.entradahorario (inicio, fin, idhoras, idaula, idgrupo, ultModificacion, diaSemana) " +
	"VALUES (?, ?, ?, ?, ?, STR_TO_DATE(?,'%Y-%d-%m'), ?)"

const DeleteEntradaHorario = "DELETE FROM app_db.entradahorario WHERE inicio = ? AND fin = ? AND idhoras = ? AND idaula = ? AND idgrupo = ?"

const DeleteEntradas = "DELETE FROM app_db.entradahorario WHERE idhoras IN ( " +
	"SELECT hora.id FROM app_db.hora WHERE idasignatura IN ( " +
	"SELECT asignatura.id FROM app_db.asignatura WHERE idT IN ( " +
	"SELECT titulacion.id FROM app_db.titulacion WHERE nombre = ? ))) " +
	"AND idgrupo IN ( " +
	"SELECT grupodocente.id FROM app_db.grupodocente WHERE numero = ? " +
	"AND idcurso IN ( " +
	"SELECT curso.id FROM app_db.curso WHERE numero = ?))"

const SearchHours = "SELECT hora.totales, hora.disponibles FROM app_db.hora WHERE hora.id=?"

const UpdateHours = "UPDATE app_db.hora SET hora.disponibles = ? WHERE hora.id=?"

const SearchEntry = "SELECT * FROM entradahorario " +
	"WHERE entradahorario.inicio=? AND entradahorario.fin=? " +
	"AND entradahorario.idhoras IN (SELECT id FROM hora WHERE hora.tipo=? AND hora.semana=? AND hora.grupo=? " +
	"AND hora.idasignatura IN (SELECT id FROM asignatura WHERE asignatura.nombre=?))"

const SelectIdNameDegree = "SELECT * FROM `titulacion`"

const SelectIdNumberYear = "SELECT curso.id, curso.numero FROM `curso` WHERE curso.idT = ?"

const SelectNameGroup = "SELECT grupodocente.numero FROM `grupodocente` WHERE grupodocente.idcurso = ?"

const SelectEntries = "SELECT b.*, asignatura.nombre FROM ( " +
	"SELECT a.*, hora.tipo, hora.grupo, hora.semana, hora.idasignatura FROM ( " +
	"SELECT entradahorario.inicio, entradahorario.fin, entradahorario.idhoras,entradahorario.diaSemana, aula.nombre FROM app_db.entradahorario " +
	"INNER JOIN app_db.aula ON entradahorario.idaula = aula.id " +
	"WHERE idgrupo IN (SELECT id FROM app_db.grupodocente " +
	"		WHERE numero = ? AND idcurso IN (SELECT id FROM app_db.curso " +
	"			WHERE numero = ? AND idT IN (SELECT id FROM app_db.titulacion " +
	"				WHERE nombre = ?)))) a " +
	"INNER JOIN app_db.hora ON a.idhoras = hora.id ) b " +
	"INNER JOIN app_db.asignatura ON b.idasignatura = asignatura.id"

const CreateDegree = "INSERT INTO `titulacion` (`id`, `nombre`) VALUES (?, ?)"
const CreateSubject = "INSERT INTO `asignatura` (`id`, `codigo`, `nombre`, `idT`) VALUES (?, ?, ?, ?)"
const CreateYear = "INSERT INTO `curso` (`id`, `numero`, `idT`) VALUES (?, ?, ?)"
const CreateGroup = "INSERT INTO `grupodocente` (`id`, `numero`, `idcurso`) VALUES (?, ?, ?)"
const CreateHour = "INSERT INTO `hora` (`id`, `disponibles`, `totales`, `tipo`, `grupo`, `semana`, `idasignatura`, `idgrupo`) VALUES (NULL, ?, ?, ?, ?, ?, ?, ?)"