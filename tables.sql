CREATE TABLE players (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    nickname VARCHAR(255),
    nationality VARCHAR(255),
    photo_url VARCHAR(255),
    base_rating DECIMAL(10, 2),
    role VARCHAR(255),
    fatigue_bps DECIMAL(10, 2),
    morale_bps DECIMAL(10, 2),
    happiness_bps DECIMAL(10, 2),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE teams (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255),
    logo_url VARCHAR(255),
    country VARCHAR(255),
    tag VARCHAR(255),
    description VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE team_players (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    team_id INTEGER,
    player_id INTEGER,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
    FOREIGN KEY (team_id) REFERENCES teams(id),
    FOREIGN KEY (player_id) REFERENCES players(id),
    UNIQUE (team_id, player_id)
);