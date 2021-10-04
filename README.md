# DubZone API

Inspired by Derek Sivers' post ["Why You Need a Database"](https://sive.rs/dbt), I set out to create a database and an API to access the database. 

I didn't know what data to hold and expose until I realized that information about weapons and loadouts for Call of Duty Warzone would be a perfect candidate, as such information is only currently accessible in a multitude of YouTube videos.

I built this in Go to further develop my knowledge of the language, and used MongoDB to gain further experience in NoSQL schemas.

_This is a personal project, and so is not currently operational for real-world usage._ Feel free to take a look at my code to see my design choices and code quality, though!

## Open Endpoints

Open endpoints require no Authentication.

* Loadouts : 
  * `GET /loadouts` --> returns all loadouts
  * `POST /loadout` --> creates a new loadout
  * `GET /loadouts/category/{cat}` --> returns all loadouts fitting a specified category
  * `GET /loadouts/weapon/{name}` --> returns all loadouts containing a specified weapon
  * `GET /loadouts/meta` --> returns all loadouts that are currently meta
* Weapons:
  * `GET /weapon/{name}` --> returns a weapon's data
  * `GET /weapons/meta` --> returns all weapons that are currently meta
  * `GET /weapons/{cat}` --> returns all weapons fitting a specified category
  * `GET /weapons/categories` --> returns all weapon categories.

