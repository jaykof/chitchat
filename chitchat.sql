CREATE DATABASE IF NOT EXISTS `chitchat` DEFAULT CHARSET utf8 COLLATE utf8_general_ci;

USE `chitchat`;

CREATE TABLE IF NOT EXISTS `users` (
    `id` SERIAL PRIMARY KEY,
    `uuid` VARCHAR(64) NOT NULL UNIQUE,
    `name` VARCHAR(255),
    `email` VARCHAR(255) NOT NULL UNIQUE,
    `password` VARCHAR(255) NOT NULL,
    `created_at` TIMESTAMP NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `sessions` (
    `id` SERIAL PRIMARY KEY,
    `uuid` VARCHAR(64) NOT NULL UNIQUE,
    `email` VARCHAR(255),
    `user_id` INTEGER REFERENCES users(id),
    `create_at` TIMESTAMP NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `threads` (
     `id` SERIAL PRIMARY KEY,
     `uuid` VARCHAR(64) NOT NULL UNIQUE,
     `topic` TEXT,
     `user_id` INTEGER REFERENCES users(id),
     `created_at` TIMESTAMP NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `posts` (
   `id` SERIAL PRIMARY KEY,
   `uuid` VARCHAR(64) NOT NULL UNIQUE,
   `body` text,
   `user_id` INTEGER REFERENCES users(id),
   `thread_id` INTEGER REFERENCES threads(id),
   `created_at` TIMESTAMP NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;