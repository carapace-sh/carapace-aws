package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53profiles_associateResourceToProfileCmd = &cobra.Command{
	Use:   "associate-resource-to-profile",
	Short: "Associates a DNS reource configuration to a Route 53 Profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53profiles_associateResourceToProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53profiles_associateResourceToProfileCmd).Standalone()

		route53profiles_associateResourceToProfileCmd.Flags().String("name", "", "Name for the resource association.")
		route53profiles_associateResourceToProfileCmd.Flags().String("profile-id", "", "ID of the Profile.")
		route53profiles_associateResourceToProfileCmd.Flags().String("resource-arn", "", "Amazon resource number, ARN, of the DNS resource.")
		route53profiles_associateResourceToProfileCmd.Flags().String("resource-properties", "", "If you are adding a DNS Firewall rule group, include also a priority.")
		route53profiles_associateResourceToProfileCmd.MarkFlagRequired("name")
		route53profiles_associateResourceToProfileCmd.MarkFlagRequired("profile-id")
		route53profiles_associateResourceToProfileCmd.MarkFlagRequired("resource-arn")
	})
	route53profilesCmd.AddCommand(route53profiles_associateResourceToProfileCmd)
}
