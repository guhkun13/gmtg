We need to evaluate the text fall into which categories inspecting their pattern.

Based on the example, we can see that there are some text categories:

1. assigning new currency to roman value
- ^(\w+) is ([IVXLCDM])$
- ^([a-z]+) is ([IVXLCDM])$

2. assigning mineral value
- ^((?:\w+ )+)(\w+) is (\d+) Credits$
- ^((?:\w+ )+)([A-Z][a-z]+) is (\d+) Credits$
- ^(([a-z]+ )+)([A-Z][a-z]+) is (\d+) Credits$

3. question about how much in new currency
4. question about how many credits in mineral value
5. question about credits comparison
6. question about new currency comparison
7. Invalid format
8. Question does not recognized
