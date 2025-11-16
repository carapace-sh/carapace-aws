package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_listGatewayInstancesCmd = &cobra.Command{
	Use:   "list-gateway-instances",
	Short: "Displays a list of instances associated with the Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_listGatewayInstancesCmd).Standalone()

	mediaconnect_listGatewayInstancesCmd.Flags().String("filter-arn", "", "Filter the list results to display only the instances associated with the selected Gateway ARN.")
	mediaconnect_listGatewayInstancesCmd.Flags().String("max-results", "", "The maximum number of results to return per API request.")
	mediaconnect_listGatewayInstancesCmd.Flags().String("next-token", "", "The token that identifies the batch of results that you want to see.")
	mediaconnectCmd.AddCommand(mediaconnect_listGatewayInstancesCmd)
}
