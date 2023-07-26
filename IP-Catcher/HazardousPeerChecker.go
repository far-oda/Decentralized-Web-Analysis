package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type NetworkPeer struct {
	ID              string   `json:"PeerID"`
	MultiAddresses  []string `json:"Maddrs"`
	Protocols       []string `json:"Protocols"`
	AgentVer        string   `json:"AgentVersion"`
	ConnDuration    string   `json:"ConnectDuration"`
	CrawlingTime    string   `json:"CrawlDuration"`
	StartVisitTime  string   `json:"VisitStartedAt"`
	EndVisitTime    string   `json:"VisitEndedAt"`
	ConnErrDetails  string   `json:"ConnectErrorStr"`
	CrawlErrDetails string   `json:"CrawlErrorStr"`
	IsPublicNode    *bool    `json:"IsExposed"`
}

func main() {
	hazardousIPs, loadErr := os.ReadFile("bad_ips.txt")
	if loadErr != nil {
		fmt.Println(loadErr)
		os.Exit(1)
	}

	dangerousIPArray := strings.Split(string(hazardousIPs), "\n")

	jsonOutput, openErr := os.Open("nebula_result.json")
	if openErr != nil {
		fmt.Println(openErr)
		os.Exit(1)
	}
	defer jsonOutput.Close()

	lineReader := bufio.NewScanner(jsonOutput)
	hazardousPeers := make(map[string]bool)

	for lineReader.Scan() {
		var networkPeer NetworkPeer
		unmarshalErr := json.Unmarshal(lineReader.Bytes(), &networkPeer)
		if unmarshalErr != nil {
			fmt.Println(unmarshalErr)
			continue
		}

		for _, multiAddr := range networkPeer.MultiAddresses {
			addrComponents := strings.Split(multiAddr, "/")

			for index, component := range addrComponents {
				ipVersions := []string{"ip4", "ip6"}
				isIPVersion := false
				for _, version := range ipVersions {
					if component == version {
						isIPVersion = true
					}
				}
				if isIPVersion {
					if index+1 < len(addrComponents) {
						extractedIP := addrComponents[index+1]

						for _, dangerousIP := range dangerousIPArray {
							if extractedIP == dangerousIP {
								peerIPKey := networkPeer.ID + "_" + extractedIP
								if _, found := hazardousPeers[peerIPKey]; !found {
									hazardousPeers[peerIPKey] = true
									fmt.Println("Warning: Hazardous peer detected. Peer ID: ", networkPeer.ID, "At IPv4: ", extractedIP)
								}
							}
						}
					}
				}
			}
		}
	}

	if lineReaderErr := lineReader.Err(); lineReaderErr != nil {
		fmt.Println("Error while reading file: ", lineReaderErr)
	}
}
