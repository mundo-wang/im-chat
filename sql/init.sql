-- noinspection SqlDialectInspectionForFile

-- noinspection SqlNoDataSourceInspectionForFile

CREATE TABLE users (
    id int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name varchar(100) NOT NULL,
    password varchar(255) NOT NULL,
    phone varchar(20) NOT NULL,
    email varchar(100) NOT NULL,
    identity varchar(50),
    client_info varchar(255) DEFAULT NULL,
    login_time datetime DEFAULT NULL,
    heartbeat_time datetime DEFAULT NULL,
    logout_time datetime DEFAULT NULL,
    is_logout tinyint(1) NOT NULL DEFAULT 0,
    device_info varchar(255) DEFAULT NULL,
    avatar varchar(255) DEFAULT NULL,
    salt varchar(255) DEFAULT NULL,
    created_at datetime DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE communities (
    id int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name varchar(100) NOT NULL,
    owner_id int unsigned NOT NULL,
    img varchar(255) DEFAULT NULL,
    description text DEFAULT NULL,
    created_at datetime DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE user_groups (
    id int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name varchar(100) NOT NULL,
    owner_id int unsigned NOT NULL,
    icon varchar(255) DEFAULT NULL,
    type tinyint NOT NULL DEFAULT 0,
    description text DEFAULT NULL,
    created_at datetime DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE contacts (
    id int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    owner_id int unsigned NOT NULL,
    target_id int unsigned NOT NULL,
    type tinyint NOT NULL DEFAULT 0,
    description text DEFAULT NULL,
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

