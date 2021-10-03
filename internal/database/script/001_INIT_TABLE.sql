PRAGMA foreign_keys = 1;

CREATE TABLE IF NOT EXISTS Inventory (
    id INTEGER PRIMARY KEY,
    name VARCHAR NOT NULL,
    price INTEGER NOT NULL DEFAULT 0,
    stock INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS Userdata (
    id INTEGER PRIMARY KEY,
    username VARCHAR UNIQUE NOT NULL,
    credit INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS Usercart (
    order_id INTEGER PRIMARY KEY AUTOINCREMENT,
    member_id,
    item_id,
    sum INTEGER NOT NULL DEFAULT 1,
    FOREIGN KEY(member_id) REFERENCES Userdata(id) ON DELETE CASCADE,
    FOREIGN KEY(item_id) REFERENCES Inventory(id) ON DELETE CASCADE
);