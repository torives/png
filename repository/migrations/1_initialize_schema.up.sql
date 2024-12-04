CREATE TABLE IF NOT EXISTS teams(
    id INTEGER PRIMARY KEY NOT NULL,
    name VARCHAR(3) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS work_types(
    id INTEGER PRIMARY KEY NOT NULL, 
    name VARCHAR(2) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS projects(
	id INTEGER NOT NULL,
	team TEXT NOT NULL,
	work_type TEXT NOT NULL,
	PRIMARY KEY(id, team, work_type),
	FOREIGN KEY(team) REFERENCES teams(name),
	FOREIGN KEY(work_type) REFERENCES work_types(name)
);

CREATE INDEX IF NOT EXISTS team_index ON projects(team);
	
CREATE INDEX IF NOT EXISTS work_type_index ON projects(team);