# Galactic Merchant Guide Trading
Welcome to the Galactic Merchant Guide Trading system! This system allows intergalactic traders to register new currency values, convert them into Roman numerals, and calculate the credit values for minerals based on these new currencies. The system supports interactions where users can ask about the value of new currencies or calculate the credit value of minerals.

## Background Story
You decided to give up on earth after the latest financial collapse left 99.99% of the earth's
population with 0.01% of the wealth. Luckily, with the scant sum of money that is left in your
account, you are able to afford to rent a spaceship, leave earth, and fly all over the galaxy to sell
common metals and dirt (which apparently is worth a lot). Buying and selling over the galaxy
requires you to convert numbers and units, and you decided to write a program to help you.The
numbers used for intergalactic transactions follows similar convention to the roman numerals and
you have painstakingly collected the appropriate translation between them. Roman numerals are
based on seven symbols:

```csharp 
Symbol   Value: 
I        1
V        5
X        10
L        50
C        100
D        500
M        1
```

Numbers are formed by combining symbols together and adding the values. For example, MMVI is
1000 + 1000 + 5 + 1 = 2006. Generally, symbols are placed in order of value, starting with the
largest values. When smaller values precede larger values, the smaller values are subtracted from
the larger values, and the result is added to the total. For example MCMXLIV = 1000 + (1000 −
100) + (50 − 10) + (5 − 1) = 1944.

The symbols "I", "X", "C", and "M" can be repeated three times in succession, but no more. (They
may appear four times if the third and fourth are separated by a smaller value, such as XXXIX.)
"D", "L", and "V" can never be repeated.

"I" can be subtracted from "V" and "X" only. "X" can be subtracted from "L" and "C" only. "C" can
be subtracted from "D" and "M" only. "V", "L", and "D" can never be subtracted.

Only one small-value symbol may be subtracted from any large-value symbol.

A number written in Arabic numerals can be broken into digits. For example, 1903 is composed of
1, 9, 0, and 3. To write the Roman numeral, each of the non-zero digits should be treated separately.
In the above example, 1,000 = M, 900 = CM, and 3 = III. Therefore, 1903 = MCMIII.

**_-- Source: Wikipedia_** (http://en.wikipedia.org/wiki/Roman_numerals)

### Sample Input:

- glob is I
- prok is V
- pish is X
- tegj is L
- glob glob Silver is 34 Credits
- glob prok Gold is 57800 Credits
- pish pish Iron is 3910 Credits
- how much is pish tegj glob glob?
- how many Credits is glob prok Silver?
- how many Credits is glob glob Gold?
- how many Credits is glob glob glob glob glob glob Gold?
- how many Credits is pish tegj glob Iron?
- Does pish tegj glob glob Iron has more Credits than glob glob Gold?
- Does glob glob Gold has less Credits than pish tegj glob glob Iron?
- Is glob prok larger than pish pish?
- Istegj glob glob smaller than glob prok?
- how much wood could a woodchuck chuck if a woodchuck could chuck wood?

### Sample Output:
- pish tegj glob glob is 42
- glob prok Silver is 68 Credits
- glob glob Gold is 28900 Credits
- Requested number is in invalid format
- pish tegj glob Iron is 8015.5 Credits
- pish tegj glob glob Iron has less Credits than glob prok Gold
- glob glob Gold has more Credits than pish tegj glob glob
- glob prok is smaller than pish pish
- tegj glob glob is larger than glob prok
- I have no idea what you are talking about

## Features
Based on the user story, this program provides several features, such as;

- **Register a New Galactic Currency**: Users can register a new galactic currency by mapping it to a Roman numeral equivalent (e.g., glob is mapped to I, prok is mapped to V).
- **Register Mineral Values**: Users can provide the value of minerals in terms of credits (e.g., glob glob Silver is 34 Credits).
- **Currency Conversion**: The system can answer how much a galactic currency value is worth by converting the currency into Roman numerals and then into an integer.
- **Credit Calculation**: Users can inquire how many credits a certain mineral is worth based on previously registered currency and mineral values.

## Usage

### Registering Galactic Currency
You can register new galactic currencies by mapping them to Roman numeral values. For example:

```python
glob is I
prok is V
pish is X
tegj is L
```

### Registering Mineral Credit Values
To register the credit value of a mineral based on galactic currency, use the following format:

```python
glob glob Silver is 34 Credits
glob prok Gold is 57800 Credits
pish pish Iron is 3910 Credits
```

Here, the system will calculate the value of the mineral (Silver, Gold, Iron) based on the currency values provided, and store the credit value for future inquiries.


### Inquiring Galactic Currency Value
You can ask how much a galactic currency value is worth by converting it to Roman numerals and then into an integer. For example:

```perl
how much is pish tegj glob glob ?
```

The system will convert `pish tegj glob glob` to the corresponding Roman numeral value (`X L I I`), then calculate the integer value (`42`).

### Inquiring Mineral Credit Value
You can also ask how many credits a specific amount of a mineral is worth. For example:

```csharp
how many Credits is glob prok Silver ?
how many Credits is pish tegj glob Iron ?
```

The system will respond with the number of credits based on the currency value and mineral credit registration.

### Example Queries
> `how much is glob prok?`
>> Output: "glob prok is 4"

> `how many Credits is glob prok Silver?` 
>> Output: "glob prok Silver is 68 Credits"

> `how many Credits is pish tegj Iron?`
>> Output: "pish tegj Iron is 782 Credits"

## How It Works
1. **Currency Registration**: Map galactic terms to Roman numerals.
1. **Credit Calculation**: Use the Roman numeral equivalents to calculate the value of minerals in credits.
1. **Query Answering**: Based on the stored values and conversions, the system can answer various user queries regarding the galactic currency or mineral credits.

### Roman Numeral Conversion

The system follows the standard Roman numeral conversion rules:

```csharp
Roman	Value
I	    1
V	    5
X	    10
L	    50
C	    100
D	    500
M	    1000
```

### Roman Numeral Rules
- **Addition**: Smaller values before larger ones are added (e.g., VI = 6).
- **Subtraction**: Smaller values before larger ones are subtracted (e.g., IV = 4).
Complete rules can be checked on 

## How to Run
You can use this system by inputting galactic currencies, minerals, and querying their values as described above.