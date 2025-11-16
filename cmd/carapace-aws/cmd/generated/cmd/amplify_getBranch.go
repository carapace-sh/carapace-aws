package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplify_getBranchCmd = &cobra.Command{
	Use:   "get-branch",
	Short: "Returns a branch for an Amplify app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplify_getBranchCmd).Standalone()

	amplify_getBranchCmd.Flags().String("app-id", "", "The unique ID for an Amplify app.")
	amplify_getBranchCmd.Flags().String("branch-name", "", "The name of the branch.")
	amplify_getBranchCmd.MarkFlagRequired("app-id")
	amplify_getBranchCmd.MarkFlagRequired("branch-name")
	amplifyCmd.AddCommand(amplify_getBranchCmd)
}
