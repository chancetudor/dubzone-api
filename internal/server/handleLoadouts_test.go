package server

import (
	"github.com/sirupsen/logrus"
	"github.com/steinfletcher/apitest"
	"net/http"
	"os"
	"testing"
)

var log = logrus.New()

func initDevLogs() {
	file, err := os.OpenFile("/Users/chancetudor/GitHub/dubzone-api/log/test_logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Error("Failed to log to file, using default stderr")
	}
}

var MockGetLoadouts = apitest.NewMock().
	Get("/loadouts").
	RespondWith().
	JSON(`[{"primary":{"weapon_name":"XM4","category":"Range","muzzle":"Suppressed","barrel":"Long Barrel","optic":"Long Sight","under_barrel":"Agent Grip","ammo":"60 round mags","meta_weapon":true},"secondary":{"weapon_name":"Colt .45","category":"Secondary","barrel":"Long Barrel","laser":"No Laser","optic":"Long Sight","under_barrel":"Agent Grip","ammo":"10 round mags"},"perk_one":"Perk1","perk_two":"Perk2","perk_three":"Perk3","lethal":"Semtex","tactical":"Stuns","meta_loadout":true},{"primary":{"weapon_name":"C58","category":"Range","muzzle":"Suppressed","barrel":"Long Barrel","optic":"Long Sight","under_barrel":"Agent Grip","ammo":"45 round mags","meta_weapon":true},"secondary":{"weapon_name":"Colt .45","category":"Secondary","barrel":"Long Barrel","laser":"No Laser","optic":"Long Sight","under_barrel":"Agent Grip","ammo":"10 round mags"},"perk_one":"Perk1","perk_two":"Perk2","perk_three":"Perk3","lethal":"Semtex","tactical":"Stuns","meta_loadout":true},{"primary":{"weapon_name":"MAC 10","category":"Close Range","muzzle":"Suppressed","barrel":"Short Barrel","laser":"Tiger Team","optic":"Red Dot Sight","under_barrel":"Agent Grip","ammo":"60 round mags","meta_weapon":true},"secondary":{"weapon_name":"Colt .45","category":"Secondary","barrel":"Long Barrel","laser":"No Laser","optic":"Long Sight","under_barrel":"Agent Grip","ammo":"10 round mags"},"perk_one":"Perk1","perk_two":"Perk2","perk_three":"Perk3","lethal":"Semtex","tactical":"Stuns"}]`).
	Status(http.StatusOK).
	End()

var MockGetMetaLoadouts = apitest.NewMock().
	Get("/loadouts").
	RespondWith().
	JSON(`[{"primary":{"weapon_name":"XM4","category":"Range","muzzle":"Suppressed","barrel":"Long Barrel","optic":"Long Sight","under_barrel":"Agent Grip","ammo":"60 round mags","meta_weapon":true},"secondary":{"weapon_name":"Colt .45","category":"Secondary","barrel":"Long Barrel","laser":"No Laser","optic":"Long Sight","under_barrel":"Agent Grip","ammo":"10 round mags"},"perk_one":"Perk1","perk_two":"Perk2","perk_three":"Perk3","lethal":"Semtex","tactical":"Stuns","meta_loadout":true},{"primary":{"weapon_name":"C58","category":"Range","muzzle":"Suppressed","barrel":"Long Barrel","optic":"Long Sight","under_barrel":"Agent Grip","ammo":"45 round mags","meta_weapon":true},"secondary":{"weapon_name":"Colt .45","category":"Secondary","barrel":"Long Barrel","laser":"No Laser","optic":"Long Sight","under_barrel":"Agent Grip","ammo":"10 round mags"},"perk_one":"Perk1","perk_two":"Perk2","perk_three":"Perk3","lethal":"Semtex","tactical":"Stuns","meta_loadout":true}]`).
	Status(http.StatusOK).
	End()

var MockCreateLoadout = apitest.NewMock().
	Post("/loadouts").
	Body(`{"primary": {"weapon_name": "Krig","category": "Range","muzzle": "Agency Suppressor","barrel": "Ranger/Mil-Spec","laser": "","optic": "3x","stock": "","under_barrel": "Field Agent","ammo": "60 Mag Normal","rear_grip": "","perk": "","meta_weapon": true},"secondary": {"weapon_name": "MAC-10","category": "SMG","muzzle": "GRU/Agency Suppressor","barrel": "Task Force","laser": "Tiger Team Spotlight","optic": "","stock": "Combat/Raider","under_barrel": "","ammo": "53 Mag Normal","rear_grip": "","perk": "","meta_weapon": true},"perk_one": "EOD","perk_two": "Overkill","perk_three": "Combat Scout","lethal": "Semtex","tactical": "Heartbeat Sensor","meta_loadout": true}`).
	RespondWith().
	JSON(`[]`).
	Status(http.StatusOK).
	End()

func TestServer_GetLoadouts(t *testing.T) {
	initDevLogs()
	log.Info("TESTING GET LOADOUTS")
	apitest.New().
		Mocks(MockGetLoadouts).
		Handler(NewServer(log).router).
		Get("/loadouts").
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestServer_GetMetaLoadouts(t *testing.T) {
	initDevLogs()
	log.Info("TESTING GET META LOADOUTS")
	apitest.New().
		Mocks(MockGetMetaLoadouts).
		Handler(NewServer(log).router).
		Get("/loadouts/meta").
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestServer_CreateLoadout(t *testing.T) {
	initDevLogs()
	log.Info("TESTING CREATE LOADOUT")
	apitest.New().
		Mocks(MockCreateLoadout).
		Handler(NewServer(log).router).
		Post("/loadouts").
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestServer_ValidateLoadout(t *testing.T) {

}
