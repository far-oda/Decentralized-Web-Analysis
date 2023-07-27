# Decentralized-Web-Analysis

The first folder **IP-Catcher** has the golang code and its files:

`bad_ips.txt` includes a list of malicious IPs identified by zenbox from virustotal when scanning various samples of a Malware botnet called **Interplanetary Storm (IPStorm)**.

`nebula_result.json` is an extraction of a crawl done using [Nebula Crawler](https://github.com/dennis-tra/nebula) with the following command: 
```nebula crawl --network IPFS --json-out ./results/```

`HazardousPeerChecker.go` is an extractor program coded by me in golang which takes a list of IPs from a file, and compares it with a JSON file (such as a result of a crawler) and looks for matches and prints out the result.

`Result.docx` Includes the hash of the samples which are related to the IPs and the results of running the Golang Code.

--------------
The Second folder **Virustotal Analysis** has the results of running the malware samples of IPStorm and the analysis of the sandbox services that it provides
Some of the results are in htm format which saved the entire page as one file to be able to access all aspects of the analysis and some are saved as pdf.

**The names are the SHA256 Hash of each sample**

Here are the links to the samples analysis on virustotal:

### Linux

[87990db7c4d3e07d4c16f45fab3efd1c34bcce27fe692f8e009794bd2ac1db40](https://www.virustotal.com/gui/file/87990db7c4d3e07d4c16f45fab3efd1c34bcce27fe692f8e009794bd2ac1db40)
[5987472321673fd54f2aebcc33ede6541897410df36a7f4a1bd984a2f724e772](https://www.virustotal.com/gui/file/5987472321673fd54f2aebcc33ede6541897410df36a7f4a1bd984a2f724e772)
[a62e88db6c56c69edb95b5032334132e50be1a8feba4c5bb7e39f5d55ecd69e3](https://www.virustotal.com/gui/file/a62e88db6c56c69edb95b5032334132e50be1a8feba4c5bb7e39f5d55ecd69e3)
[c0d7eeec4036bd2507153943a3a88dcbc399238860d8776b6581c577b7cd2c62](https://www.virustotal.com/gui/file/c0d7eeec4036bd2507153943a3a88dcbc399238860d8776b6581c577b7cd2c62)

### Windows
[fbbcd1e59ab5045a3f0ebb55fab604173da0efb11898822537aa7dea18bf092c](https://www.virustotal.com/gui/file/fbbcd1e59ab5045a3f0ebb55fab604173da0efb11898822537aa7dea18bf092c )
[c7dca4b2f45bf67275cd8530b0f43a63a37ae25c474a52e37fc00ece67e5edbf](https://www.virustotal.com/gui/file/c7dca4b2f45bf67275cd8530b0f43a63a37ae25c474a52e37fc00ece67e5edbf)
[7f3802a43a160267fc8c9140b3fa607ed5b9673b586c0b4b29690d53a95007df](https://www.virustotal.com/gui/file/7f3802a43a160267fc8c9140b3fa607ed5b9673b586c0b4b29690d53a95007df)
[984dbf6de1acd7a553e2937b879ece9ddd66fa999593f55cbd3df562252727da](https://www.virustotal.com/gui/file/984dbf6de1acd7a553e2937b879ece9ddd66fa999593f55cbd3df562252727da)
