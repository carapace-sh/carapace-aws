package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_listCustomPermissionsCmd = &cobra.Command{
	Use:   "list-custom-permissions",
	Short: "Returns a list of all the custom permissions profiles.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_listCustomPermissionsCmd).Standalone()

	quicksight_listCustomPermissionsCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the custom permissions profiles that you want to list.")
	quicksight_listCustomPermissionsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	quicksight_listCustomPermissionsCmd.Flags().String("next-token", "", "The token for the next set of results, or null if there are no more results.")
	quicksight_listCustomPermissionsCmd.MarkFlagRequired("aws-account-id")
	quicksightCmd.AddCommand(quicksight_listCustomPermissionsCmd)
}
