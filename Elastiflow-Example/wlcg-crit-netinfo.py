#!/bin/python
#
# This script queries WLCG CRIC to gather network metadata for use
# in Elastiflow IP address enrichment.
#
# However this script could also be used as an example of how to get CRIC rcsite details
#
# Shawn McKee, 25-Apr-2022, smckee@umich.edu
######################################################

# Get required packages
import json
import requests
import yaml

# Query WLCG CRIC for rcsite information.  If you have the CERN-CA-certs RPM package installed you can
# use the verified version:
# response = requests.get("https://wlcg-cric.cern.ch/api/core/rcsite/query/list/?json",verify='/etc/pki/tls/certs/CERN-bundle.pem')

# If you don't have the CERN-bundle.pem file available, set verify=False
response = requests.get(
    "https://wlcg-cric.cern.ch/api/core/rcsite/query/list/?json", verify=False
)

# Create dictionary of information for each rcsite
rcsites = json.loads(response.text)

# Examine data for a particular site
# print(json.dumps(rcsites["AGLT2"],indent=2, sort_keys=True))

# Create a networks dictionary for Elastiflow IP Address Enrichment
# https://docs.elastiflow.com/docs/config_enrich_ip/#user-defined-metadata-enrichment
networks = {}

#  This example is for AGLT2 with the state of Michigan and we need to identify those sites and networks
#  using appropriate keywards
checkfor = ["AGLT2", "Michigan"]

# Loop over all rcsites, getting information to populate the networks dictionary
for rcsite in rcsites:
    for netsite in rcsites[rcsite]["netsites"]:
        for netroute in rcsites[rcsite]["netroutes"]:
            if netroute:
                if (
                    rcsites[rcsite]["netroutes"][netroute]["netsite"] == netsite
                    or rcsites[rcsite]["netroutes"][netroute]["netsite_spare"]
                    == netsite
                ):
                    for iptype in rcsites[rcsite]["netroutes"][netroute]["networks"]:
                        for subnet in rcsites[rcsite]["netroutes"][netroute][
                            "networks"
                        ][iptype]:
                            lat = rcsites[rcsite]["latitude"]
                            long = rcsites[rcsite]["longitude"]
                            asn = rcsites[rcsite]["netroutes"][netroute]["asn"]
                            networks[subnet] = {}
                            networks[subnet]["name"] = netroute
                            networks[subnet]["metadata"] = {}
                            networks[subnet]["metadata"][".netsite"] = netsite
                            networks[subnet]["metadata"][".rcsite"] = rcsite
                            networks[subnet]["metadata"][".asn"] = asn
                            #                                                       networks[subnet]["metadata"][".geo.loc.coord"]=lat,long
                            # If site has AGLT2 or netsite has Michigan, mark internal
                            if any(
                                rcsite.find(check) > -1 for check in checkfor
                            ) or any(netsite.find(check) > -1 for check in checkfor):
                                networks[subnet]["internal"] = True
                                if any(netsite.find(check) > -1 for check in ["MSU"]):
                                    lat = 42.71633323520017
                                    long = -84.4902483345146
                                    #                                                                       networks[subnet]["metadata"][".geo.loc.coord"]=lat,long
                                    networks[subnet]["metadata"][
                                        ".geo.city.name"
                                    ] = "East Lansing"
                                else:
                                    lat = 42.2765243047681
                                    long = -83.7412221723667
                                    #                                                                       networks[subnet]["metadata"][".geo.loc.coord"]=lat,long
                                    networks[subnet]["metadata"][
                                        ".geo.city.name"
                                    ] = "Ann Arbor"
print("Writing network information to ipaddrs.yml")
with open("ipaddrs.yml", "w") as f:
    print(yaml.dump(networks), file=f)
# You may want to postprocess the ipaddrs.yml or augment it with additional non-CRIC network data
