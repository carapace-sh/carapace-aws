package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_batchGetFleetsCmd = &cobra.Command{
	Use:   "batch-get-fleets",
	Short: "Gets information about one or more compute fleets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_batchGetFleetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codebuild_batchGetFleetsCmd).Standalone()

		codebuild_batchGetFleetsCmd.Flags().String("names", "", "The names or ARNs of the compute fleets.")
		codebuild_batchGetFleetsCmd.MarkFlagRequired("names")
	})
	codebuildCmd.AddCommand(codebuild_batchGetFleetsCmd)
}
