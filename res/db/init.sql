CREATE TABLE items (
	id 					INTEGER PRIMARY KEY AUTOINCREMENT,
	name 					TEXT UNIQUE NOT NULL,
	info 					TEXT,
	uses 					INTEGER,
	price 				INTEGER,
	is_consumable 		INTEGER NOT NULL DEFAULT 1,
	effect_type 		TEXT,
	effect_value 		INT,
	effect_duration	INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE weapon_types (
	name TEXT PRIMARY KEY
);

CREATE TABLE weapon_ranks (
	name TEXT PRIMARY KEY
);

CREATE TABLE weapons (
	item_id 				INTEGER PRIMARY KEY,
	weapon_type 		TEXT NOT NULL,
	weapon_rank 		TEXT NOT NULL,
	min_range 			INTEGER NOT NULL,
	max_range 			INTEGER NOT NULL,
	can_one_handed 	INTEGER NOT NULL,
	can_two_handed 	INTEGER NOT NULL,
	to_hit 				INTEGER NOT NULL,
	to_dam 				INTEGER NOT NULL,
	bonus_ac 			INTEGER NOT NULL,
	
	FOREIGN KEY (item_id) REFERENCES items(id),
	FOREIGN KEY (weapon_type) REFERENCES weapon_types(name),
	FOREIGN KEY (weapon_rank) REFERENCES weapon_ranks(name)
);

CREATE TABLE terrains (
	id 					INTEGER PRIMARY KEY AUTOINCREMENT,
	name 					TEXT UNIQUE NOT NULL,
	info 					TEXT,
	letter  				TEXT NOT NULL,
	movement_cost 		INTEGER NOT NULL,
	walkable 			INTEGER NOT NULL,
	diggable 			INTEGER NOT NULL,
	bonus_ac 			INTEGER NOT NULL
);

CREATE TABLE objects (
	id 					INTEGER PRIMARY KEY AUTOINCREMENT,
	name					TEXT UNIQUE NOT NULL,
	info  				TEXT,
	letter 				TEXT NOT NULL,
	walkable 			INTEGER NOT NULL,
	destroyable 		INTEGER NOT NULL,
	is_building 		INTEGER NOT NULL,
	is_locked 			INTEGER NOT NULL,
	bonus_ac				INTEGER NOT NULL
);

CREATE TABLE brands (
	id 					INTEGER PRIMARY KEY AUTOINCREMENT,
	name 					TEXT UNIQUE NOT NULL,
	info			 		TEXT,
	letter 				TEXT NOT NULL,
	bonus_ac				INTEGER NOT NULL,
	max_str				INTEGER NOT NULL,
	max_dex				INTEGER NOT NULL,
	max_con				INTEGER NOT NULL,
	max_int				INTEGER NOT NULL,
	max_wis				INTEGER NOT NULL,
	max_cha				INTEGER NOT NULL,
	bonus_str 			INTEGER NOT NULL,
	bonus_dex 			INTEGER NOT NULL,
	bonus_con 			INTEGER NOT NULL,
	bonus_int 			INTEGER NOT NULL,
	bonus_wis 			INTEGER NOT NULL,
	bonus_cha 			INTEGER NOT NULL
);

CREATE TABLE genders (
	name TEXT PRIMARY KEY
);

CREATE TABLE brand_prerequisites (
	brand_id 			INTEGER PRIMARY KEY,
	level  				INTEGER NOT NULL,
	gender 				TEXT,
	
	FOREIGN KEY (brand_id) REFERENCES brands(id),
	FOREIGN KEY (gender) REFERENCES genders(name)
);

CREATE TABLE brand_weapons (
	brand_id 			INTEGER,
	weapon_type 		TEXT NOT NULL,
	weapon_rank 		TEXT NOT NULL,
	
	PRIMARY KEY (brand_id, weapon_type), -- Each weapon can only have one rank
	FOREIGN KEY (brand_id) REFERENCES brands(id),
	FOREIGN KEY (weapon_type) REFERENCES weapon_types(name),
	FOREIGN KEY (weapon_rank) REFERENCES weapon_ranks(name)
);

CREATE TABLE factions (
	name TEXT PRIMARY KEY
);

CREATE TABLE units (
	id 					INTEGER PRIMARY KEY AUTOINCREMENT,
	brand_id 			INTEGER NOT NULL,
	name 					TEXT NOT NULL,
	letter 				TEXT NOT NULL,
	level 				INTEGER NOT NULL,
	hp 					INTEGER NOT NULL,
	max_hp 				INTEGER NOT NULL,
	exp 					INTEGER NOT NULL DEFAULT 0,
	gender 				TEXT NOT NULL,
	movement 			INTEGER NOT NULL,
	base_ac				INTEGER NOT NULL,
	faction 				TEXT NOT NULL,
	str 					INTEGER NOT NULL,
	dex 					INTEGER NOT NULL,
	con 					INTEGER NOT NULL,
	int 					INTEGER NOT NULL,
	wis 					INTEGER NOT NULL,
	cha 					INTEGER NOT NULL,
	bonus_str 			INTEGER NOT NULL,
	bonus_dex 			INTEGER NOT NULL,
	bonus_con 			INTEGER NOT NULL,
	bonus_int 			INTEGER NOT NULL,
	bonus_wis 			INTEGER NOT NULL,
	bonus_cha 			INTEGER NOT NULL,
	
	FOREIGN KEY (brand_id) REFERENCES brands(id),
	FOREIGN KEY (gender) REFERENCES genders(name),
	FOREIGN KEY (faction) REFERENCES factions(name)
);

CREATE TABLE unit_recuiters (
	recuit_id 			INTEGER,
	recuiter_id 		INTEGER,
	
	PRIMARY KEY (recuit_id, recuiter_id),
	FOREIGN KEY (recuit_id) REFERENCES units(id),
	FOREIGN KEY (recuiter_id) REFERENCES units(id)
);

CREATE TABLE unit_inventories (
	id 					INTEGER PRIMARY KEY AUTOINCREMENT,
	unit_id 				INTEGER NOT NULL,
	item_id 				INTEGER NOT NULL,
	index_num 			INTEGER NOT NULL,
	
	FOREIGN KEY (unit_id) REFERENCES units(id),
	FOREIGN KEY (item_id) REFERENCES items(id)
);