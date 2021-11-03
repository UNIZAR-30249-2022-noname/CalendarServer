CREATE TABLE app_db.asignatura (
    id int PRIMARY KEY AUTO_INCREMENT,
    codigo int NOT NULL,
    nombre varchar(50)
);

CREATE TABLE app_db.titulacion (
    id int PRIMARY KEY AUTO_INCREMENT,
    nombre varchar(50)
);

CREATE TABLE app_db.perteneceA (
    idA int not null,
    idT int not null,
    FOREIGN KEY (idA) REFERENCES app_db.asignatura (id),
    FOREIGN KEY (idT) REFERENCES app_db.titulacion (id)
);

CREATE TABLE app_db.curso (
    id int PRIMARY KEY AUTO_INCREMENT,
    numero int not null
);

CREATE TABLE app_db.grupodocente (
    id int PRIMARY KEY AUTO_INCREMENT,
    numero int not null
    idcurso int not null,
    FOREIGN KEY (idcurso) REFERENCES app_db.curso (id)
);

CREATE TABLE app_db.aula (
    id int PRIMARY KEY AUTO_INCREMENT,
    numero int not null
);

CREATE TABLE app_db.hora (
    id int PRIMARY KEY AUTO_INCREMENT,
    disponibles int not null,
    totales int not null,
    tipo int not null,
    idasignatura int not null,
    idgrupo int not null,
    FOREIGN KEY (idasignatura) REFERENCES app_db.asignatura (id),
    FOREIGN KEY (idgrupo) REFERENCES app_db.grupodocente (id)
);

CREATE TABLE app_db.entradahorario (
    id int PRIMARY KEY AUTO_INCREMENT,
    inicio int not null,
    fin int not null,
    idhoras int not null,
    idaula int,
    idgrupo int,
    FOREIGN KEY (idhoras) REFERENCES app_db.hora (id),
    FOREIGN KEY (idaula) REFERENCES app_db.aula (id),
    FOREIGN KEY (idgrupo) REFERENCES app_db.grupodocente (id),
    ultModificacion date
);
