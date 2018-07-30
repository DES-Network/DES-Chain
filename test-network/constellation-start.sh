#!/bin/bash
set -u
set -e

i=$1
# for i in {1..7}
# do
    DDIR="qdata/c$i"
    mkdir -p $DDIR
    mkdir -p qdata/logs
    cp "keys/tm$i.pub" "$DDIR/tm.pub"
    cp "keys/tm$i.key" "$DDIR/tm.key"
    rm -f "$DDIR/tm.ipc"
    CMD="./bin/constellation-node --url=https://127.0.0.$i:900$i/ --port=900$i --workdir=$DDIR --socket=tm.ipc --publickeys=tm.pub --privatekeys=tm.key --othernodes=https://127.0.0.1:9001/"
    echo "$CMD >> qdata/logs/constellation$i.log 2>&1 &"
    $CMD >> "qdata/logs/constellation$i.log" 2>&1 &
# done

echo "[*] Initialized Constellation Node"
echo "[*] Waiting for confirmation that it is up..."

DOWN=true
while $DOWN; do
    sleep 0.1
    DOWN=false
	if [ ! -S "qdata/c$i/tm.ipc" ]; then
            DOWN=true
    else
            echo "It's up!"
	fi
done
