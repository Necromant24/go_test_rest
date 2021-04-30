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