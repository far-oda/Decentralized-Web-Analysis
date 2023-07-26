# Decentralized-Web-Analysis

The first folder **IP-Catcher** has the golang code and its files:

`bad_ips.txt` includes a list of malicious IPs identified by zenbox from virustotal when scanning various samples of a Malware botnet called **Interplanetary Storm (IPStorm)**.

`nebula_result.json` is an extraction of a crawl done using [Nebula Crawler](https://github.com/dennis-tra/nebula) with the following command: 
```nebula crawl --network IPFS --json-out ./results/```

`HazardousPeerChecker.go` is an extractor program coded by me in golang which takes a list of IPs from a file, and compares it with a JSON file (such as a result of a crawler) and looks for matches and prints out the result.

`Result.docx` Includes the hash of the samples which are related to the IPs and the results of running the Golang Code.

--------------
