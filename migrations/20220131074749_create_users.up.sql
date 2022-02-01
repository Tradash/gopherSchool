CREATE TABLE `users` (
                         `id` INT NOT NULL AUTO_INCREMENT,
                         `email` VARCHAR(50) NOT NULL,
                         `encrypted_password` VARCHAR(100) NOT NULL,
                         PRIMARY KEY (`id`),
                         UNIQUE INDEX `email` (`email`)
)
    COLLATE='utf8_general_ci'
    ENGINE=InnoDB
;
