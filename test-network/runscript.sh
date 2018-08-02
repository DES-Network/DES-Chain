#!/bin/bash
PRIVATE_CONFIG=qdata/c1/tm.ipc ./bin/geth --exec "loadScript(\"$1\")" attach ipc:qdata/dd1/geth.ipc
