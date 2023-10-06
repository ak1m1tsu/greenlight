ALTER TABLE movies DROP COLUMN version;
ALTER TABLE movies ADD COLUMN version uuid NOT NULL DEFAULT uuid_generate_v4();
