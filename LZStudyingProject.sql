-- Active: 1677501433605@@localhost@3306@LZStudyingProject
CREATE TABLE `Player`
(
	`id` BIGINT AUTO_INCREMENT PRIMARY KEY NOT NULL,
	`login_name`   VARCHAR(255) NOT NULL,
	`name`        varchar(255) NOT NULL,
	`icon`        MEDIUMBLOB NOT NULL,
	`sex`         VARCHAR(255) NOT NULL,
	`age`         int(11)  NOT NULL,
	`points` int(11) NOT null
);

CREATE TABLE `Commodity`
(
	`id` BIGINT AUTO_INCREMENT PRIMARY KEY NOT NULL,
	`intro` varchar(255) NOT NULL,
	`price`  DOUBLE NOT NULL,
	`limit` int(11)  NOT NULL,
	`cag` int(11)  NOT NULL
);

CREATE TABLE `Bill`
(
	`id` BIGINT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    `com_id` BIGINT  NOT NULL,
    `quantity` int(11)  NOT NULL,
    `opinion_id` BIGINT  NOT NULL,
	`order_id` BIGINT  NOT NULL
);

CREATE TABLE `Option`
(
	`id` BIGINT AUTO_INCREMENT PRIMARY KEY NOT NULL,
	`com_id` BIGINT  NOT NULL,
	`image` VARCHAR(255) NOT NULL,
	`intro` MEDIUMBLOB NOT NULL
);

CREATE TABLE `Order`
(
	`id` BIGINT AUTO_INCREMENT PRIMARY KEY NOT NULL,
	`shipping_address_id` BIGINT  NOT NULL,
	`delivery_address_id` BIGINT  NOT NULL,
	`create_time` DOUBLE NOT NULL,
	`payed_time` DOUBLE NOT NULL,
	`completed_time` DOUBLE NOT NULL,
	`state` int(11)  NOT NULL
);

CREATE TABLE `Address`
(
	`id` BIGINT AUTO_INCREMENT PRIMARY KEY NOT NULL,
	`name` varchar(255) NOT NULL,
	`sex` varchar(255) NOT NULL,
	`phone_number` VARCHAR(255) NOT NULL,
	`province` VARCHAR(255) NOT NULL,
	`city` VARCHAR(255) NOT NULL,
	`area` VARCHAR(255) NOT NULL,
	`detailed_address` VARCHAR(255) NOT NULL
);

CREATE TABLE `Cart`
(
	`player_id` BIGINT NOT NULL,
	`order_id` BIGINT NOT NULL
);

CREATE TABLE `Relationship`
(
	`player1_id` BIGINT NOT NULL,
	`player2_id` BIGINT NOT NULL,
	INDEX (`player1_id`),
	INDEX (`player2_id`)
);

 CREATE TABLE `MusicUser`
(
`music_id` VARCHAR(255) NOT NULL,
	`player_id` INT(11) NOT NULL,
	INDEX (`music_id`),
	INDEX (`player_id`)
);
CREATE TABLE `Music`
(
	`id` VARCHAR(255) PRIMARY KEY NOT NULL,
	`play_url` MEDIUMBLOB  NOT NULL,
	`type` VARCHAR(255)  NOT NULL,
	`recommend` BOOLEAN NOT NULL,
	`atime` DOUBLE NOT NULL,
	`author` VARCHAR(255) NOT NULL,
	`anime_info_id` VARCHAR(255)  NOT NULL
);


CREATE TABLE `AnimeInfo`
(
	`id` VARCHAR(255) PRIMARY KEY NOT NULL,
	`bg` VARCHAR(255)  NOT NULL,
	`year` INT(11)  NOT NULL,
	`month` INT(11) NOT NULL,
	`title` MEDIUMBLOB NOT NULL,
	`atime` DOUBLE NOT NULL,
	`desc` MEDIUMBLOB  NOT NULL,
    `logo` VARCHAR(255) not NULL 
);
