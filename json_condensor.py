import json


class Weapon:
    def __init__(self, name: str, game: str, rpm: int, bv: int, close, mid, far, loady: list):
        """Constructor"""
        self.weapon_name = name
        self.game_from = game
        self.rpm = rpm
        self.bullet_velocity = int(bv)
        self.close_dmg_profile = close
        self.mid_dmg_profile = mid
        self.far_dmg_profile = far
        self.loadouts = loady

    def __repr__(self):
        return "Weapon Name: {} From Game: {} RPM: {} Bullet Velocity: {} Close Range Damage Profile: {} " \
                "Mid-range Damage Profile: {} Loadouts: {}".format(
                    str(self.weaponname), str(self.game_from), str(self.rpm), str(self.bullet_velocity),
                    self.close_dmg_profile, self.mid_dmg_profile, self.far_dmg_profile, self.loadouts
                    )

    def __str__(self):
        return "Weapon Name: {} From Game: {} RPM: {} Bullet Velocity: {} Close Range Damage Profile: {} " \
                    "Mid-range Damage Profile: {} Loadouts: {}".format(
                        str(self.weaponname), str(self.game_from), str(self.rpm), str(self.bullet_velocity),
                        self.close_dmg_profile, self.mid_dmg_profile, self.far_dmg_profile, self.loadouts
                        )

    @property
    def weapon_name(self):
        return self.weapon_name
    
    @weapon_name.setter
    def weapon_name(self, value):
        self.weapon_name = value

    @property
    def loadouts(self):
        return self.loadouts

    @loadouts.setter
    def list(self, value):
        self.loadouts.append(value)


class DamageProfile:
    def __init__(self, min_dist, max_dist, min_stk, max_stk, min_ttk, max_ttk):
        self.mindistance = min_dist
        self.maxdistance = max_dist
        self.minstk = min_stk
        self.maxstk = max_stk
        self.minttk = min_ttk
        self.maxttk = max_ttk

    def __repr__(self):
        return "Minimum Distance: {} Maximum Distance: {} Minimum Shots to Kill: {} Maximum Shots to Kill: {} Minimum " \
               "Time to Kill: {} Maximum Time to Kill: {}".format(
                        self.mindistance, self.maxdistance, self.minstk, self.maxstk,
                        self.minttk, self.maxstk
                        )


class CloseRange(DamageProfile):
    def __init__(self, min_dist, max_dist, min_stk, max_stk, min_ttk, max_ttk):
        super().__init__(min_dist, max_dist, min_stk, max_stk, min_ttk, max_ttk)


class MidRange(DamageProfile):
    def __init__(self, min_dist, max_dist, min_stk, max_stk, min_ttk, max_ttk):
        super().__init__(min_dist, max_dist, min_stk, max_stk, min_ttk, max_ttk)


class FarRange(DamageProfile):
    def __init__(self, min_dist, max_dist, min_stk, max_stk, min_ttk, max_ttk):
        super().__init__(min_dist, max_dist, min_stk, max_stk, min_ttk, max_ttk)


class Loadout:
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
            weapon['weaponname'],
            weapon['gamefrom'],
            weapon['rpm'],
            weapon['bulletvelocity'],
            CloseRange(
                0,
                weapon['MaxDist Close'],
                weapon['STK Min Close'],
                weapon['STK Max Close'],
                weapon['TTK Min Close'],
                weapon['TTK Max Close']
            ),
            MidRange(
                weapon['MinDist Mid'],
                weapon['MaxDist Mid'],
                weapon['STK Min Mid'],
                weapon['STK Max Mid'],
                weapon['TTK Min Mid'],
                weapon['TTK Max Mid']
            ),
            FarRange(
                weapon['MinDist Far'],
                0,
                weapon['STK Min Far'],
                weapon['STK Max Far'],
                weapon['TTK Min Far'],
                weapon['TTK Max Far']
            ),
            []
        )
        condensed_weapons[w.weapon_name] = w

    with open("loadouts.json") as in_file:
        loadouts = json.load(in_file)

    for loadout in loadouts:
        l = Loadout(
            loadout['Weapon'],
            loadout['Type'],
            loadout['Muzzle'],
            loadout['Barrel'],
            loadout['Laser'],
            loadout['Optic'],
            loadout['Stock'],
            loadout['Underbarrel'],
            loadout['Ammunition'],
            loadout['Rear Grip'],
            loadout['Perk'],
        )
        condensed_weapons[l.weapon].loadouts.append(l)

    for weapon in condensed_weapons:
        print(weapon)
