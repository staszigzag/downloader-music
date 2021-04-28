CREATE TABLE users
(
    id         int          NOT NULL PRIMARY KEY,
    first_name varchar(255) NOT NULL DEFAULT '',
    last_name  varchar(255) NOT NULL DEFAULT '',
    user_name  varchar(255) NOT NULL,
    chat_id    int          NOT NULL
);

CREATE TABLE audio
(
    id         serial PRIMARY KEY,
    video_id   varchar(255) NOT NULL UNIQUE,
    name       varchar(255) NOT NULL,
    path       varchar(255) NOT NULL,
    created_at timestamptz  NOT NULL DEFAULT (now())
);

CREATE TABLE users_audio
(
    id         serial PRIMARY KEY,
    user_id    int REFERENCES users (id) NOT NULL,
    audio_id   int REFERENCES audio (id) NOT NULL,
    created_at timestamptz               NOT NULL DEFAULT (now())
);

-- ALTER TABLE users_audio
--     ADD FOREIGN KEY (user_id) REFERENCES users (id);

CREATE
    INDEX ON audio (video_id);

