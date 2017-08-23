FORMAT: 1A

# Group Players

Players is the base piece of the Give-n-Go API. It contains data about basketball players, whether they've been benched or they're starters.

### List Players [GET]

+ Response 200 (application/json)

  [
    {
      "id": "curryst01",
      "name": "Stephen Curry",
      "position": "Point Guard",
      "height": "6-3",
      "weight": 190,
      "team": "Golden State Warriors",
      "draft": "Golden State Warriors-1-7-2009",
      "titles": [
        {
          "title": "All Star",
          "count": 4
        }, {...}
      ],
      "career": {
        "games": 574,
        "pts": 22.8,
        "rebounds": 4.4,
        "assists": 6.8,
        "fieldGoal": 47.6,
        "threePoint": 43.8,
        "freeThrow": 90.1
      },
      "seasons": {
        "09": /players/curryst01/09,
        "10": /players/curryst01/10,
        "11": /players/curryst01/11,
        "12": /players/curryst01/12,
        "13": /players/curryst01/13,
        "14": /players/curryst01/14,
        ...
      }
    }
  ]

+ Response 429 (application/json)

  {
    "suspensionLength": 15092340
  }
