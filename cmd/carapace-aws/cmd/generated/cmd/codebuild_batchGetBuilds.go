package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_batchGetBuildsCmd = &cobra.Command{
	Use:   "batch-get-builds",
	Short: "Gets information about one or more builds.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_batchGetBuildsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codebuild_batchGetBuildsCmd).Standalone()

		codebuild_batchGetBuildsCmd.Flags().String("ids", "", "The IDs of the builds.")
		codebuild_batchGetBuildsCmd.MarkFlagRequired("ids")
	})
	codebuildCmd.AddCommand(codebuild_batchGetBuildsCmd)
}
