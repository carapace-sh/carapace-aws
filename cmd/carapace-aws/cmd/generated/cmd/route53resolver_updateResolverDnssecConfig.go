package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_updateResolverDnssecConfigCmd = &cobra.Command{
	Use:   "update-resolver-dnssec-config",
	Short: "Updates an existing DNSSEC validation configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_updateResolverDnssecConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53resolver_updateResolverDnssecConfigCmd).Standalone()

		route53resolver_updateResolverDnssecConfigCmd.Flags().String("resource-id", "", "The ID of the virtual private cloud (VPC) that you're updating the DNSSEC validation status for.")
		route53resolver_updateResolverDnssecConfigCmd.Flags().String("validation", "", "The new value that you are specifying for DNSSEC validation for the VPC.")
		route53resolver_updateResolverDnssecConfigCmd.MarkFlagRequired("resource-id")
		route53resolver_updateResolverDnssecConfigCmd.MarkFlagRequired("validation")
	})
	route53resolverCmd.AddCommand(route53resolver_updateResolverDnssecConfigCmd)
}
