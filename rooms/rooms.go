package rooms

type Room struct {
	Name   string
	Enter  string
	Leave  string
	Defeat string
}

var SpookyCave = Room{
	Name:   "Spooky Cave",
	Enter:  "You step cautiously into the dark cave, the air thick with an eerie silence.",
	Leave:  "You descend deeper into the cave, the darkness swallowing your footsteps.",
	Defeat: "A shadowy figure emerges from the darkness and overwhelms you.",
}

var AbandonedMine = Room{
	Name:   "Abandoned Mine",
	Enter:  "You cautiously enter the dark, echoing halls of the abandoned mine.",
	Leave:  "You proceed deeper into the mine, the echoes of your steps bouncing off the walls.",
	Defeat: "A rusted mining robot suddenly activates and attacks you.",
}

var UndergroundCastle = Room{
	Name:   "Ravine Castle",
	Enter:  "You step into the grand halls of the castle, carved into the ravine's side.",
	Leave:  "You descend further into the depths of the castle, following dimly lit corridors.",
	Defeat: "You succumb to the dangers lurking in the shadows of the underground castle.",
}

var MossyDungeon = Room{
	Name:   "Overgrown Dungeon",
	Enter:  "You enter the moss-covered dungeon, feeling a chill in the damp air.",
	Leave:  "You descend deeper into the dungeon, the stone walls closing in around you.",
	Defeat: "You meet a grim fate within the depths of the dungeon.",
}

var MurkyCavern = Room{
	Name:   "Murky Cavern",
	Enter:  "You wade into the water-filled cavern, the murky water reaching your knees.",
	Leave:  "You press on through the cavern, the water swirling around your legs.",
	Defeat: "You face a grim outcome amidst the depths of the cavern.",
}

var AncientRuins = Room{
	Name:   "Old Ruins",
	Enter:  "You explore the ancient ruins, marveling at the weathered stone structures.",
	Leave:  "You delve deeper into the ruins, uncovering secrets hidden for centuries.",
	Defeat: "You face a dire fate within the crumbling walls of the ancient ruins.",
}

var CrystalHollow = Room{
	Name:   "Crystal Hollow",
	Enter:  "You step into the hollow filled with bright, glittering crystals.",
	Leave:  "You continue deeper into the hollow, the crystals casting shimmering reflections.",
	Defeat: "You meet a grim end amidst the radiant crystals of the hollow.",
}

var MirrorMaze = Room{
	Name:   "Mirror Maze",
	Enter:  "You enter the disorienting mirror maze, reflections stretching endlessly in every direction.",
	Leave:  "You navigate through the maze, each reflection leading you deeper into confusion.",
	Defeat: "You face a daunting challenge within the labyrinth of mirrors.",
}

var TreasureRoom = Room{
	Name:   "Treasure Vault",
	Enter:  "You step into the expansive treasure room, dazzled by the sight of gold, gems, and relics.",
	Leave:  "You explore deeper into the vault, marveling at the treasures that surround you.",
	Defeat: "You face an uncertain fate within the riches of the treasure vault.",
}

var LavaPlatform = Room{
	Name:   "Volcano Platform",
	Enter:  "You arrive at the platform surrounded by bubbling lava, its heat palpable even from a distance.",
	Leave:  "You proceed cautiously across the platform, the lava's glow illuminating your path.",
	Defeat: "You confront a perilous challenge amidst the fiery depths of the volcano.",
}

var Dungeon = []Room{
	SpookyCave,
	AbandonedMine,
	UndergroundCastle,
	MossyDungeon,
	MurkyCavern,
	AncientRuins,
	CrystalHollow,
	MirrorMaze,
	TreasureRoom,
	LavaPlatform,
}
