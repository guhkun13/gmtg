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
5. question about credits comparison
6. question about new currency comparison
7. Invalid format
8. Question does not recognized
