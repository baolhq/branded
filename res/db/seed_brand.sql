INSERT INTO brands (name, info, letter, bonus_ac, 
	max_str, max_dex, max_con, max_int, max_wis, max_cha, 
	bonus_str, bonus_dex, bonus_con, bonus_wis, bonus_int, bonus_cha) VALUES
('Fighter', 'Balanced warrior skilled in physical combat', 'f', 4, 22, 20, 20, 15, 12, 12, 20, 10, 10, 0, 0, 10),
('Champion', 'Elite warrior with strong leadership and presence', 'F', 4, 45, 43, 46, 30, 30, 40, 20, 10, 10, 0, 0, 10),
('Hero', 'A skilled fighter who excels in all aspects of combat', 'F', 4, 45, 43, 46, 30, 30, 40, 20, 10, 10, 0, 0, 10),

('Cavalier', 'A mounted knight with fast movement and attack', 'c', 4, 22, 22, 18, 15, 16, 18, 20, 10, 10, 0, 0, 10),
('Paladin', 'A noble knight with both offensive and healing powers', 'C', 4, 40, 42, 45, 35, 30, 40, 20, 10, 10, 0, 0, 10),
('Dread Knight', 'A dark knight that combines power and fear', 'C', 4, 40, 40, 45, 30, 28, 38, 20, 10, 10, 0, 0, 10),

('Nomad', 'A quick and mobile archer on horseback', 'n', 4, 20, 22, 18, 15, 16, 14, 20, 10, 10, 0, 0, 10),
('Bow Knight', 'A cavalry archer with enhanced range and mobility', 'N', 4, 25, 30, 20, 15, 18, 22, 20, 10, 10, 0, 0, 10),
('Strider', 'A highly mobile archer who excels in hit and run tactics', 'N', 4, 25, 30, 20, 15, 18, 22, 20, 10, 10, 0, 0, 10),

('Pegasus Knight', 'A flying knight with speed and dexterity', 'p', 4, 20, 22, 18, 15, 16, 18, 20, 10, 10, 0, 0, 10),
('Falcon Knight', 'A high-flying knight with enhanced mobility and power', 'P', 4, 35, 38, 35, 20, 22, 25, 20, 10, 10, 0, 0, 10),
('Seraph Knight', 'A graceful flying knight with magical abilities', 'P', 4, 32, 35, 35, 25, 22, 20, 20, 10, 10, 0, 0, 10),

('Wyvern Rider', 'A powerful rider on a wyvern with high attack', 'w', 4, 22, 20, 20, 15, 16, 12, 20, 10, 10, 0, 0, 10),
('Wyvern Lord', 'A wyvern-mounted warrior with great strength and speed', 'W', 4, 40, 42, 46, 25, 20, 30, 20, 10, 10, 0, 0, 10),
('Dragon Rider', 'A formidable rider with control over dragons', 'W', 4, 42, 40, 45, 25, 22, 30, 20, 10, 10, 0, 0, 10),

('Mercenary', 'A versatile soldier skilled in both attack and defense', 'm', 4, 22, 20, 20, 15, 16, 16, 20, 10, 10, 0, 0, 10),
('Duelist', 'A swift fighter focused on one-on-one combat', 'M', 4, 30, 40, 25, 20, 22, 15, 20, 10, 10, 0, 0, 10),
('Vanguard', 'A frontline soldier who excels at leading charges', 'M', 4, 35, 38, 40, 25, 20, 30, 20, 10, 10, 0, 0, 10),

('Spellblade', 'A swordfighter who blends magic and physical attacks', 's', 4, 22, 22, 20, 18, 18, 16, 20, 10, 10, 0, 0, 10),
('Hexblade', 'A warrior who channels curses and dark magic into combat', 'S', 4, 30, 35, 40, 30, 22, 30, 20, 10, 10, 0, 0, 10),
('Battle Mage', 'A powerful mage who uses both spells and weapons', 'S', 4, 35, 32, 35, 40, 30, 30, 20, 10, 10, 0, 0, 10),

('Barbarian', 'A fierce warrior who thrives in reckless combat', 'b', 4, 22, 18, 25, 15, 16, 12, 20, 10, 10, 0, 0, 10),
('Revenger', 'A berserker who seeks vengeance through powerful attacks', 'B', 4, 45, 43, 46, 30, 28, 40, 20, 10, 10, 0, 0, 10),
('Berserker', 'A relentless warrior whose power grows with rage', 'B', 4, 45, 43, 46, 30, 28, 40, 20, 10, 10, 0, 0, 10),

('Archer', 'A ranged fighter specializing in bows and crossbows', 'a', 4, 22, 24, 18, 15, 16, 14, 20, 10, 10, 0, 0, 10),
('Marksman', 'An expert archer with exceptional aim and precision', 'A', 4, 35, 40, 25, 20, 22, 25, 20, 10, 10, 0, 0, 10),
('Sniper', 'A master of long-range combat, able to hit from great distances', 'A', 4, 35, 40, 30, 20, 25, 20, 20, 10, 10, 0, 0, 10),

('Rogue', 'A stealthy fighter skilled in surprise attacks and agility', 'r', 4, 22, 24, 18, 15, 16, 14, 20, 10, 10, 0, 0, 10),
('Assassin', 'A deadly assassin focused on swift, lethal strikes', 'R', 4, 35, 40, 30, 20, 25, 20, 20, 10, 10, 0, 0, 10),
('Sharpshooter', 'A long-range expert who specializes in pinpoint accuracy', 'R', 4, 35, 40, 30, 20, 25, 20, 20, 10, 10, 0, 0, 10),

('Mage', 'A spellcaster who wields powerful elemental magic', 'g', 4, 20, 18, 18, 40, 35, 30, 20, 10, 10, 0, 0, 10),
('Archmage', 'A master of magic who commands powerful spells', 'G', 4, 35, 30, 30, 45, 40, 35, 20, 10, 10, 0, 0, 10),
('Pyromancer', 'A mage who specializes in destructive fire magic', 'G', 4, 35, 32, 35, 45, 40, 35, 20, 10, 10, 0, 0, 10),

('Dark Mage', 'A mage who channels dark and forbidden magic', 'd', 4, 22, 18, 18, 40, 30, 30, 20, 10, 10, 0, 0, 10),
('Warlock', 'A sorcerer who commands dark magic and curses', 'D', 4, 35, 30, 35, 45, 38, 35, 20, 10, 10, 0, 0, 10),
('Necromancer', 'A master of necromantic magic and the dead', 'D', 4, 35, 30, 35, 45, 38, 35, 20, 10, 10, 0, 0, 10),

('Priest', 'A holy spellcaster who heals and supports allies', 'i', 4, 22, 18, 18, 40, 30, 30, 20, 10, 10, 0, 0, 10),
('Cleric', 'A divine healer who also wields powerful restorative magic', 'I', 4, 35, 30, 35, 45, 38, 35, 20, 10, 10, 0, 0, 10),
('Bishop', 'A divine caster who enhances healing and protection magic', 'I', 4, 35, 30, 35, 45, 38, 35, 20, 10, 10, 0, 0, 10),

('Templar', 'A holy knight who combines combat prowess and healing', 't', 4, 22, 18, 18, 40, 30, 30, 20, 10, 10, 0, 0, 10),
('Crusader', 'A holy warrior on a divine mission, wielding powerful magic', 'T', 4, 35, 30, 35, 45, 38, 35, 20, 10, 10, 0, 0, 10),
('Inquisitor', 'A fierce crusader with a relentless pursuit of justice', 'T', 35, 30, 35, 45, 38, 35, 20, 10, 10, 0, 0, 10),

('Lord', 'The noble leader, with balanced strength and leadership', '@', 4, 20, 20, 20, 15, 16, 16, 20, 10, 10, 0, 0, 10),
('Great Lord', 'A legendary leader with extraordinary strength and authority', '@', 4, 45, 43, 46, 30, 30, 40, 20, 10, 10, 0, 0, 10);