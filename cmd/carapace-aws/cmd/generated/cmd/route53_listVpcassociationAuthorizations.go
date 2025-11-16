package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_listVpcassociationAuthorizationsCmd = &cobra.Command{
	Use:   "list-vpcassociation-authorizations",
	Short: "Gets a list of the VPCs that were created by other accounts and that can be associated with a specified hosted zone because you've submitted one or more `CreateVPCAssociationAuthorization` requests.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_listVpcassociationAuthorizationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53_listVpcassociationAuthorizationsCmd).Standalone()

		route53_listVpcassociationAuthorizationsCmd.Flags().String("hosted-zone-id", "", "The ID of the hosted zone for which you want a list of VPCs that can be associated with the hosted zone.")
		route53_listVpcassociationAuthorizationsCmd.Flags().String("max-results", "", "*Optional*: An integer that specifies the maximum number of VPCs that you want Amazon Route 53 to return.")
		route53_listVpcassociationAuthorizationsCmd.Flags().String("next-token", "", "*Optional*: If a response includes a `NextToken` element, there are more VPCs that can be associated with the specified hosted zone.")
		route53_listVpcassociationAuthorizationsCmd.MarkFlagRequired("hosted-zone-id")
	})
	route53Cmd.AddCommand(route53_listVpcassociationAuthorizationsCmd)
}
