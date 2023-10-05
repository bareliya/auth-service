
DROP DATABASE IF EXISTS userdb;
CREATE DATABASE userdb
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.UTF-8'
    LC_CTYPE = 'en_US.UTF-8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;


--- this table will store  user credential
CREATE TABLE user_credentials (
     user_id SERIAL PRIMARY KEY,
     username VARCHAR(255) UNIQUE NOT NULL,
     first_name VARCHAR(255) NOT NULL,
     last_name VARCHAR(255) NOT NULL,
     hash_id VARCHAR(255) NOT NULL,
     access_token VARCHAR(255) NOT NULL
);

--- creating index on username column
CREATE INDEX idx_username ON user_credentials(username);

---this table will store admin credentials

CREATE TABLE admin_credentials (
     admin_id SERIAL PRIMARY KEY,
     admin_user VARCHAR(255) UNIQUE NOT NULL, -- Adding UNIQUE constraint here
     admin_password_hash VARCHAR(255) NOT NULL,
     is_super_admin BOOLEAN NOT NULL,
     access_token VARCHAR(255) NOT NULL,
     is_approved BOOLEAN NOT NULL
);
--- creating index on admin_user column
CREATE INDEX idx_admin_user ON admin_credentials (admin_user);

--- adding super user manually for password (superAdminPass)
INSERT INTO admin_credentials (admin_user, admin_password_hash, is_super_admin, access_token, is_approved)
VALUES ('superAdmin', 'ef4207f457cce807808b8caf64f1b82f38360fe39f941d4ab4a949aa09d935279d2ea0f678594d38510eadc942cc493868fa749cdba1fd9a7d5cfa220d60f924', true, 'D5Mvwn5rhMP624S!', true);
