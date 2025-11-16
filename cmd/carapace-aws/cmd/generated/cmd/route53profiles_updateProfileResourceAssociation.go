package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53profiles_updateProfileResourceAssociationCmd = &cobra.Command{
	Use:   "update-profile-resource-association",
	Short: "Updates the specified Route 53 Profile resourse association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53profiles_updateProfileResourceAssociationCmd).Standalone()

	route53profiles_updateProfileResourceAssociationCmd.Flags().String("name", "", "Name of the resource association.")
	route53profiles_updateProfileResourceAssociationCmd.Flags().String("profile-resource-association-id", "", "ID of the resource association.")
	route53profiles_updateProfileResourceAssociationCmd.Flags().String("resource-properties", "", "If you are adding a DNS Firewall rule group, include also a priority.")
	route53profiles_updateProfileResourceAssociationCmd.MarkFlagRequired("profile-resource-association-id")
	route53profilesCmd.AddCommand(route53profiles_updateProfileResourceAssociationCmd)
}
