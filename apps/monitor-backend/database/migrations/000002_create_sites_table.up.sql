CREATE TABLE sites (
    id BIGSERIAL PRIMARY KEY,
    url TEXT NOT NULL,
    name TEXT
);

ALTER TABLE sites ADD CONSTRAINT url_uniq UNIQUE (url);
