SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

CREATE DATABASE `car_rental` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;
USE `car_rental`;

DROP TABLE IF EXISTS `bookings`;
CREATE TABLE `bookings` (
                            `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                            `created_at` datetime(3) DEFAULT NULL,
                            `updated_at` datetime(3) DEFAULT NULL,
                            `deleted_at` datetime(3) DEFAULT NULL,
                            `uuid` varchar(191) DEFAULT NULL,
                            `access_token` longtext DEFAULT NULL,
                            `date_from` datetime(3) DEFAULT NULL,
                            `date_to` datetime(3) DEFAULT NULL,
                            `car_id` bigint(20) unsigned DEFAULT NULL,
                            PRIMARY KEY (`id`),
                            KEY `idx_bookings_deleted_at` (`deleted_at`),
                            KEY `idx_bookings_uuid` (`uuid`),
                            KEY `fk_bookings_car` (`car_id`),
                            CONSTRAINT `fk_bookings_car` FOREIGN KEY (`car_id`) REFERENCES `cars` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `cars`;
CREATE TABLE `cars` (
                        `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                        `created_at` datetime(3) DEFAULT NULL,
                        `updated_at` datetime(3) DEFAULT NULL,
                        `deleted_at` datetime(3) DEFAULT NULL,
                        `type` longtext DEFAULT NULL,
                        `uuid` longtext DEFAULT NULL,
                        PRIMARY KEY (`id`),
                        KEY `idx_cars_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;