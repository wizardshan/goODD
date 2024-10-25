-- phpMyAdmin SQL Dump
-- version 4.8.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: 2024-10-25 15:07:32
-- 服务器版本： 5.7.19
-- PHP Version: 7.4.6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `odd`
--

-- --------------------------------------------------------

--
-- 表的结构 `users`
--

CREATE TABLE `users` (
  `id` int(11) UNSIGNED NOT NULL,
  `hash_id` varchar(20) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `mobile` varchar(11) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `password` varchar(80) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `age` int(11) NOT NULL DEFAULT '0',
  `level` int(11) NOT NULL DEFAULT '0',
  `nickname` varchar(20) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `avatar` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `bio` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `amount` int(11) NOT NULL DEFAULT '0',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

--
-- 转存表中的数据 `users`
--

INSERT INTO `users` (`id`, `hash_id`, `mobile`, `password`, `age`, `level`, `nickname`, `avatar`, `bio`, `amount`, `create_time`, `update_time`) VALUES
(1, 'oKqk6tMl7z', '13000000001', '$2a$10$ZyJya29aneBi797J0oP4j.3Qj2vLkw1jlEAbBNIuBg0CaVy6CgSHS', 11, 10, '1300****001', 'avatar_default.png', '个人介绍', 800, '2024-10-11 08:00:40', '2024-10-11 08:36:53'),
(2, '02qN7SQyOb', '13000000002', '$2a$10$6DcrvLZ0wuzjI1CFKl0CpeSOOWQTbLUytCiIcGRkOZl5X3rxPVedu', 30, 0, '1300****002', 'avatar_default.png', '', 0, '2024-10-11 08:00:43', '2024-10-11 08:00:43'),
(3, 'zalLecoq7G', '13000000003', '$2a$10$kupx.Xpnsze3LYei3qqd6eaXDcuwD.Wb3zmU5p0JtUMrtVst9cQ5u', 28, 0, '1300****003', 'avatar_custom.png', '', 0, '2024-10-11 08:00:45', '2024-10-11 08:00:56');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
