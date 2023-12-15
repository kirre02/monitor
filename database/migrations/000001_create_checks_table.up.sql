CREATE TABLE checks (
    id BIGSERIAL PRIMARY KEY,
    site_id BIGINT NOT NULL,
    up BOOLEAN NOT NULL,
    checked_at TIMESTAMP WITH TIME ZONE NOT NULL,
);
