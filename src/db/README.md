## Steps
1. docker pull postgres
2. docker run --rm -p 5432:5432 --network host --add-host 0.0.0.0:0.0.0.0 -e POSTGRES_PASSWORD=mysecretpassword postgres
3. docker container ls
4. docker exec -it f3fc12e65139 bash
5. psql -U postgres
6. create DATABASE numbers;
7. \c numbers;
8. CREATE TABLE numbers(
   uuid varchar(100),
   number integer
);