CREATE TABLE IF NOT EXISTS users (
    id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(32) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    nickname VARCHAR(64) NOT NULL,
    avatar_url VARCHAR(255) DEFAULT '',
    bio VARCHAR(255) DEFAULT '',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY uk_username (username)
);

CREATE TABLE IF NOT EXISTS videos (
    id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    author_id BIGINT UNSIGNED NOT NULL,
    title VARCHAR(128) NOT NULL,
    description VARCHAR(500) DEFAULT NULL,
    video_url VARCHAR(255) NOT NULL,
    cover_url VARCHAR(255) DEFAULT NULL,
    like_count BIGINT UNSIGNED NOT NULL DEFAULT 0,
    comment_count BIGINT UNSIGNED NOT NULL DEFAULT 0,
    status TINYINT UNSIGNED NOT NULL DEFAULT 1,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    KEY idx_author_id (author_id),
    KEY idx_created_at (created_at),
    KEY idx_author_status_created (author_id, status, created_at)
);

CREATE TABLE IF NOT EXISTS follows (
    id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT UNSIGNED NOT NULL,
    follow_user_id BIGINT UNSIGNED NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY uk_user_follow (user_id, follow_user_id),
    KEY idx_follow_user_id (follow_user_id)
);

CREATE TABLE IF NOT EXISTS video_likes (
    id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT UNSIGNED NOT NULL,
    video_id BIGINT UNSIGNED NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY uk_user_video (user_id, video_id),
    KEY idx_video_id (video_id)
);

CREATE TABLE IF NOT EXISTS comments (
    id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    video_id BIGINT UNSIGNED NOT NULL,
    user_id BIGINT UNSIGNED NOT NULL,
    content VARCHAR(500) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    KEY idx_video_created (video_id, created_at),
    KEY idx_user_id (user_id)
);