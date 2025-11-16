package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_associateVpcwithHostedZoneCmd = &cobra.Command{
	Use:   "associate-vpcwith-hosted-zone",
	Short: "Associates an Amazon VPC with a private hosted zone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_associateVpcwithHostedZoneCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53_associateVpcwithHostedZoneCmd).Standalone()

		route53_associateVpcwithHostedZoneCmd.Flags().String("comment", "", "*Optional:* A comment about the association request.")
		route53_associateVpcwithHostedZoneCmd.Flags().String("hosted-zone-id", "", "The ID of the private hosted zone that you want to associate an Amazon VPC with.")
		route53_associateVpcwithHostedZoneCmd.Flags().String("vpc", "", "A complex type that contains information about the VPC that you want to associate with a private hosted zone.")
		route53_associateVpcwithHostedZoneCmd.MarkFlagRequired("hosted-zone-id")
		route53_associateVpcwithHostedZoneCmd.MarkFlagRequired("vpc")
	})
	route53Cmd.AddCommand(route53_associateVpcwithHostedZoneCmd)
}
