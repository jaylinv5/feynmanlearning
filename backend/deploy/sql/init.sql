-- 创建数据库
CREATE DATABASE IF NOT EXISTS feynman DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE feynman;

-- 用户表
CREATE TABLE IF NOT EXISTS `user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `username` varchar(50) NOT NULL,
  `password` varchar(100) NOT NULL,
  `real_name` varchar(50) DEFAULT NULL,
  `email` varchar(100) DEFAULT NULL,
  `phone` varchar(20) DEFAULT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `role` int DEFAULT '1' COMMENT '1-学生,2-教师,3-管理员',
  `grade` int DEFAULT NULL COMMENT '年级',
  `class` varchar(50) DEFAULT NULL COMMENT '班级',
  `school` varchar(100) DEFAULT NULL COMMENT '学校',
  `status` int DEFAULT '1' COMMENT '1-正常,2-禁用',
  `last_login_at` datetime DEFAULT NULL,
  `last_login_ip` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_username` (`username`),
  UNIQUE KEY `idx_email` (`email`),
  UNIQUE KEY `idx_phone` (`phone`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 知识点表
CREATE TABLE IF NOT EXISTS `knowledge_point` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `subject` varchar(20) NOT NULL COMMENT '学科:math/chinese/english',
  `grade` int NOT NULL COMMENT '年级',
  `chapter` varchar(100) NOT NULL COMMENT '章节名称',
  `chapter_order` int NOT NULL COMMENT '章节排序',
  `name` varchar(100) NOT NULL COMMENT '知识点名称',
  `code` varchar(50) NOT NULL COMMENT '知识点编码',
  `difficulty` int NOT NULL DEFAULT '2' COMMENT '难度:1-简单,2-中等,3-困难',
  `estimated_time` int NOT NULL DEFAULT '15' COMMENT '预计学习时长(分钟)',
  `content` text NOT NULL COMMENT '知识点讲解内容',
  `examples` text COMMENT '例题(JSON数组)',
  `exercises` text COMMENT '练习题(JSON数组)',
  `feynman_guide` text COMMENT '费曼引导问题(JSON数组)',
  `pre_requires` text COMMENT '前置知识点ID列表(JSON数组)',
  `status` int DEFAULT '1' COMMENT '1-草稿,2-审核中,3-上线,4-下线',
  `tags` varchar(255) DEFAULT NULL COMMENT '标签，逗号分隔',
  `created_by` bigint unsigned NOT NULL COMMENT '创建人ID',
  `updated_by` bigint unsigned DEFAULT NULL COMMENT '更新人ID',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_code` (`code`),
  KEY `idx_subject_grade` (`subject`,`grade`),
  KEY `idx_status` (`status`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 学习记录表
CREATE TABLE IF NOT EXISTS `learning_record` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `user_id` bigint unsigned NOT NULL,
  `knowledge_point_id` bigint unsigned NOT NULL,
  `status` int DEFAULT '1' COMMENT '1-学习中,2-已学完,3-已掌握',
  `study_time` int DEFAULT '0' COMMENT '学习时长(秒)',
  `progress` int DEFAULT '0' COMMENT '学习进度(0-100)',
  `first_study_at` datetime DEFAULT NULL,
  `last_study_at` datetime DEFAULT NULL,
  `completed_at` datetime DEFAULT NULL COMMENT '学完时间',
  `mastered_at` datetime DEFAULT NULL COMMENT '掌握时间',
  `total_attempts` int DEFAULT '0' COMMENT '费曼验证尝试次数',
  `best_score` int DEFAULT '0' COMMENT '费曼验证最高得分',
  `last_score` int DEFAULT '0' COMMENT '最近一次费曼验证得分',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_knowledge` (`user_id`,`knowledge_point_id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_knowledge_point_id` (`knowledge_point_id`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 学习会话表
CREATE TABLE IF NOT EXISTS `learning_session` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `user_id` bigint unsigned NOT NULL,
  `knowledge_point_id` bigint unsigned NOT NULL,
  `session_type` int NOT NULL COMMENT '1-学习模式,2-费曼教学模式',
  `status` int DEFAULT '1' COMMENT '1-进行中,2-已完成,3-已中断',
  `start_time` datetime NOT NULL,
  `end_time` datetime DEFAULT NULL,
  `duration` int DEFAULT '0' COMMENT '会话时长(秒)',
  `message_count` int DEFAULT '0' COMMENT '交互消息数量',
  `score` int DEFAULT '0' COMMENT '会话得分',
  `evaluation` text COMMENT '会话评估结果',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_knowledge_point_id` (`knowledge_point_id`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 学习消息表
CREATE TABLE IF NOT EXISTS `learning_message` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `session_id` bigint unsigned NOT NULL,
  `user_id` bigint unsigned NOT NULL,
  `role` int NOT NULL COMMENT '1-用户,2-AI',
  `content` text NOT NULL,
  `message_type` int DEFAULT '1' COMMENT '1-文本,2-图片,3-语音,4-文件',
  `meta_data` text COMMENT '元数据(JSON)',
  `is_evaluated` tinyint(1) DEFAULT '0' COMMENT '是否已评估',
  PRIMARY KEY (`id`),
  KEY `idx_session_id` (`session_id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 练习题记录表
CREATE TABLE IF NOT EXISTS `exercise_record` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `user_id` bigint unsigned NOT NULL,
  `knowledge_point_id` bigint unsigned NOT NULL,
  `exercise_id` varchar(50) NOT NULL,
  `user_answer` text,
  `correct_answer` text NOT NULL,
  `is_correct` tinyint(1) DEFAULT '0',
  `score` int DEFAULT '0',
  `time_spent` int DEFAULT '0' COMMENT '用时(秒)',
  `answered_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_knowledge_point_id` (`knowledge_point_id`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 费曼教学记录表
CREATE TABLE IF NOT EXISTS `feynman_record` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `user_id` bigint unsigned NOT NULL,
  `knowledge_point_id` bigint unsigned NOT NULL,
  `session_id` bigint unsigned NOT NULL,
  `score` int NOT NULL COMMENT '总得分(0-100)',
  `accuracy_score` int NOT NULL COMMENT '准确性得分(0-40)',
  `completeness_score` int NOT NULL COMMENT '完整性得分(0-30)',
  `logic_score` int NOT NULL COMMENT '逻辑性得分(0-20)',
  `expression_score` int NOT NULL COMMENT '表达能力得分(0-10)',
  `evaluation` text NOT NULL COMMENT '详细评估',
  `suggestions` text COMMENT '改进建议',
  `is_passed` tinyint(1) DEFAULT '0' COMMENT '是否通过',
  `weak_points` text COMMENT '薄弱点(JSON数组)',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_session_id` (`session_id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_knowledge_point_id` (`knowledge_point_id`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 插入默认管理员用户
INSERT INTO `user` (`username`, `password`, `real_name`, `role`, `status`) 
VALUES ('admin', '$2a$10$4JcH6/.tWX4q3kO7hH8k9eJf5dD7cC6bB5aA4sS3dD2fF1eE0rR9tT', '系统管理员', 3, 1);
-- 密码: admin123
