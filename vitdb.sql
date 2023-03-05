/*
SQLyog Ultimate v8.61 
MySQL - 8.0.30 : Database - vit
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`vit` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `vit`;

/*Table structure for table `admin` */

DROP TABLE IF EXISTS `admin`;

CREATE TABLE `admin` (
  `id` varchar(40) NOT NULL,
  `name` varchar(40) DEFAULT NULL,
  `email` varchar(40) DEFAULT NULL,
  `hashPassword` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `role` varchar(40) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `admin` */

insert  into `admin`(`id`,`name`,`email`,`hashPassword`,`role`) values ('257d8220-b334-4ac1-b1c3-3559f5300ef7','haider','haider@gmail.com','$2a$14$xZICMEdXeyWYjSVtwtoSSuOsCFVvNJNB0BsTmXoF4bS8HcBL7zdsi','admin');

/*Table structure for table `course` */

DROP TABLE IF EXISTS `course`;

CREATE TABLE `course` (
  `id` varchar(40) NOT NULL,
  `name` varchar(40) DEFAULT NULL,
  `faculty_id` varchar(40) DEFAULT NULL,
  `course_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `slot_id` varchar(40) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `course` */

insert  into `course`(`id`,`name`,`faculty_id`,`course_type`,`slot_id`) values ('test1212','java','[\"123\"]','THOERY','[\"A1,L2,A2\"]'),('test1213','golang','[\"123,33,321\"]','THEORY','[\"A1,L2,A2\"]');

/*Table structure for table `faculty` */

DROP TABLE IF EXISTS `faculty`;

CREATE TABLE `faculty` (
  `id` varchar(40) NOT NULL,
  `name` varchar(40) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `faculty` */

insert  into `faculty`(`id`,`name`) values ('110','akash'),('111','shahidkapoor'),('123','testFacuulty'),('321','sunil'),('33','anus'),('786','testSs');

/*Table structure for table `register` */

DROP TABLE IF EXISTS `register`;

CREATE TABLE `register` (
  `id` varchar(50) NOT NULL,
  `course_id` varchar(50) DEFAULT NULL,
  `faculty_id` varchar(40) DEFAULT NULL,
  `slot_id` varchar(100) DEFAULT NULL,
  `student_id` varchar(40) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `register` */

insert  into `register`(`id`,`course_id`,`faculty_id`,`slot_id`,`student_id`) values ('0ac01e7d-31cc-42c6-a422-d76949d7c2df','test1213','123','[\"A1\",\"L2\",\"A2\"]','2OBCE608');

/*Table structure for table `slot` */

DROP TABLE IF EXISTS `slot`;

CREATE TABLE `slot` (
  `id` varchar(40) NOT NULL,
  `day` varchar(20) DEFAULT NULL,
  `start` timestamp NULL DEFAULT NULL,
  `end` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `slot` */

insert  into `slot`(`id`,`day`,`start`,`end`) values ('A1','MON','2019-08-24 14:15:22','2019-08-24 14:15:22'),('A2','TUE','2019-08-24 14:15:22','2019-08-24 14:15:22'),('L2','WED','2019-08-24 14:15:22','2019-08-24 14:15:22');

/*Table structure for table `student` */

DROP TABLE IF EXISTS `student`;

CREATE TABLE `student` (
  `id` varchar(40) NOT NULL,
  `name` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `hashPassword` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `student` */

insert  into `student`(`id`,`name`,`hashPassword`) values ('2OBCE603','shivansh','abc321'),('2OBCE604','chotabheem','$2a$14$iDnxngOUT./cQ1UE8lVuL.has16owg8WE'),('2OBCE605','raju','$2a$14$8hqu/ELgH74HsOvIjuOt/u/9ssUjw2WHCLP1cLNhnOvy11xmhgcYy'),('2OBCE606','saqibtest1','$2a$14$/KgFfw7W7SkDeE2ki2kaXu7mKW6tmtUNNIu0tdj6n5xPEBsl6bTfq'),('2OBCE607','chintu','$2a$14$ZTvHT6xhsDrLdWKuFL6eceKeFj.UsD5AsYSr06qV5001IcXP0TsQu'),('2OBCE608','monu','$2a$14$psgneKSSPrUb/MCXTxJYq.pXjfD2qTY7/hZlUZsz7KWAfiCD65IR.'),('bsse1210','haider','test1210');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
