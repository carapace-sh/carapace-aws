package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_revokeVpcEndpointAccessCmd = &cobra.Command{
	Use:   "revoke-vpc-endpoint-access",
	Short: "Revokes access to an Amazon OpenSearch Service domain that was provided through an interface VPC endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_revokeVpcEndpointAccessCmd).Standalone()

	es_revokeVpcEndpointAccessCmd.Flags().String("account", "", "The account ID to revoke access from.")
	es_revokeVpcEndpointAccessCmd.Flags().String("domain-name", "", "The name of the OpenSearch Service domain.")
	es_revokeVpcEndpointAccessCmd.MarkFlagRequired("account")
	es_revokeVpcEndpointAccessCmd.MarkFlagRequired("domain-name")
	esCmd.AddCommand(es_revokeVpcEndpointAccessCmd)
}
