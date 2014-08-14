#!/usr/bin/env python

from __future__ import print_function
import sys

if len(sys.argv) != 3:
    print("Specify 2 strings which you want to compare.")
    sys.exit(1)

print(sys.argv[1] == sys.argv[2])
