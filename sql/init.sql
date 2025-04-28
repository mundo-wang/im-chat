-- noinspection SqlDialectInspectionForFile

-- noinspection SqlNoDataSourceInspectionForFile

CREATE TABLE users (
    id int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_code CHAR(10) NOT NULL,
    name varchar(100) NOT NULL,
    password varchar(255) NOT NULL,
    salt varchar(255) DEFAULT NULL,
    phone varchar(20) NOT NULL,
    email varchar(100) NOT NULL,
    avatar varchar(255) DEFAULT NULL,
    login_time datetime DEFAULT NULL,
    heartbeat_time datetime DEFAULT NULL,
    created_at datetime DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE communities (
    id int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    community_code CHAR(10) NOT NULL,
    name varchar(100) NOT NULL,
    owner_id int unsigned NOT NULL,
    avatar varchar(255) DEFAULT NULL,
    type tinyint NOT NULL DEFAULT 0 COMMENT '0.默认 1.兴趣爱好 2.行业交流 3.生活休闲',
    description text DEFAULT NULL,
    created_at datetime DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE contacts (
    id int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    owner_id int unsigned NOT NULL,
    target_id int unsigned NOT NULL COMMENT '对应的人/群ID',
    type tinyint NOT NULL COMMENT '1为好友，2为群组',
    created_at datetime DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE messages (
    id int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    form_id int unsigned NOT NULL,
    target_id int unsigned NOT NULL,
    type tinyint NOT NULL,
    content text DEFAULT NULL,
    pic varchar(255) DEFAULT NULL,
    url varchar(255) DEFAULT NULL,
    description varchar(255) DEFAULT NULL,
    created_at datetime DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `questions` (
  `id` INT AUTO_INCREMENT PRIMARY KEY COMMENT '主键ID',
  `title` VARCHAR(100) NOT NULL COMMENT '题目内容',
  `type` INT NOT NULL COMMENT '题型（0=选择题，1=判断题）',
  `answer` VARCHAR(100) NOT NULL COMMENT '正确答案',
  `status` INT NOT NULL DEFAULT 0 COMMENT '题目状态（0=未发布，1=已发布）',
  `agent_code` VARCHAR(100) NOT NULL COMMENT '智能体编码',
  `position_id` INT NOT NULL COMMENT '岗位名称',
  `session_ref_id` VARCHAR(100) NOT NULL COMMENT '关联的生成题目批次ID',
  `created_by` VARCHAR(100) NOT NULL COMMENT '创建人',
  `updated_by` VARCHAR(100) DEFAULT NULL COMMENT '修改人',
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_deleted` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '逻辑删除标记（0=未删除，1=已删除）'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='试题表';

CREATE TABLE `question_options` (
  `id` INT AUTO_INCREMENT PRIMARY KEY COMMENT '主键ID',
  `question_id` INT NOT NULL COMMENT '题目ID',
  `option_key` CHAR(1) NOT NULL COMMENT '选项标识（如A、B、C、D）',
  `option_text` VARCHAR(100) NOT NULL COMMENT '选项内容'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='试题选项表';

CREATE TABLE `question_session` (
  `id` INT AUTO_INCREMENT PRIMARY KEY COMMENT '主键ID',
  `session_id` VARCHAR(100) NOT NULL COMMENT 'Session唯一标识',
  `user_id` INT NOT NULL COMMENT '用户ID',
  `username` VARCHAR(100) NOT NULL COMMENT '用户名',
  `agent_code` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '智能体编码',
  `position_id` INT NOT NULL COMMENT '岗位id',
  `status` INT NOT NULL DEFAULT 0 COMMENT '生成题目状态（1=生成中，2=已生成，3=生成失败，4=已提交）',
  `created_by` varchar(100) NOT NULL COMMENT '创建人',
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='生成题目批次表';

CREATE TABLE `exam_records` (
  `id` INT AUTO_INCREMENT PRIMARY KEY COMMENT '考试记录ID',
  `user_id` INT NOT NULL COMMENT '用户ID',
  `position_id` INT NOT NULL COMMENT '岗位id',
  `score` INT DEFAULT 0 COMMENT '总得分',
  `remark` VARCHAR(500) COMMENT '考试评语',
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '提交时间',
  `is_deleted` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '逻辑删除标记（0=未删除，1=已删除）'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='考试记录表';

CREATE TABLE `answer_records` (
  `id` INT AUTO_INCREMENT PRIMARY KEY COMMENT '答题记录ID',
  `exam_id` INT NOT NULL COMMENT '考试ID',
  `question_no` INT NOT NULL COMMENT '题目编号',
  `question_id` INT NOT NULL COMMENT '题目ID',
  `user_answer` VARCHAR(100) NOT NULL COMMENT '用户答案',
  `correct_answer` VARCHAR(100) NOT NULL COMMENT '正确答案',
  `is_correct` BOOLEAN DEFAULT FALSE COMMENT '是否答对'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='答题详情记录表';