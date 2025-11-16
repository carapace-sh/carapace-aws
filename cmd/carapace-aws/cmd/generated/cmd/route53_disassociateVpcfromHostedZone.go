package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_disassociateVpcfromHostedZoneCmd = &cobra.Command{
	Use:   "disassociate-vpcfrom-hosted-zone",
	Short: "Disassociates an Amazon Virtual Private Cloud (Amazon VPC) from an Amazon Route 53 private hosted zone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_disassociateVpcfromHostedZoneCmd).Standalone()

	route53_disassociateVpcfromHostedZoneCmd.Flags().String("comment", "", "*Optional:* A comment about the disassociation request.")
	route53_disassociateVpcfromHostedZoneCmd.Flags().String("hosted-zone-id", "", "The ID of the private hosted zone that you want to disassociate a VPC from.")
	route53_disassociateVpcfromHostedZoneCmd.Flags().String("vpc", "", "A complex type that contains information about the VPC that you're disassociating from the specified hosted zone.")
	route53_disassociateVpcfromHostedZoneCmd.MarkFlagRequired("hosted-zone-id")
	route53_disassociateVpcfromHostedZoneCmd.MarkFlagRequired("vpc")
	route53Cmd.AddCommand(route53_disassociateVpcfromHostedZoneCmd)
}
