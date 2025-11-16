package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_listFoldersCmd = &cobra.Command{
	Use:   "list-folders",
	Short: "Lists all folders in an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_listFoldersCmd).Standalone()

	quicksight_listFoldersCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that contains the folder.")
	quicksight_listFoldersCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
	quicksight_listFoldersCmd.Flags().String("next-token", "", "The token for the next set of results, or null if there are no more results.")
	quicksight_listFoldersCmd.MarkFlagRequired("aws-account-id")
	quicksightCmd.AddCommand(quicksight_listFoldersCmd)
}
