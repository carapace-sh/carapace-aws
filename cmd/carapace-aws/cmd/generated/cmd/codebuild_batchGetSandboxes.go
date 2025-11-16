package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_batchGetSandboxesCmd = &cobra.Command{
	Use:   "batch-get-sandboxes",
	Short: "Gets information about the sandbox status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_batchGetSandboxesCmd).Standalone()

	codebuild_batchGetSandboxesCmd.Flags().String("ids", "", "A comma separated list of `sandboxIds` or `sandboxArns`.")
	codebuild_batchGetSandboxesCmd.MarkFlagRequired("ids")
	codebuildCmd.AddCommand(codebuild_batchGetSandboxesCmd)
}
