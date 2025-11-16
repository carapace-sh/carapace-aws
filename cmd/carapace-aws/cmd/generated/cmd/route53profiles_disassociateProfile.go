package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53profiles_disassociateProfileCmd = &cobra.Command{
	Use:   "disassociate-profile",
	Short: "Dissociates a specified Route 53 Profile from the specified VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53profiles_disassociateProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53profiles_disassociateProfileCmd).Standalone()

		route53profiles_disassociateProfileCmd.Flags().String("profile-id", "", "ID of the Profile.")
		route53profiles_disassociateProfileCmd.Flags().String("resource-id", "", "The ID of the VPC.")
		route53profiles_disassociateProfileCmd.MarkFlagRequired("profile-id")
		route53profiles_disassociateProfileCmd.MarkFlagRequired("resource-id")
	})
	route53profilesCmd.AddCommand(route53profiles_disassociateProfileCmd)
}
