-- MySQL dump 10.13  Distrib 8.0.31, for Win64 (x86_64)
--
-- Host: localhost    Database: msisensor
-- ------------------------------------------------------
-- Server version	8.0.31

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
-- Table structure for table `sample`
--

DROP TABLE IF EXISTS `sample`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sample` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `uid` bigint unsigned NOT NULL,
  `patient_name` varchar(20) NOT NULL,
  `file_name` varchar(20) DEFAULT NULL,
  `cancer_type` varchar(20) DEFAULT NULL,
  `gender` bigint DEFAULT NULL,
  `age` bigint DEFAULT '0',
  `prediction` bigint DEFAULT '2',
  PRIMARY KEY (`id`),
  KEY `idx_sample_deleted_at` (`deleted_at`),
  KEY `fk_sample_user` (`uid`),
  CONSTRAINT `fk_sample_user` FOREIGN KEY (`uid`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sample`
--

LOCK TABLES `sample` WRITE;
/*!40000 ALTER TABLE `sample` DISABLE KEYS */;
INSERT INTO `sample` VALUES (1,'2023-03-01 17:29:49.724','2023-03-01 21:05:12.188','2023-03-02 11:10:28.764',3,'Amm','sample.csv','COAD',1,22,0),(2,'2023-03-01 20:49:43.272','2023-03-01 21:05:05.692','2023-03-01 21:06:35.192',3,'Yoho','test.csv','STAD',1,22,0),(3,'2023-03-01 21:05:58.637','2023-03-02 11:10:17.176',NULL,3,'Lily','test.csv','STAD',0,56,0),(4,'2023-03-02 15:07:58.076','2023-03-02 15:08:05.890',NULL,3,'Oliver','','UCEC',1,44,2),(5,'2023-03-02 15:08:24.704','2023-03-02 15:12:12.076',NULL,3,'Andy','sample.csv','others',0,44,0),(6,'2023-03-02 15:12:01.602','2023-03-02 15:12:01.602',NULL,3,'Tom','','UCEC',1,55,2),(7,'2023-03-02 15:14:19.923','2023-03-02 15:14:19.923',NULL,3,'Sam','','UCEC',1,39,2),(8,'2023-03-02 15:14:32.102','2023-03-02 15:15:41.146',NULL,3,'Judy','','STAD',2,11,2),(9,'2023-03-02 15:14:51.775','2023-03-02 15:52:31.082',NULL,3,'Ive','sample.csv','STAD',1,87,0),(10,'2023-03-02 15:18:03.822','2023-03-02 15:38:30.501','2023-03-02 15:52:23.109',3,'Golden','test.csv','UCEC',1,43,0),(11,'2023-03-02 15:19:12.953','2023-03-02 15:19:12.953','2023-03-02 15:30:00.794',3,'','','',2,0,2),(12,'2023-03-02 15:22:08.972','2023-03-02 15:26:47.875','2023-03-02 15:36:50.776',3,'Ann','','STAD',0,90,2),(13,'2023-03-02 15:24:53.552','2023-03-02 15:24:53.552','2023-03-02 15:35:54.797',3,'Tank','','STAD',2,11,2),(14,'2023-03-02 15:28:30.424','2023-03-02 15:28:30.424','2023-03-02 15:29:54.272',3,'Yu','','COAD',2,63,2),(15,'2023-03-02 15:29:39.781','2023-03-02 15:29:39.781','2023-03-02 15:29:51.651',3,'','','',2,0,2),(16,'2023-03-02 15:29:43.296','2023-03-02 15:29:43.296','2023-03-02 15:29:56.748',3,'','','',1,0,2),(17,'2023-03-02 15:29:46.441','2023-03-02 15:29:46.441','2023-03-02 15:29:49.359',3,'','','',2,0,2),(18,'2023-03-02 15:31:24.188','2023-03-02 15:31:24.188','2023-03-02 15:35:40.870',3,'army','','STAD',2,0,2),(20,'2023-03-02 15:36:59.240','2023-03-02 15:36:59.240','2023-03-02 15:37:08.754',3,'wuwu','','',0,0,2),(21,'2023-03-02 15:37:04.515','2023-03-02 15:37:04.515','2023-03-02 15:37:06.862',3,'','','',1,0,2);
/*!40000 ALTER TABLE `sample` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `username` varchar(20) NOT NULL,
  `password` varchar(500) NOT NULL,
  `role` bigint DEFAULT '2',
  PRIMARY KEY (`id`),
  KEY `idx_user_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (3,'2023-02-21 17:33:41.681','2023-02-21 17:33:41.681',NULL,'admin','$2a$10$HZyJhnxGsrxzRJtJ3qXyrOMm1.vmVyV8Yzt287KiszLIaugDhWgfW',1),(4,'2023-02-23 13:10:14.106','2023-02-23 13:10:14.106',NULL,'admin1','$2a$10$0eOWuqbIlE5fO1sRNXElEOVymVphxgc0BNt1jl9booQUpabUlwRAe',2),(5,'2023-02-25 20:47:57.162','2023-02-25 20:47:57.162',NULL,'yangxuanhao','$2a$10$n.xKcxgt0911rodTE/s/ZOT3qzNBGiC7MD2tETdmZULRDHxcdf1RS',2);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-03-02 21:15:38
