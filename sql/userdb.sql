-- Database: userdb

-- DROP DATABASE IF EXISTS userdb;

CREATE DATABASE userdb
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.UTF-8'
    LC_CTYPE = 'en_US.UTF-8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;

CREATE TABLE user_credentials (
     user_id SERIAL PRIMARY KEY,
     username VARCHAR(255) UNIQUE NOT NULL,
     first_name VARCHAR(255) NOT NULL,
     last_name VARCHAR(255) NOT NULL,
     hash_id VARCHAR(255) NOT NULL,
     access_token VARCHAR(255) NOT NULL
);

CREATE INDEX idx_username ON user_credentials(username);
