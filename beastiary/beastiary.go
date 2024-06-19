package beastiary

import "math/rand"

type Beast struct {
	Name          string
	Level         int
	Attack        int
	EntryMessage  string
	DefeatMessage string
	WinMessage    string
}

var Kobold = Beast{
	Name:          "Kobold",
	Level:         1,
	Attack:        2,
	EntryMessage:  "The small, frail Kobold eyes you with jagged teeth, ready to defend its territory.",
	DefeatMessage: "The Kobold scampers away in fear as you prove victorious.",
	WinMessage:    "The Kobold cackles triumphantly, claiming victory over its intruder.",
}

var Ghost = Beast{
	Name:          "Ghost",
	Level:         2,
	Attack:        1,
	EntryMessage:  "A light, ethereal presence drifts into the room, emitting a sorrowful moan.",
	DefeatMessage: "The ghost dissipates into a mist, its sad wails fading away.",
	WinMessage:    "The ghost's mournful cry echoes through the room as it claims victory over you.",
}

var HauntedArmor = Beast{
	Name:          "Haunted Armor",
	Level:         3,
	Attack:        0,
	EntryMessage:  "The Haunted Armor clanks to life, its stony plates and clockwork gears grinding ominously.",
	DefeatMessage: "The Haunted Armor collapses in a heap of shattered metal and gears, its rigid form finally still.",
	WinMessage:    "The Haunted Armor stands victorious, its rigid stance unyielding against all foes.",
}

var MossyGolem = Beast{
	Name:          "Mossy Earth Golem",
	Level:         4,
	Attack:        0,
	EntryMessage:  "The moss-covered Earth Golem rises from the ground, weeds entwined in its crumbling form.",
	DefeatMessage: "The Earth Golem collapses into a heap of crumbling earth and moss, defeated.",
	WinMessage:    "The Earth Golem stands victorious, its heavy footsteps echoing through the chamber.",
}

var GiantCrab = Beast{
	Name:          "Giant Crab",
	Level:         5,
	Attack:        2,
	EntryMessage:  "A giant crab scuttles into view, its creepy eyes clicking and its giant snipping claw ready to strike.",
	DefeatMessage: "The giant crab collapses, its giant snipping claw motionless and its creepy eyes dim.",
	WinMessage:    "The giant crab clicks triumphantly, its heavy shell protecting it from all threats.",
}

var Imp = Beast{
	Name:          "Imp",
	Level:         6,
	Attack:        1,
	EntryMessage:  "An Imp flits into view, its leathery wings beating quickly as it darts around.",
	DefeatMessage: "The Imp squeals in defeat, its wings slowing as it falls to the ground.",
	WinMessage:    "The Imp cackles gleefully, its quick movements proving too elusive for you.",
}

var CrystalSentinel = Beast{
	Name:          "Crystal Sentinel",
	Level:         7,
	Attack:        0,
	EntryMessage:  "A Crystal Sentinel emerges, its translucent form shimmering as it moves with mechanical precision.",
	DefeatMessage: "The Crystal Sentinel shatters into a thousand pieces, its mechanical parts ceasing to function.",
	WinMessage:    "The Crystal Sentinel stands tall, its hulking form a testament to its unwavering strength.",
}

var Doppleganger = Beast{
	Name:          "Doppleganger",
	Level:         8,
	Attack:        2,
	EntryMessage:  "A Doppleganger materializes before you, its form shifting like a mirror image.",
	DefeatMessage: "The Doppleganger fades away into nothingness, its fleeting presence vanishing.",
	WinMessage:    "The Doppleganger laughs mockingly, its mirror images confusing and overwhelming you.",
}

var Dragon = Beast{
	Name:          "Dragon",
	Level:         9,
	Attack:        0,
	EntryMessage:  "A mighty Dragon descends, its scales shimmering in a brilliant red hue.",
	DefeatMessage: "The Dragon roars in defeat, its majestic form crashing to the ground.",
	WinMessage:    "The Dragon triumphs with a thunderous roar, its sturdy scales protecting it from harm.",
}

var Balrog = Beast{
	Name:          "Balrog",
	Level:         10,
	Attack:        1,
	EntryMessage:  "A Balrog emerges from the shadows, its flaming whip crackling in fury.",
	DefeatMessage: "The Balrog bellows in defeat, its fiery form dissipating into smoke and ash.",
	WinMessage:    "The Balrog stands triumphant, its furious gaze challenging all who oppose it.",
}

type Beastiary struct {
	beasts []Beast
}

var GetBeastiary = Beastiary{
	beasts: []Beast{
		Kobold,
		Ghost,
		HauntedArmor,
		MossyGolem,
		GiantCrab,
		Imp,
		CrystalSentinel,
		Doppleganger,
		Dragon,
		Balrog,
	},
}

func GetBeastForLevel(level int, book *Beastiary) *Beast {
	var beastForLevel = []Beast{}

	for _, beast := range book.beasts {
		if beast.Level == level {
			beastForLevel = append(beastForLevel, beast)
		}
	}

	return &beastForLevel[rand.Intn(len(beastForLevel))]
}
