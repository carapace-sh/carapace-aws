package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_batchDeleteBuildsCmd = &cobra.Command{
	Use:   "batch-delete-builds",
	Short: "Deletes one or more builds.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_batchDeleteBuildsCmd).Standalone()

	codebuild_batchDeleteBuildsCmd.Flags().String("ids", "", "The IDs of the builds to delete.")
	codebuild_batchDeleteBuildsCmd.MarkFlagRequired("ids")
	codebuildCmd.AddCommand(codebuild_batchDeleteBuildsCmd)
}
