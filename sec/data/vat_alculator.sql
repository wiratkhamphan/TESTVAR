-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jun 01, 2024 at 04:06 AM
-- Server version: 10.4.27-MariaDB
-- PHP Version: 8.0.25

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `shoplek`
--

-- --------------------------------------------------------

--
-- Table structure for table `vat_alculator`
--

CREATE TABLE `vat_alculator` (
  `id` int(11) NOT NULL,
  `price` int(11) NOT NULL,
  `vat` int(11) NOT NULL,
  `netprice` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `vat_alculator`
--

INSERT INTO `vat_alculator` (`id`, `price`, `vat`, `netprice`) VALUES
(1, 222, 0, 0),
(2, 222, 0, 0),
(3, 1111, 78, 1189),
(4, 112, 8, 120),
(5, 1121, 78, 1199),
(6, 11122, 779, 11901);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `vat_alculator`
--
ALTER TABLE `vat_alculator`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `vat_alculator`
--
ALTER TABLE `vat_alculator`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
