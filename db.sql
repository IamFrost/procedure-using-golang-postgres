DROP TABLE IF EXISTS tbl;
DROP PROCEDURE IF EXISTS drop_then_insert_data();
DROP PROCEDURE IF EXISTS insert_data();

CREATE TABLE tbl
(
	powervalue int
);

CREATE PROCEDURE drop_then_insert_data(a integer, b integer)
LANGUAGE SQL
AS $$
DROP TABLE IF EXISTS tbl;
CREATE TABLE tbl
(
    power int
);
INSERT INTO tbl VALUES (a);
INSERT INTO tbl VALUES (b);
$$;

CALL drop_then_insert_data(1, 2);

CREATE PROCEDURE insert_data(a integer, b integer)
LANGUAGE SQL
AS $$
INSERT INTO tbl VALUES (a);
INSERT INTO tbl VALUES (b);
$$;

CALL insert_data(1, 2);