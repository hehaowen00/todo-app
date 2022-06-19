CREATE TABLE IF NOT EXISTS `todos`.`users` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(64) NOT NULL,
    `password` TEXT NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `todos`.`lists` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `user_id` INT NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE (`user_id`, `name`)
    -- FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
);

CREATE TABLE IF NOT EXISTS `todos`.`todos` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `user_id` INT NOT NULL,
    `list_id` INT NOT NULL,
    `description` TEXT NOT NULL,
    `status` BOOLEAN NOT NULL DEFAULT FALSE,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
    FOREIGN KEY (`list_id`) REFERENCES `lists` (`id`)
);