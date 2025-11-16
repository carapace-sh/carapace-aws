package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_listVpcconnectionsCmd = &cobra.Command{
	Use:   "list-vpcconnections",
	Short: "Lists all of the VPC connections in the current set Amazon Web Services Region of an Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_listVpcconnectionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_listVpcconnectionsCmd).Standalone()

		quicksight_listVpcconnectionsCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID of the account that contains the VPC connections that you want to list.")
		quicksight_listVpcconnectionsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
		quicksight_listVpcconnectionsCmd.Flags().String("next-token", "", "The token for the next set of results, or null if there are no more results.")
		quicksight_listVpcconnectionsCmd.MarkFlagRequired("aws-account-id")
	})
	quicksightCmd.AddCommand(quicksight_listVpcconnectionsCmd)
}
