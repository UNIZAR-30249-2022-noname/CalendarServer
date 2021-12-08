INSERT INTO `titulacion` (`id`, `nombre`) VALUES ('1', 'Ing. Informatica');
INSERT INTO `titulacion` (`id`, `nombre`) VALUES ('2', 'Ing. Mecanica');

INSERT INTO `curso` (`id`, `numero`, `idT`) VALUES ('1', '1', '1');
INSERT INTO `curso` (`id`, `numero`, `idT`) VALUES ('2', '1', '2');
INSERT INTO `curso` (`id`, `numero`, `idT`) VALUES ('2', '2', '2');

INSERT INTO `asignatura` (`id`, `codigo`, `nombre`, `idT`) VALUES ('1', '1', 'Proyecto Software', '1');
INSERT INTO `asignatura` (`id`, `codigo`, `nombre`, `idT`) VALUES ('2', '2', 'Sistemas Operativos', '1');
INSERT INTO `asignatura` (`id`, `codigo`, `nombre`, `idT`) VALUES ('3', '3', 'Asignatura random de Mecanica', '2');

INSERT INTO `grupodocente` (`id`, `numero`, `idcurso`) VALUES ('1', '1', '1');
INSERT INTO `grupodocente` (`id`, `numero`, `idcurso`) VALUES ('2', '2', '1');
INSERT INTO `grupodocente` (`id`, `numero`, `idcurso`) VALUES ('3', '1', '2');

INSERT INTO `hora` (`id`, `disponibles`, `totales`, `tipo`, `grupo`, `semana`, `idasignatura`, `idgrupo`) VALUES ('1', '2900', '3000', '1', '', '', '1', '1');
INSERT INTO `hora` (`id`, `disponibles`, `totales`, `tipo`, `grupo`, `semana`, `idasignatura`, `idgrupo`) VALUES ('3', '1800', '2000', '2', 'mananas', 'a', '1', '1');
INSERT INTO `hora` (`id`, `disponibles`, `totales`, `tipo`, `grupo`, `semana`, `idasignatura`, `idgrupo`) VALUES ('4', '800', '1000', '3', 'niapar', '', '1', '1');
INSERT INTO `hora` (`id`, `disponibles`, `totales`, `tipo`, `grupo`, `semana`, `idasignatura`, `idgrupo`) VALUES ('2', '2500', '2500', '2', '', '', '2', '1');
INSERT INTO `hora` (`id`, `disponibles`, `totales`, `tipo`, `grupo`, `semana`, `idasignatura`, `idgrupo`) VALUES ('5', '2500', '2500', '1', '', '', '3', '2');

INSERT INTO `aula` (`id`, `numero`, `nombre`) VALUES ('1', '1', '1');
INSERT INTO `aula` (`id`, `numero`, `nombre`) VALUES ('2', '2', '2');
INSERT INTO `aula` (`id`, `numero`, `nombre`) VALUES ('3', '3', '3');

INSERT INTO `entradahorario` (`id`, `inicio`, `fin`, `diaSemana`, `idhoras`, `idaula`, `idgrupo`) VALUES ('1', '1200', '1300', '1', '1', '1', '1');
INSERT INTO `entradahorario` (`id`, `inicio`, `fin`, `diaSemana`, `idhoras`, `idaula`, `idgrupo`, `ultModificacion`) VALUES ('2', '1000', '1200', '2', '2', '2', '1', NULL);
INSERT INTO `entradahorario` (`id`, `inicio`, `fin`, `diaSemana`, `idhoras`, `idaula`, `idgrupo`, `ultModificacion`) VALUES ('3', '1500', '1700', '3', '3', '1', '1', NULL);