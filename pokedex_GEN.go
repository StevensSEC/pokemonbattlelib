// Code generated - DO NOT EDIT.
// Regenerate with `go generate`.

package pokemonbattlelib

// A map of national pokedex numbers to pokemon data.
type pData struct {
	Name string
	Type Type
}

var pokemonData = map[uint16]pData{
	1:   {Name: "Bulbasaur", Type: 2056},
	2:   {Name: "Ivysaur", Type: 2056},
	3:   {Name: "Venusaur", Type: 2056},
	4:   {Name: "Charmander", Type: 512},
	5:   {Name: "Charmeleon", Type: 512},
	6:   {Name: "Charizard", Type: 516},
	7:   {Name: "Squirtle", Type: 1024},
	8:   {Name: "Wartortle", Type: 1024},
	9:   {Name: "Blastoise", Type: 1024},
	10:  {Name: "Caterpie", Type: 64},
	11:  {Name: "Metapod", Type: 64},
	12:  {Name: "Butterfree", Type: 68},
	13:  {Name: "Weedle", Type: 72},
	14:  {Name: "Kakuna", Type: 72},
	15:  {Name: "Beedrill", Type: 72},
	16:  {Name: "Pidgey", Type: 5},
	17:  {Name: "Pidgeotto", Type: 5},
	18:  {Name: "Pidgeot", Type: 5},
	19:  {Name: "Rattata", Type: 1},
	20:  {Name: "Raticate", Type: 1},
	21:  {Name: "Spearow", Type: 5},
	22:  {Name: "Fearow", Type: 5},
	23:  {Name: "Ekans", Type: 8},
	24:  {Name: "Arbok", Type: 8},
	25:  {Name: "Pikachu", Type: 4096},
	26:  {Name: "Raichu", Type: 4096},
	27:  {Name: "Sandshrew", Type: 16},
	28:  {Name: "Sandslash", Type: 16},
	29:  {Name: "Nidoran♀", Type: 8},
	30:  {Name: "Nidorina", Type: 8},
	31:  {Name: "Nidoqueen", Type: 24},
	32:  {Name: "Nidoran♂", Type: 8},
	33:  {Name: "Nidorino", Type: 8},
	34:  {Name: "Nidoking", Type: 24},
	35:  {Name: "Clefairy", Type: 131072},
	36:  {Name: "Clefable", Type: 131072},
	37:  {Name: "Vulpix", Type: 512},
	38:  {Name: "Ninetales", Type: 512},
	39:  {Name: "Jigglypuff", Type: 131073},
	40:  {Name: "Wigglytuff", Type: 131073},
	41:  {Name: "Zubat", Type: 12},
	42:  {Name: "Golbat", Type: 12},
	43:  {Name: "Oddish", Type: 2056},
	44:  {Name: "Gloom", Type: 2056},
	45:  {Name: "Vileplume", Type: 2056},
	46:  {Name: "Paras", Type: 2112},
	47:  {Name: "Parasect", Type: 2112},
	48:  {Name: "Venonat", Type: 72},
	49:  {Name: "Venomoth", Type: 72},
	50:  {Name: "Diglett", Type: 16},
	51:  {Name: "Dugtrio", Type: 16},
	52:  {Name: "Meowth", Type: 1},
	53:  {Name: "Persian", Type: 1},
	54:  {Name: "Psyduck", Type: 1024},
	55:  {Name: "Golduck", Type: 1024},
	56:  {Name: "Mankey", Type: 2},
	57:  {Name: "Primeape", Type: 2},
	58:  {Name: "Growlithe", Type: 512},
	59:  {Name: "Arcanine", Type: 512},
	60:  {Name: "Poliwag", Type: 1024},
	61:  {Name: "Poliwhirl", Type: 1024},
	62:  {Name: "Poliwrath", Type: 1026},
	63:  {Name: "Abra", Type: 8192},
	64:  {Name: "Kadabra", Type: 8192},
	65:  {Name: "Alakazam", Type: 8192},
	66:  {Name: "Machop", Type: 2},
	67:  {Name: "Machoke", Type: 2},
	68:  {Name: "Machamp", Type: 2},
	69:  {Name: "Bellsprout", Type: 2056},
	70:  {Name: "Weepinbell", Type: 2056},
	71:  {Name: "Victreebel", Type: 2056},
	72:  {Name: "Tentacool", Type: 1032},
	73:  {Name: "Tentacruel", Type: 1032},
	74:  {Name: "Geodude", Type: 48},
	75:  {Name: "Graveler", Type: 48},
	76:  {Name: "Golem", Type: 48},
	77:  {Name: "Ponyta", Type: 512},
	78:  {Name: "Rapidash", Type: 512},
	79:  {Name: "Slowpoke", Type: 9216},
	80:  {Name: "Slowbro", Type: 9216},
	81:  {Name: "Magnemite", Type: 4352},
	82:  {Name: "Magneton", Type: 4352},
	83:  {Name: "Farfetch’d", Type: 5},
	84:  {Name: "Doduo", Type: 5},
	85:  {Name: "Dodrio", Type: 5},
	86:  {Name: "Seel", Type: 1024},
	87:  {Name: "Dewgong", Type: 17408},
	88:  {Name: "Grimer", Type: 8},
	89:  {Name: "Muk", Type: 8},
	90:  {Name: "Shellder", Type: 1024},
	91:  {Name: "Cloyster", Type: 17408},
	92:  {Name: "Gastly", Type: 136},
	93:  {Name: "Haunter", Type: 136},
	94:  {Name: "Gengar", Type: 136},
	95:  {Name: "Onix", Type: 48},
	96:  {Name: "Drowzee", Type: 8192},
	97:  {Name: "Hypno", Type: 8192},
	98:  {Name: "Krabby", Type: 1024},
	99:  {Name: "Kingler", Type: 1024},
	100: {Name: "Voltorb", Type: 4096},
	101: {Name: "Electrode", Type: 4096},
	102: {Name: "Exeggcute", Type: 10240},
	103: {Name: "Exeggutor", Type: 10240},
	104: {Name: "Cubone", Type: 16},
	105: {Name: "Marowak", Type: 16},
	106: {Name: "Hitmonlee", Type: 2},
	107: {Name: "Hitmonchan", Type: 2},
	108: {Name: "Lickitung", Type: 1},
	109: {Name: "Koffing", Type: 8},
	110: {Name: "Weezing", Type: 8},
	111: {Name: "Rhyhorn", Type: 48},
	112: {Name: "Rhydon", Type: 48},
	113: {Name: "Chansey", Type: 1},
	114: {Name: "Tangela", Type: 2048},
	115: {Name: "Kangaskhan", Type: 1},
	116: {Name: "Horsea", Type: 1024},
	117: {Name: "Seadra", Type: 1024},
	118: {Name: "Goldeen", Type: 1024},
	119: {Name: "Seaking", Type: 1024},
	120: {Name: "Staryu", Type: 1024},
	121: {Name: "Starmie", Type: 9216},
	122: {Name: "Mr. Mime", Type: 139264},
	123: {Name: "Scyther", Type: 68},
	124: {Name: "Jynx", Type: 24576},
	125: {Name: "Electabuzz", Type: 4096},
	126: {Name: "Magmar", Type: 512},
	127: {Name: "Pinsir", Type: 64},
	128: {Name: "Tauros", Type: 1},
	129: {Name: "Magikarp", Type: 1024},
	130: {Name: "Gyarados", Type: 1028},
	131: {Name: "Lapras", Type: 17408},
	132: {Name: "Ditto", Type: 1},
	133: {Name: "Eevee", Type: 1},
	134: {Name: "Vaporeon", Type: 1024},
	135: {Name: "Jolteon", Type: 4096},
	136: {Name: "Flareon", Type: 512},
	137: {Name: "Porygon", Type: 1},
	138: {Name: "Omanyte", Type: 1056},
	139: {Name: "Omastar", Type: 1056},
	140: {Name: "Kabuto", Type: 1056},
	141: {Name: "Kabutops", Type: 1056},
	142: {Name: "Aerodactyl", Type: 36},
	143: {Name: "Snorlax", Type: 1},
	144: {Name: "Articuno", Type: 16388},
	145: {Name: "Zapdos", Type: 4100},
	146: {Name: "Moltres", Type: 516},
	147: {Name: "Dratini", Type: 32768},
	148: {Name: "Dragonair", Type: 32768},
	149: {Name: "Dragonite", Type: 32772},
	150: {Name: "Mewtwo", Type: 8192},
	151: {Name: "Mew", Type: 8192},
	152: {Name: "Chikorita", Type: 2048},
	153: {Name: "Bayleef", Type: 2048},
	154: {Name: "Meganium", Type: 2048},
	155: {Name: "Cyndaquil", Type: 512},
	156: {Name: "Quilava", Type: 512},
	157: {Name: "Typhlosion", Type: 512},
	158: {Name: "Totodile", Type: 1024},
	159: {Name: "Croconaw", Type: 1024},
	160: {Name: "Feraligatr", Type: 1024},
	161: {Name: "Sentret", Type: 1},
	162: {Name: "Furret", Type: 1},
	163: {Name: "Hoothoot", Type: 5},
	164: {Name: "Noctowl", Type: 5},
	165: {Name: "Ledyba", Type: 68},
	166: {Name: "Ledian", Type: 68},
	167: {Name: "Spinarak", Type: 72},
	168: {Name: "Ariados", Type: 72},
	169: {Name: "Crobat", Type: 12},
	170: {Name: "Chinchou", Type: 5120},
	171: {Name: "Lanturn", Type: 5120},
	172: {Name: "Pichu", Type: 4096},
	173: {Name: "Cleffa", Type: 131072},
	174: {Name: "Igglybuff", Type: 131073},
	175: {Name: "Togepi", Type: 131072},
	176: {Name: "Togetic", Type: 131076},
	177: {Name: "Natu", Type: 8196},
	178: {Name: "Xatu", Type: 8196},
	179: {Name: "Mareep", Type: 4096},
	180: {Name: "Flaaffy", Type: 4096},
	181: {Name: "Ampharos", Type: 4096},
	182: {Name: "Bellossom", Type: 2048},
	183: {Name: "Marill", Type: 132096},
	184: {Name: "Azumarill", Type: 132096},
	185: {Name: "Sudowoodo", Type: 32},
	186: {Name: "Politoed", Type: 1024},
	187: {Name: "Hoppip", Type: 2052},
	188: {Name: "Skiploom", Type: 2052},
	189: {Name: "Jumpluff", Type: 2052},
	190: {Name: "Aipom", Type: 1},
	191: {Name: "Sunkern", Type: 2048},
	192: {Name: "Sunflora", Type: 2048},
	193: {Name: "Yanma", Type: 68},
	194: {Name: "Wooper", Type: 1040},
	195: {Name: "Quagsire", Type: 1040},
	196: {Name: "Espeon", Type: 8192},
	197: {Name: "Umbreon", Type: 65536},
	198: {Name: "Murkrow", Type: 65540},
	199: {Name: "Slowking", Type: 9216},
	200: {Name: "Misdreavus", Type: 128},
	201: {Name: "Unown", Type: 8192},
	202: {Name: "Wobbuffet", Type: 8192},
	203: {Name: "Girafarig", Type: 8193},
	204: {Name: "Pineco", Type: 64},
	205: {Name: "Forretress", Type: 320},
	206: {Name: "Dunsparce", Type: 1},
	207: {Name: "Gligar", Type: 20},
	208: {Name: "Steelix", Type: 272},
	209: {Name: "Snubbull", Type: 131072},
	210: {Name: "Granbull", Type: 131072},
	211: {Name: "Qwilfish", Type: 1032},
	212: {Name: "Scizor", Type: 320},
	213: {Name: "Shuckle", Type: 96},
	214: {Name: "Heracross", Type: 66},
	215: {Name: "Sneasel", Type: 81920},
	216: {Name: "Teddiursa", Type: 1},
	217: {Name: "Ursaring", Type: 1},
	218: {Name: "Slugma", Type: 512},
	219: {Name: "Magcargo", Type: 544},
	220: {Name: "Swinub", Type: 16400},
	221: {Name: "Piloswine", Type: 16400},
	222: {Name: "Corsola", Type: 1056},
	223: {Name: "Remoraid", Type: 1024},
	224: {Name: "Octillery", Type: 1024},
	225: {Name: "Delibird", Type: 16388},
	226: {Name: "Mantine", Type: 1028},
	227: {Name: "Skarmory", Type: 260},
	228: {Name: "Houndour", Type: 66048},
	229: {Name: "Houndoom", Type: 66048},
	230: {Name: "Kingdra", Type: 33792},
	231: {Name: "Phanpy", Type: 16},
	232: {Name: "Donphan", Type: 16},
	233: {Name: "Porygon2", Type: 1},
	234: {Name: "Stantler", Type: 1},
	235: {Name: "Smeargle", Type: 1},
	236: {Name: "Tyrogue", Type: 2},
	237: {Name: "Hitmontop", Type: 2},
	238: {Name: "Smoochum", Type: 24576},
	239: {Name: "Elekid", Type: 4096},
	240: {Name: "Magby", Type: 512},
	241: {Name: "Miltank", Type: 1},
	242: {Name: "Blissey", Type: 1},
	243: {Name: "Raikou", Type: 4096},
	244: {Name: "Entei", Type: 512},
	245: {Name: "Suicune", Type: 1024},
	246: {Name: "Larvitar", Type: 48},
	247: {Name: "Pupitar", Type: 48},
	248: {Name: "Tyranitar", Type: 65568},
	249: {Name: "Lugia", Type: 8196},
	250: {Name: "Ho-Oh", Type: 516},
	251: {Name: "Celebi", Type: 10240},
	252: {Name: "Treecko", Type: 2048},
	253: {Name: "Grovyle", Type: 2048},
	254: {Name: "Sceptile", Type: 2048},
	255: {Name: "Torchic", Type: 512},
	256: {Name: "Combusken", Type: 514},
	257: {Name: "Blaziken", Type: 514},
	258: {Name: "Mudkip", Type: 1024},
	259: {Name: "Marshtomp", Type: 1040},
	260: {Name: "Swampert", Type: 1040},
	261: {Name: "Poochyena", Type: 65536},
	262: {Name: "Mightyena", Type: 65536},
	263: {Name: "Zigzagoon", Type: 1},
	264: {Name: "Linoone", Type: 1},
	265: {Name: "Wurmple", Type: 64},
	266: {Name: "Silcoon", Type: 64},
	267: {Name: "Beautifly", Type: 68},
	268: {Name: "Cascoon", Type: 64},
	269: {Name: "Dustox", Type: 72},
	270: {Name: "Lotad", Type: 3072},
	271: {Name: "Lombre", Type: 3072},
	272: {Name: "Ludicolo", Type: 3072},
	273: {Name: "Seedot", Type: 2048},
	274: {Name: "Nuzleaf", Type: 67584},
	275: {Name: "Shiftry", Type: 67584},
	276: {Name: "Taillow", Type: 5},
	277: {Name: "Swellow", Type: 5},
	278: {Name: "Wingull", Type: 1028},
	279: {Name: "Pelipper", Type: 1028},
	280: {Name: "Ralts", Type: 139264},
	281: {Name: "Kirlia", Type: 139264},
	282: {Name: "Gardevoir", Type: 139264},
	283: {Name: "Surskit", Type: 1088},
	284: {Name: "Masquerain", Type: 68},
	285: {Name: "Shroomish", Type: 2048},
	286: {Name: "Breloom", Type: 2050},
	287: {Name: "Slakoth", Type: 1},
	288: {Name: "Vigoroth", Type: 1},
	289: {Name: "Slaking", Type: 1},
	290: {Name: "Nincada", Type: 80},
	291: {Name: "Ninjask", Type: 68},
	292: {Name: "Shedinja", Type: 192},
	293: {Name: "Whismur", Type: 1},
	294: {Name: "Loudred", Type: 1},
	295: {Name: "Exploud", Type: 1},
	296: {Name: "Makuhita", Type: 2},
	297: {Name: "Hariyama", Type: 2},
	298: {Name: "Azurill", Type: 131073},
	299: {Name: "Nosepass", Type: 32},
	300: {Name: "Skitty", Type: 1},
	301: {Name: "Delcatty", Type: 1},
	302: {Name: "Sableye", Type: 65664},
	303: {Name: "Mawile", Type: 131328},
	304: {Name: "Aron", Type: 288},
	305: {Name: "Lairon", Type: 288},
	306: {Name: "Aggron", Type: 288},
	307: {Name: "Meditite", Type: 8194},
	308: {Name: "Medicham", Type: 8194},
	309: {Name: "Electrike", Type: 4096},
	310: {Name: "Manectric", Type: 4096},
	311: {Name: "Plusle", Type: 4096},
	312: {Name: "Minun", Type: 4096},
	313: {Name: "Volbeat", Type: 64},
	314: {Name: "Illumise", Type: 64},
	315: {Name: "Roselia", Type: 2056},
	316: {Name: "Gulpin", Type: 8},
	317: {Name: "Swalot", Type: 8},
	318: {Name: "Carvanha", Type: 66560},
	319: {Name: "Sharpedo", Type: 66560},
	320: {Name: "Wailmer", Type: 1024},
	321: {Name: "Wailord", Type: 1024},
	322: {Name: "Numel", Type: 528},
	323: {Name: "Camerupt", Type: 528},
	324: {Name: "Torkoal", Type: 512},
	325: {Name: "Spoink", Type: 8192},
	326: {Name: "Grumpig", Type: 8192},
	327: {Name: "Spinda", Type: 1},
	328: {Name: "Trapinch", Type: 16},
	329: {Name: "Vibrava", Type: 32784},
	330: {Name: "Flygon", Type: 32784},
	331: {Name: "Cacnea", Type: 2048},
	332: {Name: "Cacturne", Type: 67584},
	333: {Name: "Swablu", Type: 5},
	334: {Name: "Altaria", Type: 32772},
	335: {Name: "Zangoose", Type: 1},
	336: {Name: "Seviper", Type: 8},
	337: {Name: "Lunatone", Type: 8224},
	338: {Name: "Solrock", Type: 8224},
	339: {Name: "Barboach", Type: 1040},
	340: {Name: "Whiscash", Type: 1040},
	341: {Name: "Corphish", Type: 1024},
	342: {Name: "Crawdaunt", Type: 66560},
	343: {Name: "Baltoy", Type: 8208},
	344: {Name: "Claydol", Type: 8208},
	345: {Name: "Lileep", Type: 2080},
	346: {Name: "Cradily", Type: 2080},
	347: {Name: "Anorith", Type: 96},
	348: {Name: "Armaldo", Type: 96},
	349: {Name: "Feebas", Type: 1024},
	350: {Name: "Milotic", Type: 1024},
	351: {Name: "Castform", Type: 1},
	352: {Name: "Kecleon", Type: 1},
	353: {Name: "Shuppet", Type: 128},
	354: {Name: "Banette", Type: 128},
	355: {Name: "Duskull", Type: 128},
	356: {Name: "Dusclops", Type: 128},
	357: {Name: "Tropius", Type: 2052},
	358: {Name: "Chimecho", Type: 8192},
	359: {Name: "Absol", Type: 65536},
	360: {Name: "Wynaut", Type: 8192},
	361: {Name: "Snorunt", Type: 16384},
	362: {Name: "Glalie", Type: 16384},
	363: {Name: "Spheal", Type: 17408},
	364: {Name: "Sealeo", Type: 17408},
	365: {Name: "Walrein", Type: 17408},
	366: {Name: "Clamperl", Type: 1024},
	367: {Name: "Huntail", Type: 1024},
	368: {Name: "Gorebyss", Type: 1024},
	369: {Name: "Relicanth", Type: 1056},
	370: {Name: "Luvdisc", Type: 1024},
	371: {Name: "Bagon", Type: 32768},
	372: {Name: "Shelgon", Type: 32768},
	373: {Name: "Salamence", Type: 32772},
	374: {Name: "Beldum", Type: 8448},
	375: {Name: "Metang", Type: 8448},
	376: {Name: "Metagross", Type: 8448},
	377: {Name: "Regirock", Type: 32},
	378: {Name: "Regice", Type: 16384},
	379: {Name: "Registeel", Type: 256},
	380: {Name: "Latias", Type: 40960},
	381: {Name: "Latios", Type: 40960},
	382: {Name: "Kyogre", Type: 1024},
	383: {Name: "Groudon", Type: 16},
	384: {Name: "Rayquaza", Type: 32772},
	385: {Name: "Jirachi", Type: 8448},
	386: {Name: "Deoxys", Type: 8192},
	387: {Name: "Turtwig", Type: 2048},
	388: {Name: "Grotle", Type: 2048},
	389: {Name: "Torterra", Type: 2064},
	390: {Name: "Chimchar", Type: 512},
	391: {Name: "Monferno", Type: 514},
	392: {Name: "Infernape", Type: 514},
	393: {Name: "Piplup", Type: 1024},
	394: {Name: "Prinplup", Type: 1024},
	395: {Name: "Empoleon", Type: 1280},
	396: {Name: "Starly", Type: 5},
	397: {Name: "Staravia", Type: 5},
	398: {Name: "Staraptor", Type: 5},
	399: {Name: "Bidoof", Type: 1},
	400: {Name: "Bibarel", Type: 1025},
	401: {Name: "Kricketot", Type: 64},
	402: {Name: "Kricketune", Type: 64},
	403: {Name: "Shinx", Type: 4096},
	404: {Name: "Luxio", Type: 4096},
	405: {Name: "Luxray", Type: 4096},
	406: {Name: "Budew", Type: 2056},
	407: {Name: "Roserade", Type: 2056},
	408: {Name: "Cranidos", Type: 32},
	409: {Name: "Rampardos", Type: 32},
	410: {Name: "Shieldon", Type: 288},
	411: {Name: "Bastiodon", Type: 288},
	412: {Name: "Burmy", Type: 64},
	413: {Name: "Wormadam", Type: 2112},
	414: {Name: "Mothim", Type: 68},
	415: {Name: "Combee", Type: 68},
	416: {Name: "Vespiquen", Type: 68},
	417: {Name: "Pachirisu", Type: 4096},
	418: {Name: "Buizel", Type: 1024},
	419: {Name: "Floatzel", Type: 1024},
	420: {Name: "Cherubi", Type: 2048},
	421: {Name: "Cherrim", Type: 2048},
	422: {Name: "Shellos", Type: 1024},
	423: {Name: "Gastrodon", Type: 1040},
	424: {Name: "Ambipom", Type: 1},
	425: {Name: "Drifloon", Type: 132},
	426: {Name: "Drifblim", Type: 132},
	427: {Name: "Buneary", Type: 1},
	428: {Name: "Lopunny", Type: 1},
	429: {Name: "Mismagius", Type: 128},
	430: {Name: "Honchkrow", Type: 65540},
	431: {Name: "Glameow", Type: 1},
	432: {Name: "Purugly", Type: 1},
	433: {Name: "Chingling", Type: 8192},
	434: {Name: "Stunky", Type: 65544},
	435: {Name: "Skuntank", Type: 65544},
	436: {Name: "Bronzor", Type: 8448},
	437: {Name: "Bronzong", Type: 8448},
	438: {Name: "Bonsly", Type: 32},
	439: {Name: "Mime Jr.", Type: 139264},
	440: {Name: "Happiny", Type: 1},
	441: {Name: "Chatot", Type: 5},
	442: {Name: "Spiritomb", Type: 65664},
	443: {Name: "Gible", Type: 32784},
	444: {Name: "Gabite", Type: 32784},
	445: {Name: "Garchomp", Type: 32784},
	446: {Name: "Munchlax", Type: 1},
	447: {Name: "Riolu", Type: 2},
	448: {Name: "Lucario", Type: 258},
	449: {Name: "Hippopotas", Type: 16},
	450: {Name: "Hippowdon", Type: 16},
	451: {Name: "Skorupi", Type: 72},
	452: {Name: "Drapion", Type: 65544},
	453: {Name: "Croagunk", Type: 10},
	454: {Name: "Toxicroak", Type: 10},
	455: {Name: "Carnivine", Type: 2048},
	456: {Name: "Finneon", Type: 1024},
	457: {Name: "Lumineon", Type: 1024},
	458: {Name: "Mantyke", Type: 1028},
	459: {Name: "Snover", Type: 18432},
	460: {Name: "Abomasnow", Type: 18432},
	461: {Name: "Weavile", Type: 81920},
	462: {Name: "Magnezone", Type: 4352},
	463: {Name: "Lickilicky", Type: 1},
	464: {Name: "Rhyperior", Type: 48},
	465: {Name: "Tangrowth", Type: 2048},
	466: {Name: "Electivire", Type: 4096},
	467: {Name: "Magmortar", Type: 512},
	468: {Name: "Togekiss", Type: 131076},
	469: {Name: "Yanmega", Type: 68},
	470: {Name: "Leafeon", Type: 2048},
	471: {Name: "Glaceon", Type: 16384},
	472: {Name: "Gliscor", Type: 20},
	473: {Name: "Mamoswine", Type: 16400},
	474: {Name: "Porygon-Z", Type: 1},
	475: {Name: "Gallade", Type: 8194},
	476: {Name: "Probopass", Type: 288},
	477: {Name: "Dusknoir", Type: 128},
	478: {Name: "Froslass", Type: 16512},
	479: {Name: "Rotom", Type: 4224},
	480: {Name: "Uxie", Type: 8192},
	481: {Name: "Mesprit", Type: 8192},
	482: {Name: "Azelf", Type: 8192},
	483: {Name: "Dialga", Type: 33024},
	484: {Name: "Palkia", Type: 33792},
	485: {Name: "Heatran", Type: 768},
	486: {Name: "Regigigas", Type: 1},
	487: {Name: "Giratina", Type: 32896},
	488: {Name: "Cresselia", Type: 8192},
	489: {Name: "Phione", Type: 1024},
	490: {Name: "Manaphy", Type: 1024},
	491: {Name: "Darkrai", Type: 65536},
	492: {Name: "Shaymin", Type: 2048},
	493: {Name: "Arceus", Type: 1},
}

// Pokemon const enum for quick lookup
const (
	PkmnBulbasaur  = 1
	PkmnIvysaur    = 2
	PkmnVenusaur   = 3
	PkmnCharmander = 4
	PkmnCharmeleon = 5
	PkmnCharizard  = 6
	PkmnSquirtle   = 7
	PkmnWartortle  = 8
	PkmnBlastoise  = 9
	PkmnCaterpie   = 10
	PkmnMetapod    = 11
	PkmnButterfree = 12
	PkmnWeedle     = 13
	PkmnKakuna     = 14
	PkmnBeedrill   = 15
	PkmnPidgey     = 16
	PkmnPidgeotto  = 17
	PkmnPidgeot    = 18
	PkmnRattata    = 19
	PkmnRaticate   = 20
	PkmnSpearow    = 21
	PkmnFearow     = 22
	PkmnEkans      = 23
	PkmnArbok      = 24
	PkmnPikachu    = 25
	PkmnRaichu     = 26
	PkmnSandshrew  = 27
	PkmnSandslash  = 28
	PkmnNidoranF   = 29
	PkmnNidorina   = 30
	PkmnNidoqueen  = 31
	PkmnNidoranM   = 32
	PkmnNidorino   = 33
	PkmnNidoking   = 34
	PkmnClefairy   = 35
	PkmnClefable   = 36
	PkmnVulpix     = 37
	PkmnNinetales  = 38
	PkmnJigglypuff = 39
	PkmnWigglytuff = 40
	PkmnZubat      = 41
	PkmnGolbat     = 42
	PkmnOddish     = 43
	PkmnGloom      = 44
	PkmnVileplume  = 45
	PkmnParas      = 46
	PkmnParasect   = 47
	PkmnVenonat    = 48
	PkmnVenomoth   = 49
	PkmnDiglett    = 50
	PkmnDugtrio    = 51
	PkmnMeowth     = 52
	PkmnPersian    = 53
	PkmnPsyduck    = 54
	PkmnGolduck    = 55
	PkmnMankey     = 56
	PkmnPrimeape   = 57
	PkmnGrowlithe  = 58
	PkmnArcanine   = 59
	PkmnPoliwag    = 60
	PkmnPoliwhirl  = 61
	PkmnPoliwrath  = 62
	PkmnAbra       = 63
	PkmnKadabra    = 64
	PkmnAlakazam   = 65
	PkmnMachop     = 66
	PkmnMachoke    = 67
	PkmnMachamp    = 68
	PkmnBellsprout = 69
	PkmnWeepinbell = 70
	PkmnVictreebel = 71
	PkmnTentacool  = 72
	PkmnTentacruel = 73
	PkmnGeodude    = 74
	PkmnGraveler   = 75
	PkmnGolem      = 76
	PkmnPonyta     = 77
	PkmnRapidash   = 78
	PkmnSlowpoke   = 79
	PkmnSlowbro    = 80
	PkmnMagnemite  = 81
	PkmnMagneton   = 82
	PkmnFarfetchd  = 83
	PkmnDoduo      = 84
	PkmnDodrio     = 85
	PkmnSeel       = 86
	PkmnDewgong    = 87
	PkmnGrimer     = 88
	PkmnMuk        = 89
	PkmnShellder   = 90
	PkmnCloyster   = 91
	PkmnGastly     = 92
	PkmnHaunter    = 93
	PkmnGengar     = 94
	PkmnOnix       = 95
	PkmnDrowzee    = 96
	PkmnHypno      = 97
	PkmnKrabby     = 98
	PkmnKingler    = 99
	PkmnVoltorb    = 100
	PkmnElectrode  = 101
	PkmnExeggcute  = 102
	PkmnExeggutor  = 103
	PkmnCubone     = 104
	PkmnMarowak    = 105
	PkmnHitmonlee  = 106
	PkmnHitmonchan = 107
	PkmnLickitung  = 108
	PkmnKoffing    = 109
	PkmnWeezing    = 110
	PkmnRhyhorn    = 111
	PkmnRhydon     = 112
	PkmnChansey    = 113
	PkmnTangela    = 114
	PkmnKangaskhan = 115
	PkmnHorsea     = 116
	PkmnSeadra     = 117
	PkmnGoldeen    = 118
	PkmnSeaking    = 119
	PkmnStaryu     = 120
	PkmnStarmie    = 121
	PkmnMrMime     = 122
	PkmnScyther    = 123
	PkmnJynx       = 124
	PkmnElectabuzz = 125
	PkmnMagmar     = 126
	PkmnPinsir     = 127
	PkmnTauros     = 128
	PkmnMagikarp   = 129
	PkmnGyarados   = 130
	PkmnLapras     = 131
	PkmnDitto      = 132
	PkmnEevee      = 133
	PkmnVaporeon   = 134
	PkmnJolteon    = 135
	PkmnFlareon    = 136
	PkmnPorygon    = 137
	PkmnOmanyte    = 138
	PkmnOmastar    = 139
	PkmnKabuto     = 140
	PkmnKabutops   = 141
	PkmnAerodactyl = 142
	PkmnSnorlax    = 143
	PkmnArticuno   = 144
	PkmnZapdos     = 145
	PkmnMoltres    = 146
	PkmnDratini    = 147
	PkmnDragonair  = 148
	PkmnDragonite  = 149
	PkmnMewtwo     = 150
	PkmnMew        = 151
	PkmnChikorita  = 152
	PkmnBayleef    = 153
	PkmnMeganium   = 154
	PkmnCyndaquil  = 155
	PkmnQuilava    = 156
	PkmnTyphlosion = 157
	PkmnTotodile   = 158
	PkmnCroconaw   = 159
	PkmnFeraligatr = 160
	PkmnSentret    = 161
	PkmnFurret     = 162
	PkmnHoothoot   = 163
	PkmnNoctowl    = 164
	PkmnLedyba     = 165
	PkmnLedian     = 166
	PkmnSpinarak   = 167
	PkmnAriados    = 168
	PkmnCrobat     = 169
	PkmnChinchou   = 170
	PkmnLanturn    = 171
	PkmnPichu      = 172
	PkmnCleffa     = 173
	PkmnIgglybuff  = 174
	PkmnTogepi     = 175
	PkmnTogetic    = 176
	PkmnNatu       = 177
	PkmnXatu       = 178
	PkmnMareep     = 179
	PkmnFlaaffy    = 180
	PkmnAmpharos   = 181
	PkmnBellossom  = 182
	PkmnMarill     = 183
	PkmnAzumarill  = 184
	PkmnSudowoodo  = 185
	PkmnPolitoed   = 186
	PkmnHoppip     = 187
	PkmnSkiploom   = 188
	PkmnJumpluff   = 189
	PkmnAipom      = 190
	PkmnSunkern    = 191
	PkmnSunflora   = 192
	PkmnYanma      = 193
	PkmnWooper     = 194
	PkmnQuagsire   = 195
	PkmnEspeon     = 196
	PkmnUmbreon    = 197
	PkmnMurkrow    = 198
	PkmnSlowking   = 199
	PkmnMisdreavus = 200
	PkmnUnown      = 201
	PkmnWobbuffet  = 202
	PkmnGirafarig  = 203
	PkmnPineco     = 204
	PkmnForretress = 205
	PkmnDunsparce  = 206
	PkmnGligar     = 207
	PkmnSteelix    = 208
	PkmnSnubbull   = 209
	PkmnGranbull   = 210
	PkmnQwilfish   = 211
	PkmnScizor     = 212
	PkmnShuckle    = 213
	PkmnHeracross  = 214
	PkmnSneasel    = 215
	PkmnTeddiursa  = 216
	PkmnUrsaring   = 217
	PkmnSlugma     = 218
	PkmnMagcargo   = 219
	PkmnSwinub     = 220
	PkmnPiloswine  = 221
	PkmnCorsola    = 222
	PkmnRemoraid   = 223
	PkmnOctillery  = 224
	PkmnDelibird   = 225
	PkmnMantine    = 226
	PkmnSkarmory   = 227
	PkmnHoundour   = 228
	PkmnHoundoom   = 229
	PkmnKingdra    = 230
	PkmnPhanpy     = 231
	PkmnDonphan    = 232
	PkmnPorygon2   = 233
	PkmnStantler   = 234
	PkmnSmeargle   = 235
	PkmnTyrogue    = 236
	PkmnHitmontop  = 237
	PkmnSmoochum   = 238
	PkmnElekid     = 239
	PkmnMagby      = 240
	PkmnMiltank    = 241
	PkmnBlissey    = 242
	PkmnRaikou     = 243
	PkmnEntei      = 244
	PkmnSuicune    = 245
	PkmnLarvitar   = 246
	PkmnPupitar    = 247
	PkmnTyranitar  = 248
	PkmnLugia      = 249
	PkmnHoOh       = 250
	PkmnCelebi     = 251
	PkmnTreecko    = 252
	PkmnGrovyle    = 253
	PkmnSceptile   = 254
	PkmnTorchic    = 255
	PkmnCombusken  = 256
	PkmnBlaziken   = 257
	PkmnMudkip     = 258
	PkmnMarshtomp  = 259
	PkmnSwampert   = 260
	PkmnPoochyena  = 261
	PkmnMightyena  = 262
	PkmnZigzagoon  = 263
	PkmnLinoone    = 264
	PkmnWurmple    = 265
	PkmnSilcoon    = 266
	PkmnBeautifly  = 267
	PkmnCascoon    = 268
	PkmnDustox     = 269
	PkmnLotad      = 270
	PkmnLombre     = 271
	PkmnLudicolo   = 272
	PkmnSeedot     = 273
	PkmnNuzleaf    = 274
	PkmnShiftry    = 275
	PkmnTaillow    = 276
	PkmnSwellow    = 277
	PkmnWingull    = 278
	PkmnPelipper   = 279
	PkmnRalts      = 280
	PkmnKirlia     = 281
	PkmnGardevoir  = 282
	PkmnSurskit    = 283
	PkmnMasquerain = 284
	PkmnShroomish  = 285
	PkmnBreloom    = 286
	PkmnSlakoth    = 287
	PkmnVigoroth   = 288
	PkmnSlaking    = 289
	PkmnNincada    = 290
	PkmnNinjask    = 291
	PkmnShedinja   = 292
	PkmnWhismur    = 293
	PkmnLoudred    = 294
	PkmnExploud    = 295
	PkmnMakuhita   = 296
	PkmnHariyama   = 297
	PkmnAzurill    = 298
	PkmnNosepass   = 299
	PkmnSkitty     = 300
	PkmnDelcatty   = 301
	PkmnSableye    = 302
	PkmnMawile     = 303
	PkmnAron       = 304
	PkmnLairon     = 305
	PkmnAggron     = 306
	PkmnMeditite   = 307
	PkmnMedicham   = 308
	PkmnElectrike  = 309
	PkmnManectric  = 310
	PkmnPlusle     = 311
	PkmnMinun      = 312
	PkmnVolbeat    = 313
	PkmnIllumise   = 314
	PkmnRoselia    = 315
	PkmnGulpin     = 316
	PkmnSwalot     = 317
	PkmnCarvanha   = 318
	PkmnSharpedo   = 319
	PkmnWailmer    = 320
	PkmnWailord    = 321
	PkmnNumel      = 322
	PkmnCamerupt   = 323
	PkmnTorkoal    = 324
	PkmnSpoink     = 325
	PkmnGrumpig    = 326
	PkmnSpinda     = 327
	PkmnTrapinch   = 328
	PkmnVibrava    = 329
	PkmnFlygon     = 330
	PkmnCacnea     = 331
	PkmnCacturne   = 332
	PkmnSwablu     = 333
	PkmnAltaria    = 334
	PkmnZangoose   = 335
	PkmnSeviper    = 336
	PkmnLunatone   = 337
	PkmnSolrock    = 338
	PkmnBarboach   = 339
	PkmnWhiscash   = 340
	PkmnCorphish   = 341
	PkmnCrawdaunt  = 342
	PkmnBaltoy     = 343
	PkmnClaydol    = 344
	PkmnLileep     = 345
	PkmnCradily    = 346
	PkmnAnorith    = 347
	PkmnArmaldo    = 348
	PkmnFeebas     = 349
	PkmnMilotic    = 350
	PkmnCastform   = 351
	PkmnKecleon    = 352
	PkmnShuppet    = 353
	PkmnBanette    = 354
	PkmnDuskull    = 355
	PkmnDusclops   = 356
	PkmnTropius    = 357
	PkmnChimecho   = 358
	PkmnAbsol      = 359
	PkmnWynaut     = 360
	PkmnSnorunt    = 361
	PkmnGlalie     = 362
	PkmnSpheal     = 363
	PkmnSealeo     = 364
	PkmnWalrein    = 365
	PkmnClamperl   = 366
	PkmnHuntail    = 367
	PkmnGorebyss   = 368
	PkmnRelicanth  = 369
	PkmnLuvdisc    = 370
	PkmnBagon      = 371
	PkmnShelgon    = 372
	PkmnSalamence  = 373
	PkmnBeldum     = 374
	PkmnMetang     = 375
	PkmnMetagross  = 376
	PkmnRegirock   = 377
	PkmnRegice     = 378
	PkmnRegisteel  = 379
	PkmnLatias     = 380
	PkmnLatios     = 381
	PkmnKyogre     = 382
	PkmnGroudon    = 383
	PkmnRayquaza   = 384
	PkmnJirachi    = 385
	PkmnDeoxys     = 386
	PkmnTurtwig    = 387
	PkmnGrotle     = 388
	PkmnTorterra   = 389
	PkmnChimchar   = 390
	PkmnMonferno   = 391
	PkmnInfernape  = 392
	PkmnPiplup     = 393
	PkmnPrinplup   = 394
	PkmnEmpoleon   = 395
	PkmnStarly     = 396
	PkmnStaravia   = 397
	PkmnStaraptor  = 398
	PkmnBidoof     = 399
	PkmnBibarel    = 400
	PkmnKricketot  = 401
	PkmnKricketune = 402
	PkmnShinx      = 403
	PkmnLuxio      = 404
	PkmnLuxray     = 405
	PkmnBudew      = 406
	PkmnRoserade   = 407
	PkmnCranidos   = 408
	PkmnRampardos  = 409
	PkmnShieldon   = 410
	PkmnBastiodon  = 411
	PkmnBurmy      = 412
	PkmnWormadam   = 413
	PkmnMothim     = 414
	PkmnCombee     = 415
	PkmnVespiquen  = 416
	PkmnPachirisu  = 417
	PkmnBuizel     = 418
	PkmnFloatzel   = 419
	PkmnCherubi    = 420
	PkmnCherrim    = 421
	PkmnShellos    = 422
	PkmnGastrodon  = 423
	PkmnAmbipom    = 424
	PkmnDrifloon   = 425
	PkmnDrifblim   = 426
	PkmnBuneary    = 427
	PkmnLopunny    = 428
	PkmnMismagius  = 429
	PkmnHonchkrow  = 430
	PkmnGlameow    = 431
	PkmnPurugly    = 432
	PkmnChingling  = 433
	PkmnStunky     = 434
	PkmnSkuntank   = 435
	PkmnBronzor    = 436
	PkmnBronzong   = 437
	PkmnBonsly     = 438
	PkmnMimeJr     = 439
	PkmnHappiny    = 440
	PkmnChatot     = 441
	PkmnSpiritomb  = 442
	PkmnGible      = 443
	PkmnGabite     = 444
	PkmnGarchomp   = 445
	PkmnMunchlax   = 446
	PkmnRiolu      = 447
	PkmnLucario    = 448
	PkmnHippopotas = 449
	PkmnHippowdon  = 450
	PkmnSkorupi    = 451
	PkmnDrapion    = 452
	PkmnCroagunk   = 453
	PkmnToxicroak  = 454
	PkmnCarnivine  = 455
	PkmnFinneon    = 456
	PkmnLumineon   = 457
	PkmnMantyke    = 458
	PkmnSnover     = 459
	PkmnAbomasnow  = 460
	PkmnWeavile    = 461
	PkmnMagnezone  = 462
	PkmnLickilicky = 463
	PkmnRhyperior  = 464
	PkmnTangrowth  = 465
	PkmnElectivire = 466
	PkmnMagmortar  = 467
	PkmnTogekiss   = 468
	PkmnYanmega    = 469
	PkmnLeafeon    = 470
	PkmnGlaceon    = 471
	PkmnGliscor    = 472
	PkmnMamoswine  = 473
	PkmnPorygonZ   = 474
	PkmnGallade    = 475
	PkmnProbopass  = 476
	PkmnDusknoir   = 477
	PkmnFroslass   = 478
	PkmnRotom      = 479
	PkmnUxie       = 480
	PkmnMesprit    = 481
	PkmnAzelf      = 482
	PkmnDialga     = 483
	PkmnPalkia     = 484
	PkmnHeatran    = 485
	PkmnRegigigas  = 486
	PkmnGiratina   = 487
	PkmnCresselia  = 488
	PkmnPhione     = 489
	PkmnManaphy    = 490
	PkmnDarkrai    = 491
	PkmnShaymin    = 492
	PkmnArceus     = 493
)

var AllMoves = []Move{
	{ID: 1, Name: "Pound", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 35, MaxPP: 35, Targets: 10, Priority: 0, Power: 40, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 2, Name: "Karate Chop", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 25, MaxPP: 25, Targets: 10, Priority: 0, Power: 50, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 1, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 3, Name: "Double Slap", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 15, Accuracy: 85, metadata: MoveMeta{MinHits: 2, MaxHits: 5, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 4, Name: "Comet Punch", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 18, Accuracy: 85, metadata: MoveMeta{MinHits: 2, MaxHits: 5, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 5, Name: "Mega Punch", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 80, Accuracy: 85, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 6, Name: "Pay Day", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 40, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 7, Name: "Fire Punch", Type: 512, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 75, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 10, FlinchChance: 0, StatChance: 0}},
	{ID: 8, Name: "Ice Punch", Type: 16384, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 75, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 10, FlinchChance: 0, StatChance: 0}},
	{ID: 9, Name: "Thunder Punch", Type: 4096, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 75, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 10, FlinchChance: 0, StatChance: 0}},
	{ID: 10, Name: "Scratch", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 35, MaxPP: 35, Targets: 10, Priority: 0, Power: 40, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 11, Name: "Vice Grip", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 30, MaxPP: 30, Targets: 10, Priority: 0, Power: 55, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 12, Name: "Guillotine", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 0, Accuracy: 30, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 13, Name: "Razor Wind", Type: 1, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 11, Priority: 0, Power: 80, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 1, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 14, Name: "Swords Dance", Type: 1, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 15, Name: "Cut", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 30, MaxPP: 30, Targets: 10, Priority: 0, Power: 50, Accuracy: 95, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 16, Name: "Gust", Type: 4, Category: MoveCategorySpecial, CurrentPP: 35, MaxPP: 35, Targets: 10, Priority: 0, Power: 40, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 17, Name: "Wing Attack", Type: 4, Category: MoveCategoryPhysical, CurrentPP: 35, MaxPP: 35, Targets: 10, Priority: 0, Power: 60, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 18, Name: "Whirlwind", Type: 1, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: -6, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 19, Name: "Fly", Type: 4, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 90, Accuracy: 95, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 20, Name: "Bind", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 15, Accuracy: 85, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 5, MaxTurns: 6, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 100, FlinchChance: 0, StatChance: 0}},
	{ID: 21, Name: "Slam", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 80, Accuracy: 75, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 22, Name: "Vine Whip", Type: 2048, Category: MoveCategoryPhysical, CurrentPP: 25, MaxPP: 25, Targets: 10, Priority: 0, Power: 45, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 23, Name: "Stomp", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 65, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 30, StatChance: 0}},
	{ID: 24, Name: "Double Kick", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 30, MaxPP: 30, Targets: 10, Priority: 0, Power: 30, Accuracy: 100, metadata: MoveMeta{MinHits: 2, MaxHits: 2, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 25, Name: "Mega Kick", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 120, Accuracy: 75, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 26, Name: "Jump Kick", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 100, Accuracy: 95, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 27, Name: "Rolling Kick", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 60, Accuracy: 85, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 30, StatChance: 0}},
	{ID: 28, Name: "Sand Attack", Type: 16, Category: MoveCategoryStatus, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 29, Name: "Headbutt", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 70, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 30, StatChance: 0}},
	{ID: 30, Name: "Horn Attack", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 25, MaxPP: 25, Targets: 10, Priority: 0, Power: 65, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 31, Name: "Fury Attack", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 15, Accuracy: 85, metadata: MoveMeta{MinHits: 2, MaxHits: 5, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 32, Name: "Horn Drill", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 0, Accuracy: 30, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 33, Name: "Tackle", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 35, MaxPP: 35, Targets: 10, Priority: 0, Power: 40, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 34, Name: "Body Slam", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 85, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 30, FlinchChance: 0, StatChance: 0}},
	{ID: 35, Name: "Wrap", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 15, Accuracy: 90, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 5, MaxTurns: 6, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 100, FlinchChance: 0, StatChance: 0}},
	{ID: 36, Name: "Take Down", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 90, Accuracy: 85, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: -25, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 37, Name: "Thrash", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 8, Priority: 0, Power: 120, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 38, Name: "Double-Edge", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 120, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: -33, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 39, Name: "Tail Whip", Type: 1, Category: MoveCategoryStatus, CurrentPP: 30, MaxPP: 30, Targets: 11, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 40, Name: "Poison Sting", Type: 8, Category: MoveCategoryPhysical, CurrentPP: 35, MaxPP: 35, Targets: 10, Priority: 0, Power: 15, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 30, FlinchChance: 0, StatChance: 0}},
	{ID: 41, Name: "Twineedle", Type: 64, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 25, Accuracy: 100, metadata: MoveMeta{MinHits: 2, MaxHits: 2, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 20, FlinchChance: 0, StatChance: 0}},
	{ID: 42, Name: "Pin Missile", Type: 64, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 25, Accuracy: 95, metadata: MoveMeta{MinHits: 2, MaxHits: 5, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 43, Name: "Leer", Type: 1, Category: MoveCategoryStatus, CurrentPP: 30, MaxPP: 30, Targets: 11, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 100}},
	{ID: 44, Name: "Bite", Type: 65536, Category: MoveCategoryPhysical, CurrentPP: 25, MaxPP: 25, Targets: 10, Priority: 0, Power: 60, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 30, StatChance: 0}},
	{ID: 45, Name: "Growl", Type: 1, Category: MoveCategoryStatus, CurrentPP: 40, MaxPP: 40, Targets: 11, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 46, Name: "Roar", Type: 1, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: -6, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 47, Name: "Sing", Type: 1, Category: MoveCategoryStatus, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 0, Accuracy: 55, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 2, MaxTurns: 4, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 48, Name: "Supersonic", Type: 1, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 0, Accuracy: 55, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 2, MaxTurns: 5, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 49, Name: "Sonic Boom", Type: 1, Category: MoveCategorySpecial, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 0, Accuracy: 90, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 50, Name: "Disable", Type: 1, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 4, MaxTurns: 4, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 51, Name: "Acid", Type: 8, Category: MoveCategorySpecial, CurrentPP: 30, MaxPP: 30, Targets: 11, Priority: 0, Power: 40, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 10}},
	{ID: 52, Name: "Ember", Type: 512, Category: MoveCategorySpecial, CurrentPP: 25, MaxPP: 25, Targets: 10, Priority: 0, Power: 40, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 10, FlinchChance: 0, StatChance: 0}},
	{ID: 53, Name: "Flamethrower", Type: 512, Category: MoveCategorySpecial, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 90, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 10, FlinchChance: 0, StatChance: 0}},
	{ID: 54, Name: "Mist", Type: 16384, Category: MoveCategoryStatus, CurrentPP: 30, MaxPP: 30, Targets: 4, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 55, Name: "Water Gun", Type: 1024, Category: MoveCategorySpecial, CurrentPP: 25, MaxPP: 25, Targets: 10, Priority: 0, Power: 40, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 56, Name: "Hydro Pump", Type: 1024, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 110, Accuracy: 80, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 57, Name: "Surf", Type: 1024, Category: MoveCategorySpecial, CurrentPP: 15, MaxPP: 15, Targets: 9, Priority: 0, Power: 90, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 58, Name: "Ice Beam", Type: 16384, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 90, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 10, FlinchChance: 0, StatChance: 0}},
	{ID: 59, Name: "Blizzard", Type: 16384, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 11, Priority: 0, Power: 110, Accuracy: 70, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 10, FlinchChance: 0, StatChance: 0}},
	{ID: 60, Name: "Psybeam", Type: 8192, Category: MoveCategorySpecial, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 65, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 2, MaxTurns: 5, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 10, FlinchChance: 0, StatChance: 0}},
	{ID: 61, Name: "Bubble Beam", Type: 1024, Category: MoveCategorySpecial, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 65, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 10}},
	{ID: 62, Name: "Aurora Beam", Type: 16384, Category: MoveCategorySpecial, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 65, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 10}},
	{ID: 63, Name: "Hyper Beam", Type: 1, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 150, Accuracy: 90, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 64, Name: "Peck", Type: 4, Category: MoveCategoryPhysical, CurrentPP: 35, MaxPP: 35, Targets: 10, Priority: 0, Power: 35, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 65, Name: "Drill Peck", Type: 4, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 80, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 66, Name: "Submission", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 80, Accuracy: 80, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: -25, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 67, Name: "Low Kick", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 68, Name: "Counter", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 1, Priority: -5, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 69, Name: "Seismic Toss", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 70, Name: "Strength", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 80, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 71, Name: "Absorb", Type: 2048, Category: MoveCategorySpecial, CurrentPP: 25, MaxPP: 25, Targets: 10, Priority: 0, Power: 20, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 50, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 72, Name: "Mega Drain", Type: 2048, Category: MoveCategorySpecial, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 40, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 50, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 73, Name: "Leech Seed", Type: 2048, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 90, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 74, Name: "Growth", Type: 1, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 75, Name: "Razor Leaf", Type: 2048, Category: MoveCategoryPhysical, CurrentPP: 25, MaxPP: 25, Targets: 11, Priority: 0, Power: 55, Accuracy: 95, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 1, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 76, Name: "Solar Beam", Type: 2048, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 120, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 77, Name: "Poison Powder", Type: 8, Category: MoveCategoryStatus, CurrentPP: 35, MaxPP: 35, Targets: 10, Priority: 0, Power: 0, Accuracy: 75, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 78, Name: "Stun Spore", Type: 2048, Category: MoveCategoryStatus, CurrentPP: 30, MaxPP: 30, Targets: 10, Priority: 0, Power: 0, Accuracy: 75, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 79, Name: "Sleep Powder", Type: 2048, Category: MoveCategoryStatus, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 0, Accuracy: 75, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 2, MaxTurns: 4, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 80, Name: "Petal Dance", Type: 2048, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 8, Priority: 0, Power: 120, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 81, Name: "String Shot", Type: 64, Category: MoveCategoryStatus, CurrentPP: 40, MaxPP: 40, Targets: 11, Priority: 0, Power: 0, Accuracy: 95, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 82, Name: "Dragon Rage", Type: 32768, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 83, Name: "Fire Spin", Type: 512, Category: MoveCategorySpecial, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 35, Accuracy: 85, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 5, MaxTurns: 6, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 100, FlinchChance: 0, StatChance: 0}},
	{ID: 84, Name: "Thunder Shock", Type: 4096, Category: MoveCategorySpecial, CurrentPP: 30, MaxPP: 30, Targets: 10, Priority: 0, Power: 40, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 10, FlinchChance: 0, StatChance: 0}},
	{ID: 85, Name: "Thunderbolt", Type: 4096, Category: MoveCategorySpecial, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 90, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 10, FlinchChance: 0, StatChance: 0}},
	{ID: 86, Name: "Thunder Wave", Type: 4096, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 0, Accuracy: 90, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 87, Name: "Thunder", Type: 4096, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 110, Accuracy: 70, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 30, FlinchChance: 0, StatChance: 0}},
	{ID: 88, Name: "Rock Throw", Type: 32, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 50, Accuracy: 90, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 89, Name: "Earthquake", Type: 16, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 9, Priority: 0, Power: 100, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 90, Name: "Fissure", Type: 16, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 0, Accuracy: 30, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 91, Name: "Dig", Type: 16, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 80, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 92, Name: "Toxic", Type: 8, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 90, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 15, MaxTurns: 15, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 93, Name: "Confusion", Type: 8192, Category: MoveCategorySpecial, CurrentPP: 25, MaxPP: 25, Targets: 10, Priority: 0, Power: 50, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 2, MaxTurns: 5, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 10, FlinchChance: 0, StatChance: 0}},
	{ID: 94, Name: "Psychic", Type: 8192, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 90, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 10}},
	{ID: 95, Name: "Hypnosis", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 0, Accuracy: 60, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 2, MaxTurns: 4, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 96, Name: "Meditate", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 40, MaxPP: 40, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 97, Name: "Agility", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 30, MaxPP: 30, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 98, Name: "Quick Attack", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 30, MaxPP: 30, Targets: 10, Priority: 1, Power: 40, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 99, Name: "Rage", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 20, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 100, Name: "Teleport", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 101, Name: "Night Shade", Type: 128, Category: MoveCategorySpecial, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 102, Name: "Mimic", Type: 1, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 103, Name: "Screech", Type: 1, Category: MoveCategoryStatus, CurrentPP: 40, MaxPP: 40, Targets: 10, Priority: 0, Power: 0, Accuracy: 85, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 104, Name: "Double Team", Type: 1, Category: MoveCategoryStatus, CurrentPP: 15, MaxPP: 15, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 105, Name: "Recover", Type: 1, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 50, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 106, Name: "Harden", Type: 1, Category: MoveCategoryStatus, CurrentPP: 30, MaxPP: 30, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 107, Name: "Minimize", Type: 1, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 108, Name: "Smokescreen", Type: 1, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 109, Name: "Confuse Ray", Type: 128, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 2, MaxTurns: 5, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 110, Name: "Withdraw", Type: 1024, Category: MoveCategoryStatus, CurrentPP: 40, MaxPP: 40, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 111, Name: "Defense Curl", Type: 1, Category: MoveCategoryStatus, CurrentPP: 40, MaxPP: 40, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 112, Name: "Barrier", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 113, Name: "Light Screen", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 30, MaxPP: 30, Targets: 4, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 114, Name: "Haze", Type: 16384, Category: MoveCategoryStatus, CurrentPP: 30, MaxPP: 30, Targets: 12, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 115, Name: "Reflect", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 4, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 116, Name: "Focus Energy", Type: 1, Category: MoveCategoryStatus, CurrentPP: 30, MaxPP: 30, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 117, Name: "Bide", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 1, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 118, Name: "Metronome", Type: 1, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 119, Name: "Mirror Move", Type: 4, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 120, Name: "Self-Destruct", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 9, Priority: 0, Power: 200, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 121, Name: "Egg Bomb", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 100, Accuracy: 75, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 122, Name: "Lick", Type: 128, Category: MoveCategoryPhysical, CurrentPP: 30, MaxPP: 30, Targets: 10, Priority: 0, Power: 30, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 30, FlinchChance: 0, StatChance: 0}},
	{ID: 123, Name: "Smog", Type: 8, Category: MoveCategorySpecial, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 30, Accuracy: 70, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 40, FlinchChance: 0, StatChance: 0}},
	{ID: 124, Name: "Sludge", Type: 8, Category: MoveCategorySpecial, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 65, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 30, FlinchChance: 0, StatChance: 0}},
	{ID: 125, Name: "Bone Club", Type: 16, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 65, Accuracy: 85, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 10, StatChance: 0}},
	{ID: 126, Name: "Fire Blast", Type: 512, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 110, Accuracy: 85, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 10, FlinchChance: 0, StatChance: 0}},
	{ID: 127, Name: "Waterfall", Type: 1024, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 80, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 20, StatChance: 0}},
	{ID: 128, Name: "Clamp", Type: 1024, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 35, Accuracy: 85, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 5, MaxTurns: 6, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 100, FlinchChance: 0, StatChance: 0}},
	{ID: 129, Name: "Swift", Type: 1, Category: MoveCategorySpecial, CurrentPP: 20, MaxPP: 20, Targets: 11, Priority: 0, Power: 60, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 130, Name: "Skull Bash", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 130, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 100, FlinchChance: 0, StatChance: 0}},
	{ID: 131, Name: "Spike Cannon", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 20, Accuracy: 100, metadata: MoveMeta{MinHits: 2, MaxHits: 5, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 132, Name: "Constrict", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 35, MaxPP: 35, Targets: 10, Priority: 0, Power: 10, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 10}},
	{ID: 133, Name: "Amnesia", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 134, Name: "Kinesis", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 0, Accuracy: 80, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 135, Name: "Soft-Boiled", Type: 1, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 50, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 136, Name: "High Jump Kick", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 130, Accuracy: 90, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 137, Name: "Glare", Type: 1, Category: MoveCategoryStatus, CurrentPP: 30, MaxPP: 30, Targets: 10, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 138, Name: "Dream Eater", Type: 8192, Category: MoveCategorySpecial, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 100, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 50, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 139, Name: "Poison Gas", Type: 8, Category: MoveCategoryStatus, CurrentPP: 40, MaxPP: 40, Targets: 11, Priority: 0, Power: 0, Accuracy: 90, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 140, Name: "Barrage", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 15, Accuracy: 85, metadata: MoveMeta{MinHits: 2, MaxHits: 5, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 141, Name: "Leech Life", Type: 64, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 80, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 50, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 142, Name: "Lovely Kiss", Type: 1, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 75, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 2, MaxTurns: 4, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 143, Name: "Sky Attack", Type: 4, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 140, Accuracy: 90, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 1, AilmentChance: 0, FlinchChance: 30, StatChance: 0}},
	{ID: 144, Name: "Transform", Type: 1, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 145, Name: "Bubble", Type: 1024, Category: MoveCategorySpecial, CurrentPP: 30, MaxPP: 30, Targets: 11, Priority: 0, Power: 40, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 10}},
	{ID: 146, Name: "Dizzy Punch", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 70, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 2, MaxTurns: 5, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 20, FlinchChance: 0, StatChance: 0}},
	{ID: 147, Name: "Spore", Type: 2048, Category: MoveCategoryStatus, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 2, MaxTurns: 4, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 148, Name: "Flash", Type: 1, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 149, Name: "Psywave", Type: 8192, Category: MoveCategorySpecial, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 150, Name: "Splash", Type: 1, Category: MoveCategoryStatus, CurrentPP: 40, MaxPP: 40, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 151, Name: "Acid Armor", Type: 8, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 152, Name: "Crabhammer", Type: 1024, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 100, Accuracy: 90, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 1, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 153, Name: "Explosion", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 9, Priority: 0, Power: 250, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 154, Name: "Fury Swipes", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 18, Accuracy: 80, metadata: MoveMeta{MinHits: 2, MaxHits: 5, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 155, Name: "Bonemerang", Type: 16, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 50, Accuracy: 90, metadata: MoveMeta{MinHits: 2, MaxHits: 2, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 156, Name: "Rest", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 157, Name: "Rock Slide", Type: 32, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 11, Priority: 0, Power: 75, Accuracy: 90, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 30, StatChance: 0}},
	{ID: 158, Name: "Hyper Fang", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 80, Accuracy: 90, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 10, StatChance: 0}},
	{ID: 159, Name: "Sharpen", Type: 1, Category: MoveCategoryStatus, CurrentPP: 30, MaxPP: 30, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 160, Name: "Conversion", Type: 1, Category: MoveCategoryStatus, CurrentPP: 30, MaxPP: 30, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 161, Name: "Tri Attack", Type: 1, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 80, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 20, FlinchChance: 0, StatChance: 0}},
	{ID: 162, Name: "Super Fang", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 90, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 163, Name: "Slash", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 70, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 1, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 164, Name: "Substitute", Type: 1, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 165, Name: "Struggle", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 1, MaxPP: 1, Targets: 8, Priority: 0, Power: 50, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: -25, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 166, Name: "Sketch", Type: 1, Category: MoveCategoryStatus, CurrentPP: 1, MaxPP: 1, Targets: 10, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 167, Name: "Triple Kick", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 10, Accuracy: 90, metadata: MoveMeta{MinHits: 3, MaxHits: 3, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 168, Name: "Thief", Type: 65536, Category: MoveCategoryPhysical, CurrentPP: 25, MaxPP: 25, Targets: 10, Priority: 0, Power: 60, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 169, Name: "Spider Web", Type: 64, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 170, Name: "Mind Reader", Type: 1, Category: MoveCategoryStatus, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 171, Name: "Nightmare", Type: 128, Category: MoveCategoryStatus, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 172, Name: "Flame Wheel", Type: 512, Category: MoveCategoryPhysical, CurrentPP: 25, MaxPP: 25, Targets: 10, Priority: 0, Power: 60, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 10, FlinchChance: 0, StatChance: 0}},
	{ID: 173, Name: "Snore", Type: 1, Category: MoveCategorySpecial, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 50, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 30, StatChance: 0}},
	{ID: 174, Name: "Curse", Type: 128, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 1, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 175, Name: "Flail", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 176, Name: "Conversion 2", Type: 1, Category: MoveCategoryStatus, CurrentPP: 30, MaxPP: 30, Targets: 10, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 177, Name: "Aeroblast", Type: 4, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 100, Accuracy: 95, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 1, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 178, Name: "Cotton Spore", Type: 2048, Category: MoveCategoryStatus, CurrentPP: 40, MaxPP: 40, Targets: 11, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 179, Name: "Reversal", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 180, Name: "Spite", Type: 128, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 181, Name: "Powder Snow", Type: 16384, Category: MoveCategorySpecial, CurrentPP: 25, MaxPP: 25, Targets: 11, Priority: 0, Power: 40, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 10, FlinchChance: 0, StatChance: 0}},
	{ID: 182, Name: "Protect", Type: 1, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 4, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 183, Name: "Mach Punch", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 30, MaxPP: 30, Targets: 10, Priority: 1, Power: 40, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 184, Name: "Scary Face", Type: 1, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 185, Name: "Feint Attack", Type: 65536, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 60, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 186, Name: "Sweet Kiss", Type: 131072, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 75, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 2, MaxTurns: 5, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 187, Name: "Belly Drum", Type: 1, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 188, Name: "Sludge Bomb", Type: 8, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 90, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 30, FlinchChance: 0, StatChance: 0}},
	{ID: 189, Name: "Mud-Slap", Type: 16, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 20, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 100}},
	{ID: 190, Name: "Octazooka", Type: 1024, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 65, Accuracy: 85, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 50}},
	{ID: 191, Name: "Spikes", Type: 16, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 6, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 192, Name: "Zap Cannon", Type: 4096, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 120, Accuracy: 50, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 100, FlinchChance: 0, StatChance: 0}},
	{ID: 193, Name: "Foresight", Type: 1, Category: MoveCategoryStatus, CurrentPP: 40, MaxPP: 40, Targets: 10, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 194, Name: "Destiny Bond", Type: 128, Category: MoveCategoryStatus, CurrentPP: 5, MaxPP: 5, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 195, Name: "Perish Song", Type: 1, Category: MoveCategoryStatus, CurrentPP: 5, MaxPP: 5, Targets: 14, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 4, MaxTurns: 4, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 196, Name: "Icy Wind", Type: 16384, Category: MoveCategorySpecial, CurrentPP: 15, MaxPP: 15, Targets: 11, Priority: 0, Power: 55, Accuracy: 95, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 100}},
	{ID: 197, Name: "Detect", Type: 2, Category: MoveCategoryStatus, CurrentPP: 5, MaxPP: 5, Targets: 7, Priority: 4, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 198, Name: "Bone Rush", Type: 16, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 25, Accuracy: 90, metadata: MoveMeta{MinHits: 2, MaxHits: 5, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 199, Name: "Lock-On", Type: 1, Category: MoveCategoryStatus, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 200, Name: "Outrage", Type: 32768, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 8, Priority: 0, Power: 120, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 201, Name: "Sandstorm", Type: 32, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 12, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 202, Name: "Giga Drain", Type: 2048, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 75, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 50, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 203, Name: "Endure", Type: 1, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 4, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 204, Name: "Charm", Type: 131072, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 205, Name: "Rollout", Type: 32, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 30, Accuracy: 90, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 206, Name: "False Swipe", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 40, MaxPP: 40, Targets: 10, Priority: 0, Power: 40, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 207, Name: "Swagger", Type: 1, Category: MoveCategoryStatus, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 0, Accuracy: 85, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 2, MaxTurns: 5, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 208, Name: "Milk Drink", Type: 1, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 50, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 209, Name: "Spark", Type: 4096, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 65, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 30, FlinchChance: 0, StatChance: 0}},
	{ID: 210, Name: "Fury Cutter", Type: 64, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 40, Accuracy: 95, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 211, Name: "Steel Wing", Type: 256, Category: MoveCategoryPhysical, CurrentPP: 25, MaxPP: 25, Targets: 10, Priority: 0, Power: 70, Accuracy: 90, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 10}},
	{ID: 212, Name: "Mean Look", Type: 1, Category: MoveCategoryStatus, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 213, Name: "Attract", Type: 1, Category: MoveCategoryStatus, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 214, Name: "Sleep Talk", Type: 1, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 215, Name: "Heal Bell", Type: 1, Category: MoveCategoryStatus, CurrentPP: 5, MaxPP: 5, Targets: 13, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 216, Name: "Return", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 217, Name: "Present", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 0, Accuracy: 90, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 218, Name: "Frustration", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 219, Name: "Safeguard", Type: 1, Category: MoveCategoryStatus, CurrentPP: 25, MaxPP: 25, Targets: 4, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 220, Name: "Pain Split", Type: 1, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 221, Name: "Sacred Fire", Type: 512, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 100, Accuracy: 95, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 50, FlinchChance: 0, StatChance: 0}},
	{ID: 222, Name: "Magnitude", Type: 16, Category: MoveCategoryPhysical, CurrentPP: 30, MaxPP: 30, Targets: 9, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 223, Name: "Dynamic Punch", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 100, Accuracy: 50, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 2, MaxTurns: 5, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 100, FlinchChance: 0, StatChance: 0}},
	{ID: 224, Name: "Megahorn", Type: 64, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 120, Accuracy: 85, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 225, Name: "Dragon Breath", Type: 32768, Category: MoveCategorySpecial, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 60, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 30, FlinchChance: 0, StatChance: 0}},
	{ID: 226, Name: "Baton Pass", Type: 1, Category: MoveCategoryStatus, CurrentPP: 40, MaxPP: 40, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 227, Name: "Encore", Type: 1, Category: MoveCategoryStatus, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 228, Name: "Pursuit", Type: 65536, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 40, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 229, Name: "Rapid Spin", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 40, MaxPP: 40, Targets: 10, Priority: 0, Power: 20, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 230, Name: "Sweet Scent", Type: 1, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 11, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 231, Name: "Iron Tail", Type: 256, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 100, Accuracy: 75, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 30}},
	{ID: 232, Name: "Metal Claw", Type: 256, Category: MoveCategoryPhysical, CurrentPP: 35, MaxPP: 35, Targets: 10, Priority: 0, Power: 50, Accuracy: 95, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 10}},
	{ID: 233, Name: "Vital Throw", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: -1, Power: 70, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 234, Name: "Morning Sun", Type: 1, Category: MoveCategoryStatus, CurrentPP: 5, MaxPP: 5, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 50, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 235, Name: "Synthesis", Type: 2048, Category: MoveCategoryStatus, CurrentPP: 5, MaxPP: 5, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 50, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 236, Name: "Moonlight", Type: 131072, Category: MoveCategoryStatus, CurrentPP: 5, MaxPP: 5, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 50, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 237, Name: "Hidden Power", Type: 1, Category: MoveCategorySpecial, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 60, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 238, Name: "Cross Chop", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 100, Accuracy: 80, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 1, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 239, Name: "Twister", Type: 32768, Category: MoveCategorySpecial, CurrentPP: 20, MaxPP: 20, Targets: 11, Priority: 0, Power: 40, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 20, StatChance: 0}},
	{ID: 240, Name: "Rain Dance", Type: 1024, Category: MoveCategoryStatus, CurrentPP: 5, MaxPP: 5, Targets: 12, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 241, Name: "Sunny Day", Type: 512, Category: MoveCategoryStatus, CurrentPP: 5, MaxPP: 5, Targets: 12, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 242, Name: "Crunch", Type: 65536, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 80, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 20}},
	{ID: 243, Name: "Mirror Coat", Type: 8192, Category: MoveCategorySpecial, CurrentPP: 20, MaxPP: 20, Targets: 1, Priority: -5, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 244, Name: "Psych Up", Type: 1, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 245, Name: "Extreme Speed", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 2, Power: 80, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 246, Name: "Ancient Power", Type: 32, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 60, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 10}},
	{ID: 247, Name: "Shadow Ball", Type: 128, Category: MoveCategorySpecial, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 80, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 20}},
	{ID: 248, Name: "Future Sight", Type: 8192, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 120, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 249, Name: "Rock Smash", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 40, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 50}},
	{ID: 250, Name: "Whirlpool", Type: 1024, Category: MoveCategorySpecial, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 35, Accuracy: 85, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 5, MaxTurns: 6, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 100, FlinchChance: 0, StatChance: 0}},
	{ID: 251, Name: "Beat Up", Type: 65536, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 6, MaxHits: 6, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 252, Name: "Fake Out", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 3, Power: 40, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 100, StatChance: 0}},
	{ID: 253, Name: "Uproar", Type: 1, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 8, Priority: 0, Power: 90, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 254, Name: "Stockpile", Type: 1, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 255, Name: "Spit Up", Type: 1, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 256, Name: "Swallow", Type: 1, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 25, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 257, Name: "Heat Wave", Type: 512, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 11, Priority: 0, Power: 95, Accuracy: 90, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 10, FlinchChance: 0, StatChance: 0}},
	{ID: 258, Name: "Hail", Type: 16384, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 12, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 259, Name: "Torment", Type: 65536, Category: MoveCategoryStatus, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 260, Name: "Flatter", Type: 65536, Category: MoveCategoryStatus, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 2, MaxTurns: 5, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 261, Name: "Will-O-Wisp", Type: 512, Category: MoveCategoryStatus, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 0, Accuracy: 85, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 262, Name: "Memento", Type: 65536, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 263, Name: "Facade", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 70, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 264, Name: "Focus Punch", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: -3, Power: 150, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 265, Name: "Smelling Salts", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 70, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 266, Name: "Follow Me", Type: 1, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 7, Priority: 2, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 267, Name: "Nature Power", Type: 1, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 268, Name: "Charge", Type: 4096, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 269, Name: "Taunt", Type: 65536, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 270, Name: "Helping Hand", Type: 1, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 3, Priority: 5, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 271, Name: "Trick", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 272, Name: "Role Play", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 273, Name: "Wish", Type: 1, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 274, Name: "Assist", Type: 1, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 275, Name: "Ingrain", Type: 2048, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 276, Name: "Superpower", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 120, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 100}},
	{ID: 277, Name: "Magic Coat", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 15, MaxPP: 15, Targets: 7, Priority: 4, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 278, Name: "Recycle", Type: 1, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 279, Name: "Revenge", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: -4, Power: 60, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 280, Name: "Brick Break", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 75, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 281, Name: "Yawn", Type: 1, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 2, MaxTurns: 2, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 282, Name: "Knock Off", Type: 65536, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 65, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 283, Name: "Endeavor", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 284, Name: "Eruption", Type: 512, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 11, Priority: 0, Power: 150, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 285, Name: "Skill Swap", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 286, Name: "Imprison", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 287, Name: "Refresh", Type: 1, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 288, Name: "Grudge", Type: 128, Category: MoveCategoryStatus, CurrentPP: 5, MaxPP: 5, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 289, Name: "Snatch", Type: 65536, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 4, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 290, Name: "Secret Power", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 70, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 30, FlinchChance: 0, StatChance: 0}},
	{ID: 291, Name: "Dive", Type: 1024, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 80, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 292, Name: "Arm Thrust", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 15, Accuracy: 100, metadata: MoveMeta{MinHits: 2, MaxHits: 5, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 293, Name: "Camouflage", Type: 1, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 294, Name: "Tail Glow", Type: 64, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 295, Name: "Luster Purge", Type: 8192, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 70, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 50}},
	{ID: 296, Name: "Mist Ball", Type: 8192, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 70, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 50}},
	{ID: 297, Name: "Feather Dance", Type: 4, Category: MoveCategoryStatus, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 298, Name: "Teeter Dance", Type: 1, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 9, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 2, MaxTurns: 5, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 299, Name: "Blaze Kick", Type: 512, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 85, Accuracy: 90, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 1, AilmentChance: 10, FlinchChance: 0, StatChance: 0}},
	{ID: 300, Name: "Mud Sport", Type: 16, Category: MoveCategoryStatus, CurrentPP: 15, MaxPP: 15, Targets: 12, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 301, Name: "Ice Ball", Type: 16384, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 30, Accuracy: 90, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 302, Name: "Needle Arm", Type: 2048, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 60, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 30, StatChance: 0}},
	{ID: 303, Name: "Slack Off", Type: 1, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 50, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 304, Name: "Hyper Voice", Type: 1, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 11, Priority: 0, Power: 90, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 305, Name: "Poison Fang", Type: 8, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 50, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 15, MaxTurns: 15, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 50, FlinchChance: 0, StatChance: 0}},
	{ID: 306, Name: "Crush Claw", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 75, Accuracy: 95, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 50}},
	{ID: 307, Name: "Blast Burn", Type: 512, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 150, Accuracy: 90, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 308, Name: "Hydro Cannon", Type: 1024, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 150, Accuracy: 90, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 309, Name: "Meteor Mash", Type: 256, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 90, Accuracy: 90, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 20}},
	{ID: 310, Name: "Astonish", Type: 128, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 30, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 30, StatChance: 0}},
	{ID: 311, Name: "Weather Ball", Type: 1, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 50, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 312, Name: "Aromatherapy", Type: 2048, Category: MoveCategoryStatus, CurrentPP: 5, MaxPP: 5, Targets: 13, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 313, Name: "Fake Tears", Type: 65536, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 314, Name: "Air Cutter", Type: 4, Category: MoveCategorySpecial, CurrentPP: 25, MaxPP: 25, Targets: 11, Priority: 0, Power: 60, Accuracy: 95, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 1, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 315, Name: "Overheat", Type: 512, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 130, Accuracy: 90, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 100}},
	{ID: 316, Name: "Odor Sleuth", Type: 1, Category: MoveCategoryStatus, CurrentPP: 40, MaxPP: 40, Targets: 10, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 317, Name: "Rock Tomb", Type: 32, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 60, Accuracy: 95, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 100}},
	{ID: 318, Name: "Silver Wind", Type: 64, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 60, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 10}},
	{ID: 319, Name: "Metal Sound", Type: 256, Category: MoveCategoryStatus, CurrentPP: 40, MaxPP: 40, Targets: 10, Priority: 0, Power: 0, Accuracy: 85, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 320, Name: "Grass Whistle", Type: 2048, Category: MoveCategoryStatus, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 0, Accuracy: 55, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 2, MaxTurns: 4, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 321, Name: "Tickle", Type: 1, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 322, Name: "Cosmic Power", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 323, Name: "Water Spout", Type: 1024, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 11, Priority: 0, Power: 150, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 324, Name: "Signal Beam", Type: 64, Category: MoveCategorySpecial, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 75, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 2, MaxTurns: 5, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 10, FlinchChance: 0, StatChance: 0}},
	{ID: 325, Name: "Shadow Punch", Type: 128, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 60, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 326, Name: "Extrasensory", Type: 8192, Category: MoveCategorySpecial, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 80, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 10, StatChance: 0}},
	{ID: 327, Name: "Sky Uppercut", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 85, Accuracy: 90, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 328, Name: "Sand Tomb", Type: 16, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 35, Accuracy: 85, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 5, MaxTurns: 6, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 100, FlinchChance: 0, StatChance: 0}},
	{ID: 329, Name: "Sheer Cold", Type: 16384, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 0, Accuracy: 30, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 330, Name: "Muddy Water", Type: 1024, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 11, Priority: 0, Power: 90, Accuracy: 85, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 30}},
	{ID: 331, Name: "Bullet Seed", Type: 2048, Category: MoveCategoryPhysical, CurrentPP: 30, MaxPP: 30, Targets: 10, Priority: 0, Power: 25, Accuracy: 100, metadata: MoveMeta{MinHits: 2, MaxHits: 5, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 332, Name: "Aerial Ace", Type: 4, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 60, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 333, Name: "Icicle Spear", Type: 16384, Category: MoveCategoryPhysical, CurrentPP: 30, MaxPP: 30, Targets: 10, Priority: 0, Power: 25, Accuracy: 100, metadata: MoveMeta{MinHits: 2, MaxHits: 5, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 334, Name: "Iron Defense", Type: 256, Category: MoveCategoryStatus, CurrentPP: 15, MaxPP: 15, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 335, Name: "Block", Type: 1, Category: MoveCategoryStatus, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 336, Name: "Howl", Type: 1, Category: MoveCategoryStatus, CurrentPP: 40, MaxPP: 40, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 337, Name: "Dragon Claw", Type: 32768, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 80, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 338, Name: "Frenzy Plant", Type: 2048, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 150, Accuracy: 90, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 339, Name: "Bulk Up", Type: 2, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 340, Name: "Bounce", Type: 4, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 85, Accuracy: 85, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 30, FlinchChance: 0, StatChance: 0}},
	{ID: 341, Name: "Mud Shot", Type: 16, Category: MoveCategorySpecial, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 55, Accuracy: 95, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 100}},
	{ID: 342, Name: "Poison Tail", Type: 8, Category: MoveCategoryPhysical, CurrentPP: 25, MaxPP: 25, Targets: 10, Priority: 0, Power: 50, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 1, AilmentChance: 10, FlinchChance: 0, StatChance: 0}},
	{ID: 343, Name: "Covet", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 25, MaxPP: 25, Targets: 10, Priority: 0, Power: 60, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 344, Name: "Volt Tackle", Type: 4096, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 120, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: -33, Healing: 0, CritRate: 0, AilmentChance: 10, FlinchChance: 0, StatChance: 0}},
	{ID: 345, Name: "Magical Leaf", Type: 2048, Category: MoveCategorySpecial, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 60, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 346, Name: "Water Sport", Type: 1024, Category: MoveCategoryStatus, CurrentPP: 15, MaxPP: 15, Targets: 12, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 347, Name: "Calm Mind", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 348, Name: "Leaf Blade", Type: 2048, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 90, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 1, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 349, Name: "Dragon Dance", Type: 32768, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 350, Name: "Rock Blast", Type: 32, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 25, Accuracy: 90, metadata: MoveMeta{MinHits: 2, MaxHits: 5, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 351, Name: "Shock Wave", Type: 4096, Category: MoveCategorySpecial, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 60, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 352, Name: "Water Pulse", Type: 1024, Category: MoveCategorySpecial, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 60, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 2, MaxTurns: 5, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 20, FlinchChance: 0, StatChance: 0}},
	{ID: 353, Name: "Doom Desire", Type: 256, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 140, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 354, Name: "Psycho Boost", Type: 8192, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 140, Accuracy: 90, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 100}},
	{ID: 355, Name: "Roost", Type: 4, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 50, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 356, Name: "Gravity", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 5, MaxPP: 5, Targets: 12, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 357, Name: "Miracle Eye", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 40, MaxPP: 40, Targets: 10, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 358, Name: "Wake-Up Slap", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 70, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 359, Name: "Hammer Arm", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 100, Accuracy: 90, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 100}},
	{ID: 360, Name: "Gyro Ball", Type: 256, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 361, Name: "Healing Wish", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 362, Name: "Brine", Type: 1024, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 65, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 363, Name: "Natural Gift", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 364, Name: "Feint", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 2, Power: 30, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 365, Name: "Pluck", Type: 4, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 60, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 366, Name: "Tailwind", Type: 4, Category: MoveCategoryStatus, CurrentPP: 15, MaxPP: 15, Targets: 4, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 367, Name: "Acupressure", Type: 1, Category: MoveCategoryStatus, CurrentPP: 30, MaxPP: 30, Targets: 5, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 368, Name: "Metal Burst", Type: 256, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 1, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 369, Name: "U-turn", Type: 64, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 70, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 370, Name: "Close Combat", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 120, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 100}},
	{ID: 371, Name: "Payback", Type: 65536, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 50, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 372, Name: "Assurance", Type: 65536, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 60, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 373, Name: "Embargo", Type: 65536, Category: MoveCategoryStatus, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 5, MaxTurns: 5, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 374, Name: "Fling", Type: 65536, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 375, Name: "Psycho Shift", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 376, Name: "Trump Card", Type: 1, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 377, Name: "Heal Block", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 15, MaxPP: 15, Targets: 11, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 5, MaxTurns: 5, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 378, Name: "Wring Out", Type: 1, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 379, Name: "Power Trick", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 380, Name: "Gastro Acid", Type: 8, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 381, Name: "Lucky Chant", Type: 1, Category: MoveCategoryStatus, CurrentPP: 30, MaxPP: 30, Targets: 4, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 382, Name: "Me First", Type: 1, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 2, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 383, Name: "Copycat", Type: 1, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 384, Name: "Power Swap", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 385, Name: "Guard Swap", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 386, Name: "Punishment", Type: 65536, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 387, Name: "Last Resort", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 140, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 388, Name: "Worry Seed", Type: 2048, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 389, Name: "Sucker Punch", Type: 65536, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 1, Power: 70, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 390, Name: "Toxic Spikes", Type: 8, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 6, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 391, Name: "Heart Swap", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 392, Name: "Aqua Ring", Type: 1024, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 393, Name: "Magnet Rise", Type: 4096, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 5, MaxTurns: 5, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 394, Name: "Flare Blitz", Type: 512, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 120, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: -33, Healing: 0, CritRate: 0, AilmentChance: 10, FlinchChance: 0, StatChance: 0}},
	{ID: 395, Name: "Force Palm", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 60, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 30, FlinchChance: 0, StatChance: 0}},
	{ID: 396, Name: "Aura Sphere", Type: 2, Category: MoveCategorySpecial, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 80, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 397, Name: "Rock Polish", Type: 32, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 398, Name: "Poison Jab", Type: 8, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 80, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 30, FlinchChance: 0, StatChance: 0}},
	{ID: 399, Name: "Dark Pulse", Type: 65536, Category: MoveCategorySpecial, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 80, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 20, StatChance: 0}},
	{ID: 400, Name: "Night Slash", Type: 65536, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 70, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 1, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 401, Name: "Aqua Tail", Type: 1024, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 90, Accuracy: 90, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 402, Name: "Seed Bomb", Type: 2048, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 80, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 403, Name: "Air Slash", Type: 4, Category: MoveCategorySpecial, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 75, Accuracy: 95, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 30, StatChance: 0}},
	{ID: 404, Name: "X-Scissor", Type: 64, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 80, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 405, Name: "Bug Buzz", Type: 64, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 90, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 10}},
	{ID: 406, Name: "Dragon Pulse", Type: 32768, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 85, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 407, Name: "Dragon Rush", Type: 32768, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 100, Accuracy: 75, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 20, StatChance: 0}},
	{ID: 408, Name: "Power Gem", Type: 32, Category: MoveCategorySpecial, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 80, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 409, Name: "Drain Punch", Type: 2, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 75, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 50, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 410, Name: "Vacuum Wave", Type: 2, Category: MoveCategorySpecial, CurrentPP: 30, MaxPP: 30, Targets: 10, Priority: 1, Power: 40, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 411, Name: "Focus Blast", Type: 2, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 120, Accuracy: 70, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 10}},
	{ID: 412, Name: "Energy Ball", Type: 2048, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 90, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 10}},
	{ID: 413, Name: "Brave Bird", Type: 4, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 120, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: -33, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 414, Name: "Earth Power", Type: 16, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 90, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 10}},
	{ID: 415, Name: "Switcheroo", Type: 65536, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 416, Name: "Giga Impact", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 150, Accuracy: 90, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 417, Name: "Nasty Plot", Type: 65536, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 418, Name: "Bullet Punch", Type: 256, Category: MoveCategoryPhysical, CurrentPP: 30, MaxPP: 30, Targets: 10, Priority: 1, Power: 40, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 419, Name: "Avalanche", Type: 16384, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: -4, Power: 60, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 420, Name: "Ice Shard", Type: 16384, Category: MoveCategoryPhysical, CurrentPP: 30, MaxPP: 30, Targets: 10, Priority: 1, Power: 40, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 421, Name: "Shadow Claw", Type: 128, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 70, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 1, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 422, Name: "Thunder Fang", Type: 4096, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 65, Accuracy: 95, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 10, FlinchChance: 10, StatChance: 0}},
	{ID: 423, Name: "Ice Fang", Type: 16384, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 65, Accuracy: 95, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 10, FlinchChance: 10, StatChance: 0}},
	{ID: 424, Name: "Fire Fang", Type: 512, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 65, Accuracy: 95, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 10, FlinchChance: 10, StatChance: 0}},
	{ID: 425, Name: "Shadow Sneak", Type: 128, Category: MoveCategoryPhysical, CurrentPP: 30, MaxPP: 30, Targets: 10, Priority: 1, Power: 40, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 426, Name: "Mud Bomb", Type: 16, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 65, Accuracy: 85, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 30}},
	{ID: 427, Name: "Psycho Cut", Type: 8192, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 70, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 1, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 428, Name: "Zen Headbutt", Type: 8192, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 80, Accuracy: 90, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 20, StatChance: 0}},
	{ID: 429, Name: "Mirror Shot", Type: 256, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 65, Accuracy: 85, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 30}},
	{ID: 430, Name: "Flash Cannon", Type: 256, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 80, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 10}},
	{ID: 431, Name: "Rock Climb", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 90, Accuracy: 85, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 2, MaxTurns: 5, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 20, FlinchChance: 0, StatChance: 0}},
	{ID: 432, Name: "Defog", Type: 4, Category: MoveCategoryStatus, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 433, Name: "Trick Room", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 5, MaxPP: 5, Targets: 12, Priority: -7, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 434, Name: "Draco Meteor", Type: 32768, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 130, Accuracy: 90, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 100}},
	{ID: 435, Name: "Discharge", Type: 4096, Category: MoveCategorySpecial, CurrentPP: 15, MaxPP: 15, Targets: 9, Priority: 0, Power: 80, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 30, FlinchChance: 0, StatChance: 0}},
	{ID: 436, Name: "Lava Plume", Type: 512, Category: MoveCategorySpecial, CurrentPP: 15, MaxPP: 15, Targets: 9, Priority: 0, Power: 80, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 30, FlinchChance: 0, StatChance: 0}},
	{ID: 437, Name: "Leaf Storm", Type: 2048, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 130, Accuracy: 90, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 100}},
	{ID: 438, Name: "Power Whip", Type: 2048, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 120, Accuracy: 85, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 439, Name: "Rock Wrecker", Type: 32, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 150, Accuracy: 90, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 440, Name: "Cross Poison", Type: 8, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 70, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 1, AilmentChance: 10, FlinchChance: 0, StatChance: 0}},
	{ID: 441, Name: "Gunk Shot", Type: 8, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 120, Accuracy: 80, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 30, FlinchChance: 0, StatChance: 0}},
	{ID: 442, Name: "Iron Head", Type: 256, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 80, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 30, StatChance: 0}},
	{ID: 443, Name: "Magnet Bomb", Type: 256, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 60, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 444, Name: "Stone Edge", Type: 32, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 100, Accuracy: 80, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 1, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 445, Name: "Captivate", Type: 1, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 11, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 446, Name: "Stealth Rock", Type: 32, Category: MoveCategoryStatus, CurrentPP: 20, MaxPP: 20, Targets: 6, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 447, Name: "Grass Knot", Type: 2048, Category: MoveCategorySpecial, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 448, Name: "Chatter", Type: 4, Category: MoveCategorySpecial, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 65, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 2, MaxTurns: 5, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 100, FlinchChance: 0, StatChance: 0}},
	{ID: 449, Name: "Judgment", Type: 1, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 100, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 450, Name: "Bug Bite", Type: 64, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 0, Power: 60, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 451, Name: "Charge Beam", Type: 4096, Category: MoveCategorySpecial, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 50, Accuracy: 90, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 70}},
	{ID: 452, Name: "Wood Hammer", Type: 2048, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 120, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: -33, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 453, Name: "Aqua Jet", Type: 1024, Category: MoveCategoryPhysical, CurrentPP: 20, MaxPP: 20, Targets: 10, Priority: 1, Power: 40, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 454, Name: "Attack Order", Type: 64, Category: MoveCategoryPhysical, CurrentPP: 15, MaxPP: 15, Targets: 10, Priority: 0, Power: 90, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 1, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 455, Name: "Defend Order", Type: 64, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 456, Name: "Heal Order", Type: 64, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 50, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 457, Name: "Head Smash", Type: 32, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 150, Accuracy: 80, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: -50, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 458, Name: "Double Hit", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 10, MaxPP: 10, Targets: 10, Priority: 0, Power: 35, Accuracy: 90, metadata: MoveMeta{MinHits: 2, MaxHits: 2, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 459, Name: "Roar of Time", Type: 32768, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 150, Accuracy: 90, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 460, Name: "Spacial Rend", Type: 32768, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 100, Accuracy: 95, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 1, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 461, Name: "Lunar Dance", Type: 8192, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 7, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 462, Name: "Crush Grip", Type: 1, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 463, Name: "Magma Storm", Type: 512, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 100, Accuracy: 75, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 5, MaxTurns: 6, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 100, FlinchChance: 0, StatChance: 0}},
	{ID: 464, Name: "Dark Void", Type: 65536, Category: MoveCategoryStatus, CurrentPP: 10, MaxPP: 10, Targets: 11, Priority: 0, Power: 0, Accuracy: 50, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 2, MaxTurns: 4, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 465, Name: "Seed Flare", Type: 2048, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 120, Accuracy: 85, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 40}},
	{ID: 466, Name: "Ominous Wind", Type: 128, Category: MoveCategorySpecial, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 60, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 10, FlinchChance: 0, StatChance: 10}},
	{ID: 467, Name: "Shadow Force", Type: 128, Category: MoveCategoryPhysical, CurrentPP: 5, MaxPP: 5, Targets: 10, Priority: 0, Power: 120, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 10001, Name: "Shadow Rush", Type: 0, Category: MoveCategoryPhysical, CurrentPP: 0, MaxPP: 0, Targets: 10, Priority: 0, Power: 55, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 10002, Name: "Shadow Blast", Type: 0, Category: MoveCategoryPhysical, CurrentPP: 0, MaxPP: 0, Targets: 10, Priority: 0, Power: 80, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 10003, Name: "Shadow Blitz", Type: 0, Category: MoveCategoryPhysical, CurrentPP: 0, MaxPP: 0, Targets: 10, Priority: 0, Power: 40, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 10004, Name: "Shadow Bolt", Type: 0, Category: MoveCategorySpecial, CurrentPP: 0, MaxPP: 0, Targets: 10, Priority: 0, Power: 75, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 10005, Name: "Shadow Break", Type: 0, Category: MoveCategoryPhysical, CurrentPP: 0, MaxPP: 0, Targets: 10, Priority: 0, Power: 75, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 10006, Name: "Shadow Chill", Type: 0, Category: MoveCategorySpecial, CurrentPP: 0, MaxPP: 0, Targets: 10, Priority: 0, Power: 75, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 10007, Name: "Shadow End", Type: 0, Category: MoveCategoryPhysical, CurrentPP: 0, MaxPP: 0, Targets: 10, Priority: 0, Power: 120, Accuracy: 60, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 10008, Name: "Shadow Fire", Type: 0, Category: MoveCategorySpecial, CurrentPP: 0, MaxPP: 0, Targets: 10, Priority: 0, Power: 75, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 10009, Name: "Shadow Rave", Type: 0, Category: MoveCategorySpecial, CurrentPP: 0, MaxPP: 0, Targets: 6, Priority: 0, Power: 70, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 10010, Name: "Shadow Storm", Type: 0, Category: MoveCategorySpecial, CurrentPP: 0, MaxPP: 0, Targets: 6, Priority: 0, Power: 95, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 10011, Name: "Shadow Wave", Type: 0, Category: MoveCategorySpecial, CurrentPP: 0, MaxPP: 0, Targets: 6, Priority: 0, Power: 50, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 10012, Name: "Shadow Down", Type: 0, Category: MoveCategoryStatus, CurrentPP: 0, MaxPP: 0, Targets: 6, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 10013, Name: "Shadow Half", Type: 0, Category: MoveCategorySpecial, CurrentPP: 0, MaxPP: 0, Targets: 12, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 10014, Name: "Shadow Hold", Type: 0, Category: MoveCategoryStatus, CurrentPP: 0, MaxPP: 0, Targets: 6, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 10015, Name: "Shadow Mist", Type: 0, Category: MoveCategoryStatus, CurrentPP: 0, MaxPP: 0, Targets: 6, Priority: 0, Power: 0, Accuracy: 100, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 10016, Name: "Shadow Panic", Type: 0, Category: MoveCategoryStatus, CurrentPP: 0, MaxPP: 0, Targets: 6, Priority: 0, Power: 0, Accuracy: 90, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 10017, Name: "Shadow Shed", Type: 0, Category: MoveCategoryStatus, CurrentPP: 0, MaxPP: 0, Targets: 12, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
	{ID: 10018, Name: "Shadow Sky", Type: 0, Category: MoveCategoryStatus, CurrentPP: 0, MaxPP: 0, Targets: 12, Priority: 0, Power: 0, Accuracy: 0, metadata: MoveMeta{MinHits: 0, MaxHits: 0, MinTurns: 0, MaxTurns: 0, Drain: 0, Healing: 0, CritRate: 0, AilmentChance: 0, FlinchChance: 0, StatChance: 0}},
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
	ItemNone Item = iota
	ItemMasterBall
	ItemUltraBall
	ItemGreatBall
	ItemPokBall
	ItemSafariBall
	ItemNetBall
	ItemDiveBall
	ItemNestBall
	ItemRepeatBall
	ItemTimerBall
	ItemLuxuryBall
	ItemPremierBall
	ItemDuskBall
	ItemHealBall
	ItemQuickBall
	ItemCherishBall
	ItemPotion
	ItemAntidote
	ItemBurnHeal
	ItemIceHeal
	ItemAwakening
	ItemParalyzeHeal
	ItemFullRestore
	ItemMaxPotion
	ItemHyperPotion
	ItemSuperPotion
	ItemFullHeal
	ItemRevive
	ItemMaxRevive
	ItemFreshWater
	ItemSodaPop
	ItemLemonade
	ItemMoomooMilk
	ItemEnergyPowder
	ItemEnergyRoot
	ItemHealPowder
	ItemRevivalHerb
	ItemEther
	ItemMaxEther
	ItemElixir
	ItemMaxElixir
	ItemLavaCookie
	ItemBerryJuice
	ItemSacredAsh
	ItemHPUp
	ItemProtein
	ItemIron
	ItemCarbos
	ItemCalcium
	ItemRareCandy
	ItemPPUp
	ItemZinc
	ItemPPMax
	ItemOldGateau
	ItemGuardSpec
	ItemDireHit
	ItemXAttack
	ItemXDefense
	ItemXSpeed
	ItemXAccuracy
	ItemXSpAtk
	ItemXSpDef
	ItemPokDoll
	ItemFluffyTail
	ItemBlueFlute
	ItemYellowFlute
	ItemRedFlute
	ItemBlackFlute
	ItemWhiteFlute
	ItemAdamantOrb
	ItemLustrousOrb
	ItemCheriBerry
	ItemChestoBerry
	ItemPechaBerry
	ItemRawstBerry
	ItemAspearBerry
	ItemLeppaBerry
	ItemOranBerry
	ItemPersimBerry
	ItemLumBerry
	ItemSitrusBerry
	ItemFigyBerry
	ItemWikiBerry
	ItemMagoBerry
	ItemAguavBerry
	ItemIapapaBerry
	ItemOccaBerry
	ItemPasshoBerry
	ItemWacanBerry
	ItemRindoBerry
	ItemYacheBerry
	ItemChopleBerry
	ItemKebiaBerry
	ItemShucaBerry
	ItemCobaBerry
	ItemPayapaBerry
	ItemTangaBerry
	ItemChartiBerry
	ItemKasibBerry
	ItemHabanBerry
	ItemColburBerry
	ItemBabiriBerry
	ItemChilanBerry
	ItemLiechiBerry
	ItemGanlonBerry
	ItemSalacBerry
	ItemPetayaBerry
	ItemApicotBerry
	ItemLansatBerry
	ItemStarfBerry
	ItemEnigmaBerry
	ItemMicleBerry
	ItemCustapBerry
	ItemJabocaBerry
	ItemRowapBerry
	ItemBrightPowder
	ItemWhiteHerb
	ItemMachoBrace
	ItemExpShare
	ItemQuickClaw
	ItemSootheBell
	ItemMentalHerb
	ItemChoiceBand
	ItemKingsRock
	ItemSilverPowder
	ItemAmuletCoin
	ItemCleanseTag
	ItemSoulDew
	ItemDeepSeaTooth
	ItemDeepSeaScale
	ItemSmokeBall
	ItemEverstone
	ItemFocusBand
	ItemLuckyEgg
	ItemScopeLens
	ItemMetalCoat
	ItemLeftovers
	ItemLightBall
	ItemSoftSand
	ItemHardStone
	ItemMiracleSeed
	ItemBlackGlasses
	ItemBlackBelt
	ItemMagnet
	ItemMysticWater
	ItemSharpBeak
	ItemPoisonBarb
	ItemNeverMeltIce
	ItemSpellTag
	ItemTwistedSpoon
	ItemCharcoal
	ItemDragonFang
	ItemSilkScarf
	ItemShellBell
	ItemSeaIncense
	ItemLaxIncense
	ItemLuckyPunch
	ItemMetalPowder
	ItemThickClub
	ItemStick
	ItemRedScarf
	ItemBlueScarf
	ItemPinkScarf
	ItemGreenScarf
	ItemYellowScarf
	ItemWideLens
	ItemMuscleBand
	ItemWiseGlasses
	ItemExpertBelt
	ItemLightClay
	ItemLifeOrb
	ItemPowerHerb
	ItemToxicOrb
	ItemFlameOrb
	ItemQuickPowder
	ItemFocusSash
	ItemZoomLens
	ItemMetronome
	ItemIronBall
	ItemLaggingTail
	ItemDestinyKnot
	ItemBlackSludge
	ItemIcyRock
	ItemSmoothRock
	ItemHeatRock
	ItemDampRock
	ItemGripClaw
	ItemChoiceScarf
	ItemStickyBarb
	ItemPowerBracer
	ItemPowerBelt
	ItemPowerLens
	ItemPowerBand
	ItemPowerAnklet
	ItemPowerWeight
	ItemShedShell
	ItemBigRoot
	ItemChoiceSpecs
	ItemFlamePlate
	ItemSplashPlate
	ItemZapPlate
	ItemMeadowPlate
	ItemIciclePlate
	ItemFistPlate
	ItemToxicPlate
	ItemEarthPlate
	ItemSkyPlate
	ItemMindPlate
	ItemInsectPlate
	ItemStonePlate
	ItemSpookyPlate
	ItemDracoPlate
	ItemDreadPlate
	ItemIronPlate
	ItemOddIncense
	ItemRockIncense
	ItemFullIncense
	ItemWaveIncense
	ItemRoseIncense
	ItemLuckIncense
	ItemPureIncense
	ItemRazorClaw
	ItemRazorFang
)

// A collection of all items in the game
var AllItems = []ItemData{
	{Name: "Master Ball", Category: 34, FlingPower: 0, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Ultra Ball", Category: 34, FlingPower: 0, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Great Ball", Category: 34, FlingPower: 0, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Poké Ball", Category: 34, FlingPower: 0, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Safari Ball", Category: 34, FlingPower: 0, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Net Ball", Category: 33, FlingPower: 0, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Dive Ball", Category: 33, FlingPower: 0, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Nest Ball", Category: 33, FlingPower: 0, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Repeat Ball", Category: 33, FlingPower: 0, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Timer Ball", Category: 33, FlingPower: 0, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Luxury Ball", Category: 33, FlingPower: 0, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Premier Ball", Category: 33, FlingPower: 0, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Dusk Ball", Category: 33, FlingPower: 0, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Heal Ball", Category: 33, FlingPower: 0, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Quick Ball", Category: 33, FlingPower: 0, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Cherish Ball", Category: 33, FlingPower: 0, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Potion", Category: 27, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Antidote", Category: 30, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Burn Heal", Category: 30, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Ice Heal", Category: 30, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Awakening", Category: 30, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Paralyze Heal", Category: 30, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Full Restore", Category: 27, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Max Potion", Category: 27, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Hyper Potion", Category: 27, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Super Potion", Category: 27, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Full Heal", Category: 30, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Revive", Category: 29, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Max Revive", Category: 29, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Fresh Water", Category: 27, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Soda Pop", Category: 27, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Lemonade", Category: 27, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Moomoo Milk", Category: 27, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Energy Powder", Category: 27, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Energy Root", Category: 27, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Heal Powder", Category: 30, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Revival Herb", Category: 29, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Ether", Category: 28, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Max Ether", Category: 28, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Elixir", Category: 28, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Max Elixir", Category: 28, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Lava Cookie", Category: 30, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Berry Juice", Category: 27, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Sacred Ash", Category: 29, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "HP Up", Category: 26, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Protein", Category: 26, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Iron", Category: 26, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Carbos", Category: 26, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Calcium", Category: 26, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Rare Candy", Category: 26, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "PP Up", Category: 26, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Zinc", Category: 26, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "PP Max", Category: 26, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Old Gateau", Category: 30, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Guard Spec.", Category: 1, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Dire Hit", Category: 1, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "X Attack", Category: 1, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "X Defense", Category: 1, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "X Speed", Category: 1, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "X Accuracy", Category: 1, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "X Sp. Atk", Category: 1, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "X Sp. Def", Category: 1, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Poké Doll", Category: 11, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Fluffy Tail", Category: 11, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Blue Flute", Category: 38, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Yellow Flute", Category: 38, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Red Flute", Category: 38, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagUsableInBattle | FlagHoldable},
	{Name: "Black Flute", Category: 11, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagHoldable},
	{Name: "White Flute", Category: 11, FlingPower: 30, FlingEffect: 0, Flags: FlagConsumable | FlagHoldable},
	{Name: "Adamant Orb", Category: 18, FlingPower: 60, FlingEffect: 0, Flags: FlagHoldable},
	{Name: "Lustrous Orb", Category: 18, FlingPower: 60, FlingEffect: 0, Flags: FlagHoldable},
	{Name: "Cheri Berry", Category: 3, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive},
	{Name: "Chesto Berry", Category: 3, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive},
	{Name: "Pecha Berry", Category: 3, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive},
	{Name: "Rawst Berry", Category: 3, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive},
	{Name: "Aspear Berry", Category: 3, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive},
	{Name: "Leppa Berry", Category: 3, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive},
	{Name: "Oran Berry", Category: 3, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive},
	{Name: "Persim Berry", Category: 3, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive},
	{Name: "Lum Berry", Category: 3, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive},
	{Name: "Sitrus Berry", Category: 3, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive},
	{Name: "Figy Berry", Category: 6, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive},
	{Name: "Wiki Berry", Category: 6, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive},
	{Name: "Mago Berry", Category: 6, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive},
	{Name: "Aguav Berry", Category: 6, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive},
	{Name: "Iapapa Berry", Category: 6, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive},
	{Name: "Occa Berry", Category: 7, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{Name: "Passho Berry", Category: 7, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{Name: "Wacan Berry", Category: 7, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{Name: "Rindo Berry", Category: 7, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{Name: "Yache Berry", Category: 7, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{Name: "Chople Berry", Category: 7, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{Name: "Kebia Berry", Category: 7, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{Name: "Shuca Berry", Category: 7, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{Name: "Coba Berry", Category: 7, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{Name: "Payapa Berry", Category: 7, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{Name: "Tanga Berry", Category: 7, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{Name: "Charti Berry", Category: 7, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{Name: "Kasib Berry", Category: 7, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{Name: "Haban Berry", Category: 7, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{Name: "Colbur Berry", Category: 7, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{Name: "Babiri Berry", Category: 7, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{Name: "Chilan Berry", Category: 7, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{Name: "Liechi Berry", Category: 5, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive | FlagConsumable | FlagHoldable},
	{Name: "Ganlon Berry", Category: 5, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive | FlagConsumable | FlagHoldable},
	{Name: "Salac Berry", Category: 5, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive | FlagConsumable | FlagHoldable},
	{Name: "Petaya Berry", Category: 5, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive | FlagConsumable | FlagHoldable},
	{Name: "Apicot Berry", Category: 5, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive | FlagConsumable | FlagHoldable},
	{Name: "Lansat Berry", Category: 5, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive | FlagConsumable | FlagHoldable},
	{Name: "Starf Berry", Category: 5, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive | FlagConsumable | FlagHoldable},
	{Name: "Enigma Berry", Category: 4, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{Name: "Micle Berry", Category: 5, FlingPower: 10, FlingEffect: 3, Flags: FlagHoldableActive | FlagConsumable | FlagHoldable},
	{Name: "Custap Berry", Category: 5, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive | FlagConsumable | FlagHoldable},
	{Name: "Jaboca Berry", Category: 4, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{Name: "Rowap Berry", Category: 4, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{Name: "Bright Powder", Category: 12, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "White Herb", Category: 12, FlingPower: 10, FlingEffect: 4, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Macho Brace", Category: 14, FlingPower: 60, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Exp. Share", Category: 16, FlingPower: 0, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Quick Claw", Category: 12, FlingPower: 80, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Soothe Bell", Category: 16, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Mental Herb", Category: 12, FlingPower: 10, FlingEffect: 4, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Choice Band", Category: 13, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "King’s Rock", Category: 12, FlingPower: 30, FlingEffect: 7, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Silver Powder", Category: 19, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Amulet Coin", Category: 16, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Cleanse Tag", Category: 16, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable},
	{Name: "Soul Dew", Category: 18, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable},
	{Name: "Deep Sea Tooth", Category: 18, FlingPower: 90, FlingEffect: 0, Flags: FlagHoldable},
	{Name: "Deep Sea Scale", Category: 18, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable},
	{Name: "Smoke Ball", Category: 12, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Everstone", Category: 16, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Focus Band", Category: 12, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Lucky Egg", Category: 16, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Scope Lens", Category: 12, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Metal Coat", Category: 19, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Leftovers", Category: 12, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Light Ball", Category: 18, FlingPower: 30, FlingEffect: 5, Flags: FlagHoldable},
	{Name: "Soft Sand", Category: 19, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Hard Stone", Category: 19, FlingPower: 100, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Miracle Seed", Category: 19, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Black Glasses", Category: 19, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Black Belt", Category: 19, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Magnet", Category: 19, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Mystic Water", Category: 19, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Sharp Beak", Category: 19, FlingPower: 50, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Poison Barb", Category: 19, FlingPower: 70, FlingEffect: 6, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Never-Melt Ice", Category: 19, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Spell Tag", Category: 19, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Twisted Spoon", Category: 19, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Charcoal", Category: 19, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Dragon Fang", Category: 19, FlingPower: 70, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Silk Scarf", Category: 19, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Shell Bell", Category: 12, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Sea Incense", Category: 19, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Lax Incense", Category: 12, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Lucky Punch", Category: 18, FlingPower: 40, FlingEffect: 0, Flags: FlagHoldable},
	{Name: "Metal Powder", Category: 18, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable},
	{Name: "Thick Club", Category: 18, FlingPower: 90, FlingEffect: 0, Flags: FlagHoldable},
	{Name: "Stick", Category: 18, FlingPower: 60, FlingEffect: 0, Flags: FlagHoldable},
	{Name: "Red Scarf", Category: 36, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{Name: "Blue Scarf", Category: 36, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{Name: "Pink Scarf", Category: 36, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{Name: "Green Scarf", Category: 36, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{Name: "Yellow Scarf", Category: 36, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldableActive},
	{Name: "Wide Lens", Category: 12, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Muscle Band", Category: 12, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Wise Glasses", Category: 12, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Expert Belt", Category: 12, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Light Clay", Category: 12, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Life Orb", Category: 12, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Power Herb", Category: 12, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Toxic Orb", Category: 15, FlingPower: 30, FlingEffect: 1, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Flame Orb", Category: 15, FlingPower: 30, FlingEffect: 2, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Quick Powder", Category: 18, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable},
	{Name: "Focus Sash", Category: 12, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Zoom Lens", Category: 12, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Metronome", Category: 12, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Iron Ball", Category: 15, FlingPower: 130, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Lagging Tail", Category: 15, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Destiny Knot", Category: 12, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Black Sludge", Category: 12, FlingPower: 30, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Icy Rock", Category: 12, FlingPower: 40, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Smooth Rock", Category: 12, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Heat Rock", Category: 12, FlingPower: 60, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Damp Rock", Category: 12, FlingPower: 60, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Grip Claw", Category: 12, FlingPower: 90, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Choice Scarf", Category: 13, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Sticky Barb", Category: 15, FlingPower: 80, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Power Bracer", Category: 14, FlingPower: 70, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Power Belt", Category: 14, FlingPower: 70, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Power Lens", Category: 14, FlingPower: 70, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Power Band", Category: 14, FlingPower: 70, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Power Anklet", Category: 14, FlingPower: 70, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Power Weight", Category: 14, FlingPower: 70, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Shed Shell", Category: 12, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Big Root", Category: 12, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Choice Specs", Category: 13, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Flame Plate", Category: 17, FlingPower: 90, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Splash Plate", Category: 17, FlingPower: 90, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Zap Plate", Category: 17, FlingPower: 90, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Meadow Plate", Category: 17, FlingPower: 90, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Icicle Plate", Category: 17, FlingPower: 90, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Fist Plate", Category: 17, FlingPower: 90, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Toxic Plate", Category: 17, FlingPower: 90, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Earth Plate", Category: 17, FlingPower: 90, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Sky Plate", Category: 17, FlingPower: 90, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Mind Plate", Category: 17, FlingPower: 90, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Insect Plate", Category: 17, FlingPower: 90, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Stone Plate", Category: 17, FlingPower: 90, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Spooky Plate", Category: 17, FlingPower: 90, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Draco Plate", Category: 17, FlingPower: 90, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Dread Plate", Category: 17, FlingPower: 90, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Iron Plate", Category: 17, FlingPower: 90, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Odd Incense", Category: 19, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Rock Incense", Category: 19, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Full Incense", Category: 15, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Wave Incense", Category: 19, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Rose Incense", Category: 19, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Luck Incense", Category: 16, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Pure Incense", Category: 16, FlingPower: 10, FlingEffect: 0, Flags: FlagHoldable},
	{Name: "Razor Claw", Category: 12, FlingPower: 80, FlingEffect: 0, Flags: FlagHoldable | FlagHoldableActive},
	{Name: "Razor Fang", Category: 12, FlingPower: 30, FlingEffect: 7, Flags: FlagHoldable | FlagHoldableActive},
}

//A table of levels mapped to the total experience at that level for each growth rate
var ExpTable = map[int]map[int]int{
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

// Table of Pokemon base stats and EV yield
var pokemonBaseStats = map[int]PokemonBaseStats{
	1:   {[6]int{45, 49, 49, 65, 65, 45}, [6]int{0, 0, 0, 1, 0, 0}},
	2:   {[6]int{60, 62, 63, 80, 80, 60}, [6]int{0, 0, 0, 1, 1, 0}},
	3:   {[6]int{80, 82, 83, 100, 100, 80}, [6]int{0, 0, 0, 2, 1, 0}},
	4:   {[6]int{39, 52, 43, 60, 50, 65}, [6]int{0, 0, 0, 0, 0, 1}},
	5:   {[6]int{58, 64, 58, 80, 65, 80}, [6]int{0, 0, 0, 1, 0, 1}},
	6:   {[6]int{78, 84, 78, 109, 85, 100}, [6]int{0, 0, 0, 3, 0, 0}},
	7:   {[6]int{44, 48, 65, 50, 64, 43}, [6]int{0, 0, 1, 0, 0, 0}},
	8:   {[6]int{59, 63, 80, 65, 80, 58}, [6]int{0, 0, 1, 0, 1, 0}},
	9:   {[6]int{79, 83, 100, 85, 105, 78}, [6]int{0, 0, 0, 0, 3, 0}},
	10:  {[6]int{45, 30, 35, 20, 20, 45}, [6]int{1, 0, 0, 0, 0, 0}},
	11:  {[6]int{50, 20, 55, 25, 25, 30}, [6]int{0, 0, 2, 0, 0, 0}},
	12:  {[6]int{60, 45, 50, 90, 80, 70}, [6]int{0, 0, 0, 2, 1, 0}},
	13:  {[6]int{40, 35, 30, 20, 20, 50}, [6]int{0, 0, 0, 0, 0, 1}},
	14:  {[6]int{45, 25, 50, 25, 25, 35}, [6]int{0, 0, 2, 0, 0, 0}},
	15:  {[6]int{65, 90, 40, 45, 80, 75}, [6]int{0, 2, 0, 0, 1, 0}},
	16:  {[6]int{40, 45, 40, 35, 35, 56}, [6]int{0, 0, 0, 0, 0, 1}},
	17:  {[6]int{63, 60, 55, 50, 50, 71}, [6]int{0, 0, 0, 0, 0, 2}},
	18:  {[6]int{83, 80, 75, 70, 70, 101}, [6]int{0, 0, 0, 0, 0, 3}},
	19:  {[6]int{30, 56, 35, 25, 35, 72}, [6]int{0, 0, 0, 0, 0, 1}},
	20:  {[6]int{55, 81, 60, 50, 70, 97}, [6]int{0, 0, 0, 0, 0, 2}},
	21:  {[6]int{40, 60, 30, 31, 31, 70}, [6]int{0, 0, 0, 0, 0, 1}},
	22:  {[6]int{65, 90, 65, 61, 61, 100}, [6]int{0, 0, 0, 0, 0, 2}},
	23:  {[6]int{35, 60, 44, 40, 54, 55}, [6]int{0, 1, 0, 0, 0, 0}},
	24:  {[6]int{60, 95, 69, 65, 79, 80}, [6]int{0, 2, 0, 0, 0, 0}},
	25:  {[6]int{35, 55, 40, 50, 50, 90}, [6]int{0, 0, 0, 0, 0, 2}},
	26:  {[6]int{60, 90, 55, 90, 80, 110}, [6]int{0, 0, 0, 0, 0, 3}},
	27:  {[6]int{50, 75, 85, 20, 30, 40}, [6]int{0, 0, 1, 0, 0, 0}},
	28:  {[6]int{75, 100, 110, 45, 55, 65}, [6]int{0, 0, 2, 0, 0, 0}},
	29:  {[6]int{55, 47, 52, 40, 40, 41}, [6]int{1, 0, 0, 0, 0, 0}},
	30:  {[6]int{70, 62, 67, 55, 55, 56}, [6]int{2, 0, 0, 0, 0, 0}},
	31:  {[6]int{90, 92, 87, 75, 85, 76}, [6]int{3, 0, 0, 0, 0, 0}},
	32:  {[6]int{46, 57, 40, 40, 40, 50}, [6]int{0, 1, 0, 0, 0, 0}},
	33:  {[6]int{61, 72, 57, 55, 55, 65}, [6]int{0, 2, 0, 0, 0, 0}},
	34:  {[6]int{81, 102, 77, 85, 75, 85}, [6]int{0, 3, 0, 0, 0, 0}},
	35:  {[6]int{70, 45, 48, 60, 65, 35}, [6]int{2, 0, 0, 0, 0, 0}},
	36:  {[6]int{95, 70, 73, 95, 90, 60}, [6]int{3, 0, 0, 0, 0, 0}},
	37:  {[6]int{38, 41, 40, 50, 65, 65}, [6]int{0, 0, 0, 0, 0, 1}},
	38:  {[6]int{73, 76, 75, 81, 100, 100}, [6]int{0, 0, 0, 0, 1, 1}},
	39:  {[6]int{115, 45, 20, 45, 25, 20}, [6]int{2, 0, 0, 0, 0, 0}},
	40:  {[6]int{140, 70, 45, 85, 50, 45}, [6]int{3, 0, 0, 0, 0, 0}},
	41:  {[6]int{40, 45, 35, 30, 40, 55}, [6]int{0, 0, 0, 0, 0, 1}},
	42:  {[6]int{75, 80, 70, 65, 75, 90}, [6]int{0, 0, 0, 0, 0, 2}},
	43:  {[6]int{45, 50, 55, 75, 65, 30}, [6]int{0, 0, 0, 1, 0, 0}},
	44:  {[6]int{60, 65, 70, 85, 75, 40}, [6]int{0, 0, 0, 2, 0, 0}},
	45:  {[6]int{75, 80, 85, 110, 90, 50}, [6]int{0, 0, 0, 3, 0, 0}},
	46:  {[6]int{35, 70, 55, 45, 55, 25}, [6]int{0, 1, 0, 0, 0, 0}},
	47:  {[6]int{60, 95, 80, 60, 80, 30}, [6]int{0, 2, 1, 0, 0, 0}},
	48:  {[6]int{60, 55, 50, 40, 55, 45}, [6]int{0, 0, 0, 0, 1, 0}},
	49:  {[6]int{70, 65, 60, 90, 75, 90}, [6]int{0, 0, 0, 1, 0, 1}},
	50:  {[6]int{10, 55, 25, 35, 45, 95}, [6]int{0, 0, 0, 0, 0, 1}},
	51:  {[6]int{35, 100, 50, 50, 70, 120}, [6]int{0, 0, 0, 0, 0, 2}},
	52:  {[6]int{40, 45, 35, 40, 40, 90}, [6]int{0, 0, 0, 0, 0, 1}},
	53:  {[6]int{65, 70, 60, 65, 65, 115}, [6]int{0, 0, 0, 0, 0, 2}},
	54:  {[6]int{50, 52, 48, 65, 50, 55}, [6]int{0, 0, 0, 1, 0, 0}},
	55:  {[6]int{80, 82, 78, 95, 80, 85}, [6]int{0, 0, 0, 2, 0, 0}},
	56:  {[6]int{40, 80, 35, 35, 45, 70}, [6]int{0, 1, 0, 0, 0, 0}},
	57:  {[6]int{65, 105, 60, 60, 70, 95}, [6]int{0, 2, 0, 0, 0, 0}},
	58:  {[6]int{55, 70, 45, 70, 50, 60}, [6]int{0, 1, 0, 0, 0, 0}},
	59:  {[6]int{90, 110, 80, 100, 80, 95}, [6]int{0, 2, 0, 0, 0, 0}},
	60:  {[6]int{40, 50, 40, 40, 40, 90}, [6]int{0, 0, 0, 0, 0, 1}},
	61:  {[6]int{65, 65, 65, 50, 50, 90}, [6]int{0, 0, 0, 0, 0, 2}},
	62:  {[6]int{90, 95, 95, 70, 90, 70}, [6]int{0, 0, 3, 0, 0, 0}},
	63:  {[6]int{25, 20, 15, 105, 55, 90}, [6]int{0, 0, 0, 1, 0, 0}},
	64:  {[6]int{40, 35, 30, 120, 70, 105}, [6]int{0, 0, 0, 2, 0, 0}},
	65:  {[6]int{55, 50, 45, 135, 95, 120}, [6]int{0, 0, 0, 3, 0, 0}},
	66:  {[6]int{70, 80, 50, 35, 35, 35}, [6]int{0, 1, 0, 0, 0, 0}},
	67:  {[6]int{80, 100, 70, 50, 60, 45}, [6]int{0, 2, 0, 0, 0, 0}},
	68:  {[6]int{90, 130, 80, 65, 85, 55}, [6]int{0, 3, 0, 0, 0, 0}},
	69:  {[6]int{50, 75, 35, 70, 30, 40}, [6]int{0, 1, 0, 0, 0, 0}},
	70:  {[6]int{65, 90, 50, 85, 45, 55}, [6]int{0, 2, 0, 0, 0, 0}},
	71:  {[6]int{80, 105, 65, 100, 70, 70}, [6]int{0, 3, 0, 0, 0, 0}},
	72:  {[6]int{40, 40, 35, 50, 100, 70}, [6]int{0, 0, 0, 0, 1, 0}},
	73:  {[6]int{80, 70, 65, 80, 120, 100}, [6]int{0, 0, 0, 0, 2, 0}},
	74:  {[6]int{40, 80, 100, 30, 30, 20}, [6]int{0, 0, 1, 0, 0, 0}},
	75:  {[6]int{55, 95, 115, 45, 45, 35}, [6]int{0, 0, 2, 0, 0, 0}},
	76:  {[6]int{80, 120, 130, 55, 65, 45}, [6]int{0, 0, 3, 0, 0, 0}},
	77:  {[6]int{50, 85, 55, 65, 65, 90}, [6]int{0, 0, 0, 0, 0, 1}},
	78:  {[6]int{65, 100, 70, 80, 80, 105}, [6]int{0, 0, 0, 0, 0, 2}},
	79:  {[6]int{90, 65, 65, 40, 40, 15}, [6]int{1, 0, 0, 0, 0, 0}},
	80:  {[6]int{95, 75, 110, 100, 80, 30}, [6]int{0, 0, 2, 0, 0, 0}},
	81:  {[6]int{25, 35, 70, 95, 55, 45}, [6]int{0, 0, 0, 1, 0, 0}},
	82:  {[6]int{50, 60, 95, 120, 70, 70}, [6]int{0, 0, 0, 2, 0, 0}},
	83:  {[6]int{52, 90, 55, 58, 62, 60}, [6]int{0, 1, 0, 0, 0, 0}},
	84:  {[6]int{35, 85, 45, 35, 35, 75}, [6]int{0, 1, 0, 0, 0, 0}},
	85:  {[6]int{60, 110, 70, 60, 60, 110}, [6]int{0, 2, 0, 0, 0, 0}},
	86:  {[6]int{65, 45, 55, 45, 70, 45}, [6]int{0, 0, 0, 0, 1, 0}},
	87:  {[6]int{90, 70, 80, 70, 95, 70}, [6]int{0, 0, 0, 0, 2, 0}},
	88:  {[6]int{80, 80, 50, 40, 50, 25}, [6]int{1, 0, 0, 0, 0, 0}},
	89:  {[6]int{105, 105, 75, 65, 100, 50}, [6]int{1, 1, 0, 0, 0, 0}},
	90:  {[6]int{30, 65, 100, 45, 25, 40}, [6]int{0, 0, 1, 0, 0, 0}},
	91:  {[6]int{50, 95, 180, 85, 45, 70}, [6]int{0, 0, 2, 0, 0, 0}},
	92:  {[6]int{30, 35, 30, 100, 35, 80}, [6]int{0, 0, 0, 1, 0, 0}},
	93:  {[6]int{45, 50, 45, 115, 55, 95}, [6]int{0, 0, 0, 2, 0, 0}},
	94:  {[6]int{60, 65, 60, 130, 75, 110}, [6]int{0, 0, 0, 3, 0, 0}},
	95:  {[6]int{35, 45, 160, 30, 45, 70}, [6]int{0, 0, 1, 0, 0, 0}},
	96:  {[6]int{60, 48, 45, 43, 90, 42}, [6]int{0, 0, 0, 0, 1, 0}},
	97:  {[6]int{85, 73, 70, 73, 115, 67}, [6]int{0, 0, 0, 0, 2, 0}},
	98:  {[6]int{30, 105, 90, 25, 25, 50}, [6]int{0, 1, 0, 0, 0, 0}},
	99:  {[6]int{55, 130, 115, 50, 50, 75}, [6]int{0, 2, 0, 0, 0, 0}},
	100: {[6]int{40, 30, 50, 55, 55, 100}, [6]int{0, 0, 0, 0, 0, 1}},
	101: {[6]int{60, 50, 70, 80, 80, 150}, [6]int{0, 0, 0, 0, 0, 2}},
	102: {[6]int{60, 40, 80, 60, 45, 40}, [6]int{0, 0, 1, 0, 0, 0}},
	103: {[6]int{95, 95, 85, 125, 75, 55}, [6]int{0, 0, 0, 2, 0, 0}},
	104: {[6]int{50, 50, 95, 40, 50, 35}, [6]int{0, 0, 1, 0, 0, 0}},
	105: {[6]int{60, 80, 110, 50, 80, 45}, [6]int{0, 0, 2, 0, 0, 0}},
	106: {[6]int{50, 120, 53, 35, 110, 87}, [6]int{0, 2, 0, 0, 0, 0}},
	107: {[6]int{50, 105, 79, 35, 110, 76}, [6]int{0, 0, 0, 0, 2, 0}},
	108: {[6]int{90, 55, 75, 60, 75, 30}, [6]int{2, 0, 0, 0, 0, 0}},
	109: {[6]int{40, 65, 95, 60, 45, 35}, [6]int{0, 0, 1, 0, 0, 0}},
	110: {[6]int{65, 90, 120, 85, 70, 60}, [6]int{0, 0, 2, 0, 0, 0}},
	111: {[6]int{80, 85, 95, 30, 30, 25}, [6]int{0, 0, 1, 0, 0, 0}},
	112: {[6]int{105, 130, 120, 45, 45, 40}, [6]int{0, 2, 0, 0, 0, 0}},
	113: {[6]int{250, 5, 5, 35, 105, 50}, [6]int{2, 0, 0, 0, 0, 0}},
	114: {[6]int{65, 55, 115, 100, 40, 60}, [6]int{0, 0, 1, 0, 0, 0}},
	115: {[6]int{105, 95, 80, 40, 80, 90}, [6]int{2, 0, 0, 0, 0, 0}},
	116: {[6]int{30, 40, 70, 70, 25, 60}, [6]int{0, 0, 0, 1, 0, 0}},
	117: {[6]int{55, 65, 95, 95, 45, 85}, [6]int{0, 0, 1, 1, 0, 0}},
	118: {[6]int{45, 67, 60, 35, 50, 63}, [6]int{0, 1, 0, 0, 0, 0}},
	119: {[6]int{80, 92, 65, 65, 80, 68}, [6]int{0, 2, 0, 0, 0, 0}},
	120: {[6]int{30, 45, 55, 70, 55, 85}, [6]int{0, 0, 0, 0, 0, 1}},
	121: {[6]int{60, 75, 85, 100, 85, 115}, [6]int{0, 0, 0, 0, 0, 2}},
	122: {[6]int{40, 45, 65, 100, 120, 90}, [6]int{0, 0, 0, 0, 2, 0}},
	123: {[6]int{70, 110, 80, 55, 80, 105}, [6]int{0, 1, 0, 0, 0, 0}},
	124: {[6]int{65, 50, 35, 115, 95, 95}, [6]int{0, 0, 0, 2, 0, 0}},
	125: {[6]int{65, 83, 57, 95, 85, 105}, [6]int{0, 0, 0, 0, 0, 2}},
	126: {[6]int{65, 95, 57, 100, 85, 93}, [6]int{0, 0, 0, 2, 0, 0}},
	127: {[6]int{65, 125, 100, 55, 70, 85}, [6]int{0, 2, 0, 0, 0, 0}},
	128: {[6]int{75, 100, 95, 40, 70, 110}, [6]int{0, 1, 0, 0, 0, 1}},
	129: {[6]int{20, 10, 55, 15, 20, 80}, [6]int{0, 0, 0, 0, 0, 1}},
	130: {[6]int{95, 125, 79, 60, 100, 81}, [6]int{0, 2, 0, 0, 0, 0}},
	131: {[6]int{130, 85, 80, 85, 95, 60}, [6]int{2, 0, 0, 0, 0, 0}},
	132: {[6]int{48, 48, 48, 48, 48, 48}, [6]int{1, 0, 0, 0, 0, 0}},
	133: {[6]int{55, 55, 50, 45, 65, 55}, [6]int{0, 0, 0, 0, 1, 0}},
	134: {[6]int{130, 65, 60, 110, 95, 65}, [6]int{2, 0, 0, 0, 0, 0}},
	135: {[6]int{65, 65, 60, 110, 95, 130}, [6]int{0, 0, 0, 0, 0, 2}},
	136: {[6]int{65, 130, 60, 95, 110, 65}, [6]int{0, 2, 0, 0, 0, 0}},
	137: {[6]int{65, 60, 70, 85, 75, 40}, [6]int{0, 0, 0, 1, 0, 0}},
	138: {[6]int{35, 40, 100, 90, 55, 35}, [6]int{0, 0, 1, 0, 0, 0}},
	139: {[6]int{70, 60, 125, 115, 70, 55}, [6]int{0, 0, 2, 0, 0, 0}},
	140: {[6]int{30, 80, 90, 55, 45, 55}, [6]int{0, 0, 1, 0, 0, 0}},
	141: {[6]int{60, 115, 105, 65, 70, 80}, [6]int{0, 2, 0, 0, 0, 0}},
	142: {[6]int{80, 105, 65, 60, 75, 130}, [6]int{0, 0, 0, 0, 0, 2}},
	143: {[6]int{160, 110, 65, 65, 110, 30}, [6]int{2, 0, 0, 0, 0, 0}},
	144: {[6]int{90, 85, 100, 95, 125, 85}, [6]int{0, 0, 0, 0, 3, 0}},
	145: {[6]int{90, 90, 85, 125, 90, 100}, [6]int{0, 0, 0, 3, 0, 0}},
	146: {[6]int{90, 100, 90, 125, 85, 90}, [6]int{0, 0, 0, 3, 0, 0}},
	147: {[6]int{41, 64, 45, 50, 50, 50}, [6]int{0, 1, 0, 0, 0, 0}},
	148: {[6]int{61, 84, 65, 70, 70, 70}, [6]int{0, 2, 0, 0, 0, 0}},
	149: {[6]int{91, 134, 95, 100, 100, 80}, [6]int{0, 3, 0, 0, 0, 0}},
	150: {[6]int{106, 110, 90, 154, 90, 130}, [6]int{0, 0, 0, 3, 0, 0}},
	151: {[6]int{100, 100, 100, 100, 100, 100}, [6]int{3, 0, 0, 0, 0, 0}},
	152: {[6]int{45, 49, 65, 49, 65, 45}, [6]int{0, 0, 0, 0, 1, 0}},
	153: {[6]int{60, 62, 80, 63, 80, 60}, [6]int{0, 0, 1, 0, 1, 0}},
	154: {[6]int{80, 82, 100, 83, 100, 80}, [6]int{0, 0, 1, 0, 2, 0}},
	155: {[6]int{39, 52, 43, 60, 50, 65}, [6]int{0, 0, 0, 0, 0, 1}},
	156: {[6]int{58, 64, 58, 80, 65, 80}, [6]int{0, 0, 0, 1, 0, 1}},
	157: {[6]int{78, 84, 78, 109, 85, 100}, [6]int{0, 0, 0, 3, 0, 0}},
	158: {[6]int{50, 65, 64, 44, 48, 43}, [6]int{0, 1, 0, 0, 0, 0}},
	159: {[6]int{65, 80, 80, 59, 63, 58}, [6]int{0, 1, 1, 0, 0, 0}},
	160: {[6]int{85, 105, 100, 79, 83, 78}, [6]int{0, 2, 1, 0, 0, 0}},
	161: {[6]int{35, 46, 34, 35, 45, 20}, [6]int{0, 1, 0, 0, 0, 0}},
	162: {[6]int{85, 76, 64, 45, 55, 90}, [6]int{0, 0, 0, 0, 0, 2}},
	163: {[6]int{60, 30, 30, 36, 56, 50}, [6]int{1, 0, 0, 0, 0, 0}},
	164: {[6]int{100, 50, 50, 86, 96, 70}, [6]int{2, 0, 0, 0, 0, 0}},
	165: {[6]int{40, 20, 30, 40, 80, 55}, [6]int{0, 0, 0, 0, 1, 0}},
	166: {[6]int{55, 35, 50, 55, 110, 85}, [6]int{0, 0, 0, 0, 2, 0}},
	167: {[6]int{40, 60, 40, 40, 40, 30}, [6]int{0, 1, 0, 0, 0, 0}},
	168: {[6]int{70, 90, 70, 60, 70, 40}, [6]int{0, 2, 0, 0, 0, 0}},
	169: {[6]int{85, 90, 80, 70, 80, 130}, [6]int{0, 0, 0, 0, 0, 3}},
	170: {[6]int{75, 38, 38, 56, 56, 67}, [6]int{1, 0, 0, 0, 0, 0}},
	171: {[6]int{125, 58, 58, 76, 76, 67}, [6]int{2, 0, 0, 0, 0, 0}},
	172: {[6]int{20, 40, 15, 35, 35, 60}, [6]int{0, 0, 0, 0, 0, 1}},
	173: {[6]int{50, 25, 28, 45, 55, 15}, [6]int{0, 0, 0, 0, 1, 0}},
	174: {[6]int{90, 30, 15, 40, 20, 15}, [6]int{1, 0, 0, 0, 0, 0}},
	175: {[6]int{35, 20, 65, 40, 65, 20}, [6]int{0, 0, 0, 0, 1, 0}},
	176: {[6]int{55, 40, 85, 80, 105, 40}, [6]int{0, 0, 0, 0, 2, 0}},
	177: {[6]int{40, 50, 45, 70, 45, 70}, [6]int{0, 0, 0, 1, 0, 0}},
	178: {[6]int{65, 75, 70, 95, 70, 95}, [6]int{0, 0, 0, 1, 0, 1}},
	179: {[6]int{55, 40, 40, 65, 45, 35}, [6]int{0, 0, 0, 1, 0, 0}},
	180: {[6]int{70, 55, 55, 80, 60, 45}, [6]int{0, 0, 0, 2, 0, 0}},
	181: {[6]int{90, 75, 85, 115, 90, 55}, [6]int{0, 0, 0, 3, 0, 0}},
	182: {[6]int{75, 80, 95, 90, 100, 50}, [6]int{0, 0, 0, 0, 3, 0}},
	183: {[6]int{70, 20, 50, 20, 50, 40}, [6]int{2, 0, 0, 0, 0, 0}},
	184: {[6]int{100, 50, 80, 60, 80, 50}, [6]int{3, 0, 0, 0, 0, 0}},
	185: {[6]int{70, 100, 115, 30, 65, 30}, [6]int{0, 0, 2, 0, 0, 0}},
	186: {[6]int{90, 75, 75, 90, 100, 70}, [6]int{0, 0, 0, 0, 3, 0}},
	187: {[6]int{35, 35, 40, 35, 55, 50}, [6]int{0, 0, 0, 0, 1, 0}},
	188: {[6]int{55, 45, 50, 45, 65, 80}, [6]int{0, 0, 0, 0, 0, 2}},
	189: {[6]int{75, 55, 70, 55, 95, 110}, [6]int{0, 0, 0, 0, 0, 3}},
	190: {[6]int{55, 70, 55, 40, 55, 85}, [6]int{0, 0, 0, 0, 0, 1}},
	191: {[6]int{30, 30, 30, 30, 30, 30}, [6]int{0, 0, 0, 1, 0, 0}},
	192: {[6]int{75, 75, 55, 105, 85, 30}, [6]int{0, 0, 0, 2, 0, 0}},
	193: {[6]int{65, 65, 45, 75, 45, 95}, [6]int{0, 0, 0, 0, 0, 1}},
	194: {[6]int{55, 45, 45, 25, 25, 15}, [6]int{1, 0, 0, 0, 0, 0}},
	195: {[6]int{95, 85, 85, 65, 65, 35}, [6]int{2, 0, 0, 0, 0, 0}},
	196: {[6]int{65, 65, 60, 130, 95, 110}, [6]int{0, 0, 0, 2, 0, 0}},
	197: {[6]int{95, 65, 110, 60, 130, 65}, [6]int{0, 0, 0, 0, 2, 0}},
	198: {[6]int{60, 85, 42, 85, 42, 91}, [6]int{0, 0, 0, 0, 0, 1}},
	199: {[6]int{95, 75, 80, 100, 110, 30}, [6]int{0, 0, 0, 0, 3, 0}},
	200: {[6]int{60, 60, 60, 85, 85, 85}, [6]int{0, 0, 0, 0, 1, 0}},
	201: {[6]int{48, 72, 48, 72, 48, 48}, [6]int{0, 1, 0, 1, 0, 0}},
	202: {[6]int{190, 33, 58, 33, 58, 33}, [6]int{2, 0, 0, 0, 0, 0}},
	203: {[6]int{70, 80, 65, 90, 65, 85}, [6]int{0, 0, 0, 2, 0, 0}},
	204: {[6]int{50, 65, 90, 35, 35, 15}, [6]int{0, 0, 1, 0, 0, 0}},
	205: {[6]int{75, 90, 140, 60, 60, 40}, [6]int{0, 0, 2, 0, 0, 0}},
	206: {[6]int{100, 70, 70, 65, 65, 45}, [6]int{1, 0, 0, 0, 0, 0}},
	207: {[6]int{65, 75, 105, 35, 65, 85}, [6]int{0, 0, 1, 0, 0, 0}},
	208: {[6]int{75, 85, 200, 55, 65, 30}, [6]int{0, 0, 2, 0, 0, 0}},
	209: {[6]int{60, 80, 50, 40, 40, 30}, [6]int{0, 1, 0, 0, 0, 0}},
	210: {[6]int{90, 120, 75, 60, 60, 45}, [6]int{0, 2, 0, 0, 0, 0}},
	211: {[6]int{65, 95, 85, 55, 55, 85}, [6]int{0, 1, 0, 0, 0, 0}},
	212: {[6]int{70, 130, 100, 55, 80, 65}, [6]int{0, 2, 0, 0, 0, 0}},
	213: {[6]int{20, 10, 230, 10, 230, 5}, [6]int{0, 0, 1, 0, 1, 0}},
	214: {[6]int{80, 125, 75, 40, 95, 85}, [6]int{0, 2, 0, 0, 0, 0}},
	215: {[6]int{55, 95, 55, 35, 75, 115}, [6]int{0, 0, 0, 0, 0, 1}},
	216: {[6]int{60, 80, 50, 50, 50, 40}, [6]int{0, 1, 0, 0, 0, 0}},
	217: {[6]int{90, 130, 75, 75, 75, 55}, [6]int{0, 2, 0, 0, 0, 0}},
	218: {[6]int{40, 40, 40, 70, 40, 20}, [6]int{0, 0, 0, 1, 0, 0}},
	219: {[6]int{60, 50, 120, 90, 80, 30}, [6]int{0, 0, 2, 0, 0, 0}},
	220: {[6]int{50, 50, 40, 30, 30, 50}, [6]int{0, 1, 0, 0, 0, 0}},
	221: {[6]int{100, 100, 80, 60, 60, 50}, [6]int{1, 1, 0, 0, 0, 0}},
	222: {[6]int{65, 55, 95, 65, 95, 35}, [6]int{0, 0, 1, 0, 1, 0}},
	223: {[6]int{35, 65, 35, 65, 35, 65}, [6]int{0, 0, 0, 1, 0, 0}},
	224: {[6]int{75, 105, 75, 105, 75, 45}, [6]int{0, 1, 0, 1, 0, 0}},
	225: {[6]int{45, 55, 45, 65, 45, 75}, [6]int{0, 0, 0, 0, 0, 1}},
	226: {[6]int{85, 40, 70, 80, 140, 70}, [6]int{0, 0, 0, 0, 2, 0}},
	227: {[6]int{65, 80, 140, 40, 70, 70}, [6]int{0, 0, 2, 0, 0, 0}},
	228: {[6]int{45, 60, 30, 80, 50, 65}, [6]int{0, 0, 0, 1, 0, 0}},
	229: {[6]int{75, 90, 50, 110, 80, 95}, [6]int{0, 0, 0, 2, 0, 0}},
	230: {[6]int{75, 95, 95, 95, 95, 85}, [6]int{0, 1, 0, 1, 1, 0}},
	231: {[6]int{90, 60, 60, 40, 40, 40}, [6]int{1, 0, 0, 0, 0, 0}},
	232: {[6]int{90, 120, 120, 60, 60, 50}, [6]int{0, 1, 1, 0, 0, 0}},
	233: {[6]int{85, 80, 90, 105, 95, 60}, [6]int{0, 0, 0, 2, 0, 0}},
	234: {[6]int{73, 95, 62, 85, 65, 85}, [6]int{0, 1, 0, 0, 0, 0}},
	235: {[6]int{55, 20, 35, 20, 45, 75}, [6]int{0, 0, 0, 0, 0, 1}},
	236: {[6]int{35, 35, 35, 35, 35, 35}, [6]int{0, 1, 0, 0, 0, 0}},
	237: {[6]int{50, 95, 95, 35, 110, 70}, [6]int{0, 0, 0, 0, 2, 0}},
	238: {[6]int{45, 30, 15, 85, 65, 65}, [6]int{0, 0, 0, 1, 0, 0}},
	239: {[6]int{45, 63, 37, 65, 55, 95}, [6]int{0, 0, 0, 0, 0, 1}},
	240: {[6]int{45, 75, 37, 70, 55, 83}, [6]int{0, 0, 0, 0, 0, 1}},
	241: {[6]int{95, 80, 105, 40, 70, 100}, [6]int{0, 0, 2, 0, 0, 0}},
	242: {[6]int{255, 10, 10, 75, 135, 55}, [6]int{3, 0, 0, 0, 0, 0}},
	243: {[6]int{90, 85, 75, 115, 100, 115}, [6]int{0, 0, 0, 1, 0, 2}},
	244: {[6]int{115, 115, 85, 90, 75, 100}, [6]int{1, 2, 0, 0, 0, 0}},
	245: {[6]int{100, 75, 115, 90, 115, 85}, [6]int{0, 0, 1, 0, 2, 0}},
	246: {[6]int{50, 64, 50, 45, 50, 41}, [6]int{0, 1, 0, 0, 0, 0}},
	247: {[6]int{70, 84, 70, 65, 70, 51}, [6]int{0, 2, 0, 0, 0, 0}},
	248: {[6]int{100, 134, 110, 95, 100, 61}, [6]int{0, 3, 0, 0, 0, 0}},
	249: {[6]int{106, 90, 130, 90, 154, 110}, [6]int{0, 0, 0, 0, 3, 0}},
	250: {[6]int{106, 130, 90, 110, 154, 90}, [6]int{0, 0, 0, 0, 3, 0}},
	251: {[6]int{100, 100, 100, 100, 100, 100}, [6]int{3, 0, 0, 0, 0, 0}},
	252: {[6]int{40, 45, 35, 65, 55, 70}, [6]int{0, 0, 0, 0, 0, 1}},
	253: {[6]int{50, 65, 45, 85, 65, 95}, [6]int{0, 0, 0, 0, 0, 2}},
	254: {[6]int{70, 85, 65, 105, 85, 120}, [6]int{0, 0, 0, 0, 0, 3}},
	255: {[6]int{45, 60, 40, 70, 50, 45}, [6]int{0, 0, 0, 1, 0, 0}},
	256: {[6]int{60, 85, 60, 85, 60, 55}, [6]int{0, 1, 0, 1, 0, 0}},
	257: {[6]int{80, 120, 70, 110, 70, 80}, [6]int{0, 3, 0, 0, 0, 0}},
	258: {[6]int{50, 70, 50, 50, 50, 40}, [6]int{0, 1, 0, 0, 0, 0}},
	259: {[6]int{70, 85, 70, 60, 70, 50}, [6]int{0, 2, 0, 0, 0, 0}},
	260: {[6]int{100, 110, 90, 85, 90, 60}, [6]int{0, 3, 0, 0, 0, 0}},
	261: {[6]int{35, 55, 35, 30, 30, 35}, [6]int{0, 1, 0, 0, 0, 0}},
	262: {[6]int{70, 90, 70, 60, 60, 70}, [6]int{0, 2, 0, 0, 0, 0}},
	263: {[6]int{38, 30, 41, 30, 41, 60}, [6]int{0, 0, 0, 0, 0, 1}},
	264: {[6]int{78, 70, 61, 50, 61, 100}, [6]int{0, 0, 0, 0, 0, 2}},
	265: {[6]int{45, 45, 35, 20, 30, 20}, [6]int{1, 0, 0, 0, 0, 0}},
	266: {[6]int{50, 35, 55, 25, 25, 15}, [6]int{0, 0, 2, 0, 0, 0}},
	267: {[6]int{60, 70, 50, 100, 50, 65}, [6]int{0, 0, 0, 3, 0, 0}},
	268: {[6]int{50, 35, 55, 25, 25, 15}, [6]int{0, 0, 2, 0, 0, 0}},
	269: {[6]int{60, 50, 70, 50, 90, 65}, [6]int{0, 0, 0, 0, 3, 0}},
	270: {[6]int{40, 30, 30, 40, 50, 30}, [6]int{0, 0, 0, 0, 1, 0}},
	271: {[6]int{60, 50, 50, 60, 70, 50}, [6]int{0, 0, 0, 0, 2, 0}},
	272: {[6]int{80, 70, 70, 90, 100, 70}, [6]int{0, 0, 0, 0, 3, 0}},
	273: {[6]int{40, 40, 50, 30, 30, 30}, [6]int{0, 0, 1, 0, 0, 0}},
	274: {[6]int{70, 70, 40, 60, 40, 60}, [6]int{0, 2, 0, 0, 0, 0}},
	275: {[6]int{90, 100, 60, 90, 60, 80}, [6]int{0, 3, 0, 0, 0, 0}},
	276: {[6]int{40, 55, 30, 30, 30, 85}, [6]int{0, 0, 0, 0, 0, 1}},
	277: {[6]int{60, 85, 60, 75, 50, 125}, [6]int{0, 0, 0, 0, 0, 2}},
	278: {[6]int{40, 30, 30, 55, 30, 85}, [6]int{0, 0, 0, 0, 0, 1}},
	279: {[6]int{60, 50, 100, 95, 70, 65}, [6]int{0, 0, 2, 0, 0, 0}},
	280: {[6]int{28, 25, 25, 45, 35, 40}, [6]int{0, 0, 0, 1, 0, 0}},
	281: {[6]int{38, 35, 35, 65, 55, 50}, [6]int{0, 0, 0, 2, 0, 0}},
	282: {[6]int{68, 65, 65, 125, 115, 80}, [6]int{0, 0, 0, 3, 0, 0}},
	283: {[6]int{40, 30, 32, 50, 52, 65}, [6]int{0, 0, 0, 0, 0, 1}},
	284: {[6]int{70, 60, 62, 100, 82, 80}, [6]int{0, 0, 0, 1, 1, 0}},
	285: {[6]int{60, 40, 60, 40, 60, 35}, [6]int{1, 0, 0, 0, 0, 0}},
	286: {[6]int{60, 130, 80, 60, 60, 70}, [6]int{0, 2, 0, 0, 0, 0}},
	287: {[6]int{60, 60, 60, 35, 35, 30}, [6]int{1, 0, 0, 0, 0, 0}},
	288: {[6]int{80, 80, 80, 55, 55, 90}, [6]int{0, 0, 0, 0, 0, 2}},
	289: {[6]int{150, 160, 100, 95, 65, 100}, [6]int{3, 0, 0, 0, 0, 0}},
	290: {[6]int{31, 45, 90, 30, 30, 40}, [6]int{0, 0, 1, 0, 0, 0}},
	291: {[6]int{61, 90, 45, 50, 50, 160}, [6]int{0, 0, 0, 0, 0, 2}},
	292: {[6]int{1, 90, 45, 30, 30, 40}, [6]int{2, 0, 0, 0, 0, 0}},
	293: {[6]int{64, 51, 23, 51, 23, 28}, [6]int{1, 0, 0, 0, 0, 0}},
	294: {[6]int{84, 71, 43, 71, 43, 48}, [6]int{2, 0, 0, 0, 0, 0}},
	295: {[6]int{104, 91, 63, 91, 73, 68}, [6]int{3, 0, 0, 0, 0, 0}},
	296: {[6]int{72, 60, 30, 20, 30, 25}, [6]int{1, 0, 0, 0, 0, 0}},
	297: {[6]int{144, 120, 60, 40, 60, 50}, [6]int{2, 0, 0, 0, 0, 0}},
	298: {[6]int{50, 20, 40, 20, 40, 20}, [6]int{1, 0, 0, 0, 0, 0}},
	299: {[6]int{30, 45, 135, 45, 90, 30}, [6]int{0, 0, 1, 0, 0, 0}},
	300: {[6]int{50, 45, 45, 35, 35, 50}, [6]int{0, 0, 0, 0, 0, 1}},
	301: {[6]int{70, 65, 65, 55, 55, 90}, [6]int{1, 0, 0, 0, 0, 1}},
	302: {[6]int{50, 75, 75, 65, 65, 50}, [6]int{0, 1, 1, 0, 0, 0}},
	303: {[6]int{50, 85, 85, 55, 55, 50}, [6]int{0, 1, 1, 0, 0, 0}},
	304: {[6]int{50, 70, 100, 40, 40, 30}, [6]int{0, 0, 1, 0, 0, 0}},
	305: {[6]int{60, 90, 140, 50, 50, 40}, [6]int{0, 0, 2, 0, 0, 0}},
	306: {[6]int{70, 110, 180, 60, 60, 50}, [6]int{0, 0, 3, 0, 0, 0}},
	307: {[6]int{30, 40, 55, 40, 55, 60}, [6]int{0, 0, 0, 0, 0, 1}},
	308: {[6]int{60, 60, 75, 60, 75, 80}, [6]int{0, 0, 0, 0, 0, 2}},
	309: {[6]int{40, 45, 40, 65, 40, 65}, [6]int{0, 0, 0, 0, 0, 1}},
	310: {[6]int{70, 75, 60, 105, 60, 105}, [6]int{0, 0, 0, 0, 0, 2}},
	311: {[6]int{60, 50, 40, 85, 75, 95}, [6]int{0, 0, 0, 0, 0, 1}},
	312: {[6]int{60, 40, 50, 75, 85, 95}, [6]int{0, 0, 0, 0, 0, 1}},
	313: {[6]int{65, 73, 75, 47, 85, 85}, [6]int{0, 0, 0, 0, 0, 1}},
	314: {[6]int{65, 47, 75, 73, 85, 85}, [6]int{0, 0, 0, 0, 0, 1}},
	315: {[6]int{50, 60, 45, 100, 80, 65}, [6]int{0, 0, 0, 2, 0, 0}},
	316: {[6]int{70, 43, 53, 43, 53, 40}, [6]int{1, 0, 0, 0, 0, 0}},
	317: {[6]int{100, 73, 83, 73, 83, 55}, [6]int{2, 0, 0, 0, 0, 0}},
	318: {[6]int{45, 90, 20, 65, 20, 65}, [6]int{0, 1, 0, 0, 0, 0}},
	319: {[6]int{70, 120, 40, 95, 40, 95}, [6]int{0, 2, 0, 0, 0, 0}},
	320: {[6]int{130, 70, 35, 70, 35, 60}, [6]int{1, 0, 0, 0, 0, 0}},
	321: {[6]int{170, 90, 45, 90, 45, 60}, [6]int{2, 0, 0, 0, 0, 0}},
	322: {[6]int{60, 60, 40, 65, 45, 35}, [6]int{0, 0, 0, 1, 0, 0}},
	323: {[6]int{70, 100, 70, 105, 75, 40}, [6]int{0, 1, 0, 1, 0, 0}},
	324: {[6]int{70, 85, 140, 85, 70, 20}, [6]int{0, 0, 2, 0, 0, 0}},
	325: {[6]int{60, 25, 35, 70, 80, 60}, [6]int{0, 0, 0, 0, 1, 0}},
	326: {[6]int{80, 45, 65, 90, 110, 80}, [6]int{0, 0, 0, 0, 2, 0}},
	327: {[6]int{60, 60, 60, 60, 60, 60}, [6]int{0, 0, 0, 1, 0, 0}},
	328: {[6]int{45, 100, 45, 45, 45, 10}, [6]int{0, 1, 0, 0, 0, 0}},
	329: {[6]int{50, 70, 50, 50, 50, 70}, [6]int{0, 1, 0, 0, 0, 1}},
	330: {[6]int{80, 100, 80, 80, 80, 100}, [6]int{0, 1, 0, 0, 0, 2}},
	331: {[6]int{50, 85, 40, 85, 40, 35}, [6]int{0, 0, 0, 1, 0, 0}},
	332: {[6]int{70, 115, 60, 115, 60, 55}, [6]int{0, 1, 0, 1, 0, 0}},
	333: {[6]int{45, 40, 60, 40, 75, 50}, [6]int{0, 0, 0, 0, 1, 0}},
	334: {[6]int{75, 70, 90, 70, 105, 80}, [6]int{0, 0, 0, 0, 2, 0}},
	335: {[6]int{73, 115, 60, 60, 60, 90}, [6]int{0, 2, 0, 0, 0, 0}},
	336: {[6]int{73, 100, 60, 100, 60, 65}, [6]int{0, 1, 0, 1, 0, 0}},
	337: {[6]int{90, 55, 65, 95, 85, 70}, [6]int{0, 0, 0, 2, 0, 0}},
	338: {[6]int{90, 95, 85, 55, 65, 70}, [6]int{0, 2, 0, 0, 0, 0}},
	339: {[6]int{50, 48, 43, 46, 41, 60}, [6]int{1, 0, 0, 0, 0, 0}},
	340: {[6]int{110, 78, 73, 76, 71, 60}, [6]int{2, 0, 0, 0, 0, 0}},
	341: {[6]int{43, 80, 65, 50, 35, 35}, [6]int{0, 1, 0, 0, 0, 0}},
	342: {[6]int{63, 120, 85, 90, 55, 55}, [6]int{0, 2, 0, 0, 0, 0}},
	343: {[6]int{40, 40, 55, 40, 70, 55}, [6]int{0, 0, 0, 0, 1, 0}},
	344: {[6]int{60, 70, 105, 70, 120, 75}, [6]int{0, 0, 0, 0, 2, 0}},
	345: {[6]int{66, 41, 77, 61, 87, 23}, [6]int{0, 0, 0, 0, 1, 0}},
	346: {[6]int{86, 81, 97, 81, 107, 43}, [6]int{0, 0, 0, 0, 2, 0}},
	347: {[6]int{45, 95, 50, 40, 50, 75}, [6]int{0, 1, 0, 0, 0, 0}},
	348: {[6]int{75, 125, 100, 70, 80, 45}, [6]int{0, 2, 0, 0, 0, 0}},
	349: {[6]int{20, 15, 20, 10, 55, 80}, [6]int{0, 0, 0, 0, 0, 1}},
	350: {[6]int{95, 60, 79, 100, 125, 81}, [6]int{0, 0, 0, 0, 2, 0}},
	351: {[6]int{70, 70, 70, 70, 70, 70}, [6]int{1, 0, 0, 0, 0, 0}},
	352: {[6]int{60, 90, 70, 60, 120, 40}, [6]int{0, 0, 0, 0, 1, 0}},
	353: {[6]int{44, 75, 35, 63, 33, 45}, [6]int{0, 1, 0, 0, 0, 0}},
	354: {[6]int{64, 115, 65, 83, 63, 65}, [6]int{0, 2, 0, 0, 0, 0}},
	355: {[6]int{20, 40, 90, 30, 90, 25}, [6]int{0, 0, 0, 0, 1, 0}},
	356: {[6]int{40, 70, 130, 60, 130, 25}, [6]int{0, 0, 1, 0, 1, 0}},
	357: {[6]int{99, 68, 83, 72, 87, 51}, [6]int{2, 0, 0, 0, 0, 0}},
	358: {[6]int{75, 50, 80, 95, 90, 65}, [6]int{0, 0, 0, 1, 1, 0}},
	359: {[6]int{65, 130, 60, 75, 60, 75}, [6]int{0, 2, 0, 0, 0, 0}},
	360: {[6]int{95, 23, 48, 23, 48, 23}, [6]int{1, 0, 0, 0, 0, 0}},
	361: {[6]int{50, 50, 50, 50, 50, 50}, [6]int{1, 0, 0, 0, 0, 0}},
	362: {[6]int{80, 80, 80, 80, 80, 80}, [6]int{2, 0, 0, 0, 0, 0}},
	363: {[6]int{70, 40, 50, 55, 50, 25}, [6]int{1, 0, 0, 0, 0, 0}},
	364: {[6]int{90, 60, 70, 75, 70, 45}, [6]int{2, 0, 0, 0, 0, 0}},
	365: {[6]int{110, 80, 90, 95, 90, 65}, [6]int{3, 0, 0, 0, 0, 0}},
	366: {[6]int{35, 64, 85, 74, 55, 32}, [6]int{0, 0, 1, 0, 0, 0}},
	367: {[6]int{55, 104, 105, 94, 75, 52}, [6]int{0, 1, 1, 0, 0, 0}},
	368: {[6]int{55, 84, 105, 114, 75, 52}, [6]int{0, 0, 0, 2, 0, 0}},
	369: {[6]int{100, 90, 130, 45, 65, 55}, [6]int{1, 0, 1, 0, 0, 0}},
	370: {[6]int{43, 30, 55, 40, 65, 97}, [6]int{0, 0, 0, 0, 0, 1}},
	371: {[6]int{45, 75, 60, 40, 30, 50}, [6]int{0, 1, 0, 0, 0, 0}},
	372: {[6]int{65, 95, 100, 60, 50, 50}, [6]int{0, 0, 2, 0, 0, 0}},
	373: {[6]int{95, 135, 80, 110, 80, 100}, [6]int{0, 3, 0, 0, 0, 0}},
	374: {[6]int{40, 55, 80, 35, 60, 30}, [6]int{0, 0, 1, 0, 0, 0}},
	375: {[6]int{60, 75, 100, 55, 80, 50}, [6]int{0, 0, 2, 0, 0, 0}},
	376: {[6]int{80, 135, 130, 95, 90, 70}, [6]int{0, 0, 3, 0, 0, 0}},
	377: {[6]int{80, 100, 200, 50, 100, 50}, [6]int{0, 0, 3, 0, 0, 0}},
	378: {[6]int{80, 50, 100, 100, 200, 50}, [6]int{0, 0, 0, 0, 3, 0}},
	379: {[6]int{80, 75, 150, 75, 150, 50}, [6]int{0, 0, 2, 0, 1, 0}},
	380: {[6]int{80, 80, 90, 110, 130, 110}, [6]int{0, 0, 0, 0, 3, 0}},
	381: {[6]int{80, 90, 80, 130, 110, 110}, [6]int{0, 0, 0, 3, 0, 0}},
	382: {[6]int{100, 100, 90, 150, 140, 90}, [6]int{0, 0, 0, 3, 0, 0}},
	383: {[6]int{100, 150, 140, 100, 90, 90}, [6]int{0, 3, 0, 0, 0, 0}},
	384: {[6]int{105, 150, 90, 150, 90, 95}, [6]int{0, 2, 0, 1, 0, 0}},
	385: {[6]int{100, 100, 100, 100, 100, 100}, [6]int{3, 0, 0, 0, 0, 0}},
	386: {[6]int{50, 150, 50, 150, 50, 150}, [6]int{0, 1, 0, 1, 0, 1}},
	387: {[6]int{55, 68, 64, 45, 55, 31}, [6]int{0, 1, 0, 0, 0, 0}},
	388: {[6]int{75, 89, 85, 55, 65, 36}, [6]int{0, 1, 1, 0, 0, 0}},
	389: {[6]int{95, 109, 105, 75, 85, 56}, [6]int{0, 2, 1, 0, 0, 0}},
	390: {[6]int{44, 58, 44, 58, 44, 61}, [6]int{0, 0, 0, 0, 0, 1}},
	391: {[6]int{64, 78, 52, 78, 52, 81}, [6]int{0, 0, 0, 1, 0, 1}},
	392: {[6]int{76, 104, 71, 104, 71, 108}, [6]int{0, 1, 0, 1, 0, 1}},
	393: {[6]int{53, 51, 53, 61, 56, 40}, [6]int{0, 0, 0, 1, 0, 0}},
	394: {[6]int{64, 66, 68, 81, 76, 50}, [6]int{0, 0, 0, 2, 0, 0}},
	395: {[6]int{84, 86, 88, 111, 101, 60}, [6]int{0, 0, 0, 3, 0, 0}},
	396: {[6]int{40, 55, 30, 30, 30, 60}, [6]int{0, 0, 0, 0, 0, 1}},
	397: {[6]int{55, 75, 50, 40, 40, 80}, [6]int{0, 0, 0, 0, 0, 2}},
	398: {[6]int{85, 120, 70, 50, 60, 100}, [6]int{0, 3, 0, 0, 0, 0}},
	399: {[6]int{59, 45, 40, 35, 40, 31}, [6]int{1, 0, 0, 0, 0, 0}},
	400: {[6]int{79, 85, 60, 55, 60, 71}, [6]int{0, 2, 0, 0, 0, 0}},
	401: {[6]int{37, 25, 41, 25, 41, 25}, [6]int{0, 0, 1, 0, 0, 0}},
	402: {[6]int{77, 85, 51, 55, 51, 65}, [6]int{0, 2, 0, 0, 0, 0}},
	403: {[6]int{45, 65, 34, 40, 34, 45}, [6]int{0, 1, 0, 0, 0, 0}},
	404: {[6]int{60, 85, 49, 60, 49, 60}, [6]int{0, 2, 0, 0, 0, 0}},
	405: {[6]int{80, 120, 79, 95, 79, 70}, [6]int{0, 3, 0, 0, 0, 0}},
	406: {[6]int{40, 30, 35, 50, 70, 55}, [6]int{0, 0, 0, 1, 0, 0}},
	407: {[6]int{60, 70, 65, 125, 105, 90}, [6]int{0, 0, 0, 3, 0, 0}},
	408: {[6]int{67, 125, 40, 30, 30, 58}, [6]int{0, 1, 0, 0, 0, 0}},
	409: {[6]int{97, 165, 60, 65, 50, 58}, [6]int{0, 2, 0, 0, 0, 0}},
	410: {[6]int{30, 42, 118, 42, 88, 30}, [6]int{0, 0, 1, 0, 0, 0}},
	411: {[6]int{60, 52, 168, 47, 138, 30}, [6]int{0, 0, 2, 0, 0, 0}},
	412: {[6]int{40, 29, 45, 29, 45, 36}, [6]int{0, 0, 0, 0, 1, 0}},
	413: {[6]int{60, 59, 85, 79, 105, 36}, [6]int{0, 0, 0, 0, 2, 0}},
	414: {[6]int{70, 94, 50, 94, 50, 66}, [6]int{0, 1, 0, 1, 0, 0}},
	415: {[6]int{30, 30, 42, 30, 42, 70}, [6]int{0, 0, 0, 0, 0, 1}},
	416: {[6]int{70, 80, 102, 80, 102, 40}, [6]int{0, 0, 1, 0, 1, 0}},
	417: {[6]int{60, 45, 70, 45, 90, 95}, [6]int{0, 0, 0, 0, 0, 1}},
	418: {[6]int{55, 65, 35, 60, 30, 85}, [6]int{0, 0, 0, 0, 0, 1}},
	419: {[6]int{85, 105, 55, 85, 50, 115}, [6]int{0, 0, 0, 0, 0, 2}},
	420: {[6]int{45, 35, 45, 62, 53, 35}, [6]int{0, 0, 0, 1, 0, 0}},
	421: {[6]int{70, 60, 70, 87, 78, 85}, [6]int{0, 0, 0, 2, 0, 0}},
	422: {[6]int{76, 48, 48, 57, 62, 34}, [6]int{1, 0, 0, 0, 0, 0}},
	423: {[6]int{111, 83, 68, 92, 82, 39}, [6]int{2, 0, 0, 0, 0, 0}},
	424: {[6]int{75, 100, 66, 60, 66, 115}, [6]int{0, 0, 0, 0, 0, 2}},
	425: {[6]int{90, 50, 34, 60, 44, 70}, [6]int{1, 0, 0, 0, 0, 0}},
	426: {[6]int{150, 80, 44, 90, 54, 80}, [6]int{2, 0, 0, 0, 0, 0}},
	427: {[6]int{55, 66, 44, 44, 56, 85}, [6]int{0, 0, 0, 0, 0, 1}},
	428: {[6]int{65, 76, 84, 54, 96, 105}, [6]int{0, 0, 0, 0, 0, 2}},
	429: {[6]int{60, 60, 60, 105, 105, 105}, [6]int{0, 0, 0, 1, 1, 0}},
	430: {[6]int{100, 125, 52, 105, 52, 71}, [6]int{0, 2, 0, 0, 0, 0}},
	431: {[6]int{49, 55, 42, 42, 37, 85}, [6]int{0, 0, 0, 0, 0, 1}},
	432: {[6]int{71, 82, 64, 64, 59, 112}, [6]int{0, 0, 0, 0, 0, 2}},
	433: {[6]int{45, 30, 50, 65, 50, 45}, [6]int{0, 0, 0, 1, 0, 0}},
	434: {[6]int{63, 63, 47, 41, 41, 74}, [6]int{0, 0, 0, 0, 0, 1}},
	435: {[6]int{103, 93, 67, 71, 61, 84}, [6]int{2, 0, 0, 0, 0, 0}},
	436: {[6]int{57, 24, 86, 24, 86, 23}, [6]int{0, 0, 1, 0, 0, 0}},
	437: {[6]int{67, 89, 116, 79, 116, 33}, [6]int{0, 0, 1, 0, 1, 0}},
	438: {[6]int{50, 80, 95, 10, 45, 10}, [6]int{0, 0, 1, 0, 0, 0}},
	439: {[6]int{20, 25, 45, 70, 90, 60}, [6]int{0, 0, 0, 0, 1, 0}},
	440: {[6]int{100, 5, 5, 15, 65, 30}, [6]int{1, 0, 0, 0, 0, 0}},
	441: {[6]int{76, 65, 45, 92, 42, 91}, [6]int{0, 1, 0, 0, 0, 0}},
	442: {[6]int{50, 92, 108, 92, 108, 35}, [6]int{0, 0, 1, 0, 1, 0}},
	443: {[6]int{58, 70, 45, 40, 45, 42}, [6]int{0, 1, 0, 0, 0, 0}},
	444: {[6]int{68, 90, 65, 50, 55, 82}, [6]int{0, 2, 0, 0, 0, 0}},
	445: {[6]int{108, 130, 95, 80, 85, 102}, [6]int{0, 3, 0, 0, 0, 0}},
	446: {[6]int{135, 85, 40, 40, 85, 5}, [6]int{1, 0, 0, 0, 0, 0}},
	447: {[6]int{40, 70, 40, 35, 40, 60}, [6]int{0, 1, 0, 0, 0, 0}},
	448: {[6]int{70, 110, 70, 115, 70, 90}, [6]int{0, 1, 0, 1, 0, 0}},
	449: {[6]int{68, 72, 78, 38, 42, 32}, [6]int{0, 0, 1, 0, 0, 0}},
	450: {[6]int{108, 112, 118, 68, 72, 47}, [6]int{0, 0, 2, 0, 0, 0}},
	451: {[6]int{40, 50, 90, 30, 55, 65}, [6]int{0, 0, 1, 0, 0, 0}},
	452: {[6]int{70, 90, 110, 60, 75, 95}, [6]int{0, 0, 2, 0, 0, 0}},
	453: {[6]int{48, 61, 40, 61, 40, 50}, [6]int{0, 1, 0, 0, 0, 0}},
	454: {[6]int{83, 106, 65, 86, 65, 85}, [6]int{0, 2, 0, 0, 0, 0}},
	455: {[6]int{74, 100, 72, 90, 72, 46}, [6]int{0, 2, 0, 0, 0, 0}},
	456: {[6]int{49, 49, 56, 49, 61, 66}, [6]int{0, 0, 0, 0, 0, 1}},
	457: {[6]int{69, 69, 76, 69, 86, 91}, [6]int{0, 0, 0, 0, 0, 2}},
	458: {[6]int{45, 20, 50, 60, 120, 50}, [6]int{0, 0, 0, 0, 1, 0}},
	459: {[6]int{60, 62, 50, 62, 60, 40}, [6]int{0, 1, 0, 0, 0, 0}},
	460: {[6]int{90, 92, 75, 92, 85, 60}, [6]int{0, 1, 0, 1, 0, 0}},
	461: {[6]int{70, 120, 65, 45, 85, 125}, [6]int{0, 1, 0, 0, 0, 1}},
	462: {[6]int{70, 70, 115, 130, 90, 60}, [6]int{0, 0, 0, 3, 0, 0}},
	463: {[6]int{110, 85, 95, 80, 95, 50}, [6]int{3, 0, 0, 0, 0, 0}},
	464: {[6]int{115, 140, 130, 55, 55, 40}, [6]int{0, 3, 0, 0, 0, 0}},
	465: {[6]int{100, 100, 125, 110, 50, 50}, [6]int{0, 0, 2, 0, 0, 0}},
	466: {[6]int{75, 123, 67, 95, 85, 95}, [6]int{0, 3, 0, 0, 0, 0}},
	467: {[6]int{75, 95, 67, 125, 95, 83}, [6]int{0, 0, 0, 3, 0, 0}},
	468: {[6]int{85, 50, 95, 120, 115, 80}, [6]int{0, 0, 0, 2, 1, 0}},
	469: {[6]int{86, 76, 86, 116, 56, 95}, [6]int{0, 2, 0, 0, 0, 0}},
	470: {[6]int{65, 110, 130, 60, 65, 95}, [6]int{0, 0, 2, 0, 0, 0}},
	471: {[6]int{65, 60, 110, 130, 95, 65}, [6]int{0, 0, 0, 2, 0, 0}},
	472: {[6]int{75, 95, 125, 45, 75, 95}, [6]int{0, 0, 2, 0, 0, 0}},
	473: {[6]int{110, 130, 80, 70, 60, 80}, [6]int{0, 3, 0, 0, 0, 0}},
	474: {[6]int{85, 80, 70, 135, 75, 90}, [6]int{0, 0, 0, 3, 0, 0}},
	475: {[6]int{68, 125, 65, 65, 115, 80}, [6]int{0, 3, 0, 0, 0, 0}},
	476: {[6]int{60, 55, 145, 75, 150, 40}, [6]int{0, 0, 1, 0, 2, 0}},
	477: {[6]int{45, 100, 135, 65, 135, 45}, [6]int{0, 0, 1, 0, 2, 0}},
	478: {[6]int{70, 80, 70, 80, 70, 110}, [6]int{0, 0, 0, 0, 0, 2}},
	479: {[6]int{50, 50, 77, 95, 77, 91}, [6]int{0, 0, 0, 1, 0, 1}},
	480: {[6]int{75, 75, 130, 75, 130, 95}, [6]int{0, 0, 2, 0, 1, 0}},
	481: {[6]int{80, 105, 105, 105, 105, 80}, [6]int{0, 1, 0, 1, 1, 0}},
	482: {[6]int{75, 125, 70, 125, 70, 115}, [6]int{0, 2, 0, 1, 0, 0}},
	483: {[6]int{100, 120, 120, 150, 100, 90}, [6]int{0, 0, 0, 3, 0, 0}},
	484: {[6]int{90, 120, 100, 150, 120, 100}, [6]int{0, 0, 0, 3, 0, 0}},
	485: {[6]int{91, 90, 106, 130, 106, 77}, [6]int{0, 0, 0, 3, 0, 0}},
	486: {[6]int{110, 160, 110, 80, 110, 100}, [6]int{0, 3, 0, 0, 0, 0}},
	487: {[6]int{150, 100, 120, 100, 120, 90}, [6]int{3, 0, 0, 0, 0, 0}},
	488: {[6]int{120, 70, 120, 75, 130, 85}, [6]int{0, 0, 0, 0, 3, 0}},
	489: {[6]int{80, 80, 80, 80, 80, 80}, [6]int{1, 0, 0, 0, 0, 0}},
	490: {[6]int{100, 100, 100, 100, 100, 100}, [6]int{3, 0, 0, 0, 0, 0}},
	491: {[6]int{70, 90, 90, 135, 90, 125}, [6]int{0, 0, 0, 2, 0, 1}},
	492: {[6]int{100, 100, 100, 100, 100, 100}, [6]int{3, 0, 0, 0, 0, 0}},
	493: {[6]int{120, 120, 120, 120, 120, 120}, [6]int{3, 0, 0, 0, 0, 0}},
}

const (
	NatureAdamant Nature = iota
	NatureBashful
	NatureBold
	NatureBrave
	NatureCalm
	NatureCareful
	NatureDocile
	NatureGentle
	NatureHardy
	NatureHasty
	NatureImpish
	NatureJolly
	NatureLax
	NatureLonely
	NatureMild
	NatureModest
	NatureNaive
	NatureNaughty
	NatureQuiet
	NatureQuirky
	NatureRash
	NatureRelaxed
	NatureSassy
	NatureSerious
	NatureTimid
)

// Get the stat modifiers that this nature gives.
func (n Nature) GetStatModifiers() (statUp, statDown int) {
	switch n {
	case NatureAdamant:
		return StatAtk, StatSpAtk
	case NatureBashful:
		return StatSpAtk, StatSpAtk
	case NatureBold:
		return StatDef, StatAtk
	case NatureBrave:
		return StatAtk, StatSpeed
	case NatureCalm:
		return StatSpDef, StatAtk
	case NatureCareful:
		return StatSpDef, StatSpAtk
	case NatureDocile:
		return StatDef, StatDef
	case NatureGentle:
		return StatSpDef, StatDef
	case NatureHardy:
		return StatAtk, StatAtk
	case NatureHasty:
		return StatSpeed, StatDef
	case NatureImpish:
		return StatDef, StatSpAtk
	case NatureJolly:
		return StatSpeed, StatSpAtk
	case NatureLax:
		return StatDef, StatSpDef
	case NatureLonely:
		return StatAtk, StatDef
	case NatureMild:
		return StatSpAtk, StatDef
	case NatureModest:
		return StatSpAtk, StatAtk
	case NatureNaive:
		return StatSpeed, StatSpDef
	case NatureNaughty:
		return StatAtk, StatSpDef
	case NatureQuiet:
		return StatSpAtk, StatSpeed
	case NatureQuirky:
		return StatSpDef, StatSpDef
	case NatureRash:
		return StatSpAtk, StatSpDef
	case NatureRelaxed:
		return StatDef, StatSpeed
	case NatureSassy:
		return StatSpDef, StatSpeed
	case NatureSerious:
		return StatSpeed, StatSpeed
	case NatureTimid:
		return StatSpeed, StatAtk
	}
	panic("Unknown nature")
}

// Get the string name of this Nature.
func (n Nature) String() string {
	switch n {
	case NatureAdamant:
		return "Adamant"
	case NatureBashful:
		return "Bashful"
	case NatureBold:
		return "Bold"
	case NatureBrave:
		return "Brave"
	case NatureCalm:
		return "Calm"
	case NatureCareful:
		return "Careful"
	case NatureDocile:
		return "Docile"
	case NatureGentle:
		return "Gentle"
	case NatureHardy:
		return "Hardy"
	case NatureHasty:
		return "Hasty"
	case NatureImpish:
		return "Impish"
	case NatureJolly:
		return "Jolly"
	case NatureLax:
		return "Lax"
	case NatureLonely:
		return "Lonely"
	case NatureMild:
		return "Mild"
	case NatureModest:
		return "Modest"
	case NatureNaive:
		return "Naive"
	case NatureNaughty:
		return "Naughty"
	case NatureQuiet:
		return "Quiet"
	case NatureQuirky:
		return "Quirky"
	case NatureRash:
		return "Rash"
	case NatureRelaxed:
		return "Relaxed"
	case NatureSassy:
		return "Sassy"
	case NatureSerious:
		return "Serious"
	case NatureTimid:
		return "Timid"
	}
	panic("Unknown nature")
}
