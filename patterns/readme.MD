
Sites to validate regex online: 
- https://regexr.com/
- https://regex101.com/

We need to evaluate the text fall into which categories inspecting their pattern.

Based on the example, we can see that there are some text categories:

1. assigning new currency to roman value
glob is I
prok is V
pish is X
tegj is L

- ^(\w+) is ([IVXLCDM])$
- ^([a-z]+) is ([IVXLCDM])$

2. assigning mineral value
glob glob Silver is 34 Credits
glob prok Gold is 57800 Credits
pish pish Iron is 3910 Credits

- ^((?:\w+ )+)(\w+) is (\d+) Credits$
- ^((?:\w+ )+)([A-Z][a-z]+) is (\d+) Credits$
- ^(([a-z]+ )+)([A-Z][a-z]+) is (\d+) Credits$

3. question about how much in new currency
how much is pish tegj glob glob ?

 - ^how much is ((?:\w+ )+)\?$
 - ^(how much is) (([a-z]+ )+)\?$
 
4. question about how many credits in mineral value
how many Credits is glob glob Gold ?
how many Credits is glob glob glob glob glob glob Gold ?
how many Credits is pish tegj glob Iron ?

 - ^how many Credits is ((?:\w+ )+)(\w+) \?$
 - ^how many Credits is (([a-z]+ )+)([A-Z][a-z]+) \?$

5. question about credits comparison
Does pish tegj glob glob Iron has more Credits than glob glob Gold ?
Does glob glob Gold has less Credits than pish tegj glob glob Iron?
 - ^Does (([a-z]+\ )+)([A-Z][a-z]+) has (less|more) Credits than (([a-z]+\ )+)([A-Z][a-z]+)\s{0,1}\?
 - ^Does ((?:[a-z]+\s+)+)([A-Z][a-z]+) has (less|more) Credits than ((?:[a-z]+\s+)+)([A-Z][a-z]+)\s?\?


6. question about new currency comparison
7. Invalid format
8. Question does not recognized


notes for optimizaton;
?: --> this is for non-capturing group, so that (([a-z]+ )+) can be grouped into 1 group instead of two\

so for currency, we can use (later)
((?:[a-z]\s+)+) to make it into 1 group instead of (([a-z]+ )+) that detected as 2 groups

comparison:
((?:[a-z]\s+)+)
(([a-z]+ )+)

pattern := `^Is ((?:[a-z]+\s?)+) (larger|smaller) than ((?:[a-z]+\s?)+)\?$`
pattern := `^Is ((?:[a-z]+\s?)+) (larger|smaller) than ((?:[a-z]+\s?)+)\s?\?$`

