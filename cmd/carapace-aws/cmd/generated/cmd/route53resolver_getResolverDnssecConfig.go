package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_getResolverDnssecConfigCmd = &cobra.Command{
	Use:   "get-resolver-dnssec-config",
	Short: "Gets DNSSEC validation information for a specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_getResolverDnssecConfigCmd).Standalone()

	route53resolver_getResolverDnssecConfigCmd.Flags().String("resource-id", "", "The ID of the virtual private cloud (VPC) for the DNSSEC validation status.")
	route53resolver_getResolverDnssecConfigCmd.MarkFlagRequired("resource-id")
	route53resolverCmd.AddCommand(route53resolver_getResolverDnssecConfigCmd)
}
