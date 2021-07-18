CREATE TABLE `book` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `book` varchar(255) DEFAULT '' COMMENT 'Book name',
    `price` int(11) unsigned DEFAULT 0 COMMENT 'book price',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;