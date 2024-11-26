CREATE TYPE roleType AS ENUM ('admin', 'moderator', 'user');

ALTER TABLE users ADD COLUMN role roleType DEFAULT 'user';