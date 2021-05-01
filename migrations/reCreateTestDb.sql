
DROP TABLE card_tables CASCADE;
DROP TABLE card_lists CASCADE;
DROP TABLE card_to_card CASCADE;
DROP TABLE card CASCADE;




CREATE TABLE IF NOT EXISTS Card_Tables (
	id SERIAL PRIMARY KEY,
  brief_description varchar(150),
	name varchar(45) NOT NULL UNIQUE
  );

  CREATE TABLE IF NOT EXISTS Card_Lists (
	id SERIAL PRIMARY KEY,
	name varchar(45) NOT NULL,
  table_id INTEGER NOT NULL,

    
        FOREIGN KEY (table_id)
        REFERENCES Card_Tables (id)
        ON DELETE CASCADE,

        UNIQUE (name, table_id)

  );

  CREATE TABLE IF NOT EXISTS Card (
	id SERIAL PRIMARY KEY,
	name varchar(45) NOT NULL,
	description TEXT,
  card_list_id INTEGER NOT NULL, 

    
        FOREIGN KEY (card_list_id)
        REFERENCES Card_Lists (id)
        ON DELETE CASCADE,

        UNIQUE(name, card_list_id)
  );


  CREATE TABLE IF NOT EXISTS card_to_card (
      key_id INTEGER NOT NULL,
      value_id INTEGER NOT NULL,

        FOREIGN KEY (key_id)
        REFERENCES Card (id)
        ON DELETE CASCADE,

        FOREIGN KEY (value_id)
        REFERENCES Card (id)
        ON DELETE CASCADE,

        UNIQUE (key_id, value_id)
  );







  

INSERT INTO Card_Tables (name) VALUES ('project1');
INSERT INTO Card_Tables (name) VALUES ('project2');



INSERT INTO Card_Lists ( name, table_id) VALUES ('TODO', 1);
INSERT INTO Card_Lists ( name, table_id) VALUES ('TODO', 2);
INSERT INTO Card_Lists ( name, table_id) VALUES ('DONE', 1);



INSERT INTO Card ( name, description, card_list_id) VALUES ('DO1', 'mnogo texta', 1);
INSERT INTO Card ( name, description, card_list_id) VALUES ('DO2', 'mnogo texta', 2);
INSERT INTO Card ( name, description, card_list_id) VALUES ('DO3', 'mnogo texta', 3);
INSERT INTO Card ( name, description, card_list_id) VALUES ('DO4', 'mnogo texta', 1);



INSERT INTO card_to_card (key_id, value_id) VALUES (1,2);
INSERT INTO card_to_card (key_id, value_id) VALUES (1,3);
INSERT INTO card_to_card (key_id, value_id) VALUES (3,2);
INSERT INTO card_to_card (key_id, value_id) VALUES (3,1);