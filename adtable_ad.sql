CREATE DATABASE  IF NOT EXISTS `adtable` /*!40100 DEFAULT CHARACTER SET utf8 */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `adtable`;
-- MySQL dump 10.13  Distrib 8.0.17, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: adtable
-- ------------------------------------------------------
-- Server version	8.0.17

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `ad`
--

DROP TABLE IF EXISTS `ad`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ad` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `ad_name` varchar(200) NOT NULL,
  `ad_value` varchar(1000) NOT NULL,
  `ad_first_photo` varchar(100) NOT NULL,
  `ad_second_photo` varchar(100) NOT NULL,
  `ad_third_photo` varchar(100) NOT NULL,
  `ad_price` int(11) NOT NULL DEFAULT '0',
  `ad_creation_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ad`
--

LOCK TABLES `ad` WRITE;
/*!40000 ALTER TABLE `ad` DISABLE KEYS */;
INSERT INTO `ad` VALUES (1,'first_ad','Congratulations! It is you first ad!','fghjkl;','dfg','dfg',0,'2020-01-11 13:06:38'),
(2,'testName','testVal','testFirst','dfg','dfg',0,'2020-01-11 13:06:38'),
(4,'test_name','test_body','qwerty','asdfgh','zxcvbn',123,'2020-01-11 19:33:40'),
(5,'wow','new test','first','second','third',321,'2020-01-11 19:34:40'),
(6,'wow','new test','first','second','third1',321,'2020-01-11 20:50:13'),
(7,'wow','new test','first','second','third2',321,'2020-01-11 20:50:18'),
(8,'wow','new test','first','second','third3',321,'2020-01-11 20:50:24'),
(9,'wow','new test','first','secon2222d','third3',321,'2020-01-11 20:50:28'),
(10,'wow','new test','fir11111st','secon2222d','third3',321,'2020-01-11 20:50:33'),
(11,'wow','new ddddddtest','fir11111st','secon2222d','third3',321,'2020-01-11 20:50:39'),
(12,'wowowowowowow','new ddddddtest','fir11111st','secon2222d','third3',321,'2020-01-11 20:50:49'),
(13,'wow','new test','first','second','third',321,'2020-01-11 20:54:36'),
(14,'wow','new test','first','second','third',321,'2020-01-11 20:55:16'),
(18,'wow','new test','first','second','third',321,'2020-01-11 21:14:29'),
(19,'wow','new test','first','second','third',321,'2020-01-11 21:26:44'),
(20,'wow','new test','first','second','third',321,'2020-01-11 21:32:12');
/*!40000 ALTER TABLE `ad` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-01-12 13:11:54
