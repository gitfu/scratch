If you have a lot of duplicate values in your JSON data, 

trying indexing it to reduce the size. 

The best way to explain is to do an example. 

Say we have data for school lunches, 

I'll do this part in Python
,so call it a dictionary.
```python
lunches={}
```

The key is the day of the month,

and the value is a list of foods 

available for lunch on that day.

Like this:

```python
# key : value
lunches ={ 7:["CheesePizza","PepperoniPizza","BuffaloChickenPizza"
,"ClassicChickenNuggets","ChickenPizzaQuesadilla","FreshSeasonalSalads"
,"ChickenClubonCiabatta","YogurtBasket","RefriedBeans","BakedSweetPotato","FrozenFruitCup"] }
```
When you have a bunch of days together, 
the duplicates are obvious.

```python
lunches={7:["CheesePizza",
"PepperoniPizza",
"BuffaloChickenPizza",
"ClassicChickenNuggets",
"ChickenPizzaQuesadilla",
"FreshSeasonalSalads",
"ChickenClubonCiabatta",
"YogurtBasket",
"RefriedBeans",
"BakedSweetPotato",
"FrozenFruitCup"],
8:["CheesePizza",
"PepperoniPizza",
"FourMeatPizza",
"Gwinnett'sBestBurger",
"ItalianTrio",
"FreshSeasonalSalads",
"DeliFreshSubs",
"SpinachDip&Chips",
"RealFruitSmoothies",
"GAGrownGreenBeans",
"TropicalFruitSalad"],
9:["CheesePizza",
"PepperoniPizza",
"BuffaloChickenPizza",
"DrumDay!",
"MashedPotatoBowls",
"FreshSeasonalSalads",
"ClassicChickenSalad&Saltines",
"YogurtParfait",
"SteamedCollardGreens",
"DriedFruit"],
10:["CheesePizza",
"PepperoniPizza",
"FlavortotheMaxSticks",
"MiniCornDogs",
"NEW!ChickenAsianBites",
"FreshSeasonalSalads",
"DeliFreshSubs",
"RealFruitSmoothies",
"SteamedCarrots",
"NEW!Minh'sEggRoll",
"Mango"],
11:["CheesePizza",
"PepperoniPizza",
"BuffaloChickenPizza",
"Chickenn'Waffles",
"PhillyChickenSub",
"FreshSeasonalSalads",
"Home-StyleCroissantSandwiches",
"CinnamonApples"],
14:["CheesePizza",
"PepperoniPizza",
"BuffaloChickenPizza",
"ChickenSpicySandwich",
"CheesyGrilledCheese",
"FreshSeasonalSalads",
"ChickenClubonCiabatta",
"YogurtBasket",
"BakedSweetPotato",
"TomatoSoup",
"FrozenFruitCup"],
15:["CheesePizza",
"PepperoniPizza",
"FourMeatPizza",
"BallParkHotDog",
"BBQPlate",
"FreshSeasonalSalads",
"DeliFreshSubs",
"RealFruitSmoothies",
"PotatoSpirals",
"Coleslaw",
"PineappleTidbits"],
16:["CheesePizza",
"PepperoniPizza",
"BuffaloChickenPizza",
"CrispyChickenTenders",
"AsianRiceBowls",
"FreshSeasonalSalads",
"SpicyChickenSalad&Saltines",
"YogurtParfait",
"Carrot&CelerySticks",
"DriedFruit"],
17:["CheesePizza",
"PepperoniPizza",
"FlavortotheMaxSticks",
"GrilledChickenSandwich",
"FiestaNachos",
"FreshSeasonalSalads",
"DeliFreshSubs",
"BlackBeanEmpanadas",
"RealFruitSmoothies",
"BlackBeans",
"Doritos",
"MandarinOranges"],
18:["CheesePizza",
"PepperoniPizza",
"BuffaloChickenPizza",
"NEW!FrenchToast&Sausage",
"NEW!SpaghettiMeatballBowl",
"FreshSeasonalSalads",
"Home-StyleCroissantSandwiches",
"TropicalFruitSalad"],
21:["CheesePizza",
"PepperoniPizza",
"BuffaloChickenPizza",
"BonelessWings",
"KoreanMeatballSub",
"FreshSeasonalSalads",
"ChickenClubonCiabatta",
"YogurtBasket",
"GAGrownGreenBeans",
"BakedSweetPotato",
"FrozenFruitCup"],
22:["CheesePizza",
"PepperoniPizza",
"FourMeatPizza",
"BaconCheeseburger",
"ChickenSoftTacos",
"FreshSeasonalSalads",
"DeliFreshSubs",
"MexiPizza",
"RealFruitSmoothies",
"NEW!HappyCorn",
"MexiRice",
"MexicanFruitCup"],
23:["CheesePizza",
"PepperoniPizza",
"BuffaloChickenPizza",
"BlackAngusSliderDogs",
"NEW!HawaiianPanini",
"FreshSeasonalSalads",
"FruityChickenSalad&Saltines",
"YogurtParfait",
"BakedBeans",
"CheesyBroccoli",
"DriedFruit"],
24:["CheesePizza",
"PepperoniPizza",
"FlavortotheMaxSticks",
"BBQChickenSandwichonCiabatta",
"LasagnaBolognese",
"FreshSeasonalSalads",
"DeliFreshSubs",
"VeggieLasagnaRoll",
"RealFruitSmoothies",
"SteamedCarrots"],
25:["CheesePizza",
"PepperoniPizza",
"BuffaloChickenPizza",
"Chickenn'Waffles",
"FreshSeasonalSalads",
"Home-StyleCroissantSandwiches",
"Plantains",
"CinnamonApples"]}
  ```
  
 
 To remove the duplicates, 
 
 make one list of only unique foods
 
 and a new dictionary with the same keys,
 
 but the values will be a list 
 
 of the index of the food in the unique food list 
 
 ```python

# New layout 

# unique food list 

unique =["CheesePizza","PepperoniPizza","BuffaloChickenPizza"
,"ClassicChickenNuggets","ChickenPizzaQuesadilla","FreshSeasonalSalads"
,"ChickenClubonCiabatta","YogurtBasket","RefriedBeans","BakedSweetPotato","FrozenFruitCup"]

idxlunches ={7:[0,1,2,3,4,5,6,7,8,9,10]}
```

Convert all of the lunches dictionary to the new format
```python
# list of unique foods 
unique=[]

# new dictionary 
idxlunches={}
# The total count of food types with duplicates
count=0
for k,v in lunches.items():
	templist=[]
  # for each food in the value list
	for i in v: 
		count+=1
		if i not in unique: unique.append(i)
		templist.append(unique.index(i))
	
	idxlunches[k]=templist

print("total items=",count)
print("unique items=",len(unique) )
```

the output
```python
total items= 155
unique items= 72
```
Over half of the foods are duplicated

Looking at character count excluding whitespace characters 

The older format with the lunches dictionary 
is 2997 characters

The new format with the unique list and idxlunches dictionary,
is only 1861 characters.

the new format is just 62% of the original data size.

##10% smaller would pique my interest.

##38% smaller makes me giddy everytime.

```python
>>> unique
['CheesePizza', 'PepperoniPizza', 'BuffaloChickenPizza', 'ClassicChickenNuggets', 'ChickenPizzaQuesadilla', 'FreshSeasonalSalads', 'ChickenClubonCiabatta', 'YogurtBasket', 'RefriedBeans', 'BakedSweetPotato', 'FrozenFruitCup', 'FourMeatPizza', "Gwinnett'sBestBurger", 'ItalianTrio', 'DeliFreshSubs', 'SpinachDip&Chips', 'RealFruitSmoothies', 'GAGrownGreenBeans', 'TropicalFruitSalad', 'DrumDay!', 'MashedPotatoBowls', 'ClassicChickenSalad&Saltines', 'YogurtParfait', 'SteamedCollardGreens', 'DriedFruit', 'FlavortotheMaxSticks', 'MiniCornDogs', 'NEW!ChickenAsianBites', 'SteamedCarrots', "NEW!Minh'sEggRoll", 'Mango', "Chickenn'Waffles", 'PhillyChickenSub', 'Home-StyleCroissantSandwiches', 'CinnamonApples', 'ChickenSpicySandwich', 'CheesyGrilledCheese', 'TomatoSoup', 'BallParkHotDog', 'BBQPlate', 'PotatoSpirals', 'Coleslaw', 'PineappleTidbits', 'CrispyChickenTenders', 'AsianRiceBowls', 'SpicyChickenSalad&Saltines', 'Carrot&CelerySticks', 'GrilledChickenSandwich', 'FiestaNachos', 'BlackBeanEmpanadas', 'BlackBeans', 'Doritos', 'MandarinOranges', 'NEW!FrenchToast&Sausage', 'NEW!SpaghettiMeatballBowl', 'BonelessWings', 'KoreanMeatballSub', 'BaconCheeseburger', 'ChickenSoftTacos', 'MexiPizza', 'NEW!HappyCorn', 'MexiRice', 'MexicanFruitCup', 'BlackAngusSliderDogs', 'NEW!HawaiianPanini', 'FruityChickenSalad&Saltines', 'BakedBeans', 'CheesyBroccoli', 'BBQChickenSandwichonCiabatta', 'LasagnaBolognese', 'VeggieLasagnaRoll', 'Plantains']

>>> idxlunches
{7: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10], 8: [0, 1, 11, 12, 13, 5, 14, 15, 16, 17, 18], 9: [0, 1, 2, 19, 20, 5, 21, 22, 23, 24], 10: [0, 1, 25, 26, 27, 5, 14, 16, 28, 29, 30], 11: [0, 1, 2, 31, 32, 5, 33, 34], 14: [0, 1, 2, 35, 36, 5, 6, 7, 9, 37, 10], 15: [0, 1, 11, 38, 39, 5, 14, 16, 40, 41, 42], 16: [0, 1, 2, 43, 44, 5, 45, 22, 46, 24], 17: [0, 1, 25, 47, 48, 5, 14, 49, 16, 50, 51, 52], 18: [0, 1, 2, 53, 54, 5, 33, 18], 21: [0, 1, 2, 55, 56, 5, 6, 7, 17, 9, 10], 22: [0, 1, 11, 57, 58, 5, 14, 59, 16, 60, 61, 62], 23: [0, 1, 2, 63, 64, 5, 65, 22, 66, 67, 24], 24: [0, 1, 25, 68, 69, 5, 14, 70, 16, 28], 25: [0, 1, 2, 31, 5, 33, 71, 34]}
```



to generate the js file

```python
>>> with open("menu.js","w+") as outfile:
        print("var unique=",json.dumps(unique),"\nvar idxlunches=",json.dumps(idxlunches),file=outfile)
```



In a webpage, I would do something like 
```js
function swap(idx){ 
	return unique[idx]
}


function swapDay(day){ 
	return idxlunches[day].map(swap)
}
```
and call today's menu with: 
```js
swapDay(new Date().getDate())
```
```
[ "CheesePizza", "PepperoniPizza", "BuffaloChickenPizza", "Chickenn'Waffles", "FreshSeasonalSalads", "Home-StyleCroissantSandwiches", "Plantains", "CinnamonApples" ]
```

This is just a short example.
I use this technique when I 
have ten of thousands of records to include in a web page.

An ordered collection such as a list or an array, 
is very fast, and even faster when accessed via index.

Doing a callback to the server everytime you need data, 
is just cheesey, and slow, and weak. 
Don't do that. 



