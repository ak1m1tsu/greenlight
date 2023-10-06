ALTER TABLE movies DROP COLUMN version;
ALTER TABLE movies ADD COLUMN version uuid DEFAULT uuid_generate_v4();
