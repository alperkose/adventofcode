  0:      1:      2:      3:      4:
 aaaa    ....    aaaa    aaaa    ....
b    c  .    c  .    c  .    c  b    c
b    c  .    c  .    c  .    c  b    c
 ....    ....    dddd    dddd    dddd
e    f  .    f  e    .  .    f  .    f
e    f  .    f  e    .  .    f  .    f
 gggg    ....    gggg    gggg    ....

  5:      6:      7:      8:      9:
 aaaa    aaaa    aaaa    aaaa    aaaa
b    .  b    .  .    c  b    c  b    c
b    .  b    .  .    c  b    c  b    c
 dddd    dddd    ....    dddd    dddd
.    f  e    f  .    f  e    f  .    f
.    f  e    f  .    f  e    f  .    f
 gggg    gggg    ....    gggg    gggg



7 - 1 = a

3 - 7 = dg


0: 6
1: 2 *
2: 5
3: 5
4: 4 *
5: 5
6: 6
7: 3 *
8: 7 *
9: 6


cdfbe
           abcdefg
acedgfb :  1111111 (8)
cdfbe   :  0111110 (5)
gcdfa   :  1011011 (2)
fbcad   :  1111010 (3)
dab     :  1101000 (7)
cefabd  :  1111110 (9)
cdfgeb  :  0111111 (6)
eafb    :  1100110 (4)
cagedb  :  1111101 (0)
ab      :  1100000 (1)


6 ch:
6 if 1 & X != 1
6 if 7 & X != 7
9 if 4 & X == 4
9 if 3 & X == 3
0 otherwise

5 ch:
3 if 1 & X == 1
5 if 9 & X == X
2 otherwise