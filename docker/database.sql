-- Sabará Mais - MySQL Database
-- Development environment

--
-- Create database
--
CREATE DATABASE IF NOT EXISTS bus;
USE bus;

--
-- Table structure `bus`
--

CREATE TABLE IF NOT EXISTS `bus` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `fare` double NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `number` varchar(255) DEFAULT NULL,
  `company_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_company_key` (`company_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=29 ;

--
-- Sample data for `bus`
--

INSERT INTO `bus` (`id`, `fare`, `name`, `number`, `company_id`) VALUES
(1, 4.8, 'BH/Sabará', '4988', 1),
(2, 6.1, 'BH/Sabará - Executivo', '4987', 1);

--
-- Table structure `company`
--

CREATE TABLE IF NOT EXISTS `company` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `description` varchar(255) DEFAULT NULL,
  `image_url` varchar(255) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=4 ;

--
-- Sample data for `company`
--

INSERT INTO `company` (`id`, `description`, `image_url`, `name`) VALUES
(1, 'Linhas Intermunicipais', 'http://rodrigobrito.net/wp-content/uploads/2017/02/cisne.jpg', 'Cisne');

--
-- Table structure `day_type`
--

CREATE TABLE IF NOT EXISTS `day_type` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `active` bit(1) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=4 ;

--
-- Sample data for `day_type`
--

INSERT INTO `day_type` (`id`, `active`, `name`) VALUES
(1, b'1', 'Dia Útil'),
(2, b'1', 'Sábado'),
(3, b'1', 'Domingo / Feriado');

-- --------------------------------------------------------

--
-- Table structure `schedule`
--

CREATE TABLE IF NOT EXISTS `schedule` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `destiny` varchar(255) DEFAULT NULL,
  `observation` varchar(255) DEFAULT NULL,
  `origin` varchar(255) DEFAULT NULL,
  `time` time DEFAULT NULL,
  `bus_id` bigint(20) DEFAULT NULL,
  `daytype_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_company_key` (`bus_id`),
  KEY `fk_daytype_key` (`daytype_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=2931 ;

--
-- Sample data for `schedule`
--

INSERT INTO `schedule` (`id`, `destiny`, `observation`, `origin`, `time`, `bus_id`, `daytype_id`) VALUES
(1, 'BH', NULL, 'Sabará', '00:10:00', 1, 1),
(2, 'BH', NULL, 'Sabará', '01:30:00', 1, 1),
(3, 'BH', NULL, 'Sabará', '03:30:00', 1, 1),
(4, 'BH', NULL, 'Sabará', '04:20:00', 1, 1),
(5, 'BH', NULL, 'Sabará', '04:40:00', 1, 1),
(6, 'BH', NULL, 'Sabará', '01:30:00', 1, 2),
(7, 'BH', NULL, 'Sabará', '00:10:00', 1, 2),
(8, 'BH', NULL, 'Sabará', '03:30:00', 1, 2),
(9, 'BH', NULL, 'Sabará', '04:25:00', 1, 2),
(10, 'BH', NULL, 'Sabará', '05:00:00', 1, 2),
(11, 'BH', NULL, 'Sabará', '00:10:00', 1, 3),
(12, 'BH', NULL, 'Sabará', '01:30:00', 1, 3),
(13, 'BH', NULL, 'Sabará', '03:30:00', 1, 3),
(14, 'BH', NULL, 'Sabará', '04:30:00', 1, 3),
(15, 'BH', NULL, 'Sabará', '05:00:00', 1, 3),
(16, 'BH', NULL, 'Sabará', '05:25:00', 1, 3);

--
-- Relationship for `bus`
--
ALTER TABLE `bus`
  ADD CONSTRAINT `fk_bus_company` FOREIGN KEY (`company_id`) REFERENCES `company` (`id`);

--
-- Relationship for `schedule`
--
ALTER TABLE `schedule`
  ADD CONSTRAINT `fk_schedule_bus` FOREIGN KEY (`bus_id`) REFERENCES `bus` (`id`),
  ADD CONSTRAINT `fk_schedule_daytype` FOREIGN KEY (`daytype_id`) REFERENCES `day_type` (`id`);