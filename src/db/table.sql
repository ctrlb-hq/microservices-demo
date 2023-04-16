create DATABASE numbers;

\c numbers;

CREATE TABLE numbers(
   uuid varchar(100),
   number integer,
   deleted_at DATE,
   id integer
);