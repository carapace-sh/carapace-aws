package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_updateEndpointAccessCmd = &cobra.Command{
	Use:   "update-endpoint-access",
	Short: "Updates an Amazon Redshift Serverless managed endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_updateEndpointAccessCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshiftServerless_updateEndpointAccessCmd).Standalone()

		redshiftServerless_updateEndpointAccessCmd.Flags().String("endpoint-name", "", "The name of the VPC endpoint to update.")
		redshiftServerless_updateEndpointAccessCmd.Flags().String("vpc-security-group-ids", "", "The list of VPC security groups associated with the endpoint after the endpoint is modified.")
		redshiftServerless_updateEndpointAccessCmd.MarkFlagRequired("endpoint-name")
	})
	redshiftServerlessCmd.AddCommand(redshiftServerless_updateEndpointAccessCmd)
}
