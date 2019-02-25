drop table if exists icons;
drop table if exists authors;

CREATE TABLE authors (
  author_id serial primary key,
  first_name varchar(16) not null,
  last_name varchar(16) not null,
  phone_number varchar(16) not null,
  age int,
  description text
);

CREATE TABLE icons (
  icon_id serial primary key,
  author_id int references authors,
  description text,
  url text
);

insert into authors (first_name, last_name, phone_number, age, description) values
('ignacy', 'radlinski', '+48698091912', 23, 'author 1'),
('agnieszka', 'miszkurka', '+48607506326', 22, 'author 2'),
('natalia', 'radlinska', '+48321321123', 29, 'author 3 - with no icons');

insert into icons (author_id, description, url) values
(1, 'icon 1', 'http://localhost:8080/whatever1'),
(2, 'icon 2', 'http://localhost:8080/whatever2'),
(2, 'icon 3', 'http://localhost:8080/whatever3'),
(1, 'icon 4', 'http://localhost:8080/whatever4'),
(default, 'icon 5 - icon without an author', 'http://localhost:8080/whatever5');

select * from authors;
select * from icons;
