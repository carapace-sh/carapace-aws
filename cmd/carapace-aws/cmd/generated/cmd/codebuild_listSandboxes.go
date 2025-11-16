package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_listSandboxesCmd = &cobra.Command{
	Use:   "list-sandboxes",
	Short: "Gets a list of sandboxes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_listSandboxesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codebuild_listSandboxesCmd).Standalone()

		codebuild_listSandboxesCmd.Flags().String("max-results", "", "The maximum number of sandbox records to be retrieved.")
		codebuild_listSandboxesCmd.Flags().String("next-token", "", "The next token, if any, to get paginated results.")
		codebuild_listSandboxesCmd.Flags().String("sort-order", "", "The order in which sandbox records should be retrieved.")
	})
	codebuildCmd.AddCommand(codebuild_listSandboxesCmd)
}
