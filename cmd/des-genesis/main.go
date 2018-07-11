// Copyright 2018 The des-network Authors
// This file is part of des-network.
//
// You should have received a copy of the GNU General Public License
// along with des-network. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
)

func main() {

	var (
		// g *core.Genesis
		//TODO: add more options
		file  = flag.String("file", "genesis.json", "Filepath for the genesis.json file")
		owner = flag.String("owner", "", "address of the owner of the chain")
		out   = flag.String("out", "des-genesis.json", "output file")
	)
	flag.Parse()

	glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(false)))
	log.Root().SetHandler(glogger)

	if *file != "" && *owner != "" {
		// TODO: (HACK) hardcoded the nonce here
		contractAddress := generateContractAddress(*owner, 0)
		g, err := makeGenesisConfig(*file)
		if err != nil {
			utils.Fatalf("%v\n", "failure trying to parse genesis file.")
		}
		// set contract address as extra data
		g.ExtraData = []byte(contractAddress)
		// rewrite address
		bytes, err := g.MarshalJSON()
		if err != nil {
			utils.Fatalf("%v: %v\n", "Couldn't write to file", *out)
		}
		err = ioutil.WriteFile(*out, bytes, 0644)
		if err != nil {
			utils.Fatalf("%v: %v\n", "Couldn't write to file", *out)
		}
	} else {
		utils.Fatalf("%v\n", "Genesis file path or owner address not given.")
	}

}

// generateContractAddress will take the address of the sender and the nonce and
// pre-determine the address of contracts
func generateContractAddress(_addr string, nonce uint) string {
	addr := common.HexToAddress(_addr)
	stuff := make([]interface{}, 0)
	stuff = append(stuff, addr.Bytes())
	stuff = append(stuff, nonce)
	bytes, err := rlp.EncodeToBytes(stuff)
	if err != nil {
		panic(err)
	}
	return "0x" + crypto.Keccak256Hash(bytes).Hex()[26:]
}

func makeGenesisConfig(genesisPath string) (*core.Genesis, error) {
	if len(genesisPath) == 0 {
		utils.Fatalf("Must supply path to genesis JSON file")
	}
	file, err := os.Open(genesisPath)
	if err != nil {
		utils.Fatalf("Failed to read genesis file: %v", err)
	}
	defer file.Close()

	genesis := new(core.Genesis)
	if err := json.NewDecoder(file).Decode(genesis); err != nil {
		utils.Fatalf("invalid genesis file: %v", err)
	}

	file.Seek(0, 0)
	return genesis, nil
}
