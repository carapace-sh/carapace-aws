package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_listCommandExecutionsForSandboxCmd = &cobra.Command{
	Use:   "list-command-executions-for-sandbox",
	Short: "Gets a list of command executions for a sandbox.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_listCommandExecutionsForSandboxCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codebuild_listCommandExecutionsForSandboxCmd).Standalone()

		codebuild_listCommandExecutionsForSandboxCmd.Flags().String("max-results", "", "The maximum number of sandbox records to be retrieved.")
		codebuild_listCommandExecutionsForSandboxCmd.Flags().String("next-token", "", "The next token, if any, to get paginated results.")
		codebuild_listCommandExecutionsForSandboxCmd.Flags().String("sandbox-id", "", "A `sandboxId` or `sandboxArn`.")
		codebuild_listCommandExecutionsForSandboxCmd.Flags().String("sort-order", "", "The order in which sandbox records should be retrieved.")
		codebuild_listCommandExecutionsForSandboxCmd.MarkFlagRequired("sandbox-id")
	})
	codebuildCmd.AddCommand(codebuild_listCommandExecutionsForSandboxCmd)
}
