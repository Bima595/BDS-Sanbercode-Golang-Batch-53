/*
 Navicat Premium Data Transfer

 Source Server         : mysql-docker
 Source Server Type    : MySQL
 Source Server Version : 80300 (8.3.0)
 Source Host           : localhost:31120
 Source Schema         : quiz3

 Target Server Type    : MySQL
 Target Server Version : 80300 (8.3.0)
 File Encoding         : 65001

 Date: 05/02/2024 19:40:43
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `category_id` int NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `update_at` timestamp NULL DEFAULT NULL,
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `image_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `article_length` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `category_id`(`category_id` ASC) USING BTREE,
  CONSTRAINT `category_id` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of article
-- ----------------------------
INSERT INTO `article` VALUES (8, 'Patrick', 4, '2024-02-05 09:57:39', '2024-02-05 09:57:39', 'This is a sample article content.asdhjakshdjkahjasdjkhd', 'https://example.com/image.jpg', 'medium');
INSERT INTO `article` VALUES (9, 'Algoritma Basisdata', 4, '2024-02-05 11:43:35', '2024-02-05 11:43:35', 'aksjdhklajsdasbdkasbckajshjklashdblkasbndkjaskjdhakjslhdlkwnkdbaskjbd', 'https://example.com/image.jpg', 'medium');
INSERT INTO `article` VALUES (10, 'Algoritma Python', 4, '2024-02-05 11:43:44', '2024-02-05 11:43:44', 'aksjdhklajsdasbdkasbckajshjklashdblkasbndkjaskjdhakjslhdlkwnkdbaskjbd', 'https://example.com/image.jpg', 'medium');
INSERT INTO `article` VALUES (12, 'Sample Title', 4, '2024-02-05 11:58:30', '2024-02-05 11:58:30', 'Sample Contentsifuwfhwjehkfjklashjkdhaskhasjkldhljkashdajklhdjkalshdjklahdjklashhdkljahsdjklsahdajkldhjkaslhdjkahdsjkhaiwudhiudxjwuiahdgaudhtyuihdgaushnhasdjasuhdudhyasghdyuagwhwsdgyhyadgh', 'https://example.com/image.jpg', 'pendek');
INSERT INTO `article` VALUES (13, 'Jane Austen’s Pride and Prejudice', 4, '2024-02-05 12:00:22', '2024-02-05 12:00:22', 'Jane Austen’s Pride and Prejudice is an 18th-century novel of manners set in rural England and portraying the relationships between the four daughters of the Bennet family and their neighbors. While accurately and vividly depicting the manners and social norms of that time, the novel also provides sharp observations on the themes of love, marriage, class, money, education, and social prestige. In this paper, the four main themes of Pride and Prejudice are analyzed. Marriage is the main topic around which the plot revolves. The author illustrates the conflict between marrying for money, which was the typical idea at the time, and marrying for love. In either case, the economic and social differences were obstacles which made it hard for young women from poor families to break out of their social circle. Each person’s position in society was determined by their class, and the relations between families also centered around differences in wealth and status. The gender differences also played an important role, as women were considered inferior to men and were practically unable to choose partners. Austen both criticizes and examines the social life of 18th-century England, advocating for marrying for love as one of the essential female rights.wqidjqwoiwqdjoiwqjdioqwjdioqjqiowjdijqijdwodjwqiodjwqiodjwqoidjwqjwqijdqdowqjdqwiojdwqojdwqiojdqwd', 'https://example.com/image.jpg', 'panjang');
INSERT INTO `article` VALUES (14, 'Jane', 4, '2024-02-05 12:11:19', '2024-02-05 12:11:19', 'Jane Austen’s Pride and Prejudice is an 18th-century novel of manners set in rural.wqidjqwoiwqdjoiwqjdioqwjdioqjqiowjdijqijdwodjwqiodjwqiodjwqoidjwqjwqijdqdowqjdqwiojdwqojdwqiojdqwd', 'https://example.com/image.jpg', 'pendek');
INSERT INTO `article` VALUES (15, 'woeeee', 4, '2024-02-05 12:25:06', '2024-02-05 12:25:06', 'Jane Austen’s Pride and Prejudice is an 18th-century novel of manners set in rural.wqidjqwoiwqdjoiwqjdioqwjdioqjqiowjdijqijdwodjwqiodjwqiodjwqoidjwqjwqijdqdowqjdqwiojdwqojdwqiojdqwd', 'https://example.com/image.jpg', 'pendek');

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of category
-- ----------------------------
INSERT INTO `category` VALUES (4, 'Pentil Jaran', '2024-02-05 09:24:24', '2024-02-05 09:24:24');

SET FOREIGN_KEY_CHECKS = 1;
