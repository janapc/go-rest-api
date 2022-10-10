create table courses(
    id serial primary key,
    name varchar,
    description varchar,
    level varchar
);

INSERT INTO courses(name, description, level) VALUES
('Iniciando com Go', 'Um curso introdutório sobre a linguagem Golang criada pelo Google e que vem ganhando cada vez mais espaço nas empresas e entre os desenvolvedores', 'basico'),
('Criando uma api rest com Go', 'Um curso aonde você irá aprender a implementa um api rest utilizando vários recursos que a linguagem possui', 'intermediario');