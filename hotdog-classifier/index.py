# Copyright (c) Alex Ellis 2017. All rights reserved.
# Licensed under the MIT license. See LICENSE file in the project root for full license information.

import sys
#from function import handler
import inception
def get_stdin():
    return sys.stdin.buffer.read()

if(__name__ == "__main__"):
    st = get_stdin()
    #handler.handle(st)
    print(inception.invoke(st))


