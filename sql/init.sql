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

