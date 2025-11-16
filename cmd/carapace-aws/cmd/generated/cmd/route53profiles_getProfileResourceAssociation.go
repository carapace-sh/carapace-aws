package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53profiles_getProfileResourceAssociationCmd = &cobra.Command{
	Use:   "get-profile-resource-association",
	Short: "Returns information about a specified Route 53 Profile resource association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53profiles_getProfileResourceAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53profiles_getProfileResourceAssociationCmd).Standalone()

		route53profiles_getProfileResourceAssociationCmd.Flags().String("profile-resource-association-id", "", "The ID of the profile resource association that you want to get information about.")
		route53profiles_getProfileResourceAssociationCmd.MarkFlagRequired("profile-resource-association-id")
	})
	route53profilesCmd.AddCommand(route53profiles_getProfileResourceAssociationCmd)
}
