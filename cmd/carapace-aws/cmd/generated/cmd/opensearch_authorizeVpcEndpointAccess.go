package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_authorizeVpcEndpointAccessCmd = &cobra.Command{
	Use:   "authorize-vpc-endpoint-access",
	Short: "Provides access to an Amazon OpenSearch Service domain through the use of an interface VPC endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_authorizeVpcEndpointAccessCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearch_authorizeVpcEndpointAccessCmd).Standalone()

		opensearch_authorizeVpcEndpointAccessCmd.Flags().String("account", "", "The Amazon Web Services account ID to grant access to.")
		opensearch_authorizeVpcEndpointAccessCmd.Flags().String("domain-name", "", "The name of the OpenSearch Service domain to provide access to.")
		opensearch_authorizeVpcEndpointAccessCmd.Flags().String("service", "", "The Amazon Web Services service SP to grant access to.")
		opensearch_authorizeVpcEndpointAccessCmd.MarkFlagRequired("domain-name")
	})
	opensearchCmd.AddCommand(opensearch_authorizeVpcEndpointAccessCmd)
}
