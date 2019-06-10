CREATE TABLE `Incidents` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `alertname` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `startsAT` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `address` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `Operations` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `surgeon_id` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `script` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `alert_id` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `Mapping` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `alert_type` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `script` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


