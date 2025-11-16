package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_listEndpointAccessCmd = &cobra.Command{
	Use:   "list-endpoint-access",
	Short: "Returns an array of `EndpointAccess` objects and relevant information.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_listEndpointAccessCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshiftServerless_listEndpointAccessCmd).Standalone()

		redshiftServerless_listEndpointAccessCmd.Flags().String("max-results", "", "An optional parameter that specifies the maximum number of results to return.")
		redshiftServerless_listEndpointAccessCmd.Flags().String("next-token", "", "If your initial `ListEndpointAccess` operation returns a `nextToken`, you can include the returned `nextToken` in following `ListEndpointAccess` operations, which returns results in the next page.")
		redshiftServerless_listEndpointAccessCmd.Flags().String("owner-account", "", "The owner Amazon Web Services account for the Amazon Redshift Serverless workgroup.")
		redshiftServerless_listEndpointAccessCmd.Flags().String("vpc-id", "", "The unique identifier of the virtual private cloud with access to Amazon Redshift Serverless.")
		redshiftServerless_listEndpointAccessCmd.Flags().String("workgroup-name", "", "The name of the workgroup associated with the VPC endpoint to return.")
	})
	redshiftServerlessCmd.AddCommand(redshiftServerless_listEndpointAccessCmd)
}
