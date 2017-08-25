If you have a lot of duplicate values in your JSON data, 
trying indexing it to reduce the size. 
The best way to explain what I mean, is to do an example. 

Say we have data for school lunches, 
I'll do this part in Python,so call it a dictionary.

The key is the day of the month,
and the value is a list of foods 
available for lunch on that day.

Like this:

```python
# key : value
7:["CheesePizza","PepperoniPizza","BuffaloChickenPizza"
,"ClassicChickenNuggets","ChickenPizzaQuesadilla","FreshSeasonalSalads"
,"ChickenClubonCiabatta","YogurtBasket","RefriedBeans","BakedSweetPotato","FrozenFruitCup"] }
```
When you have a bunch of days together, the duplicates are obvious.
Like this:
```python
lunches={
7:["CheesePizza","PepperoniPizza","BuffaloChickenPizza","ClassicChickenNuggets","ChickenPizzaQuesadilla",
	"FreshSeasonalSalads","ChickenClubonCiabatta","YogurtBasket","RefriedBeans","BakedSweetPotato","FrozenFruitCup"],
8:["CheesePizza","PepperoniPizza","FourMeatPizza","Gwinnett'sBestBurger","ItalianTrio","FreshSeasonalSalads",
	"DeliFreshSubs","SpinachDip&Chips","RealFruitSmoothies","GAGrownGreenBeans","TropicalFruitSalad"],
9:["CheesePizza","PepperoniPizza","BuffaloChickenPizza","DrumDay!","MashedPotatoBowls","FreshSeasonalSalads",
"ClassicChickenSalad&Saltines","YogurtParfait","SteamedCollardGreens","DriedFruit"],
10:["CheesePizza","PepperoniPizza","FlavortotheMaxSticks","MiniCornDogs","NEW!ChickenAsianBites",
"FreshSeasonalSalads","DeliFreshSubs","RealFruitSmoothies","SteamedCarrots","NEW!Minh'sEggRoll","Mango"],
11:["CheesePizza","PepperoniPizza","BuffaloChickenPizza","Chickenn'Waffles","PhillyChickenSub","FreshSeasonalSalads",
"Home-StyleCroissantSandwiches","CinnamonApples"],
14:["CheesePizza","PepperoniPizza","BuffaloChickenPizza","ChickenSpicySandwich","CheesyGrilledCheese","FreshSeasonalSalads",
"ChickenClubonCiabatta","YogurtBasket","BakedSweetPotato","TomatoSoup","FrozenFruitCup"],
15:["CheesePizza","PepperoniPizza","FourMeatPizza","BallParkHotDog","BBQPlate","FreshSeasonalSalads",
"DeliFreshSubs","RealFruitSmoothies","PotatoSpirals","Coleslaw","PineappleTidbits"],
16:["CheesePizza","PepperoniPizza","BuffaloChickenPizza","CrispyChickenTenders",
	"AsianRiceBowls","FreshSeasonalSalads","SpicyChickenSalad&Saltines","YogurtParfait","Carrot&CelerySticks","DriedFruit"],
17:["CheesePizza","PepperoniPizza","FlavortotheMaxSticks","GrilledChickenSandwich","FiestaNachos",
	"FreshSeasonalSalads","DeliFreshSubs","BlackBeanEmpanadas","RealFruitSmoothies","BlackBeans","Doritos","MandarinOranges"],
18:["CheesePizza","PepperoniPizza","BuffaloChickenPizza","NEW!FrenchToast&Sausage","NEW!SpaghettiMeatballBowl",
		"FreshSeasonalSalads","Home-StyleCroissantSandwiches","TropicalFruitSalad"],
21:["CheesePizza","PepperoniPizza","BuffaloChickenPizza","BonelessWings","KoreanMeatballSub",
	"FreshSeasonalSalads","ChickenClubonCiabatta","YogurtBasket","GAGrownGreenBeans","BakedSweetPotato","FrozenFruitCup"],
22:["CheesePizza","PepperoniPizza","FourMeatPizza","BaconCheeseburger","ChickenSoftTacos",
	"FreshSeasonalSalads","DeliFreshSubs","MexiPizza","RealFruitSmoothies","NEW!HappyCorn","MexiRice","MexicanFruitCup"],
23:["CheesePizza","PepperoniPizza","BuffaloChickenPizza","BlackAngusSliderDogs","NEW!HawaiianPanini",
	"FreshSeasonalSalads","FruityChickenSalad&Saltines","YogurtParfait","BakedBeans","CheesyBroccoli","DriedFruit"],
24:["CheesePizza","PepperoniPizza","FlavortotheMaxSticks","BBQChickenSandwichonCiabatta",
	"LasagnaBolognese","FreshSeasonalSalads","DeliFreshSubs","VeggieLasagnaRoll","RealFruitSmoothies","SteamedCarrots"]
,25:["CheesePizza","PepperoniPizza","BuffaloChickenPizza","Chickenn'Waffles",
	"FreshSeasonalSalads","Home-StyleCroissantSandwiches","Plantains","CinnamonApples"]
  }
  ```
  
 
 To remove the duplicates, 
 I am going to make one list of only unique foods
 and a new dictionary with thee same key,
 but the value will be the index of the food in the unique food list 
 
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

The new format with the unque list and idxlunches dictionary,
is only 1861 characters.

the new format is only 62% of the original size.

to generate the json file

```python
>>> with open("menu.json","w+") as outfile:
        print("var unique=",json.dumps(unique),"\nvar idxlunches=",json.dumps(idxlunches),file=outfile)
```
