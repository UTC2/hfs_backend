-- MySQL dump 10.19  Distrib 10.3.32-MariaDB, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: hfs_database
-- ------------------------------------------------------
-- Server version	10.3.32-MariaDB-1:10.3.32+maria~focal

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `categories`
--

DROP TABLE IF EXISTS `categories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `categories` (
                              `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                              `name` varchar(255) NOT NULL,
                              `parent_id` int(11) DEFAULT NULL,
                              `inserted_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
                              `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
                              PRIMARY KEY (`id`),
                              UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `categories`
--

LOCK TABLES `categories` WRITE;
/*!40000 ALTER TABLE `categories` DISABLE KEYS */;
/*!40000 ALTER TABLE `categories` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `cc_transactions`
--

DROP TABLE IF EXISTS `cc_transactions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `cc_transactions` (
                                   `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                                   `code` varchar(255) DEFAULT NULL,
                                   `order_id` int(11) NOT NULL,
                                   `transdate` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
                                   `processor` varchar(255) NOT NULL,
                                   `processor_trans_id` varchar(255) NOT NULL,
                                   `amount` decimal(10,0) NOT NULL,
                                   `cc_num` varchar(255) DEFAULT NULL,
                                   `cc_type` varchar(255) DEFAULT NULL,
                                   `response` text DEFAULT NULL,
                                   `inserted_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
                                   `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
                                   PRIMARY KEY (`id`),
                                   UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cc_transactions`
--

LOCK TABLES `cc_transactions` WRITE;
/*!40000 ALTER TABLE `cc_transactions` DISABLE KEYS */;
/*!40000 ALTER TABLE `cc_transactions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `coupons`
--

DROP TABLE IF EXISTS `coupons`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `coupons` (
                           `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                           `code` varchar(255) NOT NULL,
                           `description` text DEFAULT NULL,
                           `active` tinyint(1) DEFAULT 1,
                           `value` decimal(10,0) DEFAULT NULL,
                           `multiple` tinyint(1) DEFAULT 0,
                           `start_date` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
                           `end_date` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
                           `inserted_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
                           `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
                           PRIMARY KEY (`id`),
                           UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `coupons`
--

LOCK TABLES `coupons` WRITE;
/*!40000 ALTER TABLE `coupons` DISABLE KEYS */;
/*!40000 ALTER TABLE `coupons` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order_products`
--

DROP TABLE IF EXISTS `order_products`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `order_products` (
                                  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                                  `order_id` int(11) DEFAULT NULL,
                                  `sku` varchar(255) NOT NULL,
                                  `name` varchar(255) NOT NULL,
                                  `description` text DEFAULT NULL,
                                  `price` decimal(10,0) NOT NULL,
                                  `quantity` int(11) NOT NULL,
                                  `subtotal` decimal(10,0) NOT NULL,
                                  `inserted_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
                                  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
                                  PRIMARY KEY (`id`),
                                  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_products`
--

LOCK TABLES `order_products` WRITE;
/*!40000 ALTER TABLE `order_products` DISABLE KEYS */;
/*!40000 ALTER TABLE `order_products` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product_categories`
--

DROP TABLE IF EXISTS `product_categories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `product_categories` (
                                      `category_id` int(11) NOT NULL,
                                      `product_id` int(11) NOT NULL,
                                      `inserted_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
                                      `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
                                      PRIMARY KEY (`category_id`,`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product_categories`
--

LOCK TABLES `product_categories` WRITE;
/*!40000 ALTER TABLE `product_categories` DISABLE KEYS */;
/*!40000 ALTER TABLE `product_categories` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product_statuses`
--

DROP TABLE IF EXISTS `product_statuses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `product_statuses` (
                                    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                                    `name` varchar(255) NOT NULL,
                                    `inserted_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
                                    `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
                                    PRIMARY KEY (`id`),
                                    UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product_statuses`
--

LOCK TABLES `product_statuses` WRITE;
/*!40000 ALTER TABLE `product_statuses` DISABLE KEYS */;
/*!40000 ALTER TABLE `product_statuses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product_tags`
--

DROP TABLE IF EXISTS `product_tags`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `product_tags` (
                                `product_id` int(11) NOT NULL,
                                `tag_id` int(11) NOT NULL,
                                `inserted_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
                                `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
                                PRIMARY KEY (`product_id`,`tag_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product_tags`
--

LOCK TABLES `product_tags` WRITE;
/*!40000 ALTER TABLE `product_tags` DISABLE KEYS */;
/*!40000 ALTER TABLE `product_tags` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `products`
--

DROP TABLE IF EXISTS `products`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `products` (
                            `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                            `sku` varchar(255) NOT NULL,
                            `name` varchar(255) NOT NULL,
                            `description` text DEFAULT NULL,
                            `product_status_id` int(11) NOT NULL,
                            `regular_price` decimal(10,0) DEFAULT 0,
                            `discount_price` decimal(10,0) DEFAULT 0,
                            `quantity` int(11) DEFAULT 0,
                            `taxable` tinyint(1) DEFAULT 0,
                            `inserted_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
                            `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
                            PRIMARY KEY (`id`),
                            UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `products`
--

LOCK TABLES `products` WRITE;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
/*!40000 ALTER TABLE `products` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `roles`
--

DROP TABLE IF EXISTS `roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `roles` (
                         `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                         `name` varchar(255) NOT NULL,
                         `inserted_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
                         `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `roles`
--

LOCK TABLES `roles` WRITE;
/*!40000 ALTER TABLE `roles` DISABLE KEYS */;
/*!40000 ALTER TABLE `roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sales_orders`
--

DROP TABLE IF EXISTS `sales_orders`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sales_orders` (
                                `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                                `order_date` date NOT NULL,
                                `total` decimal(10,0) NOT NULL,
                                `coupon_id` int(11) DEFAULT NULL,
                                `session_id` varchar(255) NOT NULL,
                                `user_id` int(11) NOT NULL,
                                `inserted_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
                                `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
                                PRIMARY KEY (`id`),
                                UNIQUE KEY `id` (`id`),
                                KEY `fk_session_sales_order` (`session_id`),
                                CONSTRAINT `fk_session_sales_order` FOREIGN KEY (`session_id`) REFERENCES `sessions` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sales_orders`
--

LOCK TABLES `sales_orders` WRITE;
/*!40000 ALTER TABLE `sales_orders` DISABLE KEYS */;
/*!40000 ALTER TABLE `sales_orders` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sessions`
--

DROP TABLE IF EXISTS `sessions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sessions` (
                            `id` varchar(255) NOT NULL,
                            `data` text DEFAULT NULL,
                            `inserted_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
                            `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
                            PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sessions`
--

LOCK TABLES `sessions` WRITE;
/*!40000 ALTER TABLE `sessions` DISABLE KEYS */;
/*!40000 ALTER TABLE `sessions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tags`
--

DROP TABLE IF EXISTS `tags`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tags` (
                        `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                        `name` varchar(255) NOT NULL,
                        `inserted_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
                        `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tags`
--

LOCK TABLES `tags` WRITE;
/*!40000 ALTER TABLE `tags` DISABLE KEYS */;
/*!40000 ALTER TABLE `tags` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_roles`
--

DROP TABLE IF EXISTS `user_roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_roles` (
                              `user_id` int(11) NOT NULL,
                              `role_id` int(11) NOT NULL,
                              `inserted_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
                              `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
                              PRIMARY KEY (`user_id`,`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_roles`
--

LOCK TABLES `user_roles` WRITE;
/*!40000 ALTER TABLE `user_roles` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users` (
                         `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                         `email` varchar(255) NOT NULL,
                         `first_name` varchar(255) NOT NULL,
                         `last_name` varchar(255) NOT NULL,
                         `active` tinyint(1) DEFAULT 1,
                         `inserted_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
                         `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-02-09  8:46:51
