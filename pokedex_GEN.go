// Code generated - DO NOT EDIT.
// Regenerate with `go generate`.

package pokemonbattlelib

// A map of national pokedex numbers to pokemon names.
var pokemonNames = map[uint16]string{
	1:   "Bulbasaur",
	2:   "Ivysaur",
	3:   "Venusaur",
	4:   "Charmander",
	5:   "Charmeleon",
	6:   "Charizard",
	7:   "Squirtle",
	8:   "Wartortle",
	9:   "Blastoise",
	10:  "Caterpie",
	11:  "Metapod",
	12:  "Butterfree",
	13:  "Weedle",
	14:  "Kakuna",
	15:  "Beedrill",
	16:  "Pidgey",
	17:  "Pidgeotto",
	18:  "Pidgeot",
	19:  "Rattata",
	20:  "Raticate",
	21:  "Spearow",
	22:  "Fearow",
	23:  "Ekans",
	24:  "Arbok",
	25:  "Pikachu",
	26:  "Raichu",
	27:  "Sandshrew",
	28:  "Sandslash",
	29:  "Nidoran♀",
	30:  "Nidorina",
	31:  "Nidoqueen",
	32:  "Nidoran♂",
	33:  "Nidorino",
	34:  "Nidoking",
	35:  "Clefairy",
	36:  "Clefable",
	37:  "Vulpix",
	38:  "Ninetales",
	39:  "Jigglypuff",
	40:  "Wigglytuff",
	41:  "Zubat",
	42:  "Golbat",
	43:  "Oddish",
	44:  "Gloom",
	45:  "Vileplume",
	46:  "Paras",
	47:  "Parasect",
	48:  "Venonat",
	49:  "Venomoth",
	50:  "Diglett",
	51:  "Dugtrio",
	52:  "Meowth",
	53:  "Persian",
	54:  "Psyduck",
	55:  "Golduck",
	56:  "Mankey",
	57:  "Primeape",
	58:  "Growlithe",
	59:  "Arcanine",
	60:  "Poliwag",
	61:  "Poliwhirl",
	62:  "Poliwrath",
	63:  "Abra",
	64:  "Kadabra",
	65:  "Alakazam",
	66:  "Machop",
	67:  "Machoke",
	68:  "Machamp",
	69:  "Bellsprout",
	70:  "Weepinbell",
	71:  "Victreebel",
	72:  "Tentacool",
	73:  "Tentacruel",
	74:  "Geodude",
	75:  "Graveler",
	76:  "Golem",
	77:  "Ponyta",
	78:  "Rapidash",
	79:  "Slowpoke",
	80:  "Slowbro",
	81:  "Magnemite",
	82:  "Magneton",
	83:  "Farfetch’d",
	84:  "Doduo",
	85:  "Dodrio",
	86:  "Seel",
	87:  "Dewgong",
	88:  "Grimer",
	89:  "Muk",
	90:  "Shellder",
	91:  "Cloyster",
	92:  "Gastly",
	93:  "Haunter",
	94:  "Gengar",
	95:  "Onix",
	96:  "Drowzee",
	97:  "Hypno",
	98:  "Krabby",
	99:  "Kingler",
	100: "Voltorb",
	101: "Electrode",
	102: "Exeggcute",
	103: "Exeggutor",
	104: "Cubone",
	105: "Marowak",
	106: "Hitmonlee",
	107: "Hitmonchan",
	108: "Lickitung",
	109: "Koffing",
	110: "Weezing",
	111: "Rhyhorn",
	112: "Rhydon",
	113: "Chansey",
	114: "Tangela",
	115: "Kangaskhan",
	116: "Horsea",
	117: "Seadra",
	118: "Goldeen",
	119: "Seaking",
	120: "Staryu",
	121: "Starmie",
	122: "Mr. Mime",
	123: "Scyther",
	124: "Jynx",
	125: "Electabuzz",
	126: "Magmar",
	127: "Pinsir",
	128: "Tauros",
	129: "Magikarp",
	130: "Gyarados",
	131: "Lapras",
	132: "Ditto",
	133: "Eevee",
	134: "Vaporeon",
	135: "Jolteon",
	136: "Flareon",
	137: "Porygon",
	138: "Omanyte",
	139: "Omastar",
	140: "Kabuto",
	141: "Kabutops",
	142: "Aerodactyl",
	143: "Snorlax",
	144: "Articuno",
	145: "Zapdos",
	146: "Moltres",
	147: "Dratini",
	148: "Dragonair",
	149: "Dragonite",
	150: "Mewtwo",
	151: "Mew",
	152: "Chikorita",
	153: "Bayleef",
	154: "Meganium",
	155: "Cyndaquil",
	156: "Quilava",
	157: "Typhlosion",
	158: "Totodile",
	159: "Croconaw",
	160: "Feraligatr",
	161: "Sentret",
	162: "Furret",
	163: "Hoothoot",
	164: "Noctowl",
	165: "Ledyba",
	166: "Ledian",
	167: "Spinarak",
	168: "Ariados",
	169: "Crobat",
	170: "Chinchou",
	171: "Lanturn",
	172: "Pichu",
	173: "Cleffa",
	174: "Igglybuff",
	175: "Togepi",
	176: "Togetic",
	177: "Natu",
	178: "Xatu",
	179: "Mareep",
	180: "Flaaffy",
	181: "Ampharos",
	182: "Bellossom",
	183: "Marill",
	184: "Azumarill",
	185: "Sudowoodo",
	186: "Politoed",
	187: "Hoppip",
	188: "Skiploom",
	189: "Jumpluff",
	190: "Aipom",
	191: "Sunkern",
	192: "Sunflora",
	193: "Yanma",
	194: "Wooper",
	195: "Quagsire",
	196: "Espeon",
	197: "Umbreon",
	198: "Murkrow",
	199: "Slowking",
	200: "Misdreavus",
	201: "Unown",
	202: "Wobbuffet",
	203: "Girafarig",
	204: "Pineco",
	205: "Forretress",
	206: "Dunsparce",
	207: "Gligar",
	208: "Steelix",
	209: "Snubbull",
	210: "Granbull",
	211: "Qwilfish",
	212: "Scizor",
	213: "Shuckle",
	214: "Heracross",
	215: "Sneasel",
	216: "Teddiursa",
	217: "Ursaring",
	218: "Slugma",
	219: "Magcargo",
	220: "Swinub",
	221: "Piloswine",
	222: "Corsola",
	223: "Remoraid",
	224: "Octillery",
	225: "Delibird",
	226: "Mantine",
	227: "Skarmory",
	228: "Houndour",
	229: "Houndoom",
	230: "Kingdra",
	231: "Phanpy",
	232: "Donphan",
	233: "Porygon2",
	234: "Stantler",
	235: "Smeargle",
	236: "Tyrogue",
	237: "Hitmontop",
	238: "Smoochum",
	239: "Elekid",
	240: "Magby",
	241: "Miltank",
	242: "Blissey",
	243: "Raikou",
	244: "Entei",
	245: "Suicune",
	246: "Larvitar",
	247: "Pupitar",
	248: "Tyranitar",
	249: "Lugia",
	250: "Ho-Oh",
	251: "Celebi",
	252: "Treecko",
	253: "Grovyle",
	254: "Sceptile",
	255: "Torchic",
	256: "Combusken",
	257: "Blaziken",
	258: "Mudkip",
	259: "Marshtomp",
	260: "Swampert",
	261: "Poochyena",
	262: "Mightyena",
	263: "Zigzagoon",
	264: "Linoone",
	265: "Wurmple",
	266: "Silcoon",
	267: "Beautifly",
	268: "Cascoon",
	269: "Dustox",
	270: "Lotad",
	271: "Lombre",
	272: "Ludicolo",
	273: "Seedot",
	274: "Nuzleaf",
	275: "Shiftry",
	276: "Taillow",
	277: "Swellow",
	278: "Wingull",
	279: "Pelipper",
	280: "Ralts",
	281: "Kirlia",
	282: "Gardevoir",
	283: "Surskit",
	284: "Masquerain",
	285: "Shroomish",
	286: "Breloom",
	287: "Slakoth",
	288: "Vigoroth",
	289: "Slaking",
	290: "Nincada",
	291: "Ninjask",
	292: "Shedinja",
	293: "Whismur",
	294: "Loudred",
	295: "Exploud",
	296: "Makuhita",
	297: "Hariyama",
	298: "Azurill",
	299: "Nosepass",
	300: "Skitty",
	301: "Delcatty",
	302: "Sableye",
	303: "Mawile",
	304: "Aron",
	305: "Lairon",
	306: "Aggron",
	307: "Meditite",
	308: "Medicham",
	309: "Electrike",
	310: "Manectric",
	311: "Plusle",
	312: "Minun",
	313: "Volbeat",
	314: "Illumise",
	315: "Roselia",
	316: "Gulpin",
	317: "Swalot",
	318: "Carvanha",
	319: "Sharpedo",
	320: "Wailmer",
	321: "Wailord",
	322: "Numel",
	323: "Camerupt",
	324: "Torkoal",
	325: "Spoink",
	326: "Grumpig",
	327: "Spinda",
	328: "Trapinch",
	329: "Vibrava",
	330: "Flygon",
	331: "Cacnea",
	332: "Cacturne",
	333: "Swablu",
	334: "Altaria",
	335: "Zangoose",
	336: "Seviper",
	337: "Lunatone",
	338: "Solrock",
	339: "Barboach",
	340: "Whiscash",
	341: "Corphish",
	342: "Crawdaunt",
	343: "Baltoy",
	344: "Claydol",
	345: "Lileep",
	346: "Cradily",
	347: "Anorith",
	348: "Armaldo",
	349: "Feebas",
	350: "Milotic",
	351: "Castform",
	352: "Kecleon",
	353: "Shuppet",
	354: "Banette",
	355: "Duskull",
	356: "Dusclops",
	357: "Tropius",
	358: "Chimecho",
	359: "Absol",
	360: "Wynaut",
	361: "Snorunt",
	362: "Glalie",
	363: "Spheal",
	364: "Sealeo",
	365: "Walrein",
	366: "Clamperl",
	367: "Huntail",
	368: "Gorebyss",
	369: "Relicanth",
	370: "Luvdisc",
	371: "Bagon",
	372: "Shelgon",
	373: "Salamence",
	374: "Beldum",
	375: "Metang",
	376: "Metagross",
	377: "Regirock",
	378: "Regice",
	379: "Registeel",
	380: "Latias",
	381: "Latios",
	382: "Kyogre",
	383: "Groudon",
	384: "Rayquaza",
	385: "Jirachi",
	386: "Deoxys",
	387: "Turtwig",
	388: "Grotle",
	389: "Torterra",
	390: "Chimchar",
	391: "Monferno",
	392: "Infernape",
	393: "Piplup",
	394: "Prinplup",
	395: "Empoleon",
	396: "Starly",
	397: "Staravia",
	398: "Staraptor",
	399: "Bidoof",
	400: "Bibarel",
	401: "Kricketot",
	402: "Kricketune",
	403: "Shinx",
	404: "Luxio",
	405: "Luxray",
	406: "Budew",
	407: "Roserade",
	408: "Cranidos",
	409: "Rampardos",
	410: "Shieldon",
	411: "Bastiodon",
	412: "Burmy",
	413: "Wormadam",
	414: "Mothim",
	415: "Combee",
	416: "Vespiquen",
	417: "Pachirisu",
	418: "Buizel",
	419: "Floatzel",
	420: "Cherubi",
	421: "Cherrim",
	422: "Shellos",
	423: "Gastrodon",
	424: "Ambipom",
	425: "Drifloon",
	426: "Drifblim",
	427: "Buneary",
	428: "Lopunny",
	429: "Mismagius",
	430: "Honchkrow",
	431: "Glameow",
	432: "Purugly",
	433: "Chingling",
	434: "Stunky",
	435: "Skuntank",
	436: "Bronzor",
	437: "Bronzong",
	438: "Bonsly",
	439: "Mime Jr.",
	440: "Happiny",
	441: "Chatot",
	442: "Spiritomb",
	443: "Gible",
	444: "Gabite",
	445: "Garchomp",
	446: "Munchlax",
	447: "Riolu",
	448: "Lucario",
	449: "Hippopotas",
	450: "Hippowdon",
	451: "Skorupi",
	452: "Drapion",
	453: "Croagunk",
	454: "Toxicroak",
	455: "Carnivine",
	456: "Finneon",
	457: "Lumineon",
	458: "Mantyke",
	459: "Snover",
	460: "Abomasnow",
	461: "Weavile",
	462: "Magnezone",
	463: "Lickilicky",
	464: "Rhyperior",
	465: "Tangrowth",
	466: "Electivire",
	467: "Magmortar",
	468: "Togekiss",
	469: "Yanmega",
	470: "Leafeon",
	471: "Glaceon",
	472: "Gliscor",
	473: "Mamoswine",
	474: "Porygon-Z",
	475: "Gallade",
	476: "Probopass",
	477: "Dusknoir",
	478: "Froslass",
	479: "Rotom",
	480: "Uxie",
	481: "Mesprit",
	482: "Azelf",
	483: "Dialga",
	484: "Palkia",
	485: "Heatran",
	486: "Regigigas",
	487: "Giratina",
	488: "Cresselia",
	489: "Phione",
	490: "Manaphy",
	491: "Darkrai",
	492: "Shaymin",
	493: "Arceus",
}

var AllMoves = []Move{
	{ID: 1, Name: "Pound", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 35, MaxPP: 35, Targets: 10, Priority: 0, Power: 40, Accuracy: 100},
	{ID: 2, Name: "Karate Chop", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 25, MaxPP: 25, Targets: 10, Priority: 0, Power: 50, Accuracy: 100},
	{ID: 3, Name: "Double Slap", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 15, Accuracy: 85},
	{ID: 4, Name: "Comet Punch", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 18, Accuracy: 85},
	{ID: 5, Name: "Mega Punch", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 80, Accuracy: 85},
	{ID: 6, Name: "Pay Day", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 40, Accuracy: 100},
	{ID: 7, Name: "Fire Punch", Type: 512, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 75, Accuracy: 100},
	{ID: 8, Name: "Ice Punch", Type: 16384, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 75, Accuracy: 100},
	{ID: 9, Name: "Thunder Punch", Type: 4096, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 75, Accuracy: 100},
	{ID: 10, Name: "Scratch", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 35, MaxPP: 35, Targets: 10, Priority: 0, Power: 40, Accuracy: 100},
	{ID: 11, Name: "Vice Grip", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 30, MaxPP: 30, Targets: 10, Priority: 0, Power: 55, Accuracy: 100},
	{ID: 12, Name: "Guillotine", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 0, Accuracy: 30},
	{ID: 13, Name: "Razor Wind", Type: 1, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 11, Priority: 0, Power: 80, Accuracy: 100},
	{ID: 14, Name: "Swords Dance", Type: 1, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 15, Name: "Cut", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 30, MaxPP: 30, Targets: 10, Priority: 0, Power: 50, Accuracy: 95},
	{ID: 16, Name: "Gust", Type: 4, Category: MoveCategorySpecial, CurrentPP: 35, MaxPP: 35, Targets: 10, Priority: 0, Power: 40, Accuracy: 100},
	{ID: 17, Name: "Wing Attack", Type: 4, Category: MoveCategoryPhysical, CurrentPP: 35, MaxPP: 35, Targets: 10, Priority: 0, Power: 60, Accuracy: 100},
	{ID: 18, Name: "Whirlwind", Type: 1, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: -6, Power: 0, Accuracy: 0},
	{ID: 19, Name: "Fly", Type: 4, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 90, Accuracy: 95},
	{ID: 20, Name: "Bind", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 15, Accuracy: 85},
	{ID: 21, Name: "Slam", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 80, Accuracy: 75},
	{ID: 22, Name: "Vine Whip", Type: 2048, Category: MoveCategoryPhysical, CurrentPP: 25, MaxPP: 25, Targets: 10, Priority: 0, Power: 45, Accuracy: 100},
	{ID: 23, Name: "Stomp", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 65, Accuracy: 100},
	{ID: 24, Name: "Double Kick", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 30, MaxPP: 30, Targets: 10, Priority: 0, Power: 30, Accuracy: 100},
	{ID: 25, Name: "Mega Kick", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 120, Accuracy: 75},
	{ID: 26, Name: "Jump Kick", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 100, Accuracy: 95},
	{ID: 27, Name: "Rolling Kick", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 60, Accuracy: 85},
	{ID: 28, Name: "Sand Attack", Type: 16, Category: MoveCategoryStatus, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 29, Name: "Headbutt", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 70, Accuracy: 100},
	{ID: 30, Name: "Horn Attack", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 25, MaxPP: 25, Targets: 10, Priority: 0, Power: 65, Accuracy: 100},
	{ID: 31, Name: "Fury Attack", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 15, Accuracy: 85},
	{ID: 32, Name: "Horn Drill", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 0, Accuracy: 30},
	{ID: 33, Name: "Tackle", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 35, MaxPP: 35, Targets: 10, Priority: 0, Power: 40, Accuracy: 100},
	{ID: 34, Name: "Body Slam", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 85, Accuracy: 100},
	{ID: 35, Name: "Wrap", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 15, Accuracy: 90},
	{ID: 36, Name: "Take Down", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 90, Accuracy: 85},
	{ID: 37, Name: "Thrash", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 8, Priority: 0, Power: 120, Accuracy: 100},
	{ID: 38, Name: "Double-Edge", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 120, Accuracy: 100},
	{ID: 39, Name: "Tail Whip", Type: 1, Category: MoveCategoryStatus, CurrentPP: 30, MaxPP: 30, Targets: 11, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 40, Name: "Poison Sting", Type: 8, Category: MoveCategoryPhysical, CurrentPP: 35, MaxPP: 35, Targets: 10, Priority: 0, Power: 15, Accuracy: 100},
	{ID: 41, Name: "Twineedle", Type: 64, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 25, Accuracy: 100},
	{ID: 42, Name: "Pin Missile", Type: 64, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 25, Accuracy: 95},
	{ID: 43, Name: "Leer", Type: 1, Category: MoveCategoryStatus, CurrentPP: 30, MaxPP: 30, Targets: 11, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 44, Name: "Bite", Type: 65536, Category: MoveCategoryPhysical, CurrentPP: 25, MaxPP: 25, Targets: 10, Priority: 0, Power: 60, Accuracy: 100},
	{ID: 45, Name: "Growl", Type: 1, Category: MoveCategoryStatus, CurrentPP: 40, MaxPP: 40, Targets: 11, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 46, Name: "Roar", Type: 1, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: -6, Power: 0, Accuracy: 0},
	{ID: 47, Name: "Sing", Type: 1, Category: MoveCategoryStatus, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 0, Accuracy: 55},
	{ID: 48, Name: "Supersonic", Type: 1, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 0, Accuracy: 55},
	{ID: 49, Name: "Sonic Boom", Type: 1, Category: MoveCategorySpecial, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 0, Accuracy: 90},
	{ID: 50, Name: "Disable", Type: 1, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 51, Name: "Acid", Type: 8, Category: MoveCategorySpecial, CurrentPP: 30, MaxPP: 30, Targets: 11, Priority: 0, Power: 40, Accuracy: 100},
	{ID: 52, Name: "Ember", Type: 512, Category: MoveCategorySpecial, CurrentPP: 25, MaxPP: 25, Targets: 10, Priority: 0, Power: 40, Accuracy: 100},
	{ID: 53, Name: "Flamethrower", Type: 512, Category: MoveCategorySpecial, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 90, Accuracy: 100},
	{ID: 54, Name: "Mist", Type: 16384, Category: MoveCategoryStatus, CurrentPP: 30, MaxPP: 30, Targets: 4, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 55, Name: "Water Gun", Type: 1024, Category: MoveCategorySpecial, CurrentPP: 25, MaxPP: 25, Targets: 10, Priority: 0, Power: 40, Accuracy: 100},
	{ID: 56, Name: "Hydro Pump", Type: 1024, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 110, Accuracy: 80},
	{ID: 57, Name: "Surf", Type: 1024, Category: MoveCategorySpecial, CurrentPP: 15, MaxPP: 15, Targets: 9, Priority: 0, Power: 90, Accuracy: 100},
	{ID: 58, Name: "Ice Beam", Type: 16384, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 90, Accuracy: 100},
	{ID: 59, Name: "Blizzard", Type: 16384, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 11, Priority: 0, Power: 110, Accuracy: 70},
	{ID: 60, Name: "Psybeam", Type: 8192, Category: MoveCategorySpecial, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 65, Accuracy: 100},
	{ID: 61, Name: "Bubble Beam", Type: 1024, Category: MoveCategorySpecial, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 65, Accuracy: 100},
	{ID: 62, Name: "Aurora Beam", Type: 16384, Category: MoveCategorySpecial, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 65, Accuracy: 100},
	{ID: 63, Name: "Hyper Beam", Type: 1, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 150, Accuracy: 90},
	{ID: 64, Name: "Peck", Type: 4, Category: MoveCategoryPhysical, CurrentPP: 35, MaxPP: 35, Targets: 10, Priority: 0, Power: 35, Accuracy: 100},
	{ID: 65, Name: "Drill Peck", Type: 4, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 80, Accuracy: 100},
	{ID: 66, Name: "Submission", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 80, Accuracy: 80},
	{ID: 67, Name: "Low Kick", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 68, Name: "Counter", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 1, Priority: -5, Power: 0, Accuracy: 100},
	{ID: 69, Name: "Seismic Toss", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 70, Name: "Strength", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 80, Accuracy: 100},
	{ID: 71, Name: "Absorb", Type: 2048, Category: MoveCategorySpecial, CurrentPP: 25, MaxPP: 25, Targets: 10, Priority: 0, Power: 20, Accuracy: 100},
	{ID: 72, Name: "Mega Drain", Type: 2048, Category: MoveCategorySpecial, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 40, Accuracy: 100},
	{ID: 73, Name: "Leech Seed", Type: 2048, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 90},
	{ID: 74, Name: "Growth", Type: 1, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 75, Name: "Razor Leaf", Type: 2048, Category: MoveCategoryPhysical, CurrentPP: 25, MaxPP: 25, Targets: 11, Priority: 0, Power: 55, Accuracy: 95},
	{ID: 76, Name: "Solar Beam", Type: 2048, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 120, Accuracy: 100},
	{ID: 77, Name: "Poison Powder", Type: 8, Category: MoveCategoryStatus, CurrentPP: 35, MaxPP: 35, Targets: 10, Priority: 0, Power: 0, Accuracy: 75},
	{ID: 78, Name: "Stun Spore", Type: 2048, Category: MoveCategoryStatus, CurrentPP: 30, MaxPP: 30, Targets: 10, Priority: 0, Power: 0, Accuracy: 75},
	{ID: 79, Name: "Sleep Powder", Type: 2048, Category: MoveCategoryStatus, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 0, Accuracy: 75},
	{ID: 80, Name: "Petal Dance", Type: 2048, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 8, Priority: 0, Power: 120, Accuracy: 100},
	{ID: 81, Name: "String Shot", Type: 64, Category: MoveCategoryStatus, CurrentPP: 40, MaxPP: 40, Targets: 11, Priority: 0, Power: 0, Accuracy: 95},
	{ID: 82, Name: "Dragon Rage", Type: 32768, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 83, Name: "Fire Spin", Type: 512, Category: MoveCategorySpecial, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 35, Accuracy: 85},
	{ID: 84, Name: "Thunder Shock", Type: 4096, Category: MoveCategorySpecial, CurrentPP: 30, MaxPP: 30, Targets: 10, Priority: 0, Power: 40, Accuracy: 100},
	{ID: 85, Name: "Thunderbolt", Type: 4096, Category: MoveCategorySpecial, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 90, Accuracy: 100},
	{ID: 86, Name: "Thunder Wave", Type: 4096, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 0, Accuracy: 90},
	{ID: 87, Name: "Thunder", Type: 4096, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 110, Accuracy: 70},
	{ID: 88, Name: "Rock Throw", Type: 32, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 50, Accuracy: 90},
	{ID: 89, Name: "Earthquake", Type: 16, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 9, Priority: 0, Power: 100, Accuracy: 100},
	{ID: 90, Name: "Fissure", Type: 16, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 0, Accuracy: 30},
	{ID: 91, Name: "Dig", Type: 16, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 80, Accuracy: 100},
	{ID: 92, Name: "Toxic", Type: 8, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 90},
	{ID: 93, Name: "Confusion", Type: 8192, Category: MoveCategorySpecial, CurrentPP: 25, MaxPP: 25, Targets: 10, Priority: 0, Power: 50, Accuracy: 100},
	{ID: 94, Name: "Psychic", Type: 8192, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 90, Accuracy: 100},
	{ID: 95, Name: "Hypnosis", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 0, Accuracy: 60},
	{ID: 96, Name: "Meditate", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 40, MaxPP: 40, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 97, Name: "Agility", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 30, MaxPP: 30, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 98, Name: "Quick Attack", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 30, MaxPP: 30, Targets: 10, Priority: 1, Power: 40, Accuracy: 100},
	{ID: 99, Name: "Rage", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 20, Accuracy: 100},
	{ID: 100, Name: "Teleport", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 101, Name: "Night Shade", Type: 128, Category: MoveCategorySpecial, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 102, Name: "Mimic", Type: 1, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 103, Name: "Screech", Type: 1, Category: MoveCategoryStatus, CurrentPP: 40, MaxPP: 40, Targets: 10, Priority: 0, Power: 0, Accuracy: 85},
	{ID: 104, Name: "Double Team", Type: 1, Category: MoveCategoryStatus, CurrentPP: 15, MaxPP: 15, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 105, Name: "Recover", Type: 1, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 106, Name: "Harden", Type: 1, Category: MoveCategoryStatus, CurrentPP: 30, MaxPP: 30, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 107, Name: "Minimize", Type: 1, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 108, Name: "Smokescreen", Type: 1, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 109, Name: "Confuse Ray", Type: 128, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 110, Name: "Withdraw", Type: 1024, Category: MoveCategoryStatus, CurrentPP: 40, MaxPP: 40, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 111, Name: "Defense Curl", Type: 1, Category: MoveCategoryStatus, CurrentPP: 40, MaxPP: 40, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 112, Name: "Barrier", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 113, Name: "Light Screen", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 30, MaxPP: 30, Targets: 4, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 114, Name: "Haze", Type: 16384, Category: MoveCategoryStatus, CurrentPP: 30, MaxPP: 30, Targets: 12, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 115, Name: "Reflect", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 4, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 116, Name: "Focus Energy", Type: 1, Category: MoveCategoryStatus, CurrentPP: 30, MaxPP: 30, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 117, Name: "Bide", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 1, Power: 0, Accuracy: 0},
	{ID: 118, Name: "Metronome", Type: 1, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 119, Name: "Mirror Move", Type: 4, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 120, Name: "Self-Destruct", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 9, Priority: 0, Power: 200, Accuracy: 100},
	{ID: 121, Name: "Egg Bomb", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 100, Accuracy: 75},
	{ID: 122, Name: "Lick", Type: 128, Category: MoveCategoryPhysical, CurrentPP: 30, MaxPP: 30, Targets: 10, Priority: 0, Power: 30, Accuracy: 100},
	{ID: 123, Name: "Smog", Type: 8, Category: MoveCategorySpecial, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 30, Accuracy: 70},
	{ID: 124, Name: "Sludge", Type: 8, Category: MoveCategorySpecial, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 65, Accuracy: 100},
	{ID: 125, Name: "Bone Club", Type: 16, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 65, Accuracy: 85},
	{ID: 126, Name: "Fire Blast", Type: 512, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 110, Accuracy: 85},
	{ID: 127, Name: "Waterfall", Type: 1024, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 80, Accuracy: 100},
	{ID: 128, Name: "Clamp", Type: 1024, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 35, Accuracy: 85},
	{ID: 129, Name: "Swift", Type: 1, Category: MoveCategorySpecial, CurrentPP: 20, MaxPP: 20, Targets: 11, Priority: 0, Power: 60, Accuracy: 0},
	{ID: 130, Name: "Skull Bash", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 130, Accuracy: 100},
	{ID: 131, Name: "Spike Cannon", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 20, Accuracy: 100},
	{ID: 132, Name: "Constrict", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 35, MaxPP: 35, Targets: 10, Priority: 0, Power: 10, Accuracy: 100},
	{ID: 133, Name: "Amnesia", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 134, Name: "Kinesis", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 0, Accuracy: 80},
	{ID: 135, Name: "Soft-Boiled", Type: 1, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 136, Name: "High Jump Kick", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 130, Accuracy: 90},
	{ID: 137, Name: "Glare", Type: 1, Category: MoveCategoryStatus, CurrentPP: 30, MaxPP: 30, Targets: 10, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 138, Name: "Dream Eater", Type: 8192, Category: MoveCategorySpecial, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 100, Accuracy: 100},
	{ID: 139, Name: "Poison Gas", Type: 8, Category: MoveCategoryStatus, CurrentPP: 40, MaxPP: 40, Targets: 11, Priority: 0, Power: 0, Accuracy: 90},
	{ID: 140, Name: "Barrage", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 15, Accuracy: 85},
	{ID: 141, Name: "Leech Life", Type: 64, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 80, Accuracy: 100},
	{ID: 142, Name: "Lovely Kiss", Type: 1, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 75},
	{ID: 143, Name: "Sky Attack", Type: 4, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 140, Accuracy: 90},
	{ID: 144, Name: "Transform", Type: 1, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 145, Name: "Bubble", Type: 1024, Category: MoveCategorySpecial, CurrentPP: 30, MaxPP: 30, Targets: 11, Priority: 0, Power: 40, Accuracy: 100},
	{ID: 146, Name: "Dizzy Punch", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 70, Accuracy: 100},
	{ID: 147, Name: "Spore", Type: 2048, Category: MoveCategoryStatus, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 148, Name: "Flash", Type: 1, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 149, Name: "Psywave", Type: 8192, Category: MoveCategorySpecial, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 150, Name: "Splash", Type: 1, Category: MoveCategoryStatus, CurrentPP: 40, MaxPP: 40, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 151, Name: "Acid Armor", Type: 8, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 152, Name: "Crabhammer", Type: 1024, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 100, Accuracy: 90},
	{ID: 153, Name: "Explosion", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 9, Priority: 0, Power: 250, Accuracy: 100},
	{ID: 154, Name: "Fury Swipes", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 18, Accuracy: 80},
	{ID: 155, Name: "Bonemerang", Type: 16, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 50, Accuracy: 90},
	{ID: 156, Name: "Rest", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 157, Name: "Rock Slide", Type: 32, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 11, Priority: 0, Power: 75, Accuracy: 90},
	{ID: 158, Name: "Hyper Fang", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 80, Accuracy: 90},
	{ID: 159, Name: "Sharpen", Type: 1, Category: MoveCategoryStatus, CurrentPP: 30, MaxPP: 30, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 160, Name: "Conversion", Type: 1, Category: MoveCategoryStatus, CurrentPP: 30, MaxPP: 30, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 161, Name: "Tri Attack", Type: 1, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 80, Accuracy: 100},
	{ID: 162, Name: "Super Fang", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 90},
	{ID: 163, Name: "Slash", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 70, Accuracy: 100},
	{ID: 164, Name: "Substitute", Type: 1, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 165, Name: "Struggle", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 1, MaxPP: 1, Targets: 8, Priority: 0, Power: 50, Accuracy: 0},
	{ID: 166, Name: "Sketch", Type: 1, Category: MoveCategoryStatus, CurrentPP: 1, MaxPP: 1, Targets: 10, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 167, Name: "Triple Kick", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 10, Accuracy: 90},
	{ID: 168, Name: "Thief", Type: 65536, Category: MoveCategoryPhysical, CurrentPP: 25, MaxPP: 25, Targets: 10, Priority: 0, Power: 60, Accuracy: 100},
	{ID: 169, Name: "Spider Web", Type: 64, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 170, Name: "Mind Reader", Type: 1, Category: MoveCategoryStatus, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 171, Name: "Nightmare", Type: 128, Category: MoveCategoryStatus, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 172, Name: "Flame Wheel", Type: 512, Category: MoveCategoryPhysical, CurrentPP: 25, MaxPP: 25, Targets: 10, Priority: 0, Power: 60, Accuracy: 100},
	{ID: 173, Name: "Snore", Type: 1, Category: MoveCategorySpecial, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 50, Accuracy: 100},
	{ID: 174, Name: "Curse", Type: 128, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 1, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 175, Name: "Flail", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 176, Name: "Conversion 2", Type: 1, Category: MoveCategoryStatus, CurrentPP: 30, MaxPP: 30, Targets: 10, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 177, Name: "Aeroblast", Type: 4, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 100, Accuracy: 95},
	{ID: 178, Name: "Cotton Spore", Type: 2048, Category: MoveCategoryStatus, CurrentPP: 40, MaxPP: 40, Targets: 11, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 179, Name: "Reversal", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 180, Name: "Spite", Type: 128, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 181, Name: "Powder Snow", Type: 16384, Category: MoveCategorySpecial, CurrentPP: 25, MaxPP: 25, Targets: 11, Priority: 0, Power: 40, Accuracy: 100},
	{ID: 182, Name: "Protect", Type: 1, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 4, Power: 0, Accuracy: 0},
	{ID: 183, Name: "Mach Punch", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 30, MaxPP: 30, Targets: 10, Priority: 1, Power: 40, Accuracy: 100},
	{ID: 184, Name: "Scary Face", Type: 1, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 185, Name: "Feint Attack", Type: 65536, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 60, Accuracy: 0},
	{ID: 186, Name: "Sweet Kiss", Type: 131072, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 75},
	{ID: 187, Name: "Belly Drum", Type: 1, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 188, Name: "Sludge Bomb", Type: 8, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 90, Accuracy: 100},
	{ID: 189, Name: "Mud-Slap", Type: 16, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 20, Accuracy: 100},
	{ID: 190, Name: "Octazooka", Type: 1024, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 65, Accuracy: 85},
	{ID: 191, Name: "Spikes", Type: 16, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 6, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 192, Name: "Zap Cannon", Type: 4096, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 120, Accuracy: 50},
	{ID: 193, Name: "Foresight", Type: 1, Category: MoveCategoryStatus, CurrentPP: 40, MaxPP: 40, Targets: 10, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 194, Name: "Destiny Bond", Type: 128, Category: MoveCategoryStatus, CurrentPP: 5, MaxPP: 5, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 195, Name: "Perish Song", Type: 1, Category: MoveCategoryStatus, CurrentPP: 5, MaxPP: 5, Targets: 14, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 196, Name: "Icy Wind", Type: 16384, Category: MoveCategorySpecial, CurrentPP: 15, MaxPP: 15, Targets: 11, Priority: 0, Power: 55, Accuracy: 95},
	{ID: 197, Name: "Detect", Type: 2, Category: MoveCategoryStatus, CurrentPP: 5, MaxPP: 5, Targets: 7, Priority: 4, Power: 0, Accuracy: 0},
	{ID: 198, Name: "Bone Rush", Type: 16, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 25, Accuracy: 90},
	{ID: 199, Name: "Lock-On", Type: 1, Category: MoveCategoryStatus, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 200, Name: "Outrage", Type: 32768, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 8, Priority: 0, Power: 120, Accuracy: 100},
	{ID: 201, Name: "Sandstorm", Type: 32, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 12, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 202, Name: "Giga Drain", Type: 2048, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 75, Accuracy: 100},
	{ID: 203, Name: "Endure", Type: 1, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 4, Power: 0, Accuracy: 0},
	{ID: 204, Name: "Charm", Type: 131072, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 205, Name: "Rollout", Type: 32, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 30, Accuracy: 90},
	{ID: 206, Name: "False Swipe", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 40, MaxPP: 40, Targets: 10, Priority: 0, Power: 40, Accuracy: 100},
	{ID: 207, Name: "Swagger", Type: 1, Category: MoveCategoryStatus, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 0, Accuracy: 85},
	{ID: 208, Name: "Milk Drink", Type: 1, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 209, Name: "Spark", Type: 4096, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 65, Accuracy: 100},
	{ID: 210, Name: "Fury Cutter", Type: 64, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 40, Accuracy: 95},
	{ID: 211, Name: "Steel Wing", Type: 256, Category: MoveCategoryPhysical, CurrentPP: 25, MaxPP: 25, Targets: 10, Priority: 0, Power: 70, Accuracy: 90},
	{ID: 212, Name: "Mean Look", Type: 1, Category: MoveCategoryStatus, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 213, Name: "Attract", Type: 1, Category: MoveCategoryStatus, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 214, Name: "Sleep Talk", Type: 1, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 215, Name: "Heal Bell", Type: 1, Category: MoveCategoryStatus, CurrentPP: 5, MaxPP: 5, Targets: 13, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 216, Name: "Return", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 217, Name: "Present", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 0, Accuracy: 90},
	{ID: 218, Name: "Frustration", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 219, Name: "Safeguard", Type: 1, Category: MoveCategoryStatus, CurrentPP: 25, MaxPP: 25, Targets: 4, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 220, Name: "Pain Split", Type: 1, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 221, Name: "Sacred Fire", Type: 512, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 100, Accuracy: 95},
	{ID: 222, Name: "Magnitude", Type: 16, Category: MoveCategoryPhysical, CurrentPP: 30, MaxPP: 30, Targets: 9, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 223, Name: "Dynamic Punch", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 100, Accuracy: 50},
	{ID: 224, Name: "Megahorn", Type: 64, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 120, Accuracy: 85},
	{ID: 225, Name: "Dragon Breath", Type: 32768, Category: MoveCategorySpecial, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 60, Accuracy: 100},
	{ID: 226, Name: "Baton Pass", Type: 1, Category: MoveCategoryStatus, CurrentPP: 40, MaxPP: 40, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 227, Name: "Encore", Type: 1, Category: MoveCategoryStatus, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 228, Name: "Pursuit", Type: 65536, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 40, Accuracy: 100},
	{ID: 229, Name: "Rapid Spin", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 40, MaxPP: 40, Targets: 10, Priority: 0, Power: 20, Accuracy: 100},
	{ID: 230, Name: "Sweet Scent", Type: 1, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 11, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 231, Name: "Iron Tail", Type: 256, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 100, Accuracy: 75},
	{ID: 232, Name: "Metal Claw", Type: 256, Category: MoveCategoryPhysical, CurrentPP: 35, MaxPP: 35, Targets: 10, Priority: 0, Power: 50, Accuracy: 95},
	{ID: 233, Name: "Vital Throw", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: -1, Power: 70, Accuracy: 0},
	{ID: 234, Name: "Morning Sun", Type: 1, Category: MoveCategoryStatus, CurrentPP: 5, MaxPP: 5, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 235, Name: "Synthesis", Type: 2048, Category: MoveCategoryStatus, CurrentPP: 5, MaxPP: 5, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 236, Name: "Moonlight", Type: 131072, Category: MoveCategoryStatus, CurrentPP: 5, MaxPP: 5, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 237, Name: "Hidden Power", Type: 1, Category: MoveCategorySpecial, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 60, Accuracy: 100},
	{ID: 238, Name: "Cross Chop", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 100, Accuracy: 80},
	{ID: 239, Name: "Twister", Type: 32768, Category: MoveCategorySpecial, CurrentPP: 20, MaxPP: 20, Targets: 11, Priority: 0, Power: 40, Accuracy: 100},
	{ID: 240, Name: "Rain Dance", Type: 1024, Category: MoveCategoryStatus, CurrentPP: 5, MaxPP: 5, Targets: 12, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 241, Name: "Sunny Day", Type: 512, Category: MoveCategoryStatus, CurrentPP: 5, MaxPP: 5, Targets: 12, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 242, Name: "Crunch", Type: 65536, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 80, Accuracy: 100},
	{ID: 243, Name: "Mirror Coat", Type: 8192, Category: MoveCategorySpecial, CurrentPP: 20, MaxPP: 20, Targets: 1, Priority: -5, Power: 0, Accuracy: 100},
	{ID: 244, Name: "Psych Up", Type: 1, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 245, Name: "Extreme Speed", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 2, Power: 80, Accuracy: 100},
	{ID: 246, Name: "Ancient Power", Type: 32, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 60, Accuracy: 100},
	{ID: 247, Name: "Shadow Ball", Type: 128, Category: MoveCategorySpecial, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 80, Accuracy: 100},
	{ID: 248, Name: "Future Sight", Type: 8192, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 120, Accuracy: 100},
	{ID: 249, Name: "Rock Smash", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 40, Accuracy: 100},
	{ID: 250, Name: "Whirlpool", Type: 1024, Category: MoveCategorySpecial, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 35, Accuracy: 85},
	{ID: 251, Name: "Beat Up", Type: 65536, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 252, Name: "Fake Out", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 3, Power: 40, Accuracy: 100},
	{ID: 253, Name: "Uproar", Type: 1, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 8, Priority: 0, Power: 90, Accuracy: 100},
	{ID: 254, Name: "Stockpile", Type: 1, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 255, Name: "Spit Up", Type: 1, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 256, Name: "Swallow", Type: 1, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 257, Name: "Heat Wave", Type: 512, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 11, Priority: 0, Power: 95, Accuracy: 90},
	{ID: 258, Name: "Hail", Type: 16384, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 12, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 259, Name: "Torment", Type: 65536, Category: MoveCategoryStatus, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 260, Name: "Flatter", Type: 65536, Category: MoveCategoryStatus, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 261, Name: "Will-O-Wisp", Type: 512, Category: MoveCategoryStatus, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 0, Accuracy: 85},
	{ID: 262, Name: "Memento", Type: 65536, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 263, Name: "Facade", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 70, Accuracy: 100},
	{ID: 264, Name: "Focus Punch", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: -3, Power: 150, Accuracy: 100},
	{ID: 265, Name: "Smelling Salts", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 70, Accuracy: 100},
	{ID: 266, Name: "Follow Me", Type: 1, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 7, Priority: 2, Power: 0, Accuracy: 0},
	{ID: 267, Name: "Nature Power", Type: 1, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 268, Name: "Charge", Type: 4096, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 269, Name: "Taunt", Type: 65536, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 270, Name: "Helping Hand", Type: 1, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 3, Priority: 5, Power: 0, Accuracy: 0},
	{ID: 271, Name: "Trick", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 272, Name: "Role Play", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 273, Name: "Wish", Type: 1, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 274, Name: "Assist", Type: 1, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 275, Name: "Ingrain", Type: 2048, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 276, Name: "Superpower", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 120, Accuracy: 100},
	{ID: 277, Name: "Magic Coat", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 15, MaxPP: 15, Targets: 7, Priority: 4, Power: 0, Accuracy: 0},
	{ID: 278, Name: "Recycle", Type: 1, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 279, Name: "Revenge", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: -4, Power: 60, Accuracy: 100},
	{ID: 280, Name: "Brick Break", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 75, Accuracy: 100},
	{ID: 281, Name: "Yawn", Type: 1, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 282, Name: "Knock Off", Type: 65536, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 65, Accuracy: 100},
	{ID: 283, Name: "Endeavor", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 284, Name: "Eruption", Type: 512, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 11, Priority: 0, Power: 150, Accuracy: 100},
	{ID: 285, Name: "Skill Swap", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 286, Name: "Imprison", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 287, Name: "Refresh", Type: 1, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 288, Name: "Grudge", Type: 128, Category: MoveCategoryStatus, CurrentPP: 5, MaxPP: 5, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 289, Name: "Snatch", Type: 65536, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 4, Power: 0, Accuracy: 0},
	{ID: 290, Name: "Secret Power", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 70, Accuracy: 100},
	{ID: 291, Name: "Dive", Type: 1024, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 80, Accuracy: 100},
	{ID: 292, Name: "Arm Thrust", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 15, Accuracy: 100},
	{ID: 293, Name: "Camouflage", Type: 1, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 294, Name: "Tail Glow", Type: 64, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 295, Name: "Luster Purge", Type: 8192, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 70, Accuracy: 100},
	{ID: 296, Name: "Mist Ball", Type: 8192, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 70, Accuracy: 100},
	{ID: 297, Name: "Feather Dance", Type: 4, Category: MoveCategoryStatus, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 298, Name: "Teeter Dance", Type: 1, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 9, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 299, Name: "Blaze Kick", Type: 512, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 85, Accuracy: 90},
	{ID: 300, Name: "Mud Sport", Type: 16, Category: MoveCategoryStatus, CurrentPP: 15, MaxPP: 15, Targets: 12, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 301, Name: "Ice Ball", Type: 16384, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 30, Accuracy: 90},
	{ID: 302, Name: "Needle Arm", Type: 2048, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 60, Accuracy: 100},
	{ID: 303, Name: "Slack Off", Type: 1, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 304, Name: "Hyper Voice", Type: 1, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 11, Priority: 0, Power: 90, Accuracy: 100},
	{ID: 305, Name: "Poison Fang", Type: 8, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 50, Accuracy: 100},
	{ID: 306, Name: "Crush Claw", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 75, Accuracy: 95},
	{ID: 307, Name: "Blast Burn", Type: 512, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 150, Accuracy: 90},
	{ID: 308, Name: "Hydro Cannon", Type: 1024, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 150, Accuracy: 90},
	{ID: 309, Name: "Meteor Mash", Type: 256, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 90, Accuracy: 90},
	{ID: 310, Name: "Astonish", Type: 128, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 30, Accuracy: 100},
	{ID: 311, Name: "Weather Ball", Type: 1, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 50, Accuracy: 100},
	{ID: 312, Name: "Aromatherapy", Type: 2048, Category: MoveCategoryStatus, CurrentPP: 5, MaxPP: 5, Targets: 13, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 313, Name: "Fake Tears", Type: 65536, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 314, Name: "Air Cutter", Type: 4, Category: MoveCategorySpecial, CurrentPP: 25, MaxPP: 25, Targets: 11, Priority: 0, Power: 60, Accuracy: 95},
	{ID: 315, Name: "Overheat", Type: 512, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 130, Accuracy: 90},
	{ID: 316, Name: "Odor Sleuth", Type: 1, Category: MoveCategoryStatus, CurrentPP: 40, MaxPP: 40, Targets: 10, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 317, Name: "Rock Tomb", Type: 32, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 60, Accuracy: 95},
	{ID: 318, Name: "Silver Wind", Type: 64, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 60, Accuracy: 100},
	{ID: 319, Name: "Metal Sound", Type: 256, Category: MoveCategoryStatus, CurrentPP: 40, MaxPP: 40, Targets: 10, Priority: 0, Power: 0, Accuracy: 85},
	{ID: 320, Name: "Grass Whistle", Type: 2048, Category: MoveCategoryStatus, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 0, Accuracy: 55},
	{ID: 321, Name: "Tickle", Type: 1, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 322, Name: "Cosmic Power", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 323, Name: "Water Spout", Type: 1024, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 11, Priority: 0, Power: 150, Accuracy: 100},
	{ID: 324, Name: "Signal Beam", Type: 64, Category: MoveCategorySpecial, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 75, Accuracy: 100},
	{ID: 325, Name: "Shadow Punch", Type: 128, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 60, Accuracy: 0},
	{ID: 326, Name: "Extrasensory", Type: 8192, Category: MoveCategorySpecial, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 80, Accuracy: 100},
	{ID: 327, Name: "Sky Uppercut", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 85, Accuracy: 90},
	{ID: 328, Name: "Sand Tomb", Type: 16, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 35, Accuracy: 85},
	{ID: 329, Name: "Sheer Cold", Type: 16384, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 0, Accuracy: 30},
	{ID: 330, Name: "Muddy Water", Type: 1024, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 11, Priority: 0, Power: 90, Accuracy: 85},
	{ID: 331, Name: "Bullet Seed", Type: 2048, Category: MoveCategoryPhysical, CurrentPP: 30, MaxPP: 30, Targets: 10, Priority: 0, Power: 25, Accuracy: 100},
	{ID: 332, Name: "Aerial Ace", Type: 4, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 60, Accuracy: 0},
	{ID: 333, Name: "Icicle Spear", Type: 16384, Category: MoveCategoryPhysical, CurrentPP: 30, MaxPP: 30, Targets: 10, Priority: 0, Power: 25, Accuracy: 100},
	{ID: 334, Name: "Iron Defense", Type: 256, Category: MoveCategoryStatus, CurrentPP: 15, MaxPP: 15, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 335, Name: "Block", Type: 1, Category: MoveCategoryStatus, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 336, Name: "Howl", Type: 1, Category: MoveCategoryStatus, CurrentPP: 40, MaxPP: 40, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 337, Name: "Dragon Claw", Type: 32768, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 80, Accuracy: 100},
	{ID: 338, Name: "Frenzy Plant", Type: 2048, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 150, Accuracy: 90},
	{ID: 339, Name: "Bulk Up", Type: 2, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 340, Name: "Bounce", Type: 4, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 85, Accuracy: 85},
	{ID: 341, Name: "Mud Shot", Type: 16, Category: MoveCategorySpecial, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 55, Accuracy: 95},
	{ID: 342, Name: "Poison Tail", Type: 8, Category: MoveCategoryPhysical, CurrentPP: 25, MaxPP: 25, Targets: 10, Priority: 0, Power: 50, Accuracy: 100},
	{ID: 343, Name: "Covet", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 25, MaxPP: 25, Targets: 10, Priority: 0, Power: 60, Accuracy: 100},
	{ID: 344, Name: "Volt Tackle", Type: 4096, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 120, Accuracy: 100},
	{ID: 345, Name: "Magical Leaf", Type: 2048, Category: MoveCategorySpecial, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 60, Accuracy: 0},
	{ID: 346, Name: "Water Sport", Type: 1024, Category: MoveCategoryStatus, CurrentPP: 15, MaxPP: 15, Targets: 12, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 347, Name: "Calm Mind", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 348, Name: "Leaf Blade", Type: 2048, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 90, Accuracy: 100},
	{ID: 349, Name: "Dragon Dance", Type: 32768, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 350, Name: "Rock Blast", Type: 32, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 25, Accuracy: 90},
	{ID: 351, Name: "Shock Wave", Type: 4096, Category: MoveCategorySpecial, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 60, Accuracy: 0},
	{ID: 352, Name: "Water Pulse", Type: 1024, Category: MoveCategorySpecial, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 60, Accuracy: 100},
	{ID: 353, Name: "Doom Desire", Type: 256, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 140, Accuracy: 100},
	{ID: 354, Name: "Psycho Boost", Type: 8192, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 140, Accuracy: 90},
	{ID: 355, Name: "Roost", Type: 4, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 356, Name: "Gravity", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 5, MaxPP: 5, Targets: 12, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 357, Name: "Miracle Eye", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 40, MaxPP: 40, Targets: 10, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 358, Name: "Wake-Up Slap", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 70, Accuracy: 100},
	{ID: 359, Name: "Hammer Arm", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 100, Accuracy: 90},
	{ID: 360, Name: "Gyro Ball", Type: 256, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 361, Name: "Healing Wish", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 362, Name: "Brine", Type: 1024, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 65, Accuracy: 100},
	{ID: 363, Name: "Natural Gift", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 364, Name: "Feint", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 2, Power: 30, Accuracy: 100},
	{ID: 365, Name: "Pluck", Type: 4, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 60, Accuracy: 100},
	{ID: 366, Name: "Tailwind", Type: 4, Category: MoveCategoryStatus, CurrentPP: 15, MaxPP: 15, Targets: 4, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 367, Name: "Acupressure", Type: 1, Category: MoveCategoryStatus, CurrentPP: 30, MaxPP: 30, Targets: 5, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 368, Name: "Metal Burst", Type: 256, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 1, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 369, Name: "U-turn", Type: 64, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 70, Accuracy: 100},
	{ID: 370, Name: "Close Combat", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 120, Accuracy: 100},
	{ID: 371, Name: "Payback", Type: 65536, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 50, Accuracy: 100},
	{ID: 372, Name: "Assurance", Type: 65536, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 60, Accuracy: 100},
	{ID: 373, Name: "Embargo", Type: 65536, Category: MoveCategoryStatus, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 374, Name: "Fling", Type: 65536, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 375, Name: "Psycho Shift", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 376, Name: "Trump Card", Type: 1, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 377, Name: "Heal Block", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 15, MaxPP: 15, Targets: 11, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 378, Name: "Wring Out", Type: 1, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 379, Name: "Power Trick", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 380, Name: "Gastro Acid", Type: 8, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 381, Name: "Lucky Chant", Type: 1, Category: MoveCategoryStatus, CurrentPP: 30, MaxPP: 30, Targets: 4, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 382, Name: "Me First", Type: 1, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 2, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 383, Name: "Copycat", Type: 1, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 384, Name: "Power Swap", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 385, Name: "Guard Swap", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 386, Name: "Punishment", Type: 65536, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 387, Name: "Last Resort", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 140, Accuracy: 100},
	{ID: 388, Name: "Worry Seed", Type: 2048, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 389, Name: "Sucker Punch", Type: 65536, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 1, Power: 70, Accuracy: 100},
	{ID: 390, Name: "Toxic Spikes", Type: 8, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 6, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 391, Name: "Heart Swap", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 392, Name: "Aqua Ring", Type: 1024, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 393, Name: "Magnet Rise", Type: 4096, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 394, Name: "Flare Blitz", Type: 512, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 120, Accuracy: 100},
	{ID: 395, Name: "Force Palm", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 60, Accuracy: 100},
	{ID: 396, Name: "Aura Sphere", Type: 2, Category: MoveCategorySpecial, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 80, Accuracy: 0},
	{ID: 397, Name: "Rock Polish", Type: 32, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 398, Name: "Poison Jab", Type: 8, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 80, Accuracy: 100},
	{ID: 399, Name: "Dark Pulse", Type: 65536, Category: MoveCategorySpecial, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 80, Accuracy: 100},
	{ID: 400, Name: "Night Slash", Type: 65536, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 70, Accuracy: 100},
	{ID: 401, Name: "Aqua Tail", Type: 1024, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 90, Accuracy: 90},
	{ID: 402, Name: "Seed Bomb", Type: 2048, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 80, Accuracy: 100},
	{ID: 403, Name: "Air Slash", Type: 4, Category: MoveCategorySpecial, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 75, Accuracy: 95},
	{ID: 404, Name: "X-Scissor", Type: 64, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 80, Accuracy: 100},
	{ID: 405, Name: "Bug Buzz", Type: 64, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 90, Accuracy: 100},
	{ID: 406, Name: "Dragon Pulse", Type: 32768, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 85, Accuracy: 100},
	{ID: 407, Name: "Dragon Rush", Type: 32768, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 100, Accuracy: 75},
	{ID: 408, Name: "Power Gem", Type: 32, Category: MoveCategorySpecial, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 80, Accuracy: 100},
	{ID: 409, Name: "Drain Punch", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 75, Accuracy: 100},
	{ID: 410, Name: "Vacuum Wave", Type: 2, Category: MoveCategorySpecial, CurrentPP: 30, MaxPP: 30, Targets: 10, Priority: 1, Power: 40, Accuracy: 100},
	{ID: 411, Name: "Focus Blast", Type: 2, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 120, Accuracy: 70},
	{ID: 412, Name: "Energy Ball", Type: 2048, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 90, Accuracy: 100},
	{ID: 413, Name: "Brave Bird", Type: 4, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 120, Accuracy: 100},
	{ID: 414, Name: "Earth Power", Type: 16, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 90, Accuracy: 100},
	{ID: 415, Name: "Switcheroo", Type: 65536, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 416, Name: "Giga Impact", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 150, Accuracy: 90},
	{ID: 417, Name: "Nasty Plot", Type: 65536, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 418, Name: "Bullet Punch", Type: 256, Category: MoveCategoryPhysical, CurrentPP: 30, MaxPP: 30, Targets: 10, Priority: 1, Power: 40, Accuracy: 100},
	{ID: 419, Name: "Avalanche", Type: 16384, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: -4, Power: 60, Accuracy: 100},
	{ID: 420, Name: "Ice Shard", Type: 16384, Category: MoveCategoryPhysical, CurrentPP: 30, MaxPP: 30, Targets: 10, Priority: 1, Power: 40, Accuracy: 100},
	{ID: 421, Name: "Shadow Claw", Type: 128, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 70, Accuracy: 100},
	{ID: 422, Name: "Thunder Fang", Type: 4096, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 65, Accuracy: 95},
	{ID: 423, Name: "Ice Fang", Type: 16384, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 65, Accuracy: 95},
	{ID: 424, Name: "Fire Fang", Type: 512, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 65, Accuracy: 95},
	{ID: 425, Name: "Shadow Sneak", Type: 128, Category: MoveCategoryPhysical, CurrentPP: 30, MaxPP: 30, Targets: 10, Priority: 1, Power: 40, Accuracy: 100},
	{ID: 426, Name: "Mud Bomb", Type: 16, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 65, Accuracy: 85},
	{ID: 427, Name: "Psycho Cut", Type: 8192, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 70, Accuracy: 100},
	{ID: 428, Name: "Zen Headbutt", Type: 8192, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 80, Accuracy: 90},
	{ID: 429, Name: "Mirror Shot", Type: 256, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 65, Accuracy: 85},
	{ID: 430, Name: "Flash Cannon", Type: 256, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 80, Accuracy: 100},
	{ID: 431, Name: "Rock Climb", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 90, Accuracy: 85},
	{ID: 432, Name: "Defog", Type: 4, Category: MoveCategoryStatus, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 433, Name: "Trick Room", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 5, MaxPP: 5, Targets: 12, Priority: -7, Power: 0, Accuracy: 0},
	{ID: 434, Name: "Draco Meteor", Type: 32768, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 130, Accuracy: 90},
	{ID: 435, Name: "Discharge", Type: 4096, Category: MoveCategorySpecial, CurrentPP: 15, MaxPP: 15, Targets: 9, Priority: 0, Power: 80, Accuracy: 100},
	{ID: 436, Name: "Lava Plume", Type: 512, Category: MoveCategorySpecial, CurrentPP: 15, MaxPP: 15, Targets: 9, Priority: 0, Power: 80, Accuracy: 100},
	{ID: 437, Name: "Leaf Storm", Type: 2048, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 130, Accuracy: 90},
	{ID: 438, Name: "Power Whip", Type: 2048, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 120, Accuracy: 85},
	{ID: 439, Name: "Rock Wrecker", Type: 32, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 150, Accuracy: 90},
	{ID: 440, Name: "Cross Poison", Type: 8, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 70, Accuracy: 100},
	{ID: 441, Name: "Gunk Shot", Type: 8, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 120, Accuracy: 80},
	{ID: 442, Name: "Iron Head", Type: 256, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 80, Accuracy: 100},
	{ID: 443, Name: "Magnet Bomb", Type: 256, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 60, Accuracy: 0},
	{ID: 444, Name: "Stone Edge", Type: 32, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 100, Accuracy: 80},
	{ID: 445, Name: "Captivate", Type: 1, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 11, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 446, Name: "Stealth Rock", Type: 32, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 6, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 447, Name: "Grass Knot", Type: 2048, Category: MoveCategorySpecial, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 448, Name: "Chatter", Type: 4, Category: MoveCategorySpecial, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 65, Accuracy: 100},
	{ID: 449, Name: "Judgment", Type: 1, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 100, Accuracy: 100},
	{ID: 450, Name: "Bug Bite", Type: 64, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 60, Accuracy: 100},
	{ID: 451, Name: "Charge Beam", Type: 4096, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 50, Accuracy: 90},
	{ID: 452, Name: "Wood Hammer", Type: 2048, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 120, Accuracy: 100},
	{ID: 453, Name: "Aqua Jet", Type: 1024, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 1, Power: 40, Accuracy: 100},
	{ID: 454, Name: "Attack Order", Type: 64, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 90, Accuracy: 100},
	{ID: 455, Name: "Defend Order", Type: 64, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 456, Name: "Heal Order", Type: 64, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 457, Name: "Head Smash", Type: 32, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 150, Accuracy: 80},
	{ID: 458, Name: "Double Hit", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 35, Accuracy: 90},
	{ID: 459, Name: "Roar of Time", Type: 32768, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 150, Accuracy: 90},
	{ID: 460, Name: "Spacial Rend", Type: 32768, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 100, Accuracy: 95},
	{ID: 461, Name: "Lunar Dance", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 462, Name: "Crush Grip", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 463, Name: "Magma Storm", Type: 512, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 100, Accuracy: 75},
	{ID: 464, Name: "Dark Void", Type: 65536, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 11, Priority: 0, Power: 0, Accuracy: 50},
	{ID: 465, Name: "Seed Flare", Type: 2048, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 120, Accuracy: 85},
	{ID: 466, Name: "Ominous Wind", Type: 128, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 60, Accuracy: 100},
	{ID: 467, Name: "Shadow Force", Type: 128, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 120, Accuracy: 100},
	{ID: 10001, Name: "Shadow Rush", Type: 0, Category: MoveCategoryPhysical, CurrentPP: 0, MaxPP: 0, Targets: 10, Priority: 0, Power: 55, Accuracy: 100},
	{ID: 10002, Name: "Shadow Blast", Type: 0, Category: MoveCategoryPhysical, CurrentPP: 0, MaxPP: 0, Targets: 10, Priority: 0, Power: 80, Accuracy: 100},
	{ID: 10003, Name: "Shadow Blitz", Type: 0, Category: MoveCategoryPhysical, CurrentPP: 0, MaxPP: 0, Targets: 10, Priority: 0, Power: 40, Accuracy: 100},
	{ID: 10004, Name: "Shadow Bolt", Type: 0, Category: MoveCategorySpecial, CurrentPP: 0, MaxPP: 0, Targets: 10, Priority: 0, Power: 75, Accuracy: 100},
	{ID: 10005, Name: "Shadow Break", Type: 0, Category: MoveCategoryPhysical, CurrentPP: 0, MaxPP: 0, Targets: 10, Priority: 0, Power: 75, Accuracy: 100},
	{ID: 10006, Name: "Shadow Chill", Type: 0, Category: MoveCategorySpecial, CurrentPP: 0, MaxPP: 0, Targets: 10, Priority: 0, Power: 75, Accuracy: 100},
	{ID: 10007, Name: "Shadow End", Type: 0, Category: MoveCategoryPhysical, CurrentPP: 0, MaxPP: 0, Targets: 10, Priority: 0, Power: 120, Accuracy: 60},
	{ID: 10008, Name: "Shadow Fire", Type: 0, Category: MoveCategorySpecial, CurrentPP: 0, MaxPP: 0, Targets: 10, Priority: 0, Power: 75, Accuracy: 100},
	{ID: 10009, Name: "Shadow Rave", Type: 0, Category: MoveCategorySpecial, CurrentPP: 0, MaxPP: 0, Targets: 6, Priority: 0, Power: 70, Accuracy: 100},
	{ID: 10010, Name: "Shadow Storm", Type: 0, Category: MoveCategorySpecial, CurrentPP: 0, MaxPP: 0, Targets: 6, Priority: 0, Power: 95, Accuracy: 100},
	{ID: 10011, Name: "Shadow Wave", Type: 0, Category: MoveCategorySpecial, CurrentPP: 0, MaxPP: 0, Targets: 6, Priority: 0, Power: 50, Accuracy: 100},
	{ID: 10012, Name: "Shadow Down", Type: 0, Category: MoveCategoryStatus, CurrentPP: 0, MaxPP: 0, Targets: 6, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 10013, Name: "Shadow Half", Type: 0, Category: MoveCategorySpecial, CurrentPP: 0, MaxPP: 0, Targets: 12, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 10014, Name: "Shadow Hold", Type: 0, Category: MoveCategoryStatus, CurrentPP: 0, MaxPP: 0, Targets: 6, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 10015, Name: "Shadow Mist", Type: 0, Category: MoveCategoryStatus, CurrentPP: 0, MaxPP: 0, Targets: 6, Priority: 0, Power: 0, Accuracy: 100},
	{ID: 10016, Name: "Shadow Panic", Type: 0, Category: MoveCategoryStatus, CurrentPP: 0, MaxPP: 0, Targets: 6, Priority: 0, Power: 0, Accuracy: 90},
	{ID: 10017, Name: "Shadow Shed", Type: 0, Category: MoveCategoryStatus, CurrentPP: 0, MaxPP: 0, Targets: 12, Priority: 0, Power: 0, Accuracy: 0},
	{ID: 10018, Name: "Shadow Sky", Type: 0, Category: MoveCategoryStatus, CurrentPP: 0, MaxPP: 0, Targets: 12, Priority: 0, Power: 0, Accuracy: 0},
}

// Create move constant enum for quick reference
const (
	MovePound         = 1
	MoveKarateChop    = 2
	MoveDoubleSlap    = 3
	MoveCometPunch    = 4
	MoveMegaPunch     = 5
	MovePayDay        = 6
	MoveFirePunch     = 7
	MoveIcePunch      = 8
	MoveThunderPunch  = 9
	MoveScratch       = 10
	MoveViceGrip      = 11
	MoveGuillotine    = 12
	MoveRazorWind     = 13
	MoveSwordsDance   = 14
	MoveCut           = 15
	MoveGust          = 16
	MoveWingAttack    = 17
	MoveWhirlwind     = 18
	MoveFly           = 19
	MoveBind          = 20
	MoveSlam          = 21
	MoveVineWhip      = 22
	MoveStomp         = 23
	MoveDoubleKick    = 24
	MoveMegaKick      = 25
	MoveJumpKick      = 26
	MoveRollingKick   = 27
	MoveSandAttack    = 28
	MoveHeadbutt      = 29
	MoveHornAttack    = 30
	MoveFuryAttack    = 31
	MoveHornDrill     = 32
	MoveTackle        = 33
	MoveBodySlam      = 34
	MoveWrap          = 35
	MoveTakeDown      = 36
	MoveThrash        = 37
	MoveDoubleEdge    = 38
	MoveTailWhip      = 39
	MovePoisonSting   = 40
	MoveTwineedle     = 41
	MovePinMissile    = 42
	MoveLeer          = 43
	MoveBite          = 44
	MoveGrowl         = 45
	MoveRoar          = 46
	MoveSing          = 47
	MoveSupersonic    = 48
	MoveSonicBoom     = 49
	MoveDisable       = 50
	MoveAcid          = 51
	MoveEmber         = 52
	MoveFlamethrower  = 53
	MoveMist          = 54
	MoveWaterGun      = 55
	MoveHydroPump     = 56
	MoveSurf          = 57
	MoveIceBeam       = 58
	MoveBlizzard      = 59
	MovePsybeam       = 60
	MoveBubbleBeam    = 61
	MoveAuroraBeam    = 62
	MoveHyperBeam     = 63
	MovePeck          = 64
	MoveDrillPeck     = 65
	MoveSubmission    = 66
	MoveLowKick       = 67
	MoveCounter       = 68
	MoveSeismicToss   = 69
	MoveStrength      = 70
	MoveAbsorb        = 71
	MoveMegaDrain     = 72
	MoveLeechSeed     = 73
	MoveGrowth        = 74
	MoveRazorLeaf     = 75
	MoveSolarBeam     = 76
	MovePoisonPowder  = 77
	MoveStunSpore     = 78
	MoveSleepPowder   = 79
	MovePetalDance    = 80
	MoveStringShot    = 81
	MoveDragonRage    = 82
	MoveFireSpin      = 83
	MoveThunderShock  = 84
	MoveThunderbolt   = 85
	MoveThunderWave   = 86
	MoveThunder       = 87
	MoveRockThrow     = 88
	MoveEarthquake    = 89
	MoveFissure       = 90
	MoveDig           = 91
	MoveToxic         = 92
	MoveConfusion     = 93
	MovePsychic       = 94
	MoveHypnosis      = 95
	MoveMeditate      = 96
	MoveAgility       = 97
	MoveQuickAttack   = 98
	MoveRage          = 99
	MoveTeleport      = 100
	MoveNightShade    = 101
	MoveMimic         = 102
	MoveScreech       = 103
	MoveDoubleTeam    = 104
	MoveRecover       = 105
	MoveHarden        = 106
	MoveMinimize      = 107
	MoveSmokescreen   = 108
	MoveConfuseRay    = 109
	MoveWithdraw      = 110
	MoveDefenseCurl   = 111
	MoveBarrier       = 112
	MoveLightScreen   = 113
	MoveHaze          = 114
	MoveReflect       = 115
	MoveFocusEnergy   = 116
	MoveBide          = 117
	MoveMetronome     = 118
	MoveMirrorMove    = 119
	MoveSelfDestruct  = 120
	MoveEggBomb       = 121
	MoveLick          = 122
	MoveSmog          = 123
	MoveSludge        = 124
	MoveBoneClub      = 125
	MoveFireBlast     = 126
	MoveWaterfall     = 127
	MoveClamp         = 128
	MoveSwift         = 129
	MoveSkullBash     = 130
	MoveSpikeCannon   = 131
	MoveConstrict     = 132
	MoveAmnesia       = 133
	MoveKinesis       = 134
	MoveSoftBoiled    = 135
	MoveHighJumpKick  = 136
	MoveGlare         = 137
	MoveDreamEater    = 138
	MovePoisonGas     = 139
	MoveBarrage       = 140
	MoveLeechLife     = 141
	MoveLovelyKiss    = 142
	MoveSkyAttack     = 143
	MoveTransform     = 144
	MoveBubble        = 145
	MoveDizzyPunch    = 146
	MoveSpore         = 147
	MoveFlash         = 148
	MovePsywave       = 149
	MoveSplash        = 150
	MoveAcidArmor     = 151
	MoveCrabhammer    = 152
	MoveExplosion     = 153
	MoveFurySwipes    = 154
	MoveBonemerang    = 155
	MoveRest          = 156
	MoveRockSlide     = 157
	MoveHyperFang     = 158
	MoveSharpen       = 159
	MoveConversion    = 160
	MoveTriAttack     = 161
	MoveSuperFang     = 162
	MoveSlash         = 163
	MoveSubstitute    = 164
	MoveStruggle      = 165
	MoveSketch        = 166
	MoveTripleKick    = 167
	MoveThief         = 168
	MoveSpiderWeb     = 169
	MoveMindReader    = 170
	MoveNightmare     = 171
	MoveFlameWheel    = 172
	MoveSnore         = 173
	MoveCurse         = 174
	MoveFlail         = 175
	MoveConversion2   = 176
	MoveAeroblast     = 177
	MoveCottonSpore   = 178
	MoveReversal      = 179
	MoveSpite         = 180
	MovePowderSnow    = 181
	MoveProtect       = 182
	MoveMachPunch     = 183
	MoveScaryFace     = 184
	MoveFeintAttack   = 185
	MoveSweetKiss     = 186
	MoveBellyDrum     = 187
	MoveSludgeBomb    = 188
	MoveMudSlap       = 189
	MoveOctazooka     = 190
	MoveSpikes        = 191
	MoveZapCannon     = 192
	MoveForesight     = 193
	MoveDestinyBond   = 194
	MovePerishSong    = 195
	MoveIcyWind       = 196
	MoveDetect        = 197
	MoveBoneRush      = 198
	MoveLockOn        = 199
	MoveOutrage       = 200
	MoveSandstorm     = 201
	MoveGigaDrain     = 202
	MoveEndure        = 203
	MoveCharm         = 204
	MoveRollout       = 205
	MoveFalseSwipe    = 206
	MoveSwagger       = 207
	MoveMilkDrink     = 208
	MoveSpark         = 209
	MoveFuryCutter    = 210
	MoveSteelWing     = 211
	MoveMeanLook      = 212
	MoveAttract       = 213
	MoveSleepTalk     = 214
	MoveHealBell      = 215
	MoveReturn        = 216
	MovePresent       = 217
	MoveFrustration   = 218
	MoveSafeguard     = 219
	MovePainSplit     = 220
	MoveSacredFire    = 221
	MoveMagnitude     = 222
	MoveDynamicPunch  = 223
	MoveMegahorn      = 224
	MoveDragonBreath  = 225
	MoveBatonPass     = 226
	MoveEncore        = 227
	MovePursuit       = 228
	MoveRapidSpin     = 229
	MoveSweetScent    = 230
	MoveIronTail      = 231
	MoveMetalClaw     = 232
	MoveVitalThrow    = 233
	MoveMorningSun    = 234
	MoveSynthesis     = 235
	MoveMoonlight     = 236
	MoveHiddenPower   = 237
	MoveCrossChop     = 238
	MoveTwister       = 239
	MoveRainDance     = 240
	MoveSunnyDay      = 241
	MoveCrunch        = 242
	MoveMirrorCoat    = 243
	MovePsychUp       = 244
	MoveExtremeSpeed  = 245
	MoveAncientPower  = 246
	MoveShadowBall    = 247
	MoveFutureSight   = 248
	MoveRockSmash     = 249
	MoveWhirlpool     = 250
	MoveBeatUp        = 251
	MoveFakeOut       = 252
	MoveUproar        = 253
	MoveStockpile     = 254
	MoveSpitUp        = 255
	MoveSwallow       = 256
	MoveHeatWave      = 257
	MoveHail          = 258
	MoveTorment       = 259
	MoveFlatter       = 260
	MoveWillOWisp     = 261
	MoveMemento       = 262
	MoveFacade        = 263
	MoveFocusPunch    = 264
	MoveSmellingSalts = 265
	MoveFollowMe      = 266
	MoveNaturePower   = 267
	MoveCharge        = 268
	MoveTaunt         = 269
	MoveHelpingHand   = 270
	MoveTrick         = 271
	MoveRolePlay      = 272
	MoveWish          = 273
	MoveAssist        = 274
	MoveIngrain       = 275
	MoveSuperpower    = 276
	MoveMagicCoat     = 277
	MoveRecycle       = 278
	MoveRevenge       = 279
	MoveBrickBreak    = 280
	MoveYawn          = 281
	MoveKnockOff      = 282
	MoveEndeavor      = 283
	MoveEruption      = 284
	MoveSkillSwap     = 285
	MoveImprison      = 286
	MoveRefresh       = 287
	MoveGrudge        = 288
	MoveSnatch        = 289
	MoveSecretPower   = 290
	MoveDive          = 291
	MoveArmThrust     = 292
	MoveCamouflage    = 293
	MoveTailGlow      = 294
	MoveLusterPurge   = 295
	MoveMistBall      = 296
	MoveFeatherDance  = 297
	MoveTeeterDance   = 298
	MoveBlazeKick     = 299
	MoveMudSport      = 300
	MoveIceBall       = 301
	MoveNeedleArm     = 302
	MoveSlackOff      = 303
	MoveHyperVoice    = 304
	MovePoisonFang    = 305
	MoveCrushClaw     = 306
	MoveBlastBurn     = 307
	MoveHydroCannon   = 308
	MoveMeteorMash    = 309
	MoveAstonish      = 310
	MoveWeatherBall   = 311
	MoveAromatherapy  = 312
	MoveFakeTears     = 313
	MoveAirCutter     = 314
	MoveOverheat      = 315
	MoveOdorSleuth    = 316
	MoveRockTomb      = 317
	MoveSilverWind    = 318
	MoveMetalSound    = 319
	MoveGrassWhistle  = 320
	MoveTickle        = 321
	MoveCosmicPower   = 322
	MoveWaterSpout    = 323
	MoveSignalBeam    = 324
	MoveShadowPunch   = 325
	MoveExtrasensory  = 326
	MoveSkyUppercut   = 327
	MoveSandTomb      = 328
	MoveSheerCold     = 329
	MoveMuddyWater    = 330
	MoveBulletSeed    = 331
	MoveAerialAce     = 332
	MoveIcicleSpear   = 333
	MoveIronDefense   = 334
	MoveBlock         = 335
	MoveHowl          = 336
	MoveDragonClaw    = 337
	MoveFrenzyPlant   = 338
	MoveBulkUp        = 339
	MoveBounce        = 340
	MoveMudShot       = 341
	MovePoisonTail    = 342
	MoveCovet         = 343
	MoveVoltTackle    = 344
	MoveMagicalLeaf   = 345
	MoveWaterSport    = 346
	MoveCalmMind      = 347
	MoveLeafBlade     = 348
	MoveDragonDance   = 349
	MoveRockBlast     = 350
	MoveShockWave     = 351
	MoveWaterPulse    = 352
	MoveDoomDesire    = 353
	MovePsychoBoost   = 354
	MoveRoost         = 355
	MoveGravity       = 356
	MoveMiracleEye    = 357
	MoveWakeUpSlap    = 358
	MoveHammerArm     = 359
	MoveGyroBall      = 360
	MoveHealingWish   = 361
	MoveBrine         = 362
	MoveNaturalGift   = 363
	MoveFeint         = 364
	MovePluck         = 365
	MoveTailwind      = 366
	MoveAcupressure   = 367
	MoveMetalBurst    = 368
	MoveUturn         = 369
	MoveCloseCombat   = 370
	MovePayback       = 371
	MoveAssurance     = 372
	MoveEmbargo       = 373
	MoveFling         = 374
	MovePsychoShift   = 375
	MoveTrumpCard     = 376
	MoveHealBlock     = 377
	MoveWringOut      = 378
	MovePowerTrick    = 379
	MoveGastroAcid    = 380
	MoveLuckyChant    = 381
	MoveMeFirst       = 382
	MoveCopycat       = 383
	MovePowerSwap     = 384
	MoveGuardSwap     = 385
	MovePunishment    = 386
	MoveLastResort    = 387
	MoveWorrySeed     = 388
	MoveSuckerPunch   = 389
	MoveToxicSpikes   = 390
	MoveHeartSwap     = 391
	MoveAquaRing      = 392
	MoveMagnetRise    = 393
	MoveFlareBlitz    = 394
	MoveForcePalm     = 395
	MoveAuraSphere    = 396
	MoveRockPolish    = 397
	MovePoisonJab     = 398
	MoveDarkPulse     = 399
	MoveNightSlash    = 400
	MoveAquaTail      = 401
	MoveSeedBomb      = 402
	MoveAirSlash      = 403
	MoveXScissor      = 404
	MoveBugBuzz       = 405
	MoveDragonPulse   = 406
	MoveDragonRush    = 407
	MovePowerGem      = 408
	MoveDrainPunch    = 409
	MoveVacuumWave    = 410
	MoveFocusBlast    = 411
	MoveEnergyBall    = 412
	MoveBraveBird     = 413
	MoveEarthPower    = 414
	MoveSwitcheroo    = 415
	MoveGigaImpact    = 416
	MoveNastyPlot     = 417
	MoveBulletPunch   = 418
	MoveAvalanche     = 419
	MoveIceShard      = 420
	MoveShadowClaw    = 421
	MoveThunderFang   = 422
	MoveIceFang       = 423
	MoveFireFang      = 424
	MoveShadowSneak   = 425
	MoveMudBomb       = 426
	MovePsychoCut     = 427
	MoveZenHeadbutt   = 428
	MoveMirrorShot    = 429
	MoveFlashCannon   = 430
	MoveRockClimb     = 431
	MoveDefog         = 432
	MoveTrickRoom     = 433
	MoveDracoMeteor   = 434
	MoveDischarge     = 435
	MoveLavaPlume     = 436
	MoveLeafStorm     = 437
	MovePowerWhip     = 438
	MoveRockWrecker   = 439
	MoveCrossPoison   = 440
	MoveGunkShot      = 441
	MoveIronHead      = 442
	MoveMagnetBomb    = 443
	MoveStoneEdge     = 444
	MoveCaptivate     = 445
	MoveStealthRock   = 446
	MoveGrassKnot     = 447
	MoveChatter       = 448
	MoveJudgment      = 449
	MoveBugBite       = 450
	MoveChargeBeam    = 451
	MoveWoodHammer    = 452
	MoveAquaJet       = 453
	MoveAttackOrder   = 454
	MoveDefendOrder   = 455
	MoveHealOrder     = 456
	MoveHeadSmash     = 457
	MoveDoubleHit     = 458
	MoveRoarofTime    = 459
	MoveSpacialRend   = 460
	MoveLunarDance    = 461
	MoveCrushGrip     = 462
	MoveMagmaStorm    = 463
	MoveDarkVoid      = 464
	MoveSeedFlare     = 465
	MoveOminousWind   = 466
	MoveShadowForce   = 467
	MoveShadowRush    = 10001
	MoveShadowBlast   = 10002
	MoveShadowBlitz   = 10003
	MoveShadowBolt    = 10004
	MoveShadowBreak   = 10005
	MoveShadowChill   = 10006
	MoveShadowEnd     = 10007
	MoveShadowFire    = 10008
	MoveShadowRave    = 10009
	MoveShadowStorm   = 10010
	MoveShadowWave    = 10011
	MoveShadowDown    = 10012
	MoveShadowHalf    = 10013
	MoveShadowHold    = 10014
	MoveShadowMist    = 10015
	MoveShadowPanic   = 10016
	MoveShadowShed    = 10017
	MoveShadowSky     = 10018
)

// Create item constant enum for quick reference
const (
	ItemMasterBall   = 1
	ItemUltraBall    = 2
	ItemGreatBall    = 3
	ItemPokBall      = 4
	ItemSafariBall   = 5
	ItemNetBall      = 6
	ItemDiveBall     = 7
	ItemNestBall     = 8
	ItemRepeatBall   = 9
	ItemTimerBall    = 10
	ItemLuxuryBall   = 11
	ItemPremierBall  = 12
	ItemDuskBall     = 13
	ItemHealBall     = 14
	ItemQuickBall    = 15
	ItemCherishBall  = 16
	ItemPotion       = 17
	ItemAntidote     = 18
	ItemBurnHeal     = 19
	ItemIceHeal      = 20
	ItemAwakening    = 21
	ItemParalyzeHeal = 22
	ItemFullRestore  = 23
	ItemMaxPotion    = 24
	ItemHyperPotion  = 25
	ItemSuperPotion  = 26
	ItemFullHeal     = 27
	ItemRevive       = 28
	ItemMaxRevive    = 29
	ItemFreshWater   = 30
	ItemSodaPop      = 31
	ItemLemonade     = 32
	ItemMoomooMilk   = 33
	ItemEnergyPowder = 34
	ItemEnergyRoot   = 35
	ItemHealPowder   = 36
	ItemRevivalHerb  = 37
	ItemEther        = 38
	ItemMaxEther     = 39
	ItemElixir       = 40
	ItemMaxElixir    = 41
	ItemLavaCookie   = 42
	ItemBerryJuice   = 43
	ItemSacredAsh    = 44
	ItemHPUp         = 45
	ItemProtein      = 46
	ItemIron         = 47
	ItemCarbos       = 48
	ItemCalcium      = 49
	ItemRareCandy    = 50
	ItemPPUp         = 51
	ItemZinc         = 52
	ItemPPMax        = 53
	ItemOldGateau    = 54
	ItemGuardSpec    = 55
	ItemDireHit      = 56
	ItemXAttack      = 57
	ItemXDefense     = 58
	ItemXSpeed       = 59
	ItemXAccuracy    = 60
	ItemXSpAtk       = 61
	ItemXSpDef       = 62
	ItemPokDoll      = 63
	ItemFluffyTail   = 64
	ItemBlueFlute    = 65
	ItemYellowFlute  = 66
	ItemRedFlute     = 67
	ItemBlackFlute   = 68
	ItemWhiteFlute   = 69
	ItemAdamantOrb   = 112
	ItemLustrousOrb  = 113
	ItemCheriBerry   = 126
	ItemChestoBerry  = 127
	ItemPechaBerry   = 128
	ItemRawstBerry   = 129
	ItemAspearBerry  = 130
	ItemLeppaBerry   = 131
	ItemOranBerry    = 132
	ItemPersimBerry  = 133
	ItemLumBerry     = 134
	ItemSitrusBerry  = 135
	ItemFigyBerry    = 136
	ItemWikiBerry    = 137
	ItemMagoBerry    = 138
	ItemAguavBerry   = 139
	ItemIapapaBerry  = 140
	ItemOccaBerry    = 161
	ItemPasshoBerry  = 162
	ItemWacanBerry   = 163
	ItemRindoBerry   = 164
	ItemYacheBerry   = 165
	ItemChopleBerry  = 166
	ItemKebiaBerry   = 167
	ItemShucaBerry   = 168
	ItemCobaBerry    = 169
	ItemPayapaBerry  = 170
	ItemTangaBerry   = 171
	ItemChartiBerry  = 172
	ItemKasibBerry   = 173
	ItemHabanBerry   = 174
	ItemColburBerry  = 175
	ItemBabiriBerry  = 176
	ItemChilanBerry  = 177
	ItemLiechiBerry  = 178
	ItemGanlonBerry  = 179
	ItemSalacBerry   = 180
	ItemPetayaBerry  = 181
	ItemApicotBerry  = 182
	ItemLansatBerry  = 183
	ItemStarfBerry   = 184
	ItemEnigmaBerry  = 185
	ItemMicleBerry   = 186
	ItemCustapBerry  = 187
	ItemJabocaBerry  = 188
	ItemRowapBerry   = 189
	ItemBrightPowder = 190
	ItemWhiteHerb    = 191
	ItemMachoBrace   = 192
	ItemExpShare     = 193
	ItemQuickClaw    = 194
	ItemSootheBell   = 195
	ItemMentalHerb   = 196
	ItemChoiceBand   = 197
	ItemKingsRock    = 198
	ItemSilverPowder = 199
	ItemAmuletCoin   = 200
	ItemCleanseTag   = 201
	ItemSoulDew      = 202
	ItemDeepSeaTooth = 203
	ItemDeepSeaScale = 204
	ItemSmokeBall    = 205
	ItemEverstone    = 206
	ItemFocusBand    = 207
	ItemLuckyEgg     = 208
	ItemScopeLens    = 209
	ItemMetalCoat    = 210
	ItemLeftovers    = 211
	ItemLightBall    = 213
	ItemSoftSand     = 214
	ItemHardStone    = 215
	ItemMiracleSeed  = 216
	ItemBlackGlasses = 217
	ItemBlackBelt    = 218
	ItemMagnet       = 219
	ItemMysticWater  = 220
	ItemSharpBeak    = 221
	ItemPoisonBarb   = 222
	ItemNeverMeltIce = 223
	ItemSpellTag     = 224
	ItemTwistedSpoon = 225
	ItemCharcoal     = 226
	ItemDragonFang   = 227
	ItemSilkScarf    = 228
	ItemShellBell    = 230
	ItemSeaIncense   = 231
	ItemLaxIncense   = 232
	ItemLuckyPunch   = 233
	ItemMetalPowder  = 234
	ItemThickClub    = 235
	ItemStick        = 236
	ItemRedScarf     = 237
	ItemBlueScarf    = 238
	ItemPinkScarf    = 239
	ItemGreenScarf   = 240
	ItemYellowScarf  = 241
	ItemWideLens     = 242
	ItemMuscleBand   = 243
	ItemWiseGlasses  = 244
	ItemExpertBelt   = 245
	ItemLightClay    = 246
	ItemLifeOrb      = 247
	ItemPowerHerb    = 248
	ItemToxicOrb     = 249
	ItemFlameOrb     = 250
	ItemQuickPowder  = 251
	ItemFocusSash    = 252
	ItemZoomLens     = 253
	ItemMetronome    = 254
	ItemIronBall     = 255
	ItemLaggingTail  = 256
	ItemDestinyKnot  = 257
	ItemBlackSludge  = 258
	ItemIcyRock      = 259
	ItemSmoothRock   = 260
	ItemHeatRock     = 261
	ItemDampRock     = 262
	ItemGripClaw     = 263
	ItemChoiceScarf  = 264
	ItemStickyBarb   = 265
	ItemPowerBracer  = 266
	ItemPowerBelt    = 267
	ItemPowerLens    = 268
	ItemPowerBand    = 269
	ItemPowerAnklet  = 270
	ItemPowerWeight  = 271
	ItemShedShell    = 272
	ItemBigRoot      = 273
	ItemChoiceSpecs  = 274
	ItemFlamePlate   = 275
	ItemSplashPlate  = 276
	ItemZapPlate     = 277
	ItemMeadowPlate  = 278
	ItemIciclePlate  = 279
	ItemFistPlate    = 280
	ItemToxicPlate   = 281
	ItemEarthPlate   = 282
	ItemSkyPlate     = 283
	ItemMindPlate    = 284
	ItemInsectPlate  = 285
	ItemStonePlate   = 286
	ItemSpookyPlate  = 287
	ItemDracoPlate   = 288
	ItemDreadPlate   = 289
	ItemIronPlate    = 290
	ItemOddIncense   = 291
	ItemRockIncense  = 292
	ItemFullIncense  = 293
	ItemWaveIncense  = 294
	ItemRoseIncense  = 295
	ItemLuckIncense  = 296
	ItemPureIncense  = 297
	ItemRazorClaw    = 303
	ItemRazorFang    = 304
)

// A collection of all items in the game
var ALL_ITEMS = []Item{
	{ID: 1, Name: "Master Ball", Category: 34, FlingPower: 0, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 2, Name: "Ultra Ball", Category: 34, FlingPower: 0, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 3, Name: "Great Ball", Category: 34, FlingPower: 0, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 4, Name: "Poké Ball", Category: 34, FlingPower: 0, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 5, Name: "Safari Ball", Category: 34, FlingPower: 0, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 6, Name: "Net Ball", Category: 33, FlingPower: 0, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 7, Name: "Dive Ball", Category: 33, FlingPower: 0, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 8, Name: "Nest Ball", Category: 33, FlingPower: 0, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 9, Name: "Repeat Ball", Category: 33, FlingPower: 0, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 10, Name: "Timer Ball", Category: 33, FlingPower: 0, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 11, Name: "Luxury Ball", Category: 33, FlingPower: 0, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 12, Name: "Premier Ball", Category: 33, FlingPower: 0, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 13, Name: "Dusk Ball", Category: 33, FlingPower: 0, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 14, Name: "Heal Ball", Category: 33, FlingPower: 0, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 15, Name: "Quick Ball", Category: 33, FlingPower: 0, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 16, Name: "Cherish Ball", Category: 33, FlingPower: 0, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 17, Name: "Potion", Category: 27, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 18, Name: "Antidote", Category: 30, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 19, Name: "Burn Heal", Category: 30, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 20, Name: "Ice Heal", Category: 30, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 21, Name: "Awakening", Category: 30, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 22, Name: "Paralyze Heal", Category: 30, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 23, Name: "Full Restore", Category: 27, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 24, Name: "Max Potion", Category: 27, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 25, Name: "Hyper Potion", Category: 27, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 26, Name: "Super Potion", Category: 27, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 27, Name: "Full Heal", Category: 30, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 28, Name: "Revive", Category: 29, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 29, Name: "Max Revive", Category: 29, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 30, Name: "Fresh Water", Category: 27, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 31, Name: "Soda Pop", Category: 27, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 32, Name: "Lemonade", Category: 27, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 33, Name: "Moomoo Milk", Category: 27, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 34, Name: "Energy Powder", Category: 27, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 35, Name: "Energy Root", Category: 27, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 36, Name: "Heal Powder", Category: 30, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 37, Name: "Revival Herb", Category: 29, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 38, Name: "Ether", Category: 28, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 39, Name: "Max Ether", Category: 28, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 40, Name: "Elixir", Category: 28, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 41, Name: "Max Elixir", Category: 28, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 42, Name: "Lava Cookie", Category: 30, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 43, Name: "Berry Juice", Category: 27, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 44, Name: "Sacred Ash", Category: 29, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 45, Name: "HP Up", Category: 26, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 46, Name: "Protein", Category: 26, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 47, Name: "Iron", Category: 26, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 48, Name: "Carbos", Category: 26, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 49, Name: "Calcium", Category: 26, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 50, Name: "Rare Candy", Category: 26, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 51, Name: "PP Up", Category: 26, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 52, Name: "Zinc", Category: 26, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 53, Name: "PP Max", Category: 26, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 54, Name: "Old Gateau", Category: 30, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 55, Name: "Guard Spec.", Category: 1, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 56, Name: "Dire Hit", Category: 1, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 57, Name: "X Attack", Category: 1, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 58, Name: "X Defense", Category: 1, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 59, Name: "X Speed", Category: 1, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 60, Name: "X Accuracy", Category: 1, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 61, Name: "X Sp. Atk", Category: 1, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 62, Name: "X Sp. Def", Category: 1, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 63, Name: "Poké Doll", Category: 11, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 64, Name: "Fluffy Tail", Category: 11, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 65, Name: "Blue Flute", Category: 38, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 66, Name: "Yellow Flute", Category: 38, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 67, Name: "Red Flute", Category: 38, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{ID: 68, Name: "Black Flute", Category: 11, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagHoldable},
	{ID: 69, Name: "White Flute", Category: 11, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagHoldable},
	{ID: 112, Name: "Adamant Orb", Category: 18, FlingPower: 60, FlingEffect: 0, Flags: FlagHoldable},
	{ID: 113, Name: "Lustrous Orb", Category: 18, FlingPower: 60, FlingEffect: 0, Flags: FlagHoldable},
	{ID: 126, Name: "Cheri Berry", Category: 3, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive},
	{ID: 127, Name: "Chesto Berry", Category: 3, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive},
	{ID: 128, Name: "Pecha Berry", Category: 3, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive},
	{ID: 129, Name: "Rawst Berry", Category: 3, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive},
	{ID: 130, Name: "Aspear Berry", Category: 3, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive},
	{ID: 131, Name: "Leppa Berry", Category: 3, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive},
	{ID: 132, Name: "Oran Berry", Category: 3, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive},
	{ID: 133, Name: "Persim Berry", Category: 3, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive},
	{ID: 134, Name: "Lum Berry", Category: 3, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive},
	{ID: 135, Name: "Sitrus Berry", Category: 3, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive},
	{ID: 136, Name: "Figy Berry", Category: 6, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive},
	{ID: 137, Name: "Wiki Berry", Category: 6, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive},
	{ID: 138, Name: "Mago Berry", Category: 6, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive},
	{ID: 139, Name: "Aguav Berry", Category: 6, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive},
	{ID: 140, Name: "Iapapa Berry", Category: 6, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive},
	{ID: 161, Name: "Occa Berry", Category: 7, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{ID: 162, Name: "Passho Berry", Category: 7, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{ID: 163, Name: "Wacan Berry", Category: 7, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{ID: 164, Name: "Rindo Berry", Category: 7, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{ID: 165, Name: "Yache Berry", Category: 7, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{ID: 166, Name: "Chople Berry", Category: 7, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{ID: 167, Name: "Kebia Berry", Category: 7, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{ID: 168, Name: "Shuca Berry", Category: 7, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{ID: 169, Name: "Coba Berry", Category: 7, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{ID: 170, Name: "Payapa Berry", Category: 7, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{ID: 171, Name: "Tanga Berry", Category: 7, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{ID: 172, Name: "Charti Berry", Category: 7, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{ID: 173, Name: "Kasib Berry", Category: 7, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{ID: 174, Name: "Haban Berry", Category: 7, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{ID: 175, Name: "Colbur Berry", Category: 7, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{ID: 176, Name: "Babiri Berry", Category: 7, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{ID: 177, Name: "Chilan Berry", Category: 7, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{ID: 178, Name: "Liechi Berry", Category: 5, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive},
	{ID: 179, Name: "Ganlon Berry", Category: 5, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive},
	{ID: 180, Name: "Salac Berry", Category: 5, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive},
	{ID: 181, Name: "Petaya Berry", Category: 5, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive},
	{ID: 182, Name: "Apicot Berry", Category: 5, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive},
	{ID: 183, Name: "Lansat Berry", Category: 5, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive},
	{ID: 184, Name: "Starf Berry", Category: 5, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive},
	{ID: 185, Name: "Enigma Berry", Category: 4, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{ID: 186, Name: "Micle Berry", Category: 5, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive},
	{ID: 187, Name: "Custap Berry", Category: 5, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{ID: 188, Name: "Jaboca Berry", Category: 4, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{ID: 189, Name: "Rowap Berry", Category: 4, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{ID: 190, Name: "Bright Powder", Category: 12, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 191, Name: "White Herb", Category: 12, FlingPower: 10, FlingEffect: 4, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 192, Name: "Macho Brace", Category: 14, FlingPower: 60, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 193, Name: "Exp. Share", Category: 16, FlingPower: 0, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 194, Name: "Quick Claw", Category: 12, FlingPower: 80, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 195, Name: "Soothe Bell", Category: 16, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 196, Name: "Mental Herb", Category: 12, FlingPower: 10, FlingEffect: 4, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 197, Name: "Choice Band", Category: 13, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 198, Name: "King’s Rock", Category: 12, FlingPower: 30, FlingEffect: 7, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 199, Name: "Silver Powder", Category: 19, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 200, Name: "Amulet Coin", Category: 16, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 201, Name: "Cleanse Tag", Category: 16, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable},
	{ID: 202, Name: "Soul Dew", Category: 18, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable},
	{ID: 203, Name: "Deep Sea Tooth", Category: 18, FlingPower: 90, FlingEffect: 0, Flags: FlagHoldable},
	{ID: 204, Name: "Deep Sea Scale", Category: 18, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable},
	{ID: 205, Name: "Smoke Ball", Category: 12, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 206, Name: "Everstone", Category: 16, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 207, Name: "Focus Band", Category: 12, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 208, Name: "Lucky Egg", Category: 16, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 209, Name: "Scope Lens", Category: 12, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 210, Name: "Metal Coat", Category: 19, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 211, Name: "Leftovers", Category: 12, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 213, Name: "Light Ball", Category: 18, FlingPower: 30, FlingEffect: 5, Flags: FlagHoldable},
	{ID: 214, Name: "Soft Sand", Category: 19, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 215, Name: "Hard Stone", Category: 19, FlingPower: 100, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 216, Name: "Miracle Seed", Category: 19, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 217, Name: "Black Glasses", Category: 19, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 218, Name: "Black Belt", Category: 19, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 219, Name: "Magnet", Category: 19, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 220, Name: "Mystic Water", Category: 19, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 221, Name: "Sharp Beak", Category: 19, FlingPower: 50, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 222, Name: "Poison Barb", Category: 19, FlingPower: 70, FlingEffect: 6, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 223, Name: "Never-Melt Ice", Category: 19, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 224, Name: "Spell Tag", Category: 19, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 225, Name: "Twisted Spoon", Category: 19, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 226, Name: "Charcoal", Category: 19, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 227, Name: "Dragon Fang", Category: 19, FlingPower: 70, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 228, Name: "Silk Scarf", Category: 19, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 230, Name: "Shell Bell", Category: 12, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 231, Name: "Sea Incense", Category: 19, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 232, Name: "Lax Incense", Category: 12, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 233, Name: "Lucky Punch", Category: 18, FlingPower: 40, FlingEffect: 0, Flags: FlagHoldable},
	{ID: 234, Name: "Metal Powder", Category: 18, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable},
	{ID: 235, Name: "Thick Club", Category: 18, FlingPower: 90, FlingEffect: 0, Flags: FlagHoldable},
	{ID: 236, Name: "Stick", Category: 18, FlingPower: 60, FlingEffect: 0, Flags: FlagHoldable},
	{ID: 237, Name: "Red Scarf", Category: 36, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{ID: 238, Name: "Blue Scarf", Category: 36, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{ID: 239, Name: "Pink Scarf", Category: 36, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{ID: 240, Name: "Green Scarf", Category: 36, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{ID: 241, Name: "Yellow Scarf", Category: 36, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{ID: 242, Name: "Wide Lens", Category: 12, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 243, Name: "Muscle Band", Category: 12, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 244, Name: "Wise Glasses", Category: 12, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 245, Name: "Expert Belt", Category: 12, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 246, Name: "Light Clay", Category: 12, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 247, Name: "Life Orb", Category: 12, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 248, Name: "Power Herb", Category: 12, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 249, Name: "Toxic Orb", Category: 15, FlingPower: 30, FlingEffect: 1, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 250, Name: "Flame Orb", Category: 15, FlingPower: 30, FlingEffect: 2, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 251, Name: "Quick Powder", Category: 18, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable},
	{ID: 252, Name: "Focus Sash", Category: 12, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 253, Name: "Zoom Lens", Category: 12, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 254, Name: "Metronome", Category: 12, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 255, Name: "Iron Ball", Category: 15, FlingPower: 130, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 256, Name: "Lagging Tail", Category: 15, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 257, Name: "Destiny Knot", Category: 12, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 258, Name: "Black Sludge", Category: 12, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 259, Name: "Icy Rock", Category: 12, FlingPower: 40, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 260, Name: "Smooth Rock", Category: 12, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 261, Name: "Heat Rock", Category: 12, FlingPower: 60, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 262, Name: "Damp Rock", Category: 12, FlingPower: 60, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 263, Name: "Grip Claw", Category: 12, FlingPower: 90, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 264, Name: "Choice Scarf", Category: 13, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 265, Name: "Sticky Barb", Category: 15, FlingPower: 80, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 266, Name: "Power Bracer", Category: 14, FlingPower: 70, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 267, Name: "Power Belt", Category: 14, FlingPower: 70, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 268, Name: "Power Lens", Category: 14, FlingPower: 70, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 269, Name: "Power Band", Category: 14, FlingPower: 70, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 270, Name: "Power Anklet", Category: 14, FlingPower: 70, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 271, Name: "Power Weight", Category: 14, FlingPower: 70, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 272, Name: "Shed Shell", Category: 12, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 273, Name: "Big Root", Category: 12, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 274, Name: "Choice Specs", Category: 13, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 275, Name: "Flame Plate", Category: 17, FlingPower: 90, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 276, Name: "Splash Plate", Category: 17, FlingPower: 90, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 277, Name: "Zap Plate", Category: 17, FlingPower: 90, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 278, Name: "Meadow Plate", Category: 17, FlingPower: 90, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 279, Name: "Icicle Plate", Category: 17, FlingPower: 90, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 280, Name: "Fist Plate", Category: 17, FlingPower: 90, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 281, Name: "Toxic Plate", Category: 17, FlingPower: 90, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 282, Name: "Earth Plate", Category: 17, FlingPower: 90, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 283, Name: "Sky Plate", Category: 17, FlingPower: 90, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 284, Name: "Mind Plate", Category: 17, FlingPower: 90, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 285, Name: "Insect Plate", Category: 17, FlingPower: 90, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 286, Name: "Stone Plate", Category: 17, FlingPower: 90, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 287, Name: "Spooky Plate", Category: 17, FlingPower: 90, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 288, Name: "Draco Plate", Category: 17, FlingPower: 90, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 289, Name: "Dread Plate", Category: 17, FlingPower: 90, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 290, Name: "Iron Plate", Category: 17, FlingPower: 90, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 291, Name: "Odd Incense", Category: 19, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 292, Name: "Rock Incense", Category: 19, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 293, Name: "Full Incense", Category: 15, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 294, Name: "Wave Incense", Category: 19, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 295, Name: "Rose Incense", Category: 19, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 296, Name: "Luck Incense", Category: 16, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 297, Name: "Pure Incense", Category: 16, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable},
	{ID: 303, Name: "Razor Claw", Category: 12, FlingPower: 80, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{ID: 304, Name: "Razor Fang", Category: 12, FlingPower: 30, FlingEffect: 7, Flags: FlagHoldable | FlagHoldableActive},
}

//A table of levels mapped to the total experience at that level for each growth rate
var EXP_TABLE = map[int]map[int]int{
	GrowthSlow: {
		0:   0,
		1:   0,
		2:   10,
		3:   33,
		4:   80,
		5:   156,
		6:   270,
		7:   428,
		8:   640,
		9:   911,
		10:  1250,
		11:  1663,
		12:  2160,
		13:  2746,
		14:  3430,
		15:  4218,
		16:  5120,
		17:  6141,
		18:  7290,
		19:  8573,
		20:  10000,
		21:  11576,
		22:  13310,
		23:  15208,
		24:  17280,
		25:  19531,
		26:  21970,
		27:  24603,
		28:  27440,
		29:  30486,
		30:  33750,
		31:  37238,
		32:  40960,
		33:  44921,
		34:  49130,
		35:  53593,
		36:  58320,
		37:  63316,
		38:  68590,
		39:  74148,
		40:  80000,
		41:  86151,
		42:  92610,
		43:  99383,
		44:  106480,
		45:  113906,
		46:  121670,
		47:  129778,
		48:  138240,
		49:  147061,
		50:  156250,
		51:  165813,
		52:  175760,
		53:  186096,
		54:  196830,
		55:  207968,
		56:  219520,
		57:  231491,
		58:  243890,
		59:  256723,
		60:  270000,
		61:  283726,
		62:  297910,
		63:  312558,
		64:  327680,
		65:  343281,
		66:  359370,
		67:  375953,
		68:  393040,
		69:  410636,
		70:  428750,
		71:  447388,
		72:  466560,
		73:  486271,
		74:  506530,
		75:  527343,
		76:  548720,
		77:  570666,
		78:  593190,
		79:  616298,
		80:  640000,
		81:  664301,
		82:  689210,
		83:  714733,
		84:  740880,
		85:  767656,
		86:  795070,
		87:  823128,
		88:  851840,
		89:  881211,
		90:  911250,
		91:  941963,
		92:  973360,
		93:  1005446,
		94:  1038230,
		95:  1071718,
		96:  1105920,
		97:  1140841,
		98:  1176490,
		99:  1212873,
		100: 1250000,
	}, GrowthMediumFast: {
		0:   0,
		1:   0,
		2:   8,
		3:   27,
		4:   64,
		5:   125,
		6:   216,
		7:   343,
		8:   512,
		9:   729,
		10:  1000,
		11:  1331,
		12:  1728,
		13:  2197,
		14:  2744,
		15:  3375,
		16:  4096,
		17:  4913,
		18:  5832,
		19:  6859,
		20:  8000,
		21:  9261,
		22:  10648,
		23:  12167,
		24:  13824,
		25:  15625,
		26:  17576,
		27:  19683,
		28:  21952,
		29:  24389,
		30:  27000,
		31:  29791,
		32:  32768,
		33:  35937,
		34:  39304,
		35:  42875,
		36:  46656,
		37:  50653,
		38:  54872,
		39:  59319,
		40:  64000,
		41:  68921,
		42:  74088,
		43:  79507,
		44:  85184,
		45:  91125,
		46:  97336,
		47:  103823,
		48:  110592,
		49:  117649,
		50:  125000,
		51:  132651,
		52:  140608,
		53:  148877,
		54:  157464,
		55:  166375,
		56:  175616,
		57:  185193,
		58:  195112,
		59:  205379,
		60:  216000,
		61:  226981,
		62:  238328,
		63:  250047,
		64:  262144,
		65:  274625,
		66:  287496,
		67:  300763,
		68:  314432,
		69:  328509,
		70:  343000,
		71:  357911,
		72:  373248,
		73:  389017,
		74:  405224,
		75:  421875,
		76:  438976,
		77:  456533,
		78:  474552,
		79:  493039,
		80:  512000,
		81:  531441,
		82:  551368,
		83:  571787,
		84:  592704,
		85:  614125,
		86:  636056,
		87:  658503,
		88:  681472,
		89:  704969,
		90:  729000,
		91:  753571,
		92:  778688,
		93:  804357,
		94:  830584,
		95:  857375,
		96:  884736,
		97:  912673,
		98:  941192,
		99:  970299,
		100: 1000000,
	}, GrowthFast: {
		0:   0,
		1:   0,
		2:   6,
		3:   21,
		4:   51,
		5:   100,
		6:   172,
		7:   274,
		8:   409,
		9:   583,
		10:  800,
		11:  1064,
		12:  1382,
		13:  1757,
		14:  2195,
		15:  2700,
		16:  3276,
		17:  3930,
		18:  4665,
		19:  5487,
		20:  6400,
		21:  7408,
		22:  8518,
		23:  9733,
		24:  11059,
		25:  12500,
		26:  14060,
		27:  15746,
		28:  17561,
		29:  19511,
		30:  21600,
		31:  23832,
		32:  26214,
		33:  28749,
		34:  31443,
		35:  34300,
		36:  37324,
		37:  40522,
		38:  43897,
		39:  47455,
		40:  51200,
		41:  55136,
		42:  59270,
		43:  63605,
		44:  68147,
		45:  72900,
		46:  77868,
		47:  83058,
		48:  88473,
		49:  94119,
		50:  100000,
		51:  106120,
		52:  112486,
		53:  119101,
		54:  125971,
		55:  133100,
		56:  140492,
		57:  148154,
		58:  156089,
		59:  164303,
		60:  172800,
		61:  181584,
		62:  190662,
		63:  200037,
		64:  209715,
		65:  219700,
		66:  229996,
		67:  240610,
		68:  251545,
		69:  262807,
		70:  274400,
		71:  286328,
		72:  298598,
		73:  311213,
		74:  324179,
		75:  337500,
		76:  351180,
		77:  365226,
		78:  379641,
		79:  394431,
		80:  409600,
		81:  425152,
		82:  441094,
		83:  457429,
		84:  474163,
		85:  491300,
		86:  508844,
		87:  526802,
		88:  545177,
		89:  563975,
		90:  583200,
		91:  602856,
		92:  622950,
		93:  643485,
		94:  664467,
		95:  685900,
		96:  707788,
		97:  730138,
		98:  752953,
		99:  776239,
		100: 800000,
	}, GrowthMediumSlow: {
		0:   0,
		1:   0,
		2:   9,
		3:   57,
		4:   96,
		5:   135,
		6:   179,
		7:   236,
		8:   314,
		9:   419,
		10:  560,
		11:  742,
		12:  973,
		13:  1261,
		14:  1612,
		15:  2035,
		16:  2535,
		17:  3120,
		18:  3798,
		19:  4575,
		20:  5460,
		21:  6458,
		22:  7577,
		23:  8825,
		24:  10208,
		25:  11735,
		26:  13411,
		27:  15244,
		28:  17242,
		29:  19411,
		30:  21760,
		31:  24294,
		32:  27021,
		33:  29949,
		34:  33084,
		35:  36435,
		36:  40007,
		37:  43808,
		38:  47846,
		39:  52127,
		40:  56660,
		41:  61450,
		42:  66505,
		43:  71833,
		44:  77440,
		45:  83335,
		46:  89523,
		47:  96012,
		48:  102810,
		49:  109923,
		50:  117360,
		51:  125126,
		52:  133229,
		53:  141677,
		54:  150476,
		55:  159635,
		56:  169159,
		57:  179056,
		58:  189334,
		59:  199999,
		60:  211060,
		61:  222522,
		62:  234393,
		63:  246681,
		64:  259392,
		65:  272535,
		66:  286115,
		67:  300140,
		68:  314618,
		69:  329555,
		70:  344960,
		71:  360838,
		72:  377197,
		73:  394045,
		74:  411388,
		75:  429235,
		76:  447591,
		77:  466464,
		78:  485862,
		79:  505791,
		80:  526260,
		81:  547274,
		82:  568841,
		83:  590969,
		84:  613664,
		85:  636935,
		86:  660787,
		87:  685228,
		88:  710266,
		89:  735907,
		90:  762160,
		91:  789030,
		92:  816525,
		93:  844653,
		94:  873420,
		95:  902835,
		96:  932903,
		97:  963632,
		98:  995030,
		99:  1027103,
		100: 1059860,
	}, GrowthErratic: {
		0:   0,
		1:   0,
		2:   15,
		3:   52,
		4:   122,
		5:   237,
		6:   406,
		7:   637,
		8:   942,
		9:   1326,
		10:  1800,
		11:  2369,
		12:  3041,
		13:  3822,
		14:  4719,
		15:  5737,
		16:  6881,
		17:  8155,
		18:  9564,
		19:  11111,
		20:  12800,
		21:  14632,
		22:  16610,
		23:  18737,
		24:  21012,
		25:  23437,
		26:  26012,
		27:  28737,
		28:  31610,
		29:  34632,
		30:  37800,
		31:  41111,
		32:  44564,
		33:  48155,
		34:  51881,
		35:  55737,
		36:  59719,
		37:  63822,
		38:  68041,
		39:  72369,
		40:  76800,
		41:  81326,
		42:  85942,
		43:  90637,
		44:  95406,
		45:  100237,
		46:  105122,
		47:  110052,
		48:  115015,
		49:  120001,
		50:  125000,
		51:  131324,
		52:  137795,
		53:  144410,
		54:  151165,
		55:  158056,
		56:  165079,
		57:  172229,
		58:  179503,
		59:  186894,
		60:  194400,
		61:  202013,
		62:  209728,
		63:  217540,
		64:  225443,
		65:  233431,
		66:  241496,
		67:  249633,
		68:  257834,
		69:  267406,
		70:  276458,
		71:  286328,
		72:  296358,
		73:  305767,
		74:  316074,
		75:  326531,
		76:  336255,
		77:  346965,
		78:  357812,
		79:  367807,
		80:  378880,
		81:  390077,
		82:  400293,
		83:  411686,
		84:  423190,
		85:  433572,
		86:  445239,
		87:  457001,
		88:  467489,
		89:  479378,
		90:  491346,
		91:  501878,
		92:  513934,
		93:  526049,
		94:  536557,
		95:  548720,
		96:  560922,
		97:  571333,
		98:  583539,
		99:  591882,
		100: 600000,
	}, GrowthFluctuating: {
		0:   0,
		1:   0,
		2:   4,
		3:   13,
		4:   32,
		5:   65,
		6:   112,
		7:   178,
		8:   276,
		9:   393,
		10:  540,
		11:  745,
		12:  967,
		13:  1230,
		14:  1591,
		15:  1957,
		16:  2457,
		17:  3046,
		18:  3732,
		19:  4526,
		20:  5440,
		21:  6482,
		22:  7666,
		23:  9003,
		24:  10506,
		25:  12187,
		26:  14060,
		27:  16140,
		28:  18439,
		29:  20974,
		30:  23760,
		31:  26811,
		32:  30146,
		33:  33780,
		34:  37731,
		35:  42017,
		36:  46656,
		37:  50653,
		38:  55969,
		39:  60505,
		40:  66560,
		41:  71677,
		42:  78533,
		43:  84277,
		44:  91998,
		45:  98415,
		46:  107069,
		47:  114205,
		48:  123863,
		49:  131766,
		50:  142500,
		51:  151222,
		52:  163105,
		53:  172697,
		54:  185807,
		55:  196322,
		56:  210739,
		57:  222231,
		58:  238036,
		59:  250562,
		60:  267840,
		61:  281456,
		62:  300293,
		63:  315059,
		64:  335544,
		65:  351520,
		66:  373744,
		67:  390991,
		68:  415050,
		69:  433631,
		70:  459620,
		71:  479600,
		72:  507617,
		73:  529063,
		74:  559209,
		75:  582187,
		76:  614566,
		77:  639146,
		78:  673863,
		79:  700115,
		80:  737280,
		81:  765275,
		82:  804997,
		83:  834809,
		84:  877201,
		85:  908905,
		86:  954084,
		87:  987754,
		88:  1035837,
		89:  1071552,
		90:  1122660,
		91:  1160499,
		92:  1214753,
		93:  1254796,
		94:  1312322,
		95:  1354652,
		96:  1415577,
		97:  1460276,
		98:  1524731,
		99:  1571884,
		100: 1640000,
	}}

// A map of national pokedex numbers to Pokemon growth rates
var pokemonGrowthRates = map[int]int{
	1:   GrowthMediumSlow,
	2:   GrowthMediumSlow,
	3:   GrowthMediumSlow,
	4:   GrowthMediumSlow,
	5:   GrowthMediumSlow,
	6:   GrowthMediumSlow,
	7:   GrowthMediumSlow,
	8:   GrowthMediumSlow,
	9:   GrowthMediumSlow,
	10:  GrowthMediumFast,
	11:  GrowthMediumFast,
	12:  GrowthMediumFast,
	13:  GrowthMediumFast,
	14:  GrowthMediumFast,
	15:  GrowthMediumFast,
	16:  GrowthMediumSlow,
	17:  GrowthMediumSlow,
	18:  GrowthMediumSlow,
	19:  GrowthMediumFast,
	20:  GrowthMediumFast,
	21:  GrowthMediumFast,
	22:  GrowthMediumFast,
	23:  GrowthMediumFast,
	24:  GrowthMediumFast,
	25:  GrowthMediumFast,
	26:  GrowthMediumFast,
	27:  GrowthMediumFast,
	28:  GrowthMediumFast,
	29:  GrowthMediumSlow,
	30:  GrowthMediumSlow,
	31:  GrowthMediumSlow,
	32:  GrowthMediumSlow,
	33:  GrowthMediumSlow,
	34:  GrowthMediumSlow,
	35:  GrowthFast,
	36:  GrowthFast,
	37:  GrowthMediumFast,
	38:  GrowthMediumFast,
	39:  GrowthFast,
	40:  GrowthFast,
	41:  GrowthMediumFast,
	42:  GrowthMediumFast,
	43:  GrowthMediumSlow,
	44:  GrowthMediumSlow,
	45:  GrowthMediumSlow,
	46:  GrowthMediumFast,
	47:  GrowthMediumFast,
	48:  GrowthMediumFast,
	49:  GrowthMediumFast,
	50:  GrowthMediumFast,
	51:  GrowthMediumFast,
	52:  GrowthMediumFast,
	53:  GrowthMediumFast,
	54:  GrowthMediumFast,
	55:  GrowthMediumFast,
	56:  GrowthMediumFast,
	57:  GrowthMediumFast,
	58:  GrowthSlow,
	59:  GrowthSlow,
	60:  GrowthMediumSlow,
	61:  GrowthMediumSlow,
	62:  GrowthMediumSlow,
	63:  GrowthMediumSlow,
	64:  GrowthMediumSlow,
	65:  GrowthMediumSlow,
	66:  GrowthMediumSlow,
	67:  GrowthMediumSlow,
	68:  GrowthMediumSlow,
	69:  GrowthMediumSlow,
	70:  GrowthMediumSlow,
	71:  GrowthMediumSlow,
	72:  GrowthSlow,
	73:  GrowthSlow,
	74:  GrowthMediumSlow,
	75:  GrowthMediumSlow,
	76:  GrowthMediumSlow,
	77:  GrowthMediumFast,
	78:  GrowthMediumFast,
	79:  GrowthMediumFast,
	80:  GrowthMediumFast,
	81:  GrowthMediumFast,
	82:  GrowthMediumFast,
	83:  GrowthMediumFast,
	84:  GrowthMediumFast,
	85:  GrowthMediumFast,
	86:  GrowthMediumFast,
	87:  GrowthMediumFast,
	88:  GrowthMediumFast,
	89:  GrowthMediumFast,
	90:  GrowthSlow,
	91:  GrowthSlow,
	92:  GrowthMediumSlow,
	93:  GrowthMediumSlow,
	94:  GrowthMediumSlow,
	95:  GrowthMediumFast,
	96:  GrowthMediumFast,
	97:  GrowthMediumFast,
	98:  GrowthMediumFast,
	99:  GrowthMediumFast,
	100: GrowthMediumFast,
	101: GrowthMediumFast,
	102: GrowthSlow,
	103: GrowthSlow,
	104: GrowthMediumFast,
	105: GrowthMediumFast,
	106: GrowthMediumFast,
	107: GrowthMediumFast,
	108: GrowthMediumFast,
	109: GrowthMediumFast,
	110: GrowthMediumFast,
	111: GrowthSlow,
	112: GrowthSlow,
	113: GrowthFast,
	114: GrowthMediumFast,
	115: GrowthMediumFast,
	116: GrowthMediumFast,
	117: GrowthMediumFast,
	118: GrowthMediumFast,
	119: GrowthMediumFast,
	120: GrowthSlow,
	121: GrowthSlow,
	122: GrowthMediumFast,
	123: GrowthMediumFast,
	124: GrowthMediumFast,
	125: GrowthMediumFast,
	126: GrowthMediumFast,
	127: GrowthSlow,
	128: GrowthSlow,
	129: GrowthSlow,
	130: GrowthSlow,
	131: GrowthSlow,
	132: GrowthMediumFast,
	133: GrowthMediumFast,
	134: GrowthMediumFast,
	135: GrowthMediumFast,
	136: GrowthMediumFast,
	137: GrowthMediumFast,
	138: GrowthMediumFast,
	139: GrowthMediumFast,
	140: GrowthMediumFast,
	141: GrowthMediumFast,
	142: GrowthSlow,
	143: GrowthSlow,
	144: GrowthSlow,
	145: GrowthSlow,
	146: GrowthSlow,
	147: GrowthSlow,
	148: GrowthSlow,
	149: GrowthSlow,
	150: GrowthSlow,
	151: GrowthMediumSlow,
	152: GrowthMediumSlow,
	153: GrowthMediumSlow,
	154: GrowthMediumSlow,
	155: GrowthMediumSlow,
	156: GrowthMediumSlow,
	157: GrowthMediumSlow,
	158: GrowthMediumSlow,
	159: GrowthMediumSlow,
	160: GrowthMediumSlow,
	161: GrowthMediumFast,
	162: GrowthMediumFast,
	163: GrowthMediumFast,
	164: GrowthMediumFast,
	165: GrowthFast,
	166: GrowthFast,
	167: GrowthFast,
	168: GrowthFast,
	169: GrowthMediumFast,
	170: GrowthSlow,
	171: GrowthSlow,
	172: GrowthMediumFast,
	173: GrowthFast,
	174: GrowthFast,
	175: GrowthFast,
	176: GrowthFast,
	177: GrowthMediumFast,
	178: GrowthMediumFast,
	179: GrowthMediumSlow,
	180: GrowthMediumSlow,
	181: GrowthMediumSlow,
	182: GrowthMediumSlow,
	183: GrowthFast,
	184: GrowthFast,
	185: GrowthMediumFast,
	186: GrowthMediumSlow,
	187: GrowthMediumSlow,
	188: GrowthMediumSlow,
	189: GrowthMediumSlow,
	190: GrowthFast,
	191: GrowthMediumSlow,
	192: GrowthMediumSlow,
	193: GrowthMediumFast,
	194: GrowthMediumFast,
	195: GrowthMediumFast,
	196: GrowthMediumFast,
	197: GrowthMediumFast,
	198: GrowthMediumSlow,
	199: GrowthMediumFast,
	200: GrowthFast,
	201: GrowthMediumFast,
	202: GrowthMediumFast,
	203: GrowthMediumFast,
	204: GrowthMediumFast,
	205: GrowthMediumFast,
	206: GrowthMediumFast,
	207: GrowthMediumSlow,
	208: GrowthMediumFast,
	209: GrowthFast,
	210: GrowthFast,
	211: GrowthMediumFast,
	212: GrowthMediumFast,
	213: GrowthMediumSlow,
	214: GrowthSlow,
	215: GrowthMediumSlow,
	216: GrowthMediumFast,
	217: GrowthMediumFast,
	218: GrowthMediumFast,
	219: GrowthMediumFast,
	220: GrowthSlow,
	221: GrowthSlow,
	222: GrowthFast,
	223: GrowthMediumFast,
	224: GrowthMediumFast,
	225: GrowthFast,
	226: GrowthSlow,
	227: GrowthSlow,
	228: GrowthSlow,
	229: GrowthSlow,
	230: GrowthMediumFast,
	231: GrowthMediumFast,
	232: GrowthMediumFast,
	233: GrowthMediumFast,
	234: GrowthSlow,
	235: GrowthFast,
	236: GrowthMediumFast,
	237: GrowthMediumFast,
	238: GrowthMediumFast,
	239: GrowthMediumFast,
	240: GrowthMediumFast,
	241: GrowthSlow,
	242: GrowthFast,
	243: GrowthSlow,
	244: GrowthSlow,
	245: GrowthSlow,
	246: GrowthSlow,
	247: GrowthSlow,
	248: GrowthSlow,
	249: GrowthSlow,
	250: GrowthSlow,
	251: GrowthMediumSlow,
	252: GrowthMediumSlow,
	253: GrowthMediumSlow,
	254: GrowthMediumSlow,
	255: GrowthMediumSlow,
	256: GrowthMediumSlow,
	257: GrowthMediumSlow,
	258: GrowthMediumSlow,
	259: GrowthMediumSlow,
	260: GrowthMediumSlow,
	261: GrowthMediumFast,
	262: GrowthMediumFast,
	263: GrowthMediumFast,
	264: GrowthMediumFast,
	265: GrowthMediumFast,
	266: GrowthMediumFast,
	267: GrowthMediumFast,
	268: GrowthMediumFast,
	269: GrowthMediumFast,
	270: GrowthMediumSlow,
	271: GrowthMediumSlow,
	272: GrowthMediumSlow,
	273: GrowthMediumSlow,
	274: GrowthMediumSlow,
	275: GrowthMediumSlow,
	276: GrowthMediumSlow,
	277: GrowthMediumSlow,
	278: GrowthMediumFast,
	279: GrowthMediumFast,
	280: GrowthSlow,
	281: GrowthSlow,
	282: GrowthSlow,
	283: GrowthMediumFast,
	284: GrowthMediumFast,
	285: GrowthFluctuating,
	286: GrowthFluctuating,
	287: GrowthSlow,
	288: GrowthSlow,
	289: GrowthSlow,
	290: GrowthErratic,
	291: GrowthErratic,
	292: GrowthErratic,
	293: GrowthMediumSlow,
	294: GrowthMediumSlow,
	295: GrowthMediumSlow,
	296: GrowthFluctuating,
	297: GrowthFluctuating,
	298: GrowthFast,
	299: GrowthMediumFast,
	300: GrowthFast,
	301: GrowthFast,
	302: GrowthMediumSlow,
	303: GrowthFast,
	304: GrowthSlow,
	305: GrowthSlow,
	306: GrowthSlow,
	307: GrowthMediumFast,
	308: GrowthMediumFast,
	309: GrowthSlow,
	310: GrowthSlow,
	311: GrowthMediumFast,
	312: GrowthMediumFast,
	313: GrowthErratic,
	314: GrowthFluctuating,
	315: GrowthMediumSlow,
	316: GrowthFluctuating,
	317: GrowthFluctuating,
	318: GrowthSlow,
	319: GrowthSlow,
	320: GrowthFluctuating,
	321: GrowthFluctuating,
	322: GrowthMediumFast,
	323: GrowthMediumFast,
	324: GrowthMediumFast,
	325: GrowthFast,
	326: GrowthFast,
	327: GrowthFast,
	328: GrowthMediumSlow,
	329: GrowthMediumSlow,
	330: GrowthMediumSlow,
	331: GrowthMediumSlow,
	332: GrowthMediumSlow,
	333: GrowthErratic,
	334: GrowthErratic,
	335: GrowthErratic,
	336: GrowthFluctuating,
	337: GrowthFast,
	338: GrowthFast,
	339: GrowthMediumFast,
	340: GrowthMediumFast,
	341: GrowthFluctuating,
	342: GrowthFluctuating,
	343: GrowthMediumFast,
	344: GrowthMediumFast,
	345: GrowthErratic,
	346: GrowthErratic,
	347: GrowthErratic,
	348: GrowthErratic,
	349: GrowthErratic,
	350: GrowthErratic,
	351: GrowthMediumFast,
	352: GrowthMediumSlow,
	353: GrowthFast,
	354: GrowthFast,
	355: GrowthFast,
	356: GrowthFast,
	357: GrowthSlow,
	358: GrowthFast,
	359: GrowthMediumSlow,
	360: GrowthMediumFast,
	361: GrowthMediumFast,
	362: GrowthMediumFast,
	363: GrowthMediumSlow,
	364: GrowthMediumSlow,
	365: GrowthMediumSlow,
	366: GrowthErratic,
	367: GrowthErratic,
	368: GrowthErratic,
	369: GrowthSlow,
	370: GrowthFast,
	371: GrowthSlow,
	372: GrowthSlow,
	373: GrowthSlow,
	374: GrowthSlow,
	375: GrowthSlow,
	376: GrowthSlow,
	377: GrowthSlow,
	378: GrowthSlow,
	379: GrowthSlow,
	380: GrowthSlow,
	381: GrowthSlow,
	382: GrowthSlow,
	383: GrowthSlow,
	384: GrowthSlow,
	385: GrowthSlow,
	386: GrowthSlow,
	387: GrowthMediumSlow,
	388: GrowthMediumSlow,
	389: GrowthMediumSlow,
	390: GrowthMediumSlow,
	391: GrowthMediumSlow,
	392: GrowthMediumSlow,
	393: GrowthMediumSlow,
	394: GrowthMediumSlow,
	395: GrowthMediumSlow,
	396: GrowthMediumSlow,
	397: GrowthMediumSlow,
	398: GrowthMediumSlow,
	399: GrowthMediumFast,
	400: GrowthMediumFast,
	401: GrowthMediumSlow,
	402: GrowthMediumSlow,
	403: GrowthMediumSlow,
	404: GrowthMediumSlow,
	405: GrowthMediumSlow,
	406: GrowthMediumSlow,
	407: GrowthMediumSlow,
	408: GrowthErratic,
	409: GrowthErratic,
	410: GrowthErratic,
	411: GrowthErratic,
	412: GrowthMediumFast,
	413: GrowthMediumFast,
	414: GrowthMediumFast,
	415: GrowthMediumSlow,
	416: GrowthMediumSlow,
	417: GrowthMediumFast,
	418: GrowthMediumFast,
	419: GrowthMediumFast,
	420: GrowthMediumFast,
	421: GrowthMediumFast,
	422: GrowthMediumFast,
	423: GrowthMediumFast,
	424: GrowthFast,
	425: GrowthFluctuating,
	426: GrowthFluctuating,
	427: GrowthMediumFast,
	428: GrowthMediumFast,
	429: GrowthFast,
	430: GrowthMediumSlow,
	431: GrowthFast,
	432: GrowthFast,
	433: GrowthFast,
	434: GrowthMediumFast,
	435: GrowthMediumFast,
	436: GrowthMediumFast,
	437: GrowthMediumFast,
	438: GrowthMediumFast,
	439: GrowthMediumFast,
	440: GrowthFast,
	441: GrowthMediumSlow,
	442: GrowthMediumFast,
	443: GrowthSlow,
	444: GrowthSlow,
	445: GrowthSlow,
	446: GrowthSlow,
	447: GrowthMediumSlow,
	448: GrowthMediumSlow,
	449: GrowthSlow,
	450: GrowthSlow,
	451: GrowthSlow,
	452: GrowthSlow,
	453: GrowthMediumFast,
	454: GrowthMediumFast,
	455: GrowthSlow,
	456: GrowthErratic,
	457: GrowthErratic,
	458: GrowthSlow,
	459: GrowthSlow,
	460: GrowthSlow,
	461: GrowthMediumSlow,
	462: GrowthMediumFast,
	463: GrowthMediumFast,
	464: GrowthSlow,
	465: GrowthMediumFast,
	466: GrowthMediumFast,
	467: GrowthMediumFast,
	468: GrowthFast,
	469: GrowthMediumFast,
	470: GrowthMediumFast,
	471: GrowthMediumFast,
	472: GrowthMediumSlow,
	473: GrowthSlow,
	474: GrowthMediumFast,
	475: GrowthSlow,
	476: GrowthMediumFast,
	477: GrowthFast,
	478: GrowthMediumFast,
	479: GrowthMediumFast,
	480: GrowthSlow,
	481: GrowthSlow,
	482: GrowthSlow,
	483: GrowthSlow,
	484: GrowthSlow,
	485: GrowthSlow,
	486: GrowthSlow,
	487: GrowthSlow,
	488: GrowthSlow,
	489: GrowthSlow,
	490: GrowthSlow,
	491: GrowthSlow,
	492: GrowthMediumSlow,
	493: GrowthSlow,
}

// A map of national pokedex numbers to Pokemon base stats
var pokemonBaseStats = map[int][6]int{
	1:   {45, 49, 49, 65, 65, 45},
	2:   {60, 62, 63, 80, 80, 60},
	3:   {80, 82, 83, 100, 100, 80},
	4:   {39, 52, 43, 60, 50, 65},
	5:   {58, 64, 58, 80, 65, 80},
	6:   {78, 84, 78, 109, 85, 100},
	7:   {44, 48, 65, 50, 64, 43},
	8:   {59, 63, 80, 65, 80, 58},
	9:   {79, 83, 100, 85, 105, 78},
	10:  {45, 30, 35, 20, 20, 45},
	11:  {50, 20, 55, 25, 25, 30},
	12:  {60, 45, 50, 90, 80, 70},
	13:  {40, 35, 30, 20, 20, 50},
	14:  {45, 25, 50, 25, 25, 35},
	15:  {65, 90, 40, 45, 80, 75},
	16:  {40, 45, 40, 35, 35, 56},
	17:  {63, 60, 55, 50, 50, 71},
	18:  {83, 80, 75, 70, 70, 101},
	19:  {30, 56, 35, 25, 35, 72},
	20:  {55, 81, 60, 50, 70, 97},
	21:  {40, 60, 30, 31, 31, 70},
	22:  {65, 90, 65, 61, 61, 100},
	23:  {35, 60, 44, 40, 54, 55},
	24:  {60, 95, 69, 65, 79, 80},
	25:  {35, 55, 40, 50, 50, 90},
	26:  {60, 90, 55, 90, 80, 110},
	27:  {50, 75, 85, 20, 30, 40},
	28:  {75, 100, 110, 45, 55, 65},
	29:  {55, 47, 52, 40, 40, 41},
	30:  {70, 62, 67, 55, 55, 56},
	31:  {90, 92, 87, 75, 85, 76},
	32:  {46, 57, 40, 40, 40, 50},
	33:  {61, 72, 57, 55, 55, 65},
	34:  {81, 102, 77, 85, 75, 85},
	35:  {70, 45, 48, 60, 65, 35},
	36:  {95, 70, 73, 95, 90, 60},
	37:  {38, 41, 40, 50, 65, 65},
	38:  {73, 76, 75, 81, 100, 100},
	39:  {115, 45, 20, 45, 25, 20},
	40:  {140, 70, 45, 85, 50, 45},
	41:  {40, 45, 35, 30, 40, 55},
	42:  {75, 80, 70, 65, 75, 90},
	43:  {45, 50, 55, 75, 65, 30},
	44:  {60, 65, 70, 85, 75, 40},
	45:  {75, 80, 85, 110, 90, 50},
	46:  {35, 70, 55, 45, 55, 25},
	47:  {60, 95, 80, 60, 80, 30},
	48:  {60, 55, 50, 40, 55, 45},
	49:  {70, 65, 60, 90, 75, 90},
	50:  {10, 55, 25, 35, 45, 95},
	51:  {35, 100, 50, 50, 70, 120},
	52:  {40, 45, 35, 40, 40, 90},
	53:  {65, 70, 60, 65, 65, 115},
	54:  {50, 52, 48, 65, 50, 55},
	55:  {80, 82, 78, 95, 80, 85},
	56:  {40, 80, 35, 35, 45, 70},
	57:  {65, 105, 60, 60, 70, 95},
	58:  {55, 70, 45, 70, 50, 60},
	59:  {90, 110, 80, 100, 80, 95},
	60:  {40, 50, 40, 40, 40, 90},
	61:  {65, 65, 65, 50, 50, 90},
	62:  {90, 95, 95, 70, 90, 70},
	63:  {25, 20, 15, 105, 55, 90},
	64:  {40, 35, 30, 120, 70, 105},
	65:  {55, 50, 45, 135, 95, 120},
	66:  {70, 80, 50, 35, 35, 35},
	67:  {80, 100, 70, 50, 60, 45},
	68:  {90, 130, 80, 65, 85, 55},
	69:  {50, 75, 35, 70, 30, 40},
	70:  {65, 90, 50, 85, 45, 55},
	71:  {80, 105, 65, 100, 70, 70},
	72:  {40, 40, 35, 50, 100, 70},
	73:  {80, 70, 65, 80, 120, 100},
	74:  {40, 80, 100, 30, 30, 20},
	75:  {55, 95, 115, 45, 45, 35},
	76:  {80, 120, 130, 55, 65, 45},
	77:  {50, 85, 55, 65, 65, 90},
	78:  {65, 100, 70, 80, 80, 105},
	79:  {90, 65, 65, 40, 40, 15},
	80:  {95, 75, 110, 100, 80, 30},
	81:  {25, 35, 70, 95, 55, 45},
	82:  {50, 60, 95, 120, 70, 70},
	83:  {52, 90, 55, 58, 62, 60},
	84:  {35, 85, 45, 35, 35, 75},
	85:  {60, 110, 70, 60, 60, 110},
	86:  {65, 45, 55, 45, 70, 45},
	87:  {90, 70, 80, 70, 95, 70},
	88:  {80, 80, 50, 40, 50, 25},
	89:  {105, 105, 75, 65, 100, 50},
	90:  {30, 65, 100, 45, 25, 40},
	91:  {50, 95, 180, 85, 45, 70},
	92:  {30, 35, 30, 100, 35, 80},
	93:  {45, 50, 45, 115, 55, 95},
	94:  {60, 65, 60, 130, 75, 110},
	95:  {35, 45, 160, 30, 45, 70},
	96:  {60, 48, 45, 43, 90, 42},
	97:  {85, 73, 70, 73, 115, 67},
	98:  {30, 105, 90, 25, 25, 50},
	99:  {55, 130, 115, 50, 50, 75},
	100: {40, 30, 50, 55, 55, 100},
	101: {60, 50, 70, 80, 80, 150},
	102: {60, 40, 80, 60, 45, 40},
	103: {95, 95, 85, 125, 75, 55},
	104: {50, 50, 95, 40, 50, 35},
	105: {60, 80, 110, 50, 80, 45},
	106: {50, 120, 53, 35, 110, 87},
	107: {50, 105, 79, 35, 110, 76},
	108: {90, 55, 75, 60, 75, 30},
	109: {40, 65, 95, 60, 45, 35},
	110: {65, 90, 120, 85, 70, 60},
	111: {80, 85, 95, 30, 30, 25},
	112: {105, 130, 120, 45, 45, 40},
	113: {250, 5, 5, 35, 105, 50},
	114: {65, 55, 115, 100, 40, 60},
	115: {105, 95, 80, 40, 80, 90},
	116: {30, 40, 70, 70, 25, 60},
	117: {55, 65, 95, 95, 45, 85},
	118: {45, 67, 60, 35, 50, 63},
	119: {80, 92, 65, 65, 80, 68},
	120: {30, 45, 55, 70, 55, 85},
	121: {60, 75, 85, 100, 85, 115},
	122: {40, 45, 65, 100, 120, 90},
	123: {70, 110, 80, 55, 80, 105},
	124: {65, 50, 35, 115, 95, 95},
	125: {65, 83, 57, 95, 85, 105},
	126: {65, 95, 57, 100, 85, 93},
	127: {65, 125, 100, 55, 70, 85},
	128: {75, 100, 95, 40, 70, 110},
	129: {20, 10, 55, 15, 20, 80},
	130: {95, 125, 79, 60, 100, 81},
	131: {130, 85, 80, 85, 95, 60},
	132: {48, 48, 48, 48, 48, 48},
	133: {55, 55, 50, 45, 65, 55},
	134: {130, 65, 60, 110, 95, 65},
	135: {65, 65, 60, 110, 95, 130},
	136: {65, 130, 60, 95, 110, 65},
	137: {65, 60, 70, 85, 75, 40},
	138: {35, 40, 100, 90, 55, 35},
	139: {70, 60, 125, 115, 70, 55},
	140: {30, 80, 90, 55, 45, 55},
	141: {60, 115, 105, 65, 70, 80},
	142: {80, 105, 65, 60, 75, 130},
	143: {160, 110, 65, 65, 110, 30},
	144: {90, 85, 100, 95, 125, 85},
	145: {90, 90, 85, 125, 90, 100},
	146: {90, 100, 90, 125, 85, 90},
	147: {41, 64, 45, 50, 50, 50},
	148: {61, 84, 65, 70, 70, 70},
	149: {91, 134, 95, 100, 100, 80},
	150: {106, 110, 90, 154, 90, 130},
	151: {100, 100, 100, 100, 100, 100},
	152: {45, 49, 65, 49, 65, 45},
	153: {60, 62, 80, 63, 80, 60},
	154: {80, 82, 100, 83, 100, 80},
	155: {39, 52, 43, 60, 50, 65},
	156: {58, 64, 58, 80, 65, 80},
	157: {78, 84, 78, 109, 85, 100},
	158: {50, 65, 64, 44, 48, 43},
	159: {65, 80, 80, 59, 63, 58},
	160: {85, 105, 100, 79, 83, 78},
	161: {35, 46, 34, 35, 45, 20},
	162: {85, 76, 64, 45, 55, 90},
	163: {60, 30, 30, 36, 56, 50},
	164: {100, 50, 50, 86, 96, 70},
	165: {40, 20, 30, 40, 80, 55},
	166: {55, 35, 50, 55, 110, 85},
	167: {40, 60, 40, 40, 40, 30},
	168: {70, 90, 70, 60, 70, 40},
	169: {85, 90, 80, 70, 80, 130},
	170: {75, 38, 38, 56, 56, 67},
	171: {125, 58, 58, 76, 76, 67},
	172: {20, 40, 15, 35, 35, 60},
	173: {50, 25, 28, 45, 55, 15},
	174: {90, 30, 15, 40, 20, 15},
	175: {35, 20, 65, 40, 65, 20},
	176: {55, 40, 85, 80, 105, 40},
	177: {40, 50, 45, 70, 45, 70},
	178: {65, 75, 70, 95, 70, 95},
	179: {55, 40, 40, 65, 45, 35},
	180: {70, 55, 55, 80, 60, 45},
	181: {90, 75, 85, 115, 90, 55},
	182: {75, 80, 95, 90, 100, 50},
	183: {70, 20, 50, 20, 50, 40},
	184: {100, 50, 80, 60, 80, 50},
	185: {70, 100, 115, 30, 65, 30},
	186: {90, 75, 75, 90, 100, 70},
	187: {35, 35, 40, 35, 55, 50},
	188: {55, 45, 50, 45, 65, 80},
	189: {75, 55, 70, 55, 95, 110},
	190: {55, 70, 55, 40, 55, 85},
	191: {30, 30, 30, 30, 30, 30},
	192: {75, 75, 55, 105, 85, 30},
	193: {65, 65, 45, 75, 45, 95},
	194: {55, 45, 45, 25, 25, 15},
	195: {95, 85, 85, 65, 65, 35},
	196: {65, 65, 60, 130, 95, 110},
	197: {95, 65, 110, 60, 130, 65},
	198: {60, 85, 42, 85, 42, 91},
	199: {95, 75, 80, 100, 110, 30},
	200: {60, 60, 60, 85, 85, 85},
	201: {48, 72, 48, 72, 48, 48},
	202: {190, 33, 58, 33, 58, 33},
	203: {70, 80, 65, 90, 65, 85},
	204: {50, 65, 90, 35, 35, 15},
	205: {75, 90, 140, 60, 60, 40},
	206: {100, 70, 70, 65, 65, 45},
	207: {65, 75, 105, 35, 65, 85},
	208: {75, 85, 200, 55, 65, 30},
	209: {60, 80, 50, 40, 40, 30},
	210: {90, 120, 75, 60, 60, 45},
	211: {65, 95, 85, 55, 55, 85},
	212: {70, 130, 100, 55, 80, 65},
	213: {20, 10, 230, 10, 230, 5},
	214: {80, 125, 75, 40, 95, 85},
	215: {55, 95, 55, 35, 75, 115},
	216: {60, 80, 50, 50, 50, 40},
	217: {90, 130, 75, 75, 75, 55},
	218: {40, 40, 40, 70, 40, 20},
	219: {60, 50, 120, 90, 80, 30},
	220: {50, 50, 40, 30, 30, 50},
	221: {100, 100, 80, 60, 60, 50},
	222: {65, 55, 95, 65, 95, 35},
	223: {35, 65, 35, 65, 35, 65},
	224: {75, 105, 75, 105, 75, 45},
	225: {45, 55, 45, 65, 45, 75},
	226: {85, 40, 70, 80, 140, 70},
	227: {65, 80, 140, 40, 70, 70},
	228: {45, 60, 30, 80, 50, 65},
	229: {75, 90, 50, 110, 80, 95},
	230: {75, 95, 95, 95, 95, 85},
	231: {90, 60, 60, 40, 40, 40},
	232: {90, 120, 120, 60, 60, 50},
	233: {85, 80, 90, 105, 95, 60},
	234: {73, 95, 62, 85, 65, 85},
	235: {55, 20, 35, 20, 45, 75},
	236: {35, 35, 35, 35, 35, 35},
	237: {50, 95, 95, 35, 110, 70},
	238: {45, 30, 15, 85, 65, 65},
	239: {45, 63, 37, 65, 55, 95},
	240: {45, 75, 37, 70, 55, 83},
	241: {95, 80, 105, 40, 70, 100},
	242: {255, 10, 10, 75, 135, 55},
	243: {90, 85, 75, 115, 100, 115},
	244: {115, 115, 85, 90, 75, 100},
	245: {100, 75, 115, 90, 115, 85},
	246: {50, 64, 50, 45, 50, 41},
	247: {70, 84, 70, 65, 70, 51},
	248: {100, 134, 110, 95, 100, 61},
	249: {106, 90, 130, 90, 154, 110},
	250: {106, 130, 90, 110, 154, 90},
	251: {100, 100, 100, 100, 100, 100},
	252: {40, 45, 35, 65, 55, 70},
	253: {50, 65, 45, 85, 65, 95},
	254: {70, 85, 65, 105, 85, 120},
	255: {45, 60, 40, 70, 50, 45},
	256: {60, 85, 60, 85, 60, 55},
	257: {80, 120, 70, 110, 70, 80},
	258: {50, 70, 50, 50, 50, 40},
	259: {70, 85, 70, 60, 70, 50},
	260: {100, 110, 90, 85, 90, 60},
	261: {35, 55, 35, 30, 30, 35},
	262: {70, 90, 70, 60, 60, 70},
	263: {38, 30, 41, 30, 41, 60},
	264: {78, 70, 61, 50, 61, 100},
	265: {45, 45, 35, 20, 30, 20},
	266: {50, 35, 55, 25, 25, 15},
	267: {60, 70, 50, 100, 50, 65},
	268: {50, 35, 55, 25, 25, 15},
	269: {60, 50, 70, 50, 90, 65},
	270: {40, 30, 30, 40, 50, 30},
	271: {60, 50, 50, 60, 70, 50},
	272: {80, 70, 70, 90, 100, 70},
	273: {40, 40, 50, 30, 30, 30},
	274: {70, 70, 40, 60, 40, 60},
	275: {90, 100, 60, 90, 60, 80},
	276: {40, 55, 30, 30, 30, 85},
	277: {60, 85, 60, 75, 50, 125},
	278: {40, 30, 30, 55, 30, 85},
	279: {60, 50, 100, 95, 70, 65},
	280: {28, 25, 25, 45, 35, 40},
	281: {38, 35, 35, 65, 55, 50},
	282: {68, 65, 65, 125, 115, 80},
	283: {40, 30, 32, 50, 52, 65},
	284: {70, 60, 62, 100, 82, 80},
	285: {60, 40, 60, 40, 60, 35},
	286: {60, 130, 80, 60, 60, 70},
	287: {60, 60, 60, 35, 35, 30},
	288: {80, 80, 80, 55, 55, 90},
	289: {150, 160, 100, 95, 65, 100},
	290: {31, 45, 90, 30, 30, 40},
	291: {61, 90, 45, 50, 50, 160},
	292: {1, 90, 45, 30, 30, 40},
	293: {64, 51, 23, 51, 23, 28},
	294: {84, 71, 43, 71, 43, 48},
	295: {104, 91, 63, 91, 73, 68},
	296: {72, 60, 30, 20, 30, 25},
	297: {144, 120, 60, 40, 60, 50},
	298: {50, 20, 40, 20, 40, 20},
	299: {30, 45, 135, 45, 90, 30},
	300: {50, 45, 45, 35, 35, 50},
	301: {70, 65, 65, 55, 55, 90},
	302: {50, 75, 75, 65, 65, 50},
	303: {50, 85, 85, 55, 55, 50},
	304: {50, 70, 100, 40, 40, 30},
	305: {60, 90, 140, 50, 50, 40},
	306: {70, 110, 180, 60, 60, 50},
	307: {30, 40, 55, 40, 55, 60},
	308: {60, 60, 75, 60, 75, 80},
	309: {40, 45, 40, 65, 40, 65},
	310: {70, 75, 60, 105, 60, 105},
	311: {60, 50, 40, 85, 75, 95},
	312: {60, 40, 50, 75, 85, 95},
	313: {65, 73, 75, 47, 85, 85},
	314: {65, 47, 75, 73, 85, 85},
	315: {50, 60, 45, 100, 80, 65},
	316: {70, 43, 53, 43, 53, 40},
	317: {100, 73, 83, 73, 83, 55},
	318: {45, 90, 20, 65, 20, 65},
	319: {70, 120, 40, 95, 40, 95},
	320: {130, 70, 35, 70, 35, 60},
	321: {170, 90, 45, 90, 45, 60},
	322: {60, 60, 40, 65, 45, 35},
	323: {70, 100, 70, 105, 75, 40},
	324: {70, 85, 140, 85, 70, 20},
	325: {60, 25, 35, 70, 80, 60},
	326: {80, 45, 65, 90, 110, 80},
	327: {60, 60, 60, 60, 60, 60},
	328: {45, 100, 45, 45, 45, 10},
	329: {50, 70, 50, 50, 50, 70},
	330: {80, 100, 80, 80, 80, 100},
	331: {50, 85, 40, 85, 40, 35},
	332: {70, 115, 60, 115, 60, 55},
	333: {45, 40, 60, 40, 75, 50},
	334: {75, 70, 90, 70, 105, 80},
	335: {73, 115, 60, 60, 60, 90},
	336: {73, 100, 60, 100, 60, 65},
	337: {90, 55, 65, 95, 85, 70},
	338: {90, 95, 85, 55, 65, 70},
	339: {50, 48, 43, 46, 41, 60},
	340: {110, 78, 73, 76, 71, 60},
	341: {43, 80, 65, 50, 35, 35},
	342: {63, 120, 85, 90, 55, 55},
	343: {40, 40, 55, 40, 70, 55},
	344: {60, 70, 105, 70, 120, 75},
	345: {66, 41, 77, 61, 87, 23},
	346: {86, 81, 97, 81, 107, 43},
	347: {45, 95, 50, 40, 50, 75},
	348: {75, 125, 100, 70, 80, 45},
	349: {20, 15, 20, 10, 55, 80},
	350: {95, 60, 79, 100, 125, 81},
	351: {70, 70, 70, 70, 70, 70},
	352: {60, 90, 70, 60, 120, 40},
	353: {44, 75, 35, 63, 33, 45},
	354: {64, 115, 65, 83, 63, 65},
	355: {20, 40, 90, 30, 90, 25},
	356: {40, 70, 130, 60, 130, 25},
	357: {99, 68, 83, 72, 87, 51},
	358: {75, 50, 80, 95, 90, 65},
	359: {65, 130, 60, 75, 60, 75},
	360: {95, 23, 48, 23, 48, 23},
	361: {50, 50, 50, 50, 50, 50},
	362: {80, 80, 80, 80, 80, 80},
	363: {70, 40, 50, 55, 50, 25},
	364: {90, 60, 70, 75, 70, 45},
	365: {110, 80, 90, 95, 90, 65},
	366: {35, 64, 85, 74, 55, 32},
	367: {55, 104, 105, 94, 75, 52},
	368: {55, 84, 105, 114, 75, 52},
	369: {100, 90, 130, 45, 65, 55},
	370: {43, 30, 55, 40, 65, 97},
	371: {45, 75, 60, 40, 30, 50},
	372: {65, 95, 100, 60, 50, 50},
	373: {95, 135, 80, 110, 80, 100},
	374: {40, 55, 80, 35, 60, 30},
	375: {60, 75, 100, 55, 80, 50},
	376: {80, 135, 130, 95, 90, 70},
	377: {80, 100, 200, 50, 100, 50},
	378: {80, 50, 100, 100, 200, 50},
	379: {80, 75, 150, 75, 150, 50},
	380: {80, 80, 90, 110, 130, 110},
	381: {80, 90, 80, 130, 110, 110},
	382: {100, 100, 90, 150, 140, 90},
	383: {100, 150, 140, 100, 90, 90},
	384: {105, 150, 90, 150, 90, 95},
	385: {100, 100, 100, 100, 100, 100},
	386: {50, 150, 50, 150, 50, 150},
	387: {55, 68, 64, 45, 55, 31},
	388: {75, 89, 85, 55, 65, 36},
	389: {95, 109, 105, 75, 85, 56},
	390: {44, 58, 44, 58, 44, 61},
	391: {64, 78, 52, 78, 52, 81},
	392: {76, 104, 71, 104, 71, 108},
	393: {53, 51, 53, 61, 56, 40},
	394: {64, 66, 68, 81, 76, 50},
	395: {84, 86, 88, 111, 101, 60},
	396: {40, 55, 30, 30, 30, 60},
	397: {55, 75, 50, 40, 40, 80},
	398: {85, 120, 70, 50, 60, 100},
	399: {59, 45, 40, 35, 40, 31},
	400: {79, 85, 60, 55, 60, 71},
	401: {37, 25, 41, 25, 41, 25},
	402: {77, 85, 51, 55, 51, 65},
	403: {45, 65, 34, 40, 34, 45},
	404: {60, 85, 49, 60, 49, 60},
	405: {80, 120, 79, 95, 79, 70},
	406: {40, 30, 35, 50, 70, 55},
	407: {60, 70, 65, 125, 105, 90},
	408: {67, 125, 40, 30, 30, 58},
	409: {97, 165, 60, 65, 50, 58},
	410: {30, 42, 118, 42, 88, 30},
	411: {60, 52, 168, 47, 138, 30},
	412: {40, 29, 45, 29, 45, 36},
	413: {60, 59, 85, 79, 105, 36},
	414: {70, 94, 50, 94, 50, 66},
	415: {30, 30, 42, 30, 42, 70},
	416: {70, 80, 102, 80, 102, 40},
	417: {60, 45, 70, 45, 90, 95},
	418: {55, 65, 35, 60, 30, 85},
	419: {85, 105, 55, 85, 50, 115},
	420: {45, 35, 45, 62, 53, 35},
	421: {70, 60, 70, 87, 78, 85},
	422: {76, 48, 48, 57, 62, 34},
	423: {111, 83, 68, 92, 82, 39},
	424: {75, 100, 66, 60, 66, 115},
	425: {90, 50, 34, 60, 44, 70},
	426: {150, 80, 44, 90, 54, 80},
	427: {55, 66, 44, 44, 56, 85},
	428: {65, 76, 84, 54, 96, 105},
	429: {60, 60, 60, 105, 105, 105},
	430: {100, 125, 52, 105, 52, 71},
	431: {49, 55, 42, 42, 37, 85},
	432: {71, 82, 64, 64, 59, 112},
	433: {45, 30, 50, 65, 50, 45},
	434: {63, 63, 47, 41, 41, 74},
	435: {103, 93, 67, 71, 61, 84},
	436: {57, 24, 86, 24, 86, 23},
	437: {67, 89, 116, 79, 116, 33},
	438: {50, 80, 95, 10, 45, 10},
	439: {20, 25, 45, 70, 90, 60},
	440: {100, 5, 5, 15, 65, 30},
	441: {76, 65, 45, 92, 42, 91},
	442: {50, 92, 108, 92, 108, 35},
	443: {58, 70, 45, 40, 45, 42},
	444: {68, 90, 65, 50, 55, 82},
	445: {108, 130, 95, 80, 85, 102},
	446: {135, 85, 40, 40, 85, 5},
	447: {40, 70, 40, 35, 40, 60},
	448: {70, 110, 70, 115, 70, 90},
	449: {68, 72, 78, 38, 42, 32},
	450: {108, 112, 118, 68, 72, 47},
	451: {40, 50, 90, 30, 55, 65},
	452: {70, 90, 110, 60, 75, 95},
	453: {48, 61, 40, 61, 40, 50},
	454: {83, 106, 65, 86, 65, 85},
	455: {74, 100, 72, 90, 72, 46},
	456: {49, 49, 56, 49, 61, 66},
	457: {69, 69, 76, 69, 86, 91},
	458: {45, 20, 50, 60, 120, 50},
	459: {60, 62, 50, 62, 60, 40},
	460: {90, 92, 75, 92, 85, 60},
	461: {70, 120, 65, 45, 85, 125},
	462: {70, 70, 115, 130, 90, 60},
	463: {110, 85, 95, 80, 95, 50},
	464: {115, 140, 130, 55, 55, 40},
	465: {100, 100, 125, 110, 50, 50},
	466: {75, 123, 67, 95, 85, 95},
	467: {75, 95, 67, 125, 95, 83},
	468: {85, 50, 95, 120, 115, 80},
	469: {86, 76, 86, 116, 56, 95},
	470: {65, 110, 130, 60, 65, 95},
	471: {65, 60, 110, 130, 95, 65},
	472: {75, 95, 125, 45, 75, 95},
	473: {110, 130, 80, 70, 60, 80},
	474: {85, 80, 70, 135, 75, 90},
	475: {68, 125, 65, 65, 115, 80},
	476: {60, 55, 145, 75, 150, 40},
	477: {45, 100, 135, 65, 135, 45},
	478: {70, 80, 70, 80, 70, 110},
	479: {50, 50, 77, 95, 77, 91},
	480: {75, 75, 130, 75, 130, 95},
	481: {80, 105, 105, 105, 105, 80},
	482: {75, 125, 70, 125, 70, 115},
	483: {100, 120, 120, 150, 100, 90},
	484: {90, 120, 100, 150, 120, 100},
	485: {91, 90, 106, 130, 106, 77},
	486: {110, 160, 110, 80, 110, 100},
	487: {150, 100, 120, 100, 120, 90},
	488: {120, 70, 120, 75, 130, 85},
	489: {80, 80, 80, 80, 80, 80},
	490: {100, 100, 100, 100, 100, 100},
	491: {70, 90, 90, 135, 90, 125},
	492: {100, 100, 100, 100, 100, 100},
	493: {120, 120, 120, 120, 120, 120},
}
