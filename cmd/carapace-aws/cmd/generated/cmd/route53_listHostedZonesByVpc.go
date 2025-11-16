package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_listHostedZonesByVpcCmd = &cobra.Command{
	Use:   "list-hosted-zones-by-vpc",
	Short: "Lists all the private hosted zones that a specified VPC is associated with, regardless of which Amazon Web Services account or Amazon Web Services service owns the hosted zones.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_listHostedZonesByVpcCmd).Standalone()

	route53_listHostedZonesByVpcCmd.Flags().String("max-items", "", "(Optional) The maximum number of hosted zones that you want Amazon Route 53 to return.")
	route53_listHostedZonesByVpcCmd.Flags().String("next-token", "", "If the previous response included a `NextToken` element, the specified VPC is associated with more hosted zones.")
	route53_listHostedZonesByVpcCmd.Flags().String("vpcid", "", "The ID of the Amazon VPC that you want to list hosted zones for.")
	route53_listHostedZonesByVpcCmd.Flags().String("vpcregion", "", "For the Amazon VPC that you specified for `VPCId`, the Amazon Web Services Region that you created the VPC in.")
	route53_listHostedZonesByVpcCmd.MarkFlagRequired("vpcid")
	route53_listHostedZonesByVpcCmd.MarkFlagRequired("vpcregion")
	route53Cmd.AddCommand(route53_listHostedZonesByVpcCmd)
}
