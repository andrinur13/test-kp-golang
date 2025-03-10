-- -------------------------------------------------------------
-- TablePlus 6.3.2(586)
--
-- https://tableplus.com/
--
-- Database: kp-test
-- Generation Time: 2025-03-10 12:26:27.0270
-- -------------------------------------------------------------


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


CREATE TABLE `products` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `amount_price` bigint unsigned DEFAULT NULL,
  `amount_ship` bigint unsigned DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `transactions` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint unsigned NOT NULL,
  `product_id` bigint unsigned NOT NULL,
  `status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `amount_otr` bigint unsigned DEFAULT NULL,
  `amount_fee` bigint unsigned DEFAULT NULL,
  `amount_installment` bigint unsigned DEFAULT NULL,
  `amount_interest` bigint unsigned DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `product_id` (`product_id`),
  CONSTRAINT `transactions_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  CONSTRAINT `transactions_ibfk_2` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `user_tenors` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint unsigned NOT NULL,
  `tenor_in_month` bigint unsigned NOT NULL,
  `amount` bigint unsigned DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `user_tenors_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `nik` varchar(255) DEFAULT NULL,
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `full_name` varchar(255) NOT NULL,
  `legal_name` varchar(255) NOT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `born_city` varchar(255) NOT NULL,
  `born_date` date NOT NULL,
  `income` bigint unsigned NOT NULL DEFAULT '0',
  `identity_photo_path` varchar(255) NOT NULL,
  `selfie_photo_path` varchar(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `products` (`id`, `name`, `amount_price`, `amount_ship`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'Mobil Balap', 100000000, 1000000, '2025-03-10 10:17:48', '2025-03-10 10:17:48', NULL);

INSERT INTO `user_tenors` (`id`, `user_id`, `tenor_in_month`, `amount`, `created_at`, `updated_at`, `deleted_at`) VALUES
(2, 14, 1, 100000, '2025-03-10 12:02:16', '2025-03-10 12:02:16', NULL),
(3, 14, 2, 200000, '2025-03-10 12:02:16', '2025-03-10 12:02:16', NULL),
(4, 14, 3, 500000, '2025-03-10 12:02:16', '2025-03-10 12:02:16', NULL),
(5, 14, 4, 700000, '2025-03-10 12:02:16', '2025-03-10 12:02:16', NULL),
(6, 15, 1, 1000000, '2025-03-10 12:02:16', '2025-03-10 12:02:16', NULL),
(7, 15, 2, 1200000, '2025-03-10 12:02:16', '2025-03-10 12:02:16', NULL),
(8, 15, 3, 1500000, '2025-03-10 12:02:16', '2025-03-10 12:02:16', NULL),
(9, 15, 4, 2000000, '2025-03-10 12:02:16', '2025-03-10 12:02:16', NULL);

INSERT INTO `users` (`id`, `nik`, `email`, `full_name`, `legal_name`, `password`, `born_city`, `born_date`, `income`, `identity_photo_path`, `selfie_photo_path`, `created_at`, `updated_at`, `deleted_at`) VALUES
(14, '19371848484242', 'budi@gmail.com', 'Budi', 'Budi', '$2a$10$BIkx89klFICt0VOZ0BpgZ.9InDYIwB4FSx4JqLd9f8idUOSKoigeO', 'Yogyakarta', '2000-01-01', 5000000, 'identity.jpg', 'selfie.jpg', '2025-03-09 16:24:40', '2025-03-09 16:24:40', NULL),
(15, '19371848484241', 'annisa@gmail.com', 'Annisa', 'Annisa', '$2a$10$BIkx89klFICt0VOZ0BpgZ.9InDYIwB4FSx4JqLd9f8idUOSKoigeO', 'Yogyakarta', '2000-01-01', 5000000, 'identity.jpg', 'selfie.jpg', '2025-03-09 16:24:40', '2025-03-09 16:24:40', NULL);



/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;