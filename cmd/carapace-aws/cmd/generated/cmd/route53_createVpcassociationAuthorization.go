package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_createVpcassociationAuthorizationCmd = &cobra.Command{
	Use:   "create-vpcassociation-authorization",
	Short: "Authorizes the Amazon Web Services account that created a specified VPC to submit an `AssociateVPCWithHostedZone` request to associate the VPC with a specified hosted zone that was created by a different account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_createVpcassociationAuthorizationCmd).Standalone()

	route53_createVpcassociationAuthorizationCmd.Flags().String("hosted-zone-id", "", "The ID of the private hosted zone that you want to authorize associating a VPC with.")
	route53_createVpcassociationAuthorizationCmd.Flags().String("vpc", "", "A complex type that contains the VPC ID and region for the VPC that you want to authorize associating with your hosted zone.")
	route53_createVpcassociationAuthorizationCmd.MarkFlagRequired("hosted-zone-id")
	route53_createVpcassociationAuthorizationCmd.MarkFlagRequired("vpc")
	route53Cmd.AddCommand(route53_createVpcassociationAuthorizationCmd)
}
