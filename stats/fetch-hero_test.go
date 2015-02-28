package stats

import (
	"testing"
)

func TestFetchHero(t *testing.T) {
	fakeServer := makeFakeServer(stubHeroData())
	careerUrl = fakeServer.URL + "/"
	hd, err := FetchHero("brackets-1829", 57645654)

	if err != nil {
		t.Error("FetchCareer function had an error ", err)
	}

	if hd.Id != 57645654 {
		t.Error("Failed to unmarshal Id")
	}

	if hd.Stats.Damage != 222259 {
		t.Error("Failed to unmarshal Stats")
	}

	if hd.LastUpdated != 1424627763 {
		t.Error("Failed to unmashal time")
	}
}

func stubHeroData() string {
	return `{
		"id": 57645654,
		"name": "YiiIII",
		"class": "barbarian",
		"gender": 0,
		"level": 70,
		"paragonLevel": 25,
		"hardcore": true,
		"seasonal": true,
		"seasonCreated": 2,
		"skills": {
			"active": [
				{
					"skill": {
						"slug": "frenzy",
						"name": "Frenzy",
						"icon": "barbarian_frenzy",
						"level": 11,
						"categorySlug": "primary",
						"tooltipUrl": "skill/barbarian/frenzy",
						"description": "Generate: 4 Fury per attack\r\n\r\nSwing for 220% weapon damage. Frenzy's attack speed increases by 15% for 4 seconds with each swing. This effect stacks up to 5 times.",
						"simpleDescription": "Generate: 4 Fury per attack\r\n\r\nEnter a frenzied state that increases your attack speed with each hit.",
						"skillCalcId": "W"
					},
					"rune": {
						"slug": "frenzy-a",
						"type": "a",
						"name": "Maniac",
						"level": 59,
						"description": "Each Frenzy effect also increases your damage by 2.5%.\r\n\r\nFrenzy's damage turns into Fire.",
						"simpleDescription": "Increase damage dealt while frenzied.",
						"tooltipParams": "rune/frenzy/a",
						"skillCalcId": "c",
						"order": 4
					}
				},
				{
					"skill": {
						"slug": "whirlwind",
						"name": "Whirlwind",
						"icon": "barbarian_whirlwind",
						"level": 20,
						"categorySlug": "secondary",
						"tooltipUrl": "skill/barbarian/whirlwind",
						"description": "Cost: 10 Fury\r\n\r\nDeliver multiple attacks to everything in your path for 340% weapon damage.\r\n\r\nWhile whirlwinding, you move at 100% movement speed.",
						"simpleDescription": "Cost: 10 Fury\r\n\r\nUnleash a series of whirling attacks damaging everything in your path.",
						"skillCalcId": "h"
					},
					"rune": {
						"slug": "whirlwind-d",
						"type": "d",
						"name": "Wind Shear",
						"level": 44,
						"description": "Gain 1 Fury for every enemy struck.\r\n\r\nWhirlwind's damage turns into Lightning.",
						"simpleDescription": "Gain Fury for every enemy struck.",
						"tooltipParams": "rune/whirlwind/d",
						"skillCalcId": "Y",
						"order": 3
					}
				},
				{
					"skill": {
						"slug": "ground-stomp",
						"name": "Ground Stomp",
						"icon": "barbarian_groundstomp",
						"level": 4,
						"categorySlug": "defensive",
						"tooltipUrl": "skill/barbarian/ground-stomp",
						"description": "Generate: 15 Fury\r\nCooldown: 12 seconds\r\n\r\nSmash the ground, stunning all enemies within 14 yards for 4 seconds.",
						"simpleDescription": "Generate: 15 Fury\r\nCooldown: 12 seconds\r\n\r\nSmash the ground and Stun nearby enemies for 4 seconds.",
						"skillCalcId": "Y"
					},
					"rune": {
						"slug": "ground-stomp-a",
						"type": "a",
						"name": "Trembling Stomp",
						"level": 28,
						"description": "Enemies in the area also take 575% weapon damage as Fire.",
						"simpleDescription": "Damage enemies affected by Ground Stomp.",
						"tooltipParams": "rune/ground-stomp/a",
						"skillCalcId": "b",
						"order": 2
					}
				},
				{
					"skill": {
						"slug": "furious-charge",
						"name": "Furious Charge",
						"icon": "barbarian_furiouscharge",
						"level": 21,
						"categorySlug": "might",
						"tooltipUrl": "skill/barbarian/furious-charge",
						"description": "Cost: 1 Charge\r\nGenerate: 15 Fury\r\n\r\nRush forward knocking back and dealing 600% weapon damage to enemies along your path.\r\n\r\nYou gain a charge every 10 seconds and can have up to 1 charge stored at a time.",
						"simpleDescription": "Cost: 1 Charge\r\nGenerate: 15 Fury\r\n\r\nCharge all enemies in your path and knock them back.\r\n\r\nYou gain a charge every 10 seconds and can have up to 1 charge stored at a time.",
						"skillCalcId": "S"
					},
					"rune": {
						"slug": "furious-charge-e",
						"type": "e",
						"name": "Merciless Assault",
						"level": 33,
						"description": "Recharge time is reduced by 2 seconds for every enemy hit. This effect can reduce the recharge time by up to 10 seconds.",
						"simpleDescription": "Each enemy charged reduces the recharge time of Furious Charge.",
						"tooltipParams": "rune/furious-charge/e",
						"skillCalcId": "Z",
						"order": 1
					}
				},
				{
					"skill": {
						"slug": "war-cry",
						"name": "War Cry",
						"icon": "x1_barbarian_warcry_v2",
						"level": 28,
						"categorySlug": "tactics",
						"tooltipUrl": "skill/barbarian/war-cry",
						"description": "Generate: 20 Fury\r\nCooldown: 20 seconds\r\n\r\nUnleash a rallying cry to increase Armor for you and all allies within 100 yards by 20% for 60 seconds.",
						"simpleDescription": "Generate: 20 Fury\r\nCooldown: 20 seconds\r\n\r\nA rallying cry increases Armor for you and your allies for 60 seconds.",
						"skillCalcId": "k"
					},
					"rune": {
						"slug": "war-cry-e",
						"type": "e",
						"name": "Invigorate",
						"level": 41,
						"description": "Increase maximum Life by 10% and Life regeneration by 8315 per second while affected by War Cry.",
						"simpleDescription": "Increase your maximum Life and Life regeneration while affected by War Cry.",
						"tooltipParams": "rune/war-cry/e",
						"skillCalcId": "b",
						"order": 2
					}
				},
				{
					"skill": {
						"slug": "call-of-the-ancients",
						"name": "Call of the Ancients",
						"icon": "barbarian_calloftheancients",
						"level": 25,
						"categorySlug": "rage",
						"tooltipUrl": "skill/barbarian/call-of-the-ancients",
						"description": "Cooldown: 120 seconds\r\n\r\nSummon the ancient Barbarians Talic, Korlic, and Madawc to fight alongside you for 20 seconds. Each deals 270% weapon damage per swing in addition to bonus abilities.\r\n\r\n Talic wields a sword and shield and uses Whirlwind and Leap Attack.\r\n Korlic wields a massive polearm and uses Cleave and Furious Charge.\r\n Madawc dual-wields axes and uses Weapon Throw and Seismic Slam.",
						"simpleDescription": "Cooldown: 120 seconds\r\n\r\nSummon three ancestral heroes to fight alongside you for 20 seconds.",
						"skillCalcId": "j"
					},
					"rune": {
						"slug": "call-of-the-ancients-d",
						"type": "d",
						"name": "Duty to the Clan",
						"level": 37,
						"description": "Enemies hit by the Ancients are Chilled for 2 seconds and have 10% increased chance to be Critically Hit.\r\n\r\nThe Ancients' damage turns into Cold.",
						"simpleDescription": "Enemies hit by the Ancients are Chilled and take increased damage.",
						"tooltipParams": "rune/call-of-the-ancients/d",
						"skillCalcId": "Z",
						"order": 1
					}
				}
			],
			"passive": [
				{
					"skill": {
						"slug": "boon-of-bulkathos",
						"name": "Boon of Bul-Kathos",
						"icon": "barbarian_passive_boonofbulkathos",
						"level": 60,
						"tooltipUrl": "skill/barbarian/boon-of-bulkathos",
						"description": "Reduce the cooldowns of your:\r\n Earthquake by 15 seconds.\r\n Call of the Ancients by 30 seconds.\r\n Wrath of the Berserker by 30 seconds.",
						"flavor": "The aid of the Immortal King shall come to those who have earned it.",
						"skillCalcId": "S"
					}
				},
				{
					"skill": {
						"slug": "nerves-of-steel",
						"name": "Nerves of Steel",
						"icon": "barbarian_passive_nervesofsteel",
						"level": 13,
						"tooltipUrl": "skill/barbarian/nerves-of-steel",
						"description": "Increase your Armor by 50% of your Vitality.",
						"flavor": "\"The trials begin with childhood; skinning ferocious beasts, climbing windswept cliffs and carrying weapons heavy enough to make a southern soldier weep. Is it any wonder that they never give ground?\" —Sir Aric of Duncraig",
						"skillCalcId": "b"
					}
				},
				{
					"skill": {
						"slug": "tough-as-nails",
						"name": "Tough as Nails",
						"icon": "barbarian_passive_toughasnails",
						"level": 30,
						"tooltipUrl": "skill/barbarian/tough-as-nails",
						"description": "Increase Armor by 25%.\r\nIncrease Thorns damage dealt by 50%.",
						"flavor": "After a while the scars serve as their own layer of armor.",
						"skillCalcId": "V"
					}
				},
				{
					"skill": {
						"slug": "weapons-master",
						"name": "Weapons Master",
						"icon": "barbarian_passive_weaponsmaster",
						"level": 16,
						"tooltipUrl": "skill/barbarian/weapons-master",
						"description": "Gain a bonus based on the weapon type of your main hand weapon:\r\nSwords/Daggers: 8% increased damage\r\nMaces/Axes: 5% Critical Hit Chance\r\nPolearms/Spears: 8% attack speed\r\nMighty Weapons: 2 Fury per hit",
						"flavor": "\"My men ambushed the traveler in the woods, but it was we who were surprised. Such keen senses, swift motions, flawless strikes. We never stood a chance.\" —Redjack the Bandit",
						"skillCalcId": "Y"
					}
				}
			]
		},
		"items": {
			"head": {
				"id": "Unique_Helm_Set_12_x1",
				"name": "Crown of the Invoker",
				"icon": "unique_helm_set_12_x1_demonhunter_male",
				"displayColor": "green",
				"tooltipParams": "item/CnQIkt-l6Q4SBwgEFbETQ4kdgYHGxB1mIwZQHay801MdmwYAyx07J28kHXGLd_AwixI4vgJAAEgNUBJYBGC-AmorCgwIABDc8s6lgYCAwD0SGwiZ2JmfDRIHCAQVgyeUsTCPEjgAQAFYBJABAoABRrUBf_lOXRidsefrDlACWACgAZ2x5-sO",
				"craftedBy": []
			},
			"torso": {
				"id": "ChestArmor_208",
				"name": "Baleful Plate",
				"icon": "chestarmor_208_barbarian_male",
				"displayColor": "yellow",
				"tooltipParams": "item/CtABCOmTnvwMEgcIBBUxIhlgHXMjBlAdBwVfMh1oa9k0HZKU3AgdcIt38B3WPL1JIgsIABXM_gEAGAogIDCJEjiTBUAASA1QEGCTBWorCgwIABChzu-tgYCAwCUSGwixnNyxDRIHCAQVFZ_YqzCLEjgAQAFYBJABAmorCgwIABDilb_VgoCAgBcSGwjh1KuVCxIHCAQVaPWLXzCPEjgAQAFYBJABAmorCgwIABD-ze-tgYCAwCUSGwjH87eZCBIHCAQVFZ_YqzCPEjgAQAFYBJABAhjiiMmBDFACWAA",
				"craftedBy": [
					{
						"id": "T12_Armor_Chest",
						"slug": "sovereign-ascended-armor",
						"name": "Sovereign Ascended Armor",
						"cost": 108000,
						"reagents": [
							{
								"quantity": 8,
								"item": {
									"id": "Crafting_AssortedParts_01",
									"name": "Reusable Parts",
									"icon": "crafting_assortedparts_01_demonhunter_male",
									"displayColor": "white",
									"tooltipParams": "item/reusable-parts"
								}
							},
							{
								"quantity": 6,
								"item": {
									"id": "Crafting_Magic_01",
									"name": "Arcane Dust",
									"icon": "crafting_magic_01_demonhunter_male",
									"displayColor": "blue",
									"tooltipParams": "item/arcane-dust"
								}
							},
							{
								"quantity": 3,
								"item": {
									"id": "Crafting_Rare_01",
									"name": "Veiled Crystal",
									"icon": "crafting_rare_01_demonhunter_male",
									"displayColor": "yellow",
									"tooltipParams": "item/veiled-crystal"
								}
							}
						],
						"itemProduced": {
							"id": "ChestArmor_208",
							"name": "Sovereign Ascended Armor",
							"icon": "chestarmor_208_demonhunter_male",
							"displayColor": "yellow",
							"tooltipParams": "recipe/sovereign-ascended-armor"
						}
					}
				]
			},
			"feet": {
				"id": "Unique_Boots_007_x1",
				"name": "Fire Walkers",
				"icon": "unique_boots_007_x1_demonhunter_male",
				"displayColor": "orange",
				"tooltipParams": "item/CkII3Jyqtw0SBwgEFfvkpnMd6-LkBR2pIDa6HeQTmN0dOydvJB0_8gnYMIsSONACQABIDVASWARg0AKAAUa1AaBFm38Y9pO3ug5QAlgA",
				"craftedBy": []
			},
			"hands": {
				"id": "Unique_Gloves_014_x1",
				"name": "Magefist",
				"icon": "unique_gloves_014_x1_demonhunter_male",
				"displayColor": "orange",
				"tooltipParams": "item/CkII6Ymf8wgSBwgEFaLQXCId6XAvux2aBgDLHZKU3AgdfnZVoR2AgcbEMIsSOMECQABIDVASWARgwQKAAUW1AchbkqQYxMLmpQRQAlgA",
				"craftedBy": []
			},
			"shoulders": {
				"id": "Unique_Shoulder_007_x1",
				"name": "Corruption",
				"icon": "unique_shoulder_007_x1_demonhunter_male",
				"displayColor": "orange",
				"tooltipParams": "item/Cj4IjPLJ2AYSBwgEFVpVM4sdV0rVkh3q1ZI_HSCNN1QdhAOh1h0bXc9EHa3TmTkwjxo4gAJAAEgUUBJYBGCAAhjLquXMDlACWAA",
				"recipe": {
					"id": "T12_Legendary_Unique_Shoulder_007_x1",
					"slug": "corruption",
					"name": "Corruption",
					"cost": 172800,
					"reagents": [
						{
							"quantity": 25,
							"item": {
								"id": "Crafting_AssortedParts_01",
								"name": "Reusable Parts",
								"icon": "crafting_assortedparts_01_demonhunter_male",
								"displayColor": "white",
								"tooltipParams": "item/reusable-parts"
							}
						},
						{
							"quantity": 12,
							"item": {
								"id": "Crafting_Magic_01",
								"name": "Arcane Dust",
								"icon": "crafting_magic_01_demonhunter_male",
								"displayColor": "blue",
								"tooltipParams": "item/arcane-dust"
							}
						},
						{
							"quantity": 7,
							"item": {
								"id": "Crafting_Rare_01",
								"name": "Veiled Crystal",
								"icon": "crafting_rare_01_demonhunter_male",
								"displayColor": "yellow",
								"tooltipParams": "item/veiled-crystal"
							}
						},
						{
							"quantity": 2,
							"item": {
								"id": "Crafting_Legendary_01",
								"name": "Forgotten Soul",
								"icon": "crafting_legendary_01_demonhunter_male",
								"displayColor": "orange",
								"tooltipParams": "item/forgotten-soul"
							}
						}
					],
					"itemProduced": {
						"id": "Unique_Shoulder_007_x1",
						"name": "Corruption",
						"icon": "unique_shoulder_007_x1_demonhunter_male",
						"displayColor": "orange",
						"tooltipParams": "recipe/corruption"
					}
				},
				"craftedBy": [
					{
						"id": "T12_Legendary_Unique_Shoulder_007_x1",
						"slug": "corruption",
						"name": "Corruption",
						"cost": 172800,
						"reagents": [
							{
								"quantity": 25,
								"item": {
									"id": "Crafting_AssortedParts_01",
									"name": "Reusable Parts",
									"icon": "crafting_assortedparts_01_demonhunter_male",
									"displayColor": "white",
									"tooltipParams": "item/reusable-parts"
								}
							},
							{
								"quantity": 12,
								"item": {
									"id": "Crafting_Magic_01",
									"name": "Arcane Dust",
									"icon": "crafting_magic_01_demonhunter_male",
									"displayColor": "blue",
									"tooltipParams": "item/arcane-dust"
								}
							},
							{
								"quantity": 7,
								"item": {
									"id": "Crafting_Rare_01",
									"name": "Veiled Crystal",
									"icon": "crafting_rare_01_demonhunter_male",
									"displayColor": "yellow",
									"tooltipParams": "item/veiled-crystal"
								}
							},
							{
								"quantity": 2,
								"item": {
									"id": "Crafting_Legendary_01",
									"name": "Forgotten Soul",
									"icon": "crafting_legendary_01_demonhunter_male",
									"displayColor": "orange",
									"tooltipParams": "item/forgotten-soul"
								}
							}
						],
						"itemProduced": {
							"id": "Unique_Shoulder_007_x1",
							"name": "Corruption",
							"icon": "unique_shoulder_007_x1_demonhunter_male",
							"displayColor": "orange",
							"tooltipParams": "recipe/corruption"
						}
					}
				]
			},
			"legs": {
				"id": "Unique_Pants_002_x1",
				"name": "Hammer Jammers",
				"icon": "unique_pants_002_x1_demonhunter_male",
				"displayColor": "orange",
				"tooltipParams": "item/CrQBCMy6j8gHEgcIBBW1x9bRHY-QylAd_GC7hx1DTkRRHevi5AUdqSA2uh0_8gnYMIsSOJgDQABQElgEYJgDaioKDAgAENeO4LuBgICAPhIaCI6m8ycSBwgEFRWf2KswixI4AEABWASQAQJqKwoMCAAQopDgu4GAgIA-EhsIv8mjwQoSBwgEFRWf2KswixI4AEABWASQAQKAAUalAfxgu4etATsnbyS1AZ-S1aW4Ab69x7QBwAEBGJXhyeIFUAJYAA",
				"craftedBy": []
			},
			"bracers": {
				"id": "Bracers_208",
				"name": "Garrison Sanctum",
				"icon": "bracers_208_demonhunter_male",
				"displayColor": "yellow",
				"tooltipParams": "item/CkIImPnH4A4SBwgEFdssyssdaGvZNB2SlNwIHQKZ58sdOXXajB3WPL1JIgsIARWKQgMAGCQgHjCJEjiYA0AAUA5gmAMYycyuwwZQAlgA",
				"craftedBy": [
					{
						"id": "T12_Armor_Bracers",
						"slug": "sovereign-ascended-bracers",
						"name": "Sovereign Ascended Bracers",
						"cost": 18360,
						"reagents": [
							{
								"quantity": 8,
								"item": {
									"id": "Crafting_AssortedParts_01",
									"name": "Reusable Parts",
									"icon": "crafting_assortedparts_01_demonhunter_male",
									"displayColor": "white",
									"tooltipParams": "item/reusable-parts"
								}
							},
							{
								"quantity": 60,
								"item": {
									"id": "Crafting_Magic_01",
									"name": "Arcane Dust",
									"icon": "crafting_magic_01_demonhunter_male",
									"displayColor": "blue",
									"tooltipParams": "item/arcane-dust"
								}
							},
							{
								"quantity": 3,
								"item": {
									"id": "Crafting_Rare_01",
									"name": "Veiled Crystal",
									"icon": "crafting_rare_01_demonhunter_male",
									"displayColor": "yellow",
									"tooltipParams": "item/veiled-crystal"
								}
							}
						],
						"itemProduced": {
							"id": "Bracers_208",
							"name": "Sovereign Ascended Bracers",
							"icon": "bracers_208_demonhunter_male",
							"displayColor": "yellow",
							"tooltipParams": "recipe/sovereign-ascended-bracers"
						}
					}
				]
			},
			"mainHand": {
				"id": "Unique_Sword_1H_007_x1",
				"name": "Sever",
				"icon": "unique_sword_1h_007_x1_demonhunter_male",
				"displayColor": "orange",
				"tooltipParams": "item/CjsI_qO3kgESBwgEFfvYKE4dgYHGxB3jb8yEHZ5Pg9wdaklp6jCLEjiiAkAAUBJYBGCiAoABRrUBm9nZ7xj248biCVACWAA",
				"craftedBy": []
			},
			"offHand": {
				"id": "Unique_Shield_011_x1",
				"name": "Wall of Man",
				"icon": "unique_shield_011_x1_demonhunter_male",
				"displayColor": "orange",
				"tooltipParams": "item/CjwI6YSQxwkSBwgEFSjjrQkdcxjs8B1-VrMuHZsGAMsdTT8Y9h2BgcbEHcStOVwwjxI4iQJAAFASWARgiQIY0IzvmgFQAlgA",
				"recipe": {
					"id": "T12_Legendary_Unique_Shield_011_x1",
					"slug": "wall-of-man",
					"name": "Wall of Man",
					"cost": 64800,
					"reagents": [
						{
							"quantity": 40,
							"item": {
								"id": "Crafting_AssortedParts_01",
								"name": "Reusable Parts",
								"icon": "crafting_assortedparts_01_demonhunter_male",
								"displayColor": "white",
								"tooltipParams": "item/reusable-parts"
							}
						},
						{
							"quantity": 38,
							"item": {
								"id": "Crafting_Magic_01",
								"name": "Arcane Dust",
								"icon": "crafting_magic_01_demonhunter_male",
								"displayColor": "blue",
								"tooltipParams": "item/arcane-dust"
							}
						},
						{
							"quantity": 30,
							"item": {
								"id": "Crafting_Rare_01",
								"name": "Veiled Crystal",
								"icon": "crafting_rare_01_demonhunter_male",
								"displayColor": "yellow",
								"tooltipParams": "item/veiled-crystal"
							}
						},
						{
							"quantity": 2,
							"item": {
								"id": "Crafting_Legendary_01",
								"name": "Forgotten Soul",
								"icon": "crafting_legendary_01_demonhunter_male",
								"displayColor": "orange",
								"tooltipParams": "item/forgotten-soul"
							}
						},
						{
							"quantity": 5,
							"item": {
								"id": "Crafting_Looted_Reagent_01",
								"name": "Death's Breath",
								"icon": "crafting_looted_reagent_01_demonhunter_male",
								"displayColor": "yellow",
								"tooltipParams": "item/deaths-breath"
							}
						}
					],
					"itemProduced": {
						"id": "Unique_Shield_011_x1",
						"name": "Wall of Man",
						"icon": "unique_shield_011_x1_demonhunter_male",
						"displayColor": "orange",
						"tooltipParams": "recipe/wall-of-man"
					}
				},
				"craftedBy": [
					{
						"id": "T12_Legendary_Unique_Shield_011_x1",
						"slug": "wall-of-man",
						"name": "Wall of Man",
						"cost": 64800,
						"reagents": [
							{
								"quantity": 40,
								"item": {
									"id": "Crafting_AssortedParts_01",
									"name": "Reusable Parts",
									"icon": "crafting_assortedparts_01_demonhunter_male",
									"displayColor": "white",
									"tooltipParams": "item/reusable-parts"
								}
							},
							{
								"quantity": 38,
								"item": {
									"id": "Crafting_Magic_01",
									"name": "Arcane Dust",
									"icon": "crafting_magic_01_demonhunter_male",
									"displayColor": "blue",
									"tooltipParams": "item/arcane-dust"
								}
							},
							{
								"quantity": 30,
								"item": {
									"id": "Crafting_Rare_01",
									"name": "Veiled Crystal",
									"icon": "crafting_rare_01_demonhunter_male",
									"displayColor": "yellow",
									"tooltipParams": "item/veiled-crystal"
								}
							},
							{
								"quantity": 2,
								"item": {
									"id": "Crafting_Legendary_01",
									"name": "Forgotten Soul",
									"icon": "crafting_legendary_01_demonhunter_male",
									"displayColor": "orange",
									"tooltipParams": "item/forgotten-soul"
								}
							},
							{
								"quantity": 5,
								"item": {
									"id": "Crafting_Looted_Reagent_01",
									"name": "Death's Breath",
									"icon": "crafting_looted_reagent_01_demonhunter_male",
									"displayColor": "yellow",
									"tooltipParams": "item/deaths-breath"
								}
							}
						],
						"itemProduced": {
							"id": "Unique_Shield_011_x1",
							"name": "Wall of Man",
							"icon": "unique_shield_011_x1_demonhunter_male",
							"displayColor": "orange",
							"tooltipParams": "recipe/wall-of-man"
						}
					}
				]
			},
			"waist": {
				"id": "P2_Unique_BarbBelt_006",
				"name": "The Undisputed Champion",
				"icon": "p2_unique_barbbelt_006_demonhunter_male",
				"displayColor": "orange",
				"tooltipParams": "item/ClYI3YCs1AQSBwgEFer1v3QdYZdTUB0KcJEpHSCNN1QdOydvJB2pIDa6MIsSOLUCQABQElgEYLUCgAFGpQEKcJEprQFDTkRRtQGV9-R9uAGvwLvCC8ABBBjU1__LDlACWAA",
				"craftedBy": []
			},
			"rightFinger": {
				"id": "Ring_25",
				"name": "Assault Ring",
				"icon": "ring_25_demonhunter_male",
				"displayColor": "yellow",
				"tooltipParams": "item/CkkI-eP1hwYSBwgEFTZVXUQdaGvZNB13B8T0HZavuewdSjwevR1RUwyyHavTmTkiCwgAFbj-AQAYAiAcMI8SOOUDQABQEFgEYOUDGOzU6sUIUAJYAA",
				"recipe": {
					"id": "T12_Jeweler_Ring",
					"slug": "sovereign-ring",
					"name": "Sovereign Ring",
					"cost": 10000,
					"reagents": [
						{
							"quantity": 2,
							"item": {
								"id": "Diamond_15",
								"name": "Marquise Diamond",
								"icon": "diamond_15_demonhunter_male",
								"displayColor": "blue",
								"tooltipParams": "recipe/marquise-diamond"
							}
						}
					],
					"itemProduced": {
						"id": "Ring_25",
						"name": "Sovereign Ring",
						"icon": "ring_25_demonhunter_male",
						"displayColor": "yellow",
						"tooltipParams": "recipe/sovereign-ring"
					}
				},
				"craftedBy": [
					{
						"id": "T12_Jeweler_Ring",
						"slug": "sovereign-ring",
						"name": "Sovereign Ring",
						"cost": 10000,
						"reagents": [
							{
								"quantity": 2,
								"item": {
									"id": "Diamond_15",
									"name": "Marquise Diamond",
									"icon": "diamond_15_demonhunter_male",
									"displayColor": "blue",
									"tooltipParams": "recipe/marquise-diamond"
								}
							}
						],
						"itemProduced": {
							"id": "Ring_25",
							"name": "Sovereign Ring",
							"icon": "ring_25_demonhunter_male",
							"displayColor": "yellow",
							"tooltipParams": "recipe/sovereign-ring"
						}
					}
				]
			},
			"leftFinger": {
				"id": "Ring_25",
				"name": "Eternity Spine",
				"icon": "ring_25_demonhunter_male",
				"displayColor": "yellow",
				"tooltipParams": "item/CnYIsuDFxgcSBwgEFTZVXUQdApnnyx1oa9k0HajWdfwdZiMGUB1tzYkeHWsHd84iCwgBFYlCAwAYKCAKMIkSOMIDQABQEGDCA2otCgwIABDk-M6lgYCAwD0SHQjR0YrtAxIHCAQVKxe0wTCLEjgAQABQElgEkAECGOzU6sUIUAJYAA",
				"craftedBy": [
					{
						"id": "T12_Jeweler_Ring",
						"slug": "sovereign-ring",
						"name": "Sovereign Ring",
						"cost": 10000,
						"reagents": [
							{
								"quantity": 2,
								"item": {
									"id": "Diamond_15",
									"name": "Marquise Diamond",
									"icon": "diamond_15_demonhunter_male",
									"displayColor": "blue",
									"tooltipParams": "recipe/marquise-diamond"
								}
							}
						],
						"itemProduced": {
							"id": "Ring_25",
							"name": "Sovereign Ring",
							"icon": "ring_25_demonhunter_male",
							"displayColor": "yellow",
							"tooltipParams": "recipe/sovereign-ring"
						}
					}
				]
			},
			"neck": {
				"id": "Amulet17",
				"name": "Prestige Demise",
				"icon": "amulet17_demonhunter_male",
				"displayColor": "yellow",
				"tooltipParams": "item/CkcIhZ_Z_QMSBwgEFbDFRGQdgIHGxB3A3IcbHZoGAMsdaUlp6h1PGawMHavTmTkiCwgBFWxCAwAYGiAeMIkSOPkCQABQEGD5AhjglqbEDFACWAA",
				"craftedBy": [
					{
						"id": "T12_Jeweler_Amulet",
						"slug": "sovereign-amulet",
						"name": "Sovereign Amulet",
						"cost": 10000,
						"reagents": [
							{
								"quantity": 1,
								"item": {
									"id": "Ruby_15",
									"name": "Marquise Ruby",
									"icon": "ruby_15_demonhunter_male",
									"displayColor": "blue",
									"tooltipParams": "recipe/marquise-ruby"
								}
							},
							{
								"quantity": 1,
								"item": {
									"id": "Emerald_15",
									"name": "Marquise Emerald",
									"icon": "emerald_15_demonhunter_male",
									"displayColor": "blue",
									"tooltipParams": "recipe/marquise-emerald"
								}
							},
							{
								"quantity": 1,
								"item": {
									"id": "Amethyst_15",
									"name": "Marquise Amethyst",
									"icon": "amethyst_15_demonhunter_male",
									"displayColor": "blue",
									"tooltipParams": "recipe/marquise-amethyst"
								}
							},
							{
								"quantity": 1,
								"item": {
									"id": "Topaz_15",
									"name": "Marquise Topaz",
									"icon": "topaz_15_demonhunter_male",
									"displayColor": "blue",
									"tooltipParams": "recipe/marquise-topaz"
								}
							}
						],
						"itemProduced": {
							"id": "Amulet17",
							"name": "Sovereign Amulet",
							"icon": "amulet17_demonhunter_male",
							"displayColor": "yellow",
							"tooltipParams": "recipe/sovereign-amulet"
						}
					}
				]
			}
		},
		"stats": {
			"life": 516716,
			"damage": 222259,
			"toughness": 8426970,
			"healing": 914194,
			"attackSpeed": 1.399999976158142,
			"armor": 12543,
			"strength": 7009,
			"dexterity": 77,
			"vitality": 5164,
			"intelligence": 77,
			"physicalResist": 665,
			"fireResist": 791,
			"coldResist": 791,
			"lightningResist": 792,
			"poisonResist": 665,
			"arcaneResist": 665,
			"critDamage": 2.55,
			"blockChance": 0.31,
			"blockAmountMin": 13663,
			"blockAmountMax": 17074,
			"damageIncrease": 0,
			"critChance": 0.14000000075,
			"damageReduction": 0,
			"thorns": 6213,
			"lifeSteal": 0,
			"lifePerKill": 0,
			"goldFind": 1.25,
			"magicFind": 0,
			"lifeOnHit": 28410,
			"primaryResource": 112,
			"secondaryResource": 0
		},
		"kills": {
			"elites": 1380
		},
		"dead": false,
		"last-updated": 1424627763
	}`
}
