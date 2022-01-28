-- MySQL dump 10.13  Distrib 5.5.62, for linux-glibc2.12 (x86_64)
--
-- Host: localhost    Database: hfs_database
-- ------------------------------------------------------
-- Server version	5.5.62

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `hfs_sale`
--

DROP TABLE IF EXISTS `hfs_sale`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `hfs_sale` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `order_id` varchar(512) CHARACTER SET utf8 DEFAULT NULL,
  `product_id` int(11) DEFAULT NULL,
  `modelTransaction` varchar(255) CHARACTER SET utf8 DEFAULT NULL,
  `feeRental` varchar(255) CHARACTER SET utf8 DEFAULT NULL,
  `startDateRental` date DEFAULT NULL,
  `endDateRental` date DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8 DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `device_id` (`product_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=latin1 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `hfs_sale`
--

LOCK TABLES `hfs_sale` WRITE;
/*!40000 ALTER TABLE `hfs_sale` DISABLE KEYS */;
/*!40000 ALTER TABLE `hfs_sale` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `hfs_user`
--

DROP TABLE IF EXISTS `hfs_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `hfs_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `groupid` text NOT NULL,
  `default_lang_id` tinyint(11) DEFAULT NULL,
  `username` varchar(32) NOT NULL,
  `password` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `city` varchar(255) NOT NULL,
  `district` varchar(255) NOT NULL,
  `Commune` varchar(255) NOT NULL,
  `birthday` date DEFAULT NULL,
  `joindate` bigint(11) DEFAULT NULL,
  `lastvisit` bigint(11) DEFAULT NULL,
  `ipaddress` varchar(64) DEFAULT '',
  `active` tinyint(1) NOT NULL DEFAULT '0',
  `activecode` varchar(255) DEFAULT NULL,
  `address` text,
  `groups` text,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `hfs_user`
--

LOCK TABLES `hfs_user` WRITE;
/*!40000 ALTER TABLE `hfs_user` DISABLE KEYS */;
/*!40000 ALTER TABLE `hfs_user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-01-28  6:25:28
