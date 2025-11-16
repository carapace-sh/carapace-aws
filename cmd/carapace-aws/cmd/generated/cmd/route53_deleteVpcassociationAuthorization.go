package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_deleteVpcassociationAuthorizationCmd = &cobra.Command{
	Use:   "delete-vpcassociation-authorization",
	Short: "Removes authorization to submit an `AssociateVPCWithHostedZone` request to associate a specified VPC with a hosted zone that was created by a different account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_deleteVpcassociationAuthorizationCmd).Standalone()

	route53_deleteVpcassociationAuthorizationCmd.Flags().String("hosted-zone-id", "", "When removing authorization to associate a VPC that was created by one Amazon Web Services account with a hosted zone that was created with a different Amazon Web Services account, the ID of the hosted zone.")
	route53_deleteVpcassociationAuthorizationCmd.Flags().String("vpc", "", "When removing authorization to associate a VPC that was created by one Amazon Web Services account with a hosted zone that was created with a different Amazon Web Services account, a complex type that includes the ID and region of the VPC.")
	route53_deleteVpcassociationAuthorizationCmd.MarkFlagRequired("hosted-zone-id")
	route53_deleteVpcassociationAuthorizationCmd.MarkFlagRequired("vpc")
	route53Cmd.AddCommand(route53_deleteVpcassociationAuthorizationCmd)
}
