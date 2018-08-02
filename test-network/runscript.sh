#!/bin/bash
PRIVATE_CONFIG=qdata/c$1/tm.ipc ./bin/geth --exec "loadScript(\"$2\")" attach ipc:qdata/dd$1/geth.ipc
