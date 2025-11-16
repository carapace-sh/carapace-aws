#!/bin/sh
sed -i \
    -e 's/carapace.Gen(\(.*\)).Standalone()/carapace.Gen(\1).PreRun(func(cmd *cobra.Command, args []string) {\n\t\0/' \
    -e 's/.*.AddCommand(.*)/\t})\n\0/' \
    *.go
# TODO exclude root
