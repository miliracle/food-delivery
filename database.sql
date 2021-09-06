DROP TABLE IF EXISTS `restaurants`;
CREATE TABLE `restaurants` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `owner_id` int(11) DEFAULT NULL,
  `name` varchar(50) NOT NULL,
  `addr` varchar(255) NOT NULL,
  `city_id` int(11) DEFAULT NULL,
  `lat` double DEFAULT NULL,
  `lng` double DEFAULT NULL,
  `cover` json DEFAULT NULL,
  `logo` json DEFAULT NULL,
  `shipping_fee_per_km` double DEFAULT '0',
  `status` int(11) NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `owner_id` (`owner_id`) USING BTREE,
  KEY `city_id` (`city_id`) USING BTREE,
  KEY `status` (`status`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `images`;
CREATE TABLE `images` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `url` text,
  `width` int(11) DEFAULT NULL,
  `height` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;