# Reading CRIC Network Data

This directory holds some example python3 code which queries WLCG's CRIC rcsite information and extracts some network metadata for use in Elastiflow (see https://docs.elastiflow.com/).    **This code can be used as an example of how to query CRIC for network data**.

Note there are two examples of the 'request', one which uses verify in the SSL connection and requires you have the CERN-CA-certs rpm installed and one which sets verify=False.

The output of this example script is a file name ipaddrs.yml, which contains the metadata for various IP subnets (both IPv4 and IPv6).
