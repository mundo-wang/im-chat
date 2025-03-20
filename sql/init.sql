CREATE TABLE communities (
    id int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name varchar(255),
    owner_id int unsigned DEFAULT NULL,
    img varchar(255),
    description varchar(255),
    created_at datetime DEFAULT NULL,
    updated_at datetime DEFAULT NULL,
    deleted_at datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE contact (
    id int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    owner_id int unsigned DEFAULT NULL,
    target_id int unsigned DEFAULT NULL,
    type int DEFAULT NULL,
    description varchar(255),
    created_at datetime DEFAULT NULL,
    updated_at datetime DEFAULT NULL,
    deleted_at datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE group_basic (
    id int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name varchar(255),
    owner_id int unsigned DEFAULT NULL,
    icon varchar(255),
    type int DEFAULT NULL,
    description varchar(255),
    created_at datetime DEFAULT NULL,
    updated_at datetime DEFAULT NULL,
    deleted_at datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE message (
    id int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    form_id varchar(255),
    target_id varchar(255),
    type varchar(255),
    media int DEFAULT NULL,
    content varchar(255),
    pic varchar(255),
    url varchar(255),
    description varchar(255),
    amount int DEFAULT NULL,
    created_at datetime DEFAULT NULL,
    updated_at datetime DEFAULT NULL,
    deleted_at datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE user_basic (
    id int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name varchar(255),
    pass_word varchar(255),
    phone varchar(255),
    email varchar(255),
    identity varchar(255),
    client_ip varchar(255),
    client_port varchar(255),
    login_time datetime DEFAULT NULL,
    heartbeat_time datetime DEFAULT NULL,
    login_out_time datetime DEFAULT NULL,
    is_logout tinyint(1) DEFAULT NULL,
    device_info varchar(255),
    salt varchar(255),
    avatar varchar(255) DEFAULT NULL,
    created_at datetime DEFAULT NULL,
    updated_at datetime DEFAULT NULL,
    deleted_at datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
