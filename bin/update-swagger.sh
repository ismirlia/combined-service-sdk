#!/bin/bash

# Run from top-level project directory.
swaggerdir="../service-broker/swagger/swagger.yaml"

# get swagger directory
# printf "Enter swagger file directory, i.e., ~/User/Projects/service-broker/swagger\n   "
# read swaggerdir
# if [[ "$swaggerdir" != *"swagger.yaml" ]]; then
#     if [[ "$swaggerdir" != *"/" ]]; then
#         swaggerdir+="/"
#     fi
#     swaggerdir+="swagger.yaml"
# fi

# Remove old ppc-aas directory
if [ -d "ppc-aas" ]; then
    printf "Removing old Power-Go-Client swagger files\n"
    rm -rf ppc-aas
fi
mkdir ppc-aas

# Generate swagger files
printf "Generating new Power-Go-Client swagger file\n"
generatecmd="swagger generate client -f $swaggerdir -t ppc-aas"
eval $generatecmd

# Remove unused files
printf "\nRemoving unused Power-Go-Client swagger files\n"
rm ppc-aas/models/user_access.go
rm ppc-aas/models/principal.go