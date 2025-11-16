package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_authorizeVpcEndpointAccessCmd = &cobra.Command{
	Use:   "authorize-vpc-endpoint-access",
	Short: "Provides access to an Amazon OpenSearch Service domain through the use of an interface VPC endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_authorizeVpcEndpointAccessCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(es_authorizeVpcEndpointAccessCmd).Standalone()

		es_authorizeVpcEndpointAccessCmd.Flags().String("account", "", "The account ID to grant access to.")
		es_authorizeVpcEndpointAccessCmd.Flags().String("domain-name", "", "The name of the OpenSearch Service domain to provide access to.")
		es_authorizeVpcEndpointAccessCmd.MarkFlagRequired("account")
		es_authorizeVpcEndpointAccessCmd.MarkFlagRequired("domain-name")
	})
	esCmd.AddCommand(es_authorizeVpcEndpointAccessCmd)
}
