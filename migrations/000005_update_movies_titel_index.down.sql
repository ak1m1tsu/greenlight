DROP INDEX movies_title_idx CASCADE;
CREATE INDEX movies_title_idx ON movies USING GIN (to_tsvector('simple', title));
