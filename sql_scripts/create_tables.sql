CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE ads
(
    id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name        TEXT      NOT NULL,
    description TEXT      NOT NULL,
    price       INTEGER   NOT NULL,
    photos_urls TEXT[]    NOT NULL,
    created_at  TIMESTAMP NOT NULL
);
