package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53profiles_disassociateResourceFromProfileCmd = &cobra.Command{
	Use:   "disassociate-resource-from-profile",
	Short: "Dissoaciated a specified resource, from the Route 53 Profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53profiles_disassociateResourceFromProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53profiles_disassociateResourceFromProfileCmd).Standalone()

		route53profiles_disassociateResourceFromProfileCmd.Flags().String("profile-id", "", "The ID of the Profile.")
		route53profiles_disassociateResourceFromProfileCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
		route53profiles_disassociateResourceFromProfileCmd.MarkFlagRequired("profile-id")
		route53profiles_disassociateResourceFromProfileCmd.MarkFlagRequired("resource-arn")
	})
	route53profilesCmd.AddCommand(route53profiles_disassociateResourceFromProfileCmd)
}
