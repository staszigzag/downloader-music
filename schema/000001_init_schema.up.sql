CREATE TABLE users
(
    id      serial PRIMARY KEY,
    name    varchar(255) NOT NULL,
    chat_id int          NOT NULL
);

CREATE TABLE audio
(
    id         serial PRIMARY KEY,
    video_id   varchar(255) NOT NULL,
    name       varchar(255) NOT NULL,
    path       varchar(255) NOT NULL,
    created_at timestamptz  NOT NULL DEFAULT (now())
);

CREATE TABLE users_audio
(
    id       serial PRIMARY KEY,
    user_id  int REFERENCES users (id) NOT NULL,
    audio_id int REFERENCES audio (id) NOT NULL
);

-- ALTER TABLE users_audio
--     ADD FOREIGN KEY (user_id) REFERENCES users (id);

CREATE
    INDEX ON audio (video_id);

