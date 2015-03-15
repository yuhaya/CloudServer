# ************************************************************
# Sequel Pro SQL dump
# Version 4096
#
# http://www.sequelpro.com/
# http://code.google.com/p/sequel-pro/
#
# Host: 127.0.0.1 (MySQL 5.6.21)
# Database: cloud
# Generation Time: 2015-03-15 02:40:35 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table ittr_admins
# ------------------------------------------------------------

DROP TABLE IF EXISTS `ittr_admins`;

CREATE TABLE `ittr_admins` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `guid` varchar(50) NOT NULL,
  `username` varchar(50) NOT NULL,
  `password` varchar(50) NOT NULL,
  `realname` varchar(50) DEFAULT NULL,
  `school_guid` varchar(50) NOT NULL,
  `create_time` datetime NOT NULL,
  `super` tinyint(4) NOT NULL,
  `enabled` tinyint(4) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `guid` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table ittr_attendances
# ------------------------------------------------------------

DROP TABLE IF EXISTS `ittr_attendances`;

CREATE TABLE `ittr_attendances` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `card` varchar(50) NOT NULL,
  `time` datetime NOT NULL,
  `type` tinyint(4) NOT NULL,
  `school_guid` varchar(50) NOT NULL,
  `auto` tinyint(4) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table ittr_card_receiver
# ------------------------------------------------------------

DROP TABLE IF EXISTS `ittr_card_receiver`;

CREATE TABLE `ittr_card_receiver` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `card` varchar(50) NOT NULL,
  `guid` varchar(50) NOT NULL,
  `type` tinyint(4) NOT NULL,
  `school_guid` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table ittr_cards
# ------------------------------------------------------------

DROP TABLE IF EXISTS `ittr_cards`;

CREATE TABLE `ittr_cards` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `guid` varchar(50) NOT NULL,
  `kind` tinyint(4) NOT NULL,
  `school_guid` varchar(50) NOT NULL,
  `family_guid` varchar(50) NOT NULL,
  `enabled` tinyint(4) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `guid` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table ittr_class_teacher
# ------------------------------------------------------------

DROP TABLE IF EXISTS `ittr_class_teacher`;

CREATE TABLE `ittr_class_teacher` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `class_guid` varchar(50) NOT NULL,
  `teacher_guid` varchar(50) NOT NULL,
  `adviser` tinyint(4) NOT NULL,
  `school_guid` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table ittr_classes
# ------------------------------------------------------------

DROP TABLE IF EXISTS `ittr_classes`;

CREATE TABLE `ittr_classes` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `guid` varchar(50) NOT NULL,
  `name` varchar(20) NOT NULL,
  `school_guid` varchar(50) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `guid` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table ittr_devices
# ------------------------------------------------------------

DROP TABLE IF EXISTS `ittr_devices`;

CREATE TABLE `ittr_devices` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `guid` varchar(50) NOT NULL,
  `device` varchar(50) NOT NULL,
  `kind` tinyint(4) NOT NULL,
  `vmp` varchar(10) NOT NULL,
  `school_guid` varchar(50) NOT NULL,
  `group` tinyint(4) NOT NULL,
  `description` varchar(255) NOT NULL,
  `status` tinyint(4) NOT NULL,
  `enabled` tinyint(4) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `guid` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table ittr_families
# ------------------------------------------------------------

DROP TABLE IF EXISTS `ittr_families`;

CREATE TABLE `ittr_families` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `guid` varchar(50) NOT NULL,
  `first_guid` varchar(50) NOT NULL,
  `school_guid` varchar(50) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `guid` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table ittr_family_member
# ------------------------------------------------------------

DROP TABLE IF EXISTS `ittr_family_member`;

CREATE TABLE `ittr_family_member` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `family_guid` varchar(50) NOT NULL,
  `member_guid` varchar(50) NOT NULL,
  `type` tinyint(4) NOT NULL,
  `school_guid` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table ittr_family_relation
# ------------------------------------------------------------

DROP TABLE IF EXISTS `ittr_family_relation`;

CREATE TABLE `ittr_family_relation` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_guid` varchar(50) NOT NULL,
  `stu_guid` varchar(50) NOT NULL,
  `relation` varchar(50) NOT NULL,
  `school_guid` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table ittr_grade_class
# ------------------------------------------------------------

DROP TABLE IF EXISTS `ittr_grade_class`;

CREATE TABLE `ittr_grade_class` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `grade_guid` varchar(50) NOT NULL,
  `class_guid` varchar(50) NOT NULL,
  `school_guid` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table ittr_grades
# ------------------------------------------------------------

DROP TABLE IF EXISTS `ittr_grades`;

CREATE TABLE `ittr_grades` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `guid` varchar(50) NOT NULL,
  `name` varchar(50) NOT NULL,
  `rating` bigint(20) unsigned NOT NULL,
  `school_guid` varchar(50) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `guid` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table ittr_member_card
# ------------------------------------------------------------

DROP TABLE IF EXISTS `ittr_member_card`;

CREATE TABLE `ittr_member_card` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `card` varchar(50) NOT NULL,
  `guid` varchar(50) NOT NULL,
  `type` tinyint(4) NOT NULL,
  `school_guid` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table ittr_schools
# ------------------------------------------------------------

DROP TABLE IF EXISTS `ittr_schools`;

CREATE TABLE `ittr_schools` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `guid` varchar(50) NOT NULL,
  `name` varchar(30) NOT NULL,
  `spell` varchar(50) NOT NULL,
  `province` varchar(20) NOT NULL,
  `city` varchar(20) NOT NULL,
  `county` varchar(20) NOT NULL,
  `address` varchar(80) DEFAULT NULL,
  `update_time` datetime NOT NULL,
  `door_num` smallint(6) NOT NULL,
  `enabled` tinyint(3) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `guid` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table ittr_students
# ------------------------------------------------------------

DROP TABLE IF EXISTS `ittr_students`;

CREATE TABLE `ittr_students` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `guid` varchar(50) NOT NULL,
  `sid` varchar(50) DEFAULT NULL,
  `realname` varchar(10) NOT NULL,
  `spell` varchar(10) NOT NULL,
  `gender` tinyint(4) NOT NULL,
  `picture` varchar(255) DEFAULT NULL,
  `birthday` date NOT NULL,
  `school_guid` varchar(50) NOT NULL,
  `grade_guid` varchar(50) NOT NULL,
  `classname` varchar(20) DEFAULT NULL,
  `enrol_time` date NOT NULL,
  `create_time` datetime NOT NULL,
  `class_guid` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `guid` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table ittr_system
# ------------------------------------------------------------

DROP TABLE IF EXISTS `ittr_system`;

CREATE TABLE `ittr_system` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `key` varchar(50) NOT NULL,
  `value` varchar(250) NOT NULL,
  `school_guid` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table ittr_teachers
# ------------------------------------------------------------

DROP TABLE IF EXISTS `ittr_teachers`;

CREATE TABLE `ittr_teachers` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `guid` varchar(50) NOT NULL,
  `name` varchar(50) NOT NULL,
  `phone` varchar(20) NOT NULL,
  `gender` tinyint(4) NOT NULL,
  `school_guid` varchar(50) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `guid` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table ittr_users
# ------------------------------------------------------------

DROP TABLE IF EXISTS `ittr_users`;

CREATE TABLE `ittr_users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `guid` varchar(50) NOT NULL,
  `phone` varchar(20) DEFAULT NULL,
  `realname` varchar(10) DEFAULT NULL,
  `password` varchar(50) NOT NULL,
  `spell` varchar(10) DEFAULT NULL,
  `gender` tinyint(4) NOT NULL,
  `id_card` varchar(20) DEFAULT NULL,
  `picture` varchar(100) DEFAULT NULL,
  `school_guid` varchar(50) NOT NULL,
  `create_time` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `guid` (`guid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
