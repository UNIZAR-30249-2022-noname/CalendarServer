package consultas

const Titulacion1 = "INSERT INTO `titulacion` (`id`, `nombre`) VALUES ('1', 'Ing. Informatica')"
const Titulacion2 = "INSERT INTO `titulacion` (`id`, `nombre`) VALUES ('2', 'Ing. Mecanica')"

const Curso1 = "INSERT INTO `curso` (`id`, `numero`) VALUES ('1', '1')"
const Curso2 = "INSERT INTO `curso` (`id`, `numero`) VALUES ('2', '2')"

const Asignatura1 = "INSERT INTO `asignatura` (`id`, `codigo`, `nombre`, `idT`) VALUES ('1', '1', 'Proyecto Software', '1')"
const Asignatura2 = "INSERT INTO `asignatura` (`id`, `codigo`, `nombre`, `idT`) VALUES ('2', '2', 'Sistemas Operativos', '1')"

const Grupodocente1 = "INSERT INTO `grupodocente` (`id`, `numero`, `idcurso`) VALUES ('1', '1', '1')"
const Grupodocente2 = "INSERT INTO `grupodocente` (`id`, `numero`, `idcurso`) VALUES ('2', '2', '1')"

const Hora1 = "INSERT INTO `hora` (`id`, `disponibles`, `totales`, `tipo`, `grupo`, `semana`, `idasignatura`, `idgrupo`) VALUES ('1', '30', '30', '1', '', '', '1', '1')"
const Hora2 = "INSERT INTO `hora` (`id`, `disponibles`, `totales`, `tipo`, `grupo`, `semana`, `idasignatura`, `idgrupo`) VALUES ('2', '25', '25', '2', '', '', '2', '1')"

const Aula1 = "INSERT INTO `aula` (`id`, `numero`, `nombre`) VALUES ('1', '1', '1')"
const TruncHora = "DELETE from app_db.hora WHERE id > 0;"
const TruncGrupo = "DELETE from app_db.grupodocente WHERE id > 0;"
const TruncAsignatura = "DELETE from app_db.asignatura WHERE id > 0;"
const TruncCurso = "DELETE from app_db.curso WHERE id > 0;"
const TruncTitulacion = "DELETE from app_db.titulacion WHERE id > 0;"
const TruncAula = "DELETE from app_db.aula WHERE id > 0;"
const TruncEntry = "DELETE from app_db.entradahorario WHERE id > 0;"
