#!/usr/bin/env python

from __future__ import print_function
import sys

if len(sys.argv) != 3:
    print("Specify 2 strings which you want to compare.")
    sys.exit(1)

result = (sys.argv[1] == sys.argv[2])
color = "32" if result else "31"
print("\033[{0};1m{1}\033[0m".format(color, result))
