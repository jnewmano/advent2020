from sympy.ntheory.modular import crt

things = open('input.txt', 'r').readlines()
buses = things[1].strip().split(',')

mods = []
values = []
for i, v in enumerate(buses):
    if v == 'x':
        continue
    values.append(int(v))
    mods.append(-int(i))

results = crt(values, mods)

print(results[0])
