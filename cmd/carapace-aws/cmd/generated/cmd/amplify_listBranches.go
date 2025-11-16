package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplify_listBranchesCmd = &cobra.Command{
	Use:   "list-branches",
	Short: "Lists the branches of an Amplify app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplify_listBranchesCmd).Standalone()

	amplify_listBranchesCmd.Flags().String("app-id", "", "The unique ID for an Amplify app.")
	amplify_listBranchesCmd.Flags().String("max-results", "", "The maximum number of records to list in a single response.")
	amplify_listBranchesCmd.Flags().String("next-token", "", "A pagination token.")
	amplify_listBranchesCmd.MarkFlagRequired("app-id")
	amplifyCmd.AddCommand(amplify_listBranchesCmd)
}
