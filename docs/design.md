# turn based strategy backend engine

this engine can be used for various time periods, both historical and fantasy, there's flexibility in altering the story line, but engine is fairly prescriptive when you get down to battle mode.  basically a user controls a faction or country and fights against other factions for territory and dominance.  there can be diplomacy.  maps are fixed and run until everyone quits or there's one person standing, these maps can last many years, so long as it's active.


## gameplay

* several aspects: battle preparation and battle

### battle preparation

there will be some level of preparation needed for battle, a user could go into battle without any preparation, but will have very little chance of success

* unit training
* supply and funding
* supply line
* marching to attack area
* scouting, intelligence gathering
* unit organization
* leadership selection
* diplomacy
* policy
* strategy

if the user is defending then preparation could consist of:

* fortifications
* traps

battle preparation will be a very important aspect of the game, careful consideration may decide the outcome of the game, a more superior force may be destroyed by a lesser force, a force with more people may also be destroyed by a smaller force, these are all posibilities so planning is critical

### battles

users will play the overall commander and will have full omniscient view and control over

* chain of command, including organization
* troop arrangement
* orders
* assessment of current
* tracking of progress
* fortifications can be created
* intelligence gathering
* care of injured

## maps

* squares or polygons
* terrain such as mountain, plain, hill, water, forest, steppe, desert, plateau, swamp, wasteland
* fog of war, but special intelligence can provide some predictionse
* random generation
* varying sizes

### terrain


## units

### core attributes

ranges 0 - 1000

attribute | fixed | relative | description | notes
--- | --- | --- | ---
age | | | |
rank | | | |
size | x | | |
hitpoints | | | |
intelligence | x | | |
strength | | | |
vision | | | |
charisma | x | | |
dexterity | | | |
experience | | | |
endurance | | | |
emotional intelligence | x | | |
psychology | | | |
luck | x | | |
loyalty | | x | |
morale | | | |
speed | x | | |
health | | | |
recovery | | | |
discipline | | | |
communication | x | | |
obey | | | |
strategy | | | |
respect | | x | |
knowledge | | | |
constitution | | | |
learning | | | |
stamina | | | |
agility | | | |
wisdom | | | |
movement | | | |

_fixed attribute: doesn't change _
_relative attribute: varies for each individual _

### abilities/special attributes

attribute | description | notes
--- | --- | ---
horse | |
hand to hand combat | |
short range | |
medium range | |
long range | |
crazy range | |
healing | |

specialty type, e.g. horse, gun, infantry, medic

* there should be other attributes as well for specific skills like shooting an arrow or firing a cannon.
* also should have a weapons list of attributes

### melee weapon attributes

this accommodates for close range weapons like swords, knives, lances

attribute | description | notes
--- | --- | ---
name | |
weight | |
size | |
speed | |
damage | |
attack radius | |


### ranged weapon attributes

bow and arrows, crossbows, pistols, rifles, cannons, mortars, grenades

attribute | description | notes
--- | --- | ---
name | |
range | |
power | |
speed | |
load speed | |
attack radius | |
damage | |
weight | |
size | size of the weapon, small - super large |
caliber | |
accuracy | |


### heroes

over time, as units gain experience, they will be given special bonuses on attack and defense, if you should increase rank for units, they may bring special abilities to the brigades.

attribute | description | notes
--- | --- | ---
battles | # of battles |
won | |
lost | |
streak | |
last won? | boolean for status of last battle |

### army sizes

operation | description | total units | commander
--- | --- | --- | ---
squad | 7-14 troops | 7-14 | sergeant
platoon | 3-4 squads | 21-56 | lieutenant
company | 2-4 platoons | 42-224 | captain
battalion | 2-5 companies | 84-1120 | lieutenant colonel
brigade/regiment | 3 battalions | 252-3360 | brigadier general/colonel
division | 2-3 brigrades/regiments | 504-10080 | major general
corps | 2-7 divisions | 1008-70560 | lieutenant general
field army | 2-5 corps | 2016-352800 | general

https://www.britannica.com/topic/brigade
https://www.orlandosentinel.com/news/os-xpm-2003-03-29-0303290257-story.html

### troop organization

rank | description | commands
--- | --- | ---
soldier | | x
corporal | | 7-14
sergeant | | 16 - 50
regimental sergeant | |
lieutenant | |
captain | |
major | |
lieutenant colonel | | 650
colonel | |
brigadier general | |
major general | |
lieutenant general | |
general | |
field marshal | |



### medical

medical units will help the wounded, but this should play a smaller role, i guess wounded will typically be unable to fight in the next battle unless the damage is light.

### unit discipline

before and after battles, troops will still need to be maintained, there will be a certain level of desertion, soldiers will get into disciplinary action, and general sickness will
abound.  it will be important to maintain soldier training and discipline.

## weather

weather will mostly impact terrain and troop movement/visibility, morale will also be impacted.  for maps, weather will not be consistent across each space.

## supply

### soldiers

### food/water

### ammunition

### medical supplies

### camp infra

tents, cooking utensils

## battle

### formations

## spoils

spoils can be ransacked on the country side and after battles, but this will have a negative repurcussion on the people's stability.

## notes

* technology shouldnt be constantly evolving making battles, strategy, and experience more important, should be a much more gradual approach, focus on battles and preparation for battles
* economy should be based on cities, industries, and population, shouldnt require manually mining resources like starcraft, this should something that happens in the background
* each faction should have different attributes like populations, income, army competence
* what should the multiplayer angle be since this is turn based, some players will block the whole game's progress?  perhaps if users dont progress within a certain time, then a default set of actions will be enacted
* hero units will appear over time since all units have varying degrees of ability, it is up to the user to recognize these units and promote them or develop them, there will be an additional intangible field for each unit which adds some increase in abilities
