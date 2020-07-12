-- MySQL dump 10.13  Distrib 8.0.19, for Linux (x86_64)
--
-- Host: localhost    Database: donate
-- ------------------------------------------------------
-- Server version	8.0.19

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
-- Table structure for table `blood_type`
--

DROP TABLE IF EXISTS `blood_type`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `blood_type` (
  `id` int NOT NULL AUTO_INCREMENT,
  `type` varchar(3) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `blood_type`
--

LOCK TABLES `blood_type` WRITE;
/*!40000 ALTER TABLE `blood_type` DISABLE KEYS */;
INSERT INTO `blood_type` VALUES (1,'O-'),(2,'O+'),(3,'A-'),(4,'A+'),(5,'B-'),(6,'B+'),(7,'AB-'),(8,'AB+');
/*!40000 ALTER TABLE `blood_type` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `city`
--

DROP TABLE IF EXISTS `city`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `city` (
  `id` int NOT NULL AUTO_INCREMENT,
  `city` varchar(15) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `city`
--

LOCK TABLES `city` WRITE;
/*!40000 ALTER TABLE `city` DISABLE KEYS */;
INSERT INTO `city` VALUES (1,'Cochabamba'),(2,'La Paz'),(3,'Santa Cruz'),(4,'Oruro'),(5,'Potosi'),(6,'Tarija'),(7,'Chuquisaca'),(8,'Beni'),(9,'Pando');
/*!40000 ALTER TABLE `city` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `compatibility`
--

DROP TABLE IF EXISTS `compatibility`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `compatibility` (
  `recipient_blood_type_id` int NOT NULL,
  `donor_blood_type_id` int NOT NULL,
  PRIMARY KEY (`recipient_blood_type_id`,`donor_blood_type_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `compatibility`
--

LOCK TABLES `compatibility` WRITE;
/*!40000 ALTER TABLE `compatibility` DISABLE KEYS */;
INSERT INTO `compatibility` VALUES (1,1),(2,2),(2,4),(2,6),(2,8),(3,3),(4,4),(4,8),(5,5),(6,6),(6,8),(7,7),(8,8);
/*!40000 ALTER TABLE `compatibility` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `donor`
--

DROP TABLE IF EXISTS `donor`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `donor` (
  `id` int NOT NULL AUTO_INCREMENT,
  `blood_type_id` int NOT NULL,
  `name` varchar(250) DEFAULT NULL,
  `cell` varchar(20) NOT NULL,
  `email` varchar(100) DEFAULT NULL,
  `city_id` int DEFAULT NULL,
  `verified` tinyint(1) DEFAULT '1',
  `public` tinyint(1) DEFAULT '1',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `donor`
--

LOCK TABLES `donor` WRITE;
/*!40000 ALTER TABLE `donor` DISABLE KEYS */;
INSERT INTO `donor` VALUES (1,2,'Daniela Goitia','77951654','',1,1,1,'2020-07-04 17:13:55','2020-07-04 17:13:55',NULL),(2,2,'','75521253','',1,1,1,'2020-07-04 17:13:55','2020-07-04 17:13:55',NULL),(3,4,'','70287182','',2,1,1,'2020-07-04 17:13:55','2020-07-04 17:13:55',NULL),(4,4,'Daniel Coca Mallon','71389604','',3,1,1,'2020-07-04 17:13:55','2020-07-04 17:13:55',NULL),(5,2,'Camaleon Gzf','69086590','',3,1,1,'2020-07-04 17:13:55','2020-07-04 17:13:55',NULL),(6,2,'','78523417','',3,1,1,'2020-07-04 17:13:55','2020-07-04 17:13:55',NULL),(7,2,'Nataly Garzon','70867858','',3,1,1,'2020-07-04 17:13:55','2020-07-04 17:13:55',NULL),(8,2,'Iver Gonzales','76360350','',3,1,1,'2020-07-04 17:13:55','2020-07-04 17:13:55',NULL),(9,4,'Abraham Pho','62085356','',NULL,1,1,'2020-07-04 17:13:55','2020-07-04 17:13:55',NULL),(10,4,'','62085356','',NULL,1,1,'2020-07-04 17:13:55','2020-07-04 17:13:55',NULL),(11,4,'','62085356','',NULL,1,1,'2020-07-04 17:13:55','2020-07-04 17:13:55',NULL),(12,4,'','70253631','',NULL,1,1,'2020-07-04 17:13:55','2020-07-04 17:13:55',NULL),(13,4,'','62085356','',NULL,1,1,'2020-07-04 17:13:55','2020-07-04 17:13:55',NULL),(14,4,'Sebastian Salinas','70253631','',NULL,1,1,'2020-07-04 17:13:55','2020-07-04 17:13:55',NULL),(15,4,'Andy Guz','727608','',NULL,1,1,'2020-07-04 17:13:55','2020-07-04 17:13:55',NULL),(16,4,'','70287182','',NULL,1,1,'2020-07-04 17:13:55','2020-07-04 17:13:55',NULL);
/*!40000 ALTER TABLE `donor` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `donor_recipient`
--

DROP TABLE IF EXISTS `donor_recipient`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `donor_recipient` (
  `donor_id` int NOT NULL,
  `recipient_id` int NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`donor_id`,`recipient_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `donor_recipient`
--

LOCK TABLES `donor_recipient` WRITE;
/*!40000 ALTER TABLE `donor_recipient` DISABLE KEYS */;
/*!40000 ALTER TABLE `donor_recipient` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `recipient`
--

DROP TABLE IF EXISTS `recipient`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `recipient` (
  `id` int NOT NULL AUTO_INCREMENT,
  `blood_type_id` int NOT NULL,
  `name` varchar(250) DEFAULT NULL,
  `cell_numbers` varchar(100) NOT NULL,
  `email` varchar(100) DEFAULT NULL,
  `photo_path` varchar(150) DEFAULT NULL,
  `city_id` int DEFAULT NULL,
  `verified` tinyint(1) DEFAULT '1',
  `public` tinyint(1) DEFAULT '1',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=89 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `recipient`
--

LOCK TABLES `recipient` WRITE;
/*!40000 ALTER TABLE `recipient` DISABLE KEYS */;
INSERT INTO `recipient` VALUES (1,4,'Jubenal Butron','44583137','',NULL,1,1,1,'2020-07-04 17:05:59','2020-07-04 17:05:59',NULL),(2,5,'Delia Alvarez B.','70386272','',NULL,1,1,1,'2020-07-04 17:05:59','2020-07-04 17:05:59',NULL),(3,6,'','76443838','',NULL,1,1,1,'2020-07-04 17:05:59','2020-07-04 17:05:59',NULL),(4,2,'Edwin Zurita Flores','72733446','',NULL,1,1,1,'2020-07-04 17:05:59','2020-07-04 17:05:59',NULL),(5,2,'Majail Tapia Monroy','72242948','',NULL,1,1,1,'2020-07-04 17:05:59','2020-07-04 17:05:59',NULL),(6,2,'','74646564','',NULL,1,1,1,'2020-07-04 17:05:59','2020-07-04 17:05:59',NULL),(7,2,'Jaime Albarracin','67464397','',NULL,1,1,1,'2020-07-04 17:05:59','2020-07-04 17:05:59',NULL),(8,2,'Edwin Zurita Flores','72733446','',NULL,1,1,1,'2020-07-04 17:05:59','2020-07-04 17:05:59',NULL),(9,2,'Maria Del Carmen Cala','72782737','',NULL,1,1,1,'2020-07-04 17:05:59','2020-07-04 17:05:59',NULL),(10,2,'Luz Vasquez Gonzales','70300775','',NULL,1,1,1,'2020-07-04 17:05:59','2020-07-04 17:05:59',NULL),(11,2,'Hugo Zeballos Claros','77494624','',NULL,1,1,1,'2020-07-04 17:05:59','2020-07-04 17:05:59',NULL),(12,2,'Luis Alberto Garcia Gongora','72259872','',NULL,1,1,1,'2020-07-04 17:05:59','2020-07-04 17:05:59',NULL),(13,2,'Noel Mamani Villca','67751491','',NULL,1,1,1,'2020-07-04 17:06:00','2020-07-04 17:06:00',NULL),(14,2,'Vanessa Pardo Vasquez','70300775','',NULL,1,1,1,'2020-07-04 17:06:00','2020-07-04 17:06:00',NULL),(15,2,'David Aguada Semo','71720433','',NULL,1,1,1,'2020-07-04 17:06:00','2020-07-04 17:06:00',NULL),(16,2,'Jorge Luizaga','70766082','',NULL,1,1,1,'2020-07-04 17:06:00','2020-07-04 17:06:00',NULL),(17,2,'Alfredo Morato Maldonado','77995363','',NULL,1,1,1,'2020-07-04 17:06:00','2020-07-04 17:06:00',NULL),(18,2,'','73760882','',NULL,1,1,1,'2020-07-04 17:06:00','2020-07-04 17:06:00',NULL),(19,2,'Juan Carlos Marcani Sejas','72202426','',NULL,1,1,1,'2020-07-04 17:06:00','2020-07-04 17:06:00',NULL),(20,4,'','70630958','',NULL,2,1,1,'2020-07-04 17:06:00','2020-07-04 17:06:00',NULL),(21,4,'Fabiola Linares De Peñaranda','77725594','',NULL,2,1,1,'2020-07-04 17:06:00','2020-07-04 17:06:00',NULL),(22,2,'Gonzalo Alberto Peñarands Vargas','77725594','',NULL,2,1,1,'2020-07-04 17:06:00','2020-07-04 17:06:00',NULL),(23,2,'','77511153','',NULL,2,1,1,'2020-07-04 17:06:00','2020-07-04 17:06:00',NULL),(24,2,'Harold Martinez Rivera','79600677','',NULL,2,1,1,'2020-07-04 17:06:00','2020-07-04 17:06:00',NULL),(25,2,'Karen Velasco Rodal','62600524','',NULL,2,1,1,'2020-07-04 17:06:00','2020-07-04 17:06:00',NULL),(26,4,'Luis Severich','76803348','',NULL,3,1,1,'2020-07-04 17:06:00','2020-07-04 17:06:00',NULL),(27,4,'','73121782','',NULL,3,1,1,'2020-07-04 17:06:00','2020-07-04 17:06:00',NULL),(28,4,'Susana Tonores Fariñas','72119653','',NULL,3,1,1,'2020-07-04 17:06:00','2020-07-04 17:06:00',NULL),(29,4,'Rodrigo Gil Camacho','70251603','',NULL,3,1,1,'2020-07-04 17:06:00','2020-07-04 17:06:00',NULL),(30,2,'Dolores Garcia','67890388','',NULL,3,1,1,'2020-07-04 17:06:00','2020-07-04 17:06:00',NULL),(31,2,'Raul Choque','68922617','',NULL,3,1,1,'2020-07-04 17:06:00','2020-07-04 17:06:00',NULL),(32,2,'Santsder Choque Mamani','78550518','',NULL,3,1,1,'2020-07-04 17:06:00','2020-07-04 17:06:00',NULL),(33,2,'Iblen Gomez Duran','63464496','',NULL,3,1,1,'2020-07-04 17:06:00','2020-07-04 17:06:00',NULL),(34,2,'Addy Maria Agilera Algarañas','77347657','',NULL,3,1,1,'2020-07-04 17:06:00','2020-07-04 17:06:00',NULL),(35,2,'Damiana Loza Poma','76563196','',NULL,3,1,1,'2020-07-04 17:06:00','2020-07-04 17:06:00',NULL),(36,2,'Jorge Altamirano Campoverde','77741774','',NULL,4,1,1,'2020-07-04 17:06:00','2020-07-04 17:06:00',NULL),(37,2,'Felix Colque','72435353','',NULL,5,1,1,'2020-07-04 17:06:00','2020-07-04 17:06:00',NULL),(38,4,'Luis Severich','76803348','',NULL,NULL,1,1,'2020-07-04 17:06:00','2020-07-04 17:06:00',NULL),(39,4,'Federico Gutierrez','72210665','',NULL,NULL,1,1,'2020-07-04 17:06:00','2020-07-04 17:06:00',NULL),(40,4,'Tito Patiño Menacho','76341027','',NULL,NULL,1,1,'2020-07-04 17:06:00','2020-07-04 17:06:00',NULL),(41,4,'Chingolo Pereira','75363070','',NULL,NULL,1,1,'2020-07-04 17:06:00','2020-07-04 17:06:00',NULL),(42,4,'Margot Arauco','72109082','',NULL,NULL,1,1,'2020-07-04 17:06:00','2020-07-04 17:06:00',NULL),(43,4,'Tito Patiño','76341027','',NULL,NULL,1,1,'2020-07-04 17:06:00','2020-07-04 17:06:00',NULL),(44,4,'Maria Veizaga','76382224','',NULL,NULL,1,1,'2020-07-04 17:06:00','2020-07-04 17:06:00',NULL),(45,4,'Alberto Murrillo Guzman','76931423','',NULL,NULL,1,1,'2020-07-04 17:06:00','2020-07-04 17:06:00',NULL),(46,4,'Natividad Muñoz','71059477','',NULL,NULL,1,1,'2020-07-04 17:06:01','2020-07-04 17:06:01',NULL),(47,4,'Victor Luiz Semizo Vaca','70287844','',NULL,NULL,1,1,'2020-07-04 17:06:01','2020-07-04 17:06:01',NULL),(48,4,'Maria Sonia Soliz Franco','76332989','',NULL,NULL,1,1,'2020-07-04 17:06:01','2020-07-04 17:06:01',NULL),(49,4,'Maria Eugenia Vargas Fernandez','70313888','',NULL,NULL,1,1,'2020-07-04 17:06:01','2020-07-04 17:06:01',NULL),(50,4,'Carlos','73121782','',NULL,NULL,1,1,'2020-07-04 17:06:01','2020-07-04 17:06:01',NULL),(51,4,'Pedro Castro','79882846','',NULL,NULL,1,1,'2020-07-04 17:06:01','2020-07-04 17:06:01',NULL),(52,4,'Omar Lastra Matienzo','75019599','',NULL,NULL,1,1,'2020-07-04 17:06:01','2020-07-04 17:06:01',NULL),(53,5,'','79351124','',NULL,NULL,1,1,'2020-07-04 17:06:01','2020-07-04 17:06:01',NULL),(54,6,'Luis','70711571','',NULL,NULL,1,1,'2020-07-04 17:06:01','2020-07-04 17:06:01',NULL),(55,6,'Rolando Escalante','72697065','',NULL,NULL,1,1,'2020-07-04 17:06:01','2020-07-04 17:06:01',NULL),(56,1,'','79351124','',NULL,NULL,1,1,'2020-07-04 17:06:01','2020-07-04 17:06:01',NULL),(57,1,'','79351124','',NULL,NULL,1,1,'2020-07-04 17:06:01','2020-07-04 17:06:01',NULL),(58,1,'Cristian Tomicha','71359123','',NULL,NULL,1,1,'2020-07-04 17:06:01','2020-07-04 17:06:01',NULL),(59,1,'Paul Melgar Cadario','77834671','',NULL,NULL,1,1,'2020-07-04 17:06:01','2020-07-04 17:06:01',NULL),(60,1,'','70052070','',NULL,NULL,1,1,'2020-07-04 17:06:01','2020-07-04 17:06:01',NULL),(61,2,'José Ruben Espinoza','71775155','',NULL,NULL,1,1,'2020-07-04 17:06:01','2020-07-04 17:06:01',NULL),(62,2,'Ernesto Romero','71531354','',NULL,NULL,1,1,'2020-07-04 17:06:01','2020-07-04 17:06:01',NULL),(63,2,'','78305151','',NULL,NULL,1,1,'2020-07-04 17:06:01','2020-07-04 17:06:01',NULL),(64,2,'Oliver Apaza Marca','77568066','',NULL,NULL,1,1,'2020-07-04 17:06:01','2020-07-04 17:06:01',NULL),(65,2,'Tatiana Cornejo','65713906','',NULL,NULL,1,1,'2020-07-04 17:06:01','2020-07-04 17:06:01',NULL),(66,2,'Bianka Miranda','76101282','',NULL,NULL,1,1,'2020-07-04 17:06:01','2020-07-04 17:06:01',NULL),(67,2,'Fabiola Nostas Villarroel','72612614','',NULL,NULL,1,1,'2020-07-04 17:06:01','2020-07-04 17:06:01',NULL),(68,2,'','78305151','',NULL,NULL,1,1,'2020-07-04 17:06:01','2020-07-04 17:06:01',NULL),(69,2,'Octavio Yucra','73133359','',NULL,NULL,1,1,'2020-07-04 17:06:01','2020-07-04 17:06:01',NULL),(70,2,'Rolando Inturias','71316151','',NULL,NULL,1,1,'2020-07-04 17:06:01','2020-07-04 17:06:01',NULL),(71,2,'Doris Albino Estrada','69314891','',NULL,NULL,1,1,'2020-07-04 17:06:01','2020-07-04 17:06:01',NULL),(72,2,'','77502120','',NULL,NULL,1,1,'2020-07-04 17:06:01','2020-07-04 17:06:01',NULL),(73,2,'Mauricio','69416516','',NULL,NULL,1,1,'2020-07-04 17:06:01','2020-07-04 17:06:01',NULL),(74,2,'','71685763','',NULL,NULL,1,1,'2020-07-04 17:06:01','2020-07-04 17:06:01',NULL),(75,2,'Virginia Cruz Puneira','63500444','',NULL,NULL,1,1,'2020-07-04 17:06:01','2020-07-04 17:06:01',NULL),(76,2,'Adela Fernandez Contreras','70025116','',NULL,NULL,1,1,'2020-07-04 17:06:01','2020-07-04 17:06:01',NULL),(77,2,'Amilcar Chambi','70057770','',NULL,NULL,1,1,'2020-07-04 17:06:01','2020-07-04 17:06:01',NULL),(78,2,'Luz Pamela Katari Valencia','75211019','',NULL,NULL,1,1,'2020-07-04 17:06:01','2020-07-04 17:06:01',NULL),(79,2,'Aurelio Mendoza','68032037','',NULL,NULL,1,1,'2020-07-04 17:06:01','2020-07-04 17:06:01',NULL),(80,2,'Edgar Yañez','72208799','',NULL,NULL,1,1,'2020-07-04 17:06:01','2020-07-04 17:06:01',NULL),(81,2,'Delfin Villarroel','79070332','',NULL,NULL,1,1,'2020-07-04 17:06:01','2020-07-04 17:06:01',NULL),(82,2,'','60501886','',NULL,NULL,1,1,'2020-07-04 17:06:01','2020-07-04 17:06:01',NULL),(83,2,'Moises Choi','79028080','',NULL,NULL,1,1,'2020-07-04 17:06:02','2020-07-04 17:06:02',NULL),(84,2,'Nestor Cháves Ortega','70047842','',NULL,NULL,1,1,'2020-07-04 17:06:02','2020-07-04 17:06:02',NULL),(85,2,'Luis Fernando Flores','60501886','',NULL,NULL,1,1,'2020-07-04 17:06:02','2020-07-04 17:06:02',NULL),(86,4,'Yesi Rojas G','71726777','yesi.rojas.g@gmail.com','none',2,1,1,'2020-07-04 17:06:02','2020-07-04 17:06:02',NULL),(87,2,'Mauricio Murillo','69416516','',NULL,NULL,1,1,'2020-07-04 17:06:02','2020-07-04 17:06:02',NULL),(88,2,'Carlos Miguel Dominguezaltamirano','78767040','',NULL,NULL,1,1,'2020-07-04 17:06:02','2020-07-04 17:06:02',NULL);
/*!40000 ALTER TABLE `recipient` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-07-04 22:59:58