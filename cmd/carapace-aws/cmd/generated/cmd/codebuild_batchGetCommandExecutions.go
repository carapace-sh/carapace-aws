package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_batchGetCommandExecutionsCmd = &cobra.Command{
	Use:   "batch-get-command-executions",
	Short: "Gets information about the command executions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_batchGetCommandExecutionsCmd).Standalone()

	codebuild_batchGetCommandExecutionsCmd.Flags().String("command-execution-ids", "", "A comma separated list of `commandExecutionIds`.")
	codebuild_batchGetCommandExecutionsCmd.Flags().String("sandbox-id", "", "A `sandboxId` or `sandboxArn`.")
	codebuild_batchGetCommandExecutionsCmd.MarkFlagRequired("command-execution-ids")
	codebuild_batchGetCommandExecutionsCmd.MarkFlagRequired("sandbox-id")
	codebuildCmd.AddCommand(codebuild_batchGetCommandExecutionsCmd)
}
