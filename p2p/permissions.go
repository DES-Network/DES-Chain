package p2p

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p/discover"
	"github.com/ethereum/go-ethereum/private"
)

const (
	NODE_NAME_LENGTH    = 32
	PERMISSIONED_CONFIG = "permissioned-nodes.json"
)

// DES: required for handling permissions from contract
type PermissioningClient struct {
	r *private.RegulatorClient
}


// NewPermissioningClient will create a permissioning client
func NewPermissioningClient() *PermissioningClient {
	client := private.NewRegulatorClient()
	return &PermissioningClient{client}
}

// check if a given node is permissioned to connect to the change
func (p *PermissioningClient) IsNodePermissioned(nodename string, currentNode string, datadir string, direction string) bool {

	var permissionedList []string
	nodes := parsePermissionedNodes(datadir)

	// DES: checking if contract deployed, else operate like Quorum
	if !p.r.IsDeployed() {
		log.Trace("Contract not yet deployed, take permissioned list for granted")
		for _, v := range nodes {
			permissionedList = append(permissionedList, v.ID.String())
		}
	} else {
		log.Trace("Contract deployed now, filter nodes using whitelist")
		for _, v := range nodes {
			isWhitelisted, err := p.r.IsWhitelisted(v.ID.String())
			if err != nil {
				log.Error("Error checking permissioned whitelist", "error", err)
				break
			}
			if isWhitelisted {
				permissionedList = append(permissionedList, v.ID.String())
			}
		}
	}
	// DES - end
	log.Trace("isNodePermissioned", "permissionedList", permissionedList)
	for _, v := range permissionedList {
		if v == nodename {
			log.Debug("isNodePermissioned", "connection", direction, "nodename", nodename[:NODE_NAME_LENGTH], "ALLOWED-BY", currentNode[:NODE_NAME_LENGTH])
			return true
		}
	}
	log.Debug("isNodePermissioned", "connection", direction, "nodename", nodename[:NODE_NAME_LENGTH], "DENIED-BY", currentNode[:NODE_NAME_LENGTH])
	return false
}

//this is a shameless copy from the config.go. It is a duplication of the code
//for the timebeing to allow reload of the permissioned nodes while the server is running

func parsePermissionedNodes(DataDir string) []*discover.Node {

	log.Debug("parsePermissionedNodes", "DataDir", DataDir, "file", PERMISSIONED_CONFIG)

	path := filepath.Join(DataDir, PERMISSIONED_CONFIG)
	if _, err := os.Stat(path); err != nil {
		log.Error("Read Error for permissioned-nodes.json file. This is because 'permissioned' flag is specified but no permissioned-nodes.json file is present.", "err", err)
		return nil
	}
	// Load the nodes from the config file
	blob, err := ioutil.ReadFile(path)
	if err != nil {
		log.Error("parsePermissionedNodes: Failed to access nodes", "err", err)
		return nil
	}

	nodelist := []string{}
	if err := json.Unmarshal(blob, &nodelist); err != nil {
		log.Error("parsePermissionedNodes: Failed to load nodes", "err", err)
		return nil
	}
	// Interpret the list as a discovery node array
	var nodes []*discover.Node
	for _, url := range nodelist {
		if url == "" {
			log.Error("parsePermissionedNodes: Node URL blank")
			continue
		}
		node, err := discover.ParseNode(url)
		if err != nil {
			log.Error("parsePermissionedNodes: Node URL", "url", url, "err", err)
			continue
		}
		nodes = append(nodes, node)
	}
	return nodes
}
