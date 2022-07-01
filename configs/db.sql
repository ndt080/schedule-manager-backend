CREATE TABLE IF NOT EXISTS USERS
(
    id            SERIAL PRIMARY KEY,
    username      TEXT NOT NULL,
    email         TEXT NOT NULL,
    password_hash TEXT NOT NULL,
    is_verified   BOOL DEFAULT false,
    image         TEXT DEFAULT 'https://i.ibb.co/zm3YCP3/avatar.png',
    status        TEXT DEFAULT 'user'
);

--#####################################################WORKSPACE
CREATE TABLE IF NOT EXISTS WORKSPACE
(
    id          SERIAL PRIMARY KEY,
    name        TEXT NOT NULL,
    description TEXT,
    image       TEXT DEFAULT 'https://i.ibb.co/JmYr3ys/5.png',
    created_at  DATE DEFAULT CURRENT_DATE
);

CREATE TABLE IF NOT EXISTS WORKSPACE_MEMBER
(
    id        SERIAL PRIMARY KEY,
    workspace INTEGER REFERENCES WORKSPACE (id),
    member    INTEGER REFERENCES USERS (id),
    status    TEXT DEFAULT 'member'
);

--#####################################################TASK
CREATE TABLE IF NOT EXISTS WORKSPACE_TASK
(
    id          SERIAL PRIMARY KEY,
    workspace   INTEGER REFERENCES WORKSPACE (id),
    name        TEXT NOT NULL,
    description TEXT,
    creator     INTEGER REFERENCES USERS (id)
);

--#####################################################SCHEDULE
CREATE TABLE IF NOT EXISTS WORKSPACE_SCHEDULE
(
    id        SERIAL PRIMARY KEY,
    start     TIMESTAMP NOT NULL,
    workspace INTEGER REFERENCES WORKSPACE (id)
);

CREATE TABLE IF NOT EXISTS WORKSPACE_SCHEDULE_RECORD
(
    id             SERIAL PRIMARY KEY,
    schedule       INTEGER REFERENCES WORKSPACE_SCHEDULE (id),
    description    TEXT NOT NULL,
    start_datetime TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    end_datetime   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    task           INTEGER REFERENCES WORKSPACE_TASK (id)
);

--#####################################################QUEUE
CREATE TABLE IF NOT EXISTS WORKSPACE_QUEUE
(
    id        SERIAL PRIMARY KEY,
    name      TEXT NOT NULL,
    workspace INTEGER REFERENCES WORKSPACE (id)
);

CREATE TABLE IF NOT EXISTS WORKSPACE_QUEUE_MEMBER
(
    id     SERIAL PRIMARY KEY,
    queue  INTEGER REFERENCES WORKSPACE_QUEUE (id),
    member INTEGER REFERENCES USERS (id)
);