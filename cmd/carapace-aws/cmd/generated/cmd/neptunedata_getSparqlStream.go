package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptunedata_getSparqlStreamCmd = &cobra.Command{
	Use:   "get-sparql-stream",
	Short: "Gets a stream for an RDF graph.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptunedata_getSparqlStreamCmd).Standalone()

	neptunedata_getSparqlStreamCmd.Flags().String("commit-num", "", "The commit number of the starting record to read from the change-log stream.")
	neptunedata_getSparqlStreamCmd.Flags().String("encoding", "", "If set to TRUE, Neptune compresses the response using gzip encoding.")
	neptunedata_getSparqlStreamCmd.Flags().String("iterator-type", "", "Can be one of:")
	neptunedata_getSparqlStreamCmd.Flags().String("limit", "", "Specifies the maximum number of records to return.")
	neptunedata_getSparqlStreamCmd.Flags().String("op-num", "", "The operation sequence number within the specified commit to start reading from in the change-log stream data.")
	neptunedataCmd.AddCommand(neptunedata_getSparqlStreamCmd)
}
