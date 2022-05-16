### How loc.json works
```
{   "name": "dark_cave_entrance", // Name of the place
        "current_description": 0, // Current active description
        "descriptions": [ // List of descriptions
            "You are standing in a cold, dark cave. You are looking directly into the darkness. You can feel the cold breeze on your neck indicating to you that the entrance is behind you."
        ],
        "connected_locations": { // Locations you can get to from here. False means you cannot get to it yet, and the player will not see it listed.
            "dark_cave_interior": false
        }
    }
```