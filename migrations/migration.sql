DROP TABLE IF EXISTS pokemon;
DROP TABLE IF EXISTS battle;

CREATE TABLE IF NOT EXISTS pokemon(
   pokemon_id SERIAL PRIMARY KEY,
   name varchar(255) NOT NULL,
   battle_id int NOT NULL,
   scores int NOT NULL
);

CREATE TABLE IF NOT EXISTS battle(
   battle_id SERIAL PRIMARY KEY,
   winner varchar(255) NOT NULL,
   start_time timestamp NOT NULL,
   end_time timestamp NOT NULL
);