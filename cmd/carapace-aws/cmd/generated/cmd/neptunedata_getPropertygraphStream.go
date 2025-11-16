package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptunedata_getPropertygraphStreamCmd = &cobra.Command{
	Use:   "get-propertygraph-stream",
	Short: "Gets a stream for a property graph.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptunedata_getPropertygraphStreamCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptunedata_getPropertygraphStreamCmd).Standalone()

		neptunedata_getPropertygraphStreamCmd.Flags().String("commit-num", "", "The commit number of the starting record to read from the change-log stream.")
		neptunedata_getPropertygraphStreamCmd.Flags().String("encoding", "", "If set to TRUE, Neptune compresses the response using gzip encoding.")
		neptunedata_getPropertygraphStreamCmd.Flags().String("iterator-type", "", "Can be one of:")
		neptunedata_getPropertygraphStreamCmd.Flags().String("limit", "", "Specifies the maximum number of records to return.")
		neptunedata_getPropertygraphStreamCmd.Flags().String("op-num", "", "The operation sequence number within the specified commit to start reading from in the change-log stream data.")
	})
	neptunedataCmd.AddCommand(neptunedata_getPropertygraphStreamCmd)
}
