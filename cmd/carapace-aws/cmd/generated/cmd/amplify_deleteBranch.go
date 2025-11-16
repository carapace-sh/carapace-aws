package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplify_deleteBranchCmd = &cobra.Command{
	Use:   "delete-branch",
	Short: "Deletes a branch for an Amplify app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplify_deleteBranchCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplify_deleteBranchCmd).Standalone()

		amplify_deleteBranchCmd.Flags().String("app-id", "", "The unique ID for an Amplify app.")
		amplify_deleteBranchCmd.Flags().String("branch-name", "", "The name of the branch.")
		amplify_deleteBranchCmd.MarkFlagRequired("app-id")
		amplify_deleteBranchCmd.MarkFlagRequired("branch-name")
	})
	amplifyCmd.AddCommand(amplify_deleteBranchCmd)
}
