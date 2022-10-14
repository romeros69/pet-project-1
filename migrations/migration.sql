create table if not exists books (
    id serial primary key,
    tittle varchar(255) not null ,
    author varchar(255)
);

insert into books (tittle, author) values ('How to begin kaif', 'Romich');
insert into books (tittle, author) values ('Tanks', 'Alex');
insert into books (tittle, author) values ('Shedevr', 'Evgen');
insert into books (tittle, author) values ('lolololol', 'Alexey');
insert into books (tittle, author) values ('Checkof my brorher', 'Lopes Jenifer');
