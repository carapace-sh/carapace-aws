package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53profiles_listProfileResourceAssociationsCmd = &cobra.Command{
	Use:   "list-profile-resource-associations",
	Short: "Lists all the resource associations for the specified Route 53 Profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53profiles_listProfileResourceAssociationsCmd).Standalone()

	route53profiles_listProfileResourceAssociationsCmd.Flags().String("max-results", "", "The maximum number of objects that you want to return for this request.")
	route53profiles_listProfileResourceAssociationsCmd.Flags().String("next-token", "", "For the first call to this list request, omit this value.")
	route53profiles_listProfileResourceAssociationsCmd.Flags().String("profile-id", "", "The ID of the Profile.")
	route53profiles_listProfileResourceAssociationsCmd.Flags().String("resource-type", "", "ID of a resource if you want information on only one type.")
	route53profiles_listProfileResourceAssociationsCmd.MarkFlagRequired("profile-id")
	route53profilesCmd.AddCommand(route53profiles_listProfileResourceAssociationsCmd)
}
