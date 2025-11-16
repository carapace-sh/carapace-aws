package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sns_deleteEndpointCmd = &cobra.Command{
	Use:   "delete-endpoint",
	Short: "Deletes the endpoint for a device and mobile app from Amazon SNS.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sns_deleteEndpointCmd).Standalone()

	sns_deleteEndpointCmd.Flags().String("endpoint-arn", "", "`EndpointArn` of endpoint to delete.")
	sns_deleteEndpointCmd.MarkFlagRequired("endpoint-arn")
	snsCmd.AddCommand(sns_deleteEndpointCmd)
}
