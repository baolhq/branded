INSERT INTO genders (name) VALUES ('Unknown'), ('Male'), ('Female');

INSERT INTO factions (name) VALUES ('Other'), ('Party'), ('Enemy'), ('Ally');

INSERT INTO weapon_types (name) VALUES ('Sword'), ('Axe'), ('Lance'), ('Staff'), ('Bow'), ('Crossbow'), ('Arcane'), ('Divine'), ('Shadow');

INSERT INTO weapon_ranks (name) VALUES ('D'), ('C'), ('B'), ('A'), ('S');

INSERT INTO terrains (name, info, letter, movement_cost, walkable, diggable, bonus_ac) VALUES
('Plain', 'Flat land with no cover', '.', 1, 1, 1, 0),
('Water', 'Deep water, slows movement', '~', 3, 1, 0, 2),
('Lava', 'Molten rock, nearly impassable', '%', 99, 0, 0, 4),
('Snow', 'Slippery terrain, hinders movement', ',', 2, 1, 1, 0),
('Forest', 'Dense woods, provides cover', '^', 2, 1, 1, 2),
('Mountain', 'Steep terrain, highly defensive', '`', 3, 1, 1, 3),
('Desert', 'Harsh land, minor cover', '*', 2, 1, 1, 1);

INSERT INTO objects (name, info, letter, walkable, destroyable, is_building, is_locked, bonus_ac) VALUES
('sink', '', '#', 0, 1, 0, 0, 0),
('table', 'Titled table, provide cover', '#', 1, 1, 0, 0, 1),
('chest', 'Chest filled with treasure', '-', 0, 1, 0, 1, 0),
('gold', 'Rare piece of gold', '$', 1, 0, 0, 0, 0);