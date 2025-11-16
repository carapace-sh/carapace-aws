package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_listNamespacesCmd = &cobra.Command{
	Use:   "list-namespaces",
	Short: "Lists the namespaces for the specified Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_listNamespacesCmd).Standalone()

	quicksight_listNamespacesCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that contains the Quick Sight namespaces that you want to list.")
	quicksight_listNamespacesCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	quicksight_listNamespacesCmd.Flags().String("next-token", "", "A unique pagination token that can be used in a subsequent request.")
	quicksight_listNamespacesCmd.MarkFlagRequired("aws-account-id")
	quicksightCmd.AddCommand(quicksight_listNamespacesCmd)
}
