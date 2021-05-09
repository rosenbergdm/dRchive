-- For sqlite3
CREATE TABLE files (
  filepath TEXT PRIMARY KEY ASC,
  mtime INTEGER,
  lastactive
  INTEGER,
  hash TEXT,
  VERSION INTEGER DEFAULT 1
);

