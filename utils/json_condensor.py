import json
from pprint import pprint
import jsonpickle


class Weapon(object):
	def __init__(self, name: str, game: str, rpm: int, bv: int, dp, loady: list):
		"""Constructor"""
		self.weapon_name = name
		self.game_from = game
		self.rpm = rpm
		self.bullet_velocity = int(bv)
		self.damage_profile = dp
		self.loadouts = loady

	def __repr__(self):
		return "Weapon Name: {} From Game: {} RPM: {} Bullet Velocity: {} Close Range Damage Profile: {} " \
			   "Mid-range Damage Profile: {} Loadouts: {}".format(
			str(self.weapon_name), str(self.game_from), str(self.rpm), str(self.bullet_velocity),
			self.close_dmg_profile, self.mid_dmg_profile, self.far_dmg_profile, self.loadouts
		)

	def __str__(self):
		return "Weapon Name: {} From Game: {} RPM: {} Bullet Velocity: {} Close Range Damage Profile: {} " \
			   "Mid-range Damage Profile: {} Loadouts: {}".format(
			str(self.weapon_name), str(self.game_from), str(self.rpm), str(self.bullet_velocity),
			self.close_dmg_profile, self.mid_dmg_profile, self.far_dmg_profile, self.loadouts
		)


class DamageProfile(object):
	def __init__(self, close_range_dp, mid_range_dp, far_range_dp):
		self.CloseRange = close_range_dp
		self.MidRange = mid_range_dp
		self.FarRange = far_range_dp

	def __repr__(self):
		return "Minimum Distance: {} Maximum Distance: {} Minimum Shots to Kill: {} Maximum Shots to Kill: {} Minimum " \
			   "Time to Kill: {} Maximum Time to Kill: {}".format(
			self.mindistance, self.maxdistance, self.minstk, self.maxstk,
			self.minttk, self.maxstk
		)


class Range(DamageProfile, object):
	def __init__(self, min_dist, max_dist, min_stk, max_stk, min_ttk, max_ttk):
		self.mindistance = min_dist
		self.maxdistance = max_dist
		self.minstk = min_stk
		self.maxstk = max_stk
		self.minttk = min_ttk
		self.maxttk = max_ttk



class Loadout(object):
	def __init__(self, wep, cat, muzzle, barrel, laser, optic, stock, underbarrel, ammo, reargrip, perk):
		self.weapon = wep
		self.category = cat
		self.muzzle = muzzle
		self.barrel = barrel
		self.laser = laser
		self.optic = optic
		self.stock = stock
		self.underbarrel = underbarrel
		self.ammo = ammo
		self.rear_grip = reargrip
		self.perk = perk

	def __repr__(self):
		return "Category: {} Muzzle: {} Barrel: {} Laser: {} Optic: {} Stock: {} Underbarrel: {} Ammo: {} Rear Grip: " \
			   "{}  Perk: {}".format(
			self.category, self.muzzle, self.barrel, self.laser,
			self.optic, self.stock, self.underbarrel, self.ammo, self.rear_grip, self.perk
		)


if __name__ == "__main__":
	condensed_weapons = {}

	with open('dmgprofiles.json') as in_file:
		weapons = json.load(in_file)

	for weapon in weapons:
		w = Weapon(
			str(weapon['weaponname']).upper(),
			str(weapon['gamefrom']),
			int(weapon['rpm']),
			int(weapon['bulletvelocity']),
			DamageProfile(
				# close range
				Range(
					0,
					weapon['MaxDist Close'],
					weapon['STK Min Close'],
					weapon['STK Max Close'],
					weapon['TTK Min Close'],
					weapon['TTK Max Close']
				),
				# mid range
				Range(
					weapon['MinDist Mid'],
					weapon['MaxDist Mid'],
					weapon['STK Min Mid'],
					weapon['STK Max Mid'],
					weapon['TTK Min Mid'],
					weapon['TTK Max Mid']
				),
				# far range
				Range(
					weapon['MinDist Far'],
					0,
					weapon['STK Min Far'],
					weapon['STK Max Far'],
					weapon['TTK Min Far'],
					weapon['TTK Max Far']
				)
			),
			[]
		)
		condensed_weapons[w.weapon_name] = w

	with open("loadouts.json") as in_file:
		loadouts = json.load(in_file)

	for loadout in loadouts:
		l = Loadout(
			str(loadout['weapon']).upper(),
			str(loadout['type']).upper(),
			str(loadout['muzzle']),
			str(loadout['barrel']),
			str(loadout['laser']),
			str(loadout['optic']),
			str(loadout['stock']),
			str(loadout['underbarrel']),
			str(loadout['ammunition']),
			str(loadout['reargrip']),
			str(loadout['perk']),
		)
		condensed_weapons[l.weapon].loadouts.append(l)

	with open('weaponsList.json', 'w') as write_file:
		for weapon in condensed_weapons:
			json.dump(jsonpickle.encode(condensed_weapons[weapon], unpicklable=False), write_file)
