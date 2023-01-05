-- MySQL dump 10.13  Distrib 8.0.30, for Linux (x86_64)
--
-- Host: localhost    Database: test_abera
-- ------------------------------------------------------
-- Server version	8.0.30-0ubuntu0.20.04.2

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `kecamatans`
--

DROP TABLE IF EXISTS `kecamatans`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `kecamatans` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `kota_id` bigint NOT NULL,
  `kecamatan_name` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `kecamatan_FK` (`kota_id`),
  CONSTRAINT `kecamatan_FK` FOREIGN KEY (`kota_id`) REFERENCES `kota` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `kecamatans`
--

LOCK TABLES `kecamatans` WRITE;
/*!40000 ALTER TABLE `kecamatans` DISABLE KEYS */;
/*!40000 ALTER TABLE `kecamatans` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `kelurahans`
--

DROP TABLE IF EXISTS `kelurahans`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `kelurahans` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `kecamatan_id` bigint NOT NULL,
  `kelurahan_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `keluarahan_FK` (`kecamatan_id`),
  CONSTRAINT `keluarahan_FK` FOREIGN KEY (`kecamatan_id`) REFERENCES `kecamatans` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `kelurahans`
--

LOCK TABLES `kelurahans` WRITE;
/*!40000 ALTER TABLE `kelurahans` DISABLE KEYS */;
/*!40000 ALTER TABLE `kelurahans` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `kota`
--

DROP TABLE IF EXISTS `kota`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `kota` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `kota_name` varchar(100) DEFAULT NULL,
  `provinsi_id` bigint NOT NULL,
  PRIMARY KEY (`id`),
  KEY `kota_FK` (`provinsi_id`),
  CONSTRAINT `kota_FK` FOREIGN KEY (`provinsi_id`) REFERENCES `provinsis` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `kota`
--

LOCK TABLES `kota` WRITE;
/*!40000 ALTER TABLE `kota` DISABLE KEYS */;
/*!40000 ALTER TABLE `kota` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `member_details`
--

DROP TABLE IF EXISTS `member_details`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `member_details` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `nik` varchar(100) DEFAULT NULL,
  `phone_number` varchar(100) DEFAULT NULL,
  `address` text,
  `provinsi_id` bigint NOT NULL,
  `kota_id` bigint NOT NULL,
  `kelurahan_id` bigint NOT NULL,
  `kecamatan_id` bigint NOT NULL,
  PRIMARY KEY (`id`),
  KEY `member_details_FK` (`user_id`),
  CONSTRAINT `member_details_FK` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `member_details`
--

LOCK TABLES `member_details` WRITE;
/*!40000 ALTER TABLE `member_details` DISABLE KEYS */;
INSERT INTO `member_details` VALUES (3,1,'12345678','09832749234','ciputat banten',1,1,1,1);
/*!40000 ALTER TABLE `member_details` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `provinsis`
--

DROP TABLE IF EXISTS `provinsis`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `provinsis` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `provinsi_name` varchar(100) DEFAULT NULL,
  `provinsi_code` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `provinsis`
--

LOCK TABLES `provinsis` WRITE;
/*!40000 ALTER TABLE `provinsis` DISABLE KEYS */;
INSERT INTO `provinsis` VALUES (2,'test2','123'),(3,'test2','123');
/*!40000 ALTER TABLE `provinsis` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `full_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `user_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `email` varchar(100) DEFAULT NULL,
  `password` varchar(100) DEFAULT NULL,
  `profile_pict` varchar(100) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'sujai','sujai','sujai@sujai.com','$2a$04$Pfq5itpTurefBaK1fQ0wC./U0DoUuxmIBbo4lQ.VvOkCrIO6QMINa','images/1-avatar.jpeg','2022-08-10 23:24:10','2022-08-12 13:55:44');
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

-- Dump completed on 2022-08-12 14:52:56
