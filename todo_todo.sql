-- MySQL dump 10.13  Distrib 8.0.25, for Win64 (x86_64)
--
-- Host: localhost    Database: todo
-- ------------------------------------------------------
-- Server version	8.0.25

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
-- Table structure for table `todo`
--

DROP TABLE IF EXISTS `todo`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `todo` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(45) DEFAULT NULL,
  `description` varchar(200) DEFAULT NULL,
  `category` varchar(45) DEFAULT NULL,
  `progress` varchar(45) DEFAULT NULL,
  `deadline` datetime DEFAULT NULL,
  `status` varchar(45) DEFAULT NULL,
  `createdTime` datetime DEFAULT NULL,
  `updatedTime` datetime DEFAULT NULL,
  `remainingDay` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `todo`
--

LOCK TABLES `todo` WRITE;
/*!40000 ALTER TABLE `todo` DISABLE KEYS */;
INSERT INTO `todo` VALUES (1,'title1','description1 acilll','work','in progress','2021-06-22 12:42:31','urgent','2021-06-03 15:30:23',NULL,18),(3,'title3','description3 acil','home','in progress','2021-06-05 12:42:31','urgent','2021-05-03 15:31:57','2021-06-03 15:40:31',1),(4,'title4','description4','work','done','2021-06-12 12:42:31','normal','2021-06-02 15:43:16',NULL,0),(5,'title5','description5 acil','work','in progress','2021-06-15 12:42:31','urgent','2021-06-02 16:29:54','2021-06-02 21:12:14',12),(6,'title6','description6 acil','home','in progress','2021-06-15 12:42:31','urgent','2021-06-02 17:02:13','2021-06-02 18:16:39',12),(7,'title6','description6 acil','home','done','2021-06-15 12:42:31','urgent','2021-06-02 18:14:16','2021-06-02 18:17:01',0),(8,'title5','description5','home','in progress','2021-06-15 12:42:31','normal','2021-06-02 18:31:34',NULL,12),(9,'title5','description5','hobby','in progress','2021-06-15 12:42:31','normal','2021-06-02 18:58:25',NULL,12),(10,'title5','description5','work','in progress','2021-06-15 12:42:31','normal','2021-06-02 18:58:46',NULL,12),(11,'title512','description5','work','in progress','2021-06-15 12:42:31','normal','2021-06-02 19:37:54',NULL,12),(12,'title51asdasd2','descriasdsadption5','work','in progress','2021-06-15 12:42:31','normal','2021-06-02 19:39:04',NULL,12),(13,'title2','descr5','work','in progress','2021-06-15 12:42:31','normal','2021-06-02 19:42:24',NULL,12),(14,'title2','descr5','work','in progress','2021-06-15 12:42:31','normal','2021-06-02 19:43:19',NULL,12),(15,'title2','descr5','work','in progress','2021-06-15 12:42:31','normal','2021-06-02 20:26:06',NULL,12),(16,'title2','descr5','work','done','2021-05-15 12:42:31','normal','2021-06-02 20:27:28',NULL,0);
/*!40000 ALTER TABLE `todo` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-06-03  0:29:20
