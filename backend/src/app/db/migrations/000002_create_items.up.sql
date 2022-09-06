BEGIN;
CREATE TABLE IF NOT EXISTS items(
    id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    category_id bigint(20) unsigned NOT NULL,
    title VARCHAR(255) UNIQUE NOT NULL,
    description text,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME DEFAULT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`category_id`) REFERENCES categories(id)
);
COMMIT;