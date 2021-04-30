

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