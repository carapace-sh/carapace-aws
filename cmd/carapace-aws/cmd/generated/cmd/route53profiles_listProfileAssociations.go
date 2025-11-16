package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53profiles_listProfileAssociationsCmd = &cobra.Command{
	Use:   "list-profile-associations",
	Short: "Lists all the VPCs that the specified Route 53 Profile is associated with.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53profiles_listProfileAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53profiles_listProfileAssociationsCmd).Standalone()

		route53profiles_listProfileAssociationsCmd.Flags().String("max-results", "", "The maximum number of objects that you want to return for this request.")
		route53profiles_listProfileAssociationsCmd.Flags().String("next-token", "", "For the first call to this list request, omit this value.")
		route53profiles_listProfileAssociationsCmd.Flags().String("profile-id", "", "ID of the Profile.")
		route53profiles_listProfileAssociationsCmd.Flags().String("resource-id", "", "ID of the VPC.")
	})
	route53profilesCmd.AddCommand(route53profiles_listProfileAssociationsCmd)
}
