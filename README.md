# Golang rest api

**Go Developer Test**

The idea of this project is to allow us to evaluate how you organize your work, set the project
structure, architecture, write understandable code and solve Go specific issues.

**Objective:** build the Monkey Island RESTFul API.

Monkey island is filled with toys: cuddly toys, toy weapons and ghosts. The cuddly toys are
further divided in two factions, dogs and monkeys.


**Requirements**
- Create a RESTful API that allows you to create, view, list, update and delete dogs,
monkeys and weapons.
   - E.g. A list of dogs should be available at /api/cuddly_toys/dogs/
   - E.g. A list of cuddly toys should be available at /api/cuddly_toys/
   - E.g. A single weapon should be available at /api/weapons/1 (if a weapon
with that ID exists in the database)
- The API should also generate a random list of ghosts whenever /ghosts/ is
requested, with a maximum of 10 ghosts.
- The API should have a single point of entry (e.g. all requests to /api/ and further
should be redirected to the same file) and communicate exclusively in JSON.
- The API should be backed by a database, keeping track of the dogs, monkeys and
weapons.
- The database class may only be loaded when a resource is requested that needs
database access (i.e. everything except ghosts).
- Every toy has a unique ID and a name (ghost names are random strings of < 8
alphanumeric characters)
- Every cuddly toy has an energy level and every weapon has a power level.
Basic documentation should be provided which includes information on all the requests (e.g.
"POST /dogs/ with a JSON object containing 'name' and 'energy_level' to create a new
dog")
