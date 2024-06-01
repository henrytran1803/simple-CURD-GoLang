CREATE TABLE `roles` (
                         `id` int(11) NOT NULL AUTO_INCREMENT,
                         `name` varchar(50) NOT NULL,
                         `status` int NOT NULL DEFAULT '1',
                         `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                         `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                         PRIMARY KEY (`id`)
);
CREATE TABLE `users` (
                         `id` int(11) NOT NULL AUTO_INCREMENT,
                         `role_id` int(11) NULL,
                         `name` varchar(50) NOT NULL,
                         `addr` varchar(255) NOT NULL,
                         `email` varchar(255) DEFAULT NULL,
                         `status` int NOT NULL DEFAULT '1',
                         `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                         `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                         PRIMARY KEY (`id`),
                         FOREIGN KEY (`role_id`) REFERENCES `roles`(`id`)
);
CREATE TABLE `account` (
                           `id` int(11) NOT NULL,
                           `username` varchar(50) NOT NULL,
                           `pass` varchar(50) NOT NULL,
                           `status` int NOT NULL DEFAULT '1',
                           `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                           `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                           PRIMARY KEY (`id`),
                           FOREIGN KEY (`id`) REFERENCES `users`(`id`)
);
