CREATE OR REPLACE FUNCTION trigger_modified_date_timestamp()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.modified_date = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION trigger_created_date_timestamp()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.created_date = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TABLE users
(
    id            serial PRIMARY KEY,
    created_date  TIMESTAMP DEFAULT NOW(),
    modified_date TIMESTAMP DEFAULT NOW(),
    deleted_date  TIMESTAMP NULL,
    nickname      TEXT,
    PRIMARY KEY (id)
);

CREATE TRIGGER set_modified_date
    BEFORE UPDATE
    ON users
    FOR EACH ROW
EXECUTE PROCEDURE trigger_modified_date_timestamp();

CREATE TRIGGER set_created_date
    BEFORE INSERT
    ON users
    FOR EACH ROW
EXECUTE PROCEDURE trigger_created_date_timestamp();


CREATE TABLE follows
(
    id          serial PRIMARY KEY,
    follower_id INT,
    followee_id INT,
    FOREIGN KEY (follower_id)
        REFERENCES users (id),
    FOREIGN KEY (followee_id)
        REFERENCES users (id)
);

CREATE TRIGGER set_modified_date
    BEFORE UPDATE
    ON follows
    FOR EACH ROW
EXECUTE PROCEDURE trigger_modified_date_timestamp();
