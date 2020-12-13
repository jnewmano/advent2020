import collections
import math
import re
import sys
from sympy.ntheory.modular import crt

f = open('input.txt', 'r') 
lines = f.readlines()
buses = lines[1].strip().split(',')

mods = []
values = []
for i, v in enumerate(buses):
  if v == 'x':
    continue
  values.append(int(v))
  mods.append(-int(i))

results = crt(values, mods)

print(results[0])
