--#####################################################USER
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
    workspace INTEGER REFERENCES WORKSPACE (id) ON DELETE CASCADE,
    member    INTEGER REFERENCES USERS (id) ON DELETE CASCADE,
    status    TEXT DEFAULT 'member'
);

--#####################################################WORKSPACE-TASK
CREATE TABLE IF NOT EXISTS WORKSPACE_TASK
(
    id          SERIAL PRIMARY KEY,
    workspace   INTEGER REFERENCES WORKSPACE (id) ON DELETE CASCADE,
    name        TEXT    NOT NULL,
    description TEXT,
    creator     INTEGER REFERENCES USERS (id) ON DELETE SET NULL
);

--#####################################################WORKSPACE-SCHEDULE
CREATE TABLE IF NOT EXISTS WORKSPACE_SCHEDULE
(
    id        SERIAL PRIMARY KEY,
    start     TIMESTAMP NOT NULL,
    workspace INTEGER REFERENCES WORKSPACE (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS WORKSPACE_SCHEDULE_RECORD
(
    id             SERIAL PRIMARY KEY,
    schedule       INTEGER REFERENCES WORKSPACE_SCHEDULE (id) ON DELETE CASCADE,
    description    TEXT    NOT NULL,
    start_datetime TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    end_datetime   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    task           INTEGER REFERENCES WORKSPACE_TASK (id) ON DELETE SET NULL
);

--#####################################################WORKSPACE-QUEUE
CREATE TABLE IF NOT EXISTS WORKSPACE_QUEUE
(
    id        SERIAL PRIMARY KEY,
    name      TEXT NOT NULL,
    workspace INTEGER REFERENCES WORKSPACE (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS WORKSPACE_QUEUE_MEMBER
(
    id     SERIAL PRIMARY KEY,
    queue  INTEGER REFERENCES WORKSPACE_QUEUE (id) ON DELETE CASCADE,
    member INTEGER REFERENCES USERS (id) ON DELETE CASCADE
);