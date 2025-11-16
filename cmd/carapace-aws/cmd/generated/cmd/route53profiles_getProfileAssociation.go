package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53profiles_getProfileAssociationCmd = &cobra.Command{
	Use:   "get-profile-association",
	Short: "Retrieves a Route 53 Profile association for a VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53profiles_getProfileAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53profiles_getProfileAssociationCmd).Standalone()

		route53profiles_getProfileAssociationCmd.Flags().String("profile-association-id", "", "The identifier of the association you want to get information about.")
		route53profiles_getProfileAssociationCmd.MarkFlagRequired("profile-association-id")
	})
	route53profilesCmd.AddCommand(route53profiles_getProfileAssociationCmd)
}
