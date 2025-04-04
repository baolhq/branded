package data

func UseItem(u *Unit, i *Item) {
	if i.EffectDuration > 0 {
		// Short-term buff with countdown
		ApplyShortTermBuff(u, i)
	} else {
		// Permanent buff (no countdown)
		ApplyPermanentBuff(u, i)
	}
}

func ApplyShortTermBuff(u *Unit, i *Item) {
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

func ApplyPermanentBuff(u *Unit, i *Item) {
	switch i.EffectType {
	case "heal_hp":
		u.Hp += i.EffectValue
		if u.Hp > u.MaxHp {
			u.Hp = u.MaxHp
		}
	case "lighten":
		LightenSurrounding(u.PosX, u.PosY, i.EffectValue)
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

func LightenSurrounding(x, y, r int) {
	// TODO: Implement fog of war logic
}
