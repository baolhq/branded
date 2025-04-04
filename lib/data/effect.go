package data

func UseItem(u *Unit, i *Item, c *Chapter) {
	if i.EffectDuration > 0 {
		// Short-term buff with countdown
		ApplyShortTermBuff(u, i, c)
	} else {
		// Permanent buff or consumables (no countdown)
		ApplyPermanentBuff(u, i, c)
	}
}

func ApplyShortTermBuff(u *Unit, i *Item, c *Chapter) {
	switch i.EffectType {
	case "boost_str":
		u.Str += i.EffectValue
		u.ActiveEffects["boost_str"] = i.EffectDuration
	case "boost_dex":
		u.Dex += i.EffectValue
		u.ActiveEffects["boost_dex"] = i.EffectDuration
	case "boost_con":
		u.Con += i.EffectValue
		u.ActiveEffects["boost_con"] = i.EffectDuration
	case "boost_int":
		u.Int += i.EffectValue
		u.ActiveEffects["boost_int"] = i.EffectDuration
	case "boost_wis":
		u.Wis += i.EffectValue
		u.ActiveEffects["boost_wis"] = i.EffectDuration
	case "boost_cha":
		u.Cha += i.EffectValue
		u.ActiveEffects["boost_cha"] = i.EffectDuration
	case "boost_ac":
		u.BaseAc += i.EffectValue
		u.ActiveEffects["boost_ac"] = i.EffectDuration
	}
}

func ApplyPermanentBuff(u *Unit, i *Item, c *Chapter) {
	switch i.EffectType {
	case "heal_hp":
		u.Hp += i.EffectValue
		if u.Hp > u.MaxHp {
			u.Hp = u.MaxHp
		}
	case "lighten":
		LightenSurrounding(u.PosX, u.PosY, i.EffectValue, c)
	case "boost_str":
		u.Str += i.EffectValue
	case "boost_dex":
		u.Dex += i.EffectValue
	case "boost_con":
		u.Con += i.EffectValue
	case "boost_int":
		u.Int += i.EffectValue
	case "boost_wis":
		u.Wis += i.EffectValue
	case "boost_cha":
		u.Cha += i.EffectValue
	case "boost_ac":
		u.BaseAc += i.EffectValue
	}
}

func LightenSurrounding(x, y, r int, c *Chapter) {
	// Lighten the surrounding area within radius r
	for dx := -r; dx <= r; dx++ {
		for dy := -r; dy <= r; dy++ {
			// Calculate the distance from the center
			if dx*dx+dy*dy <= r*r {
				// Ensure the coordinates are within map bounds
				nx, ny := x+dx, y+dy
				if nx >= 0 && ny >= 0 && nx < len(c.Map) && ny < len(c.Map[0]) {
					c.Map[nx][ny].Visible = true
				}
			}
		}
	}
}
