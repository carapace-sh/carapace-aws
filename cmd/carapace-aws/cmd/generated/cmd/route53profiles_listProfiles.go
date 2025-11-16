package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53profiles_listProfilesCmd = &cobra.Command{
	Use:   "list-profiles",
	Short: "Lists all the Route 53 Profiles associated with your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53profiles_listProfilesCmd).Standalone()

	route53profiles_listProfilesCmd.Flags().String("max-results", "", "The maximum number of objects that you want to return for this request.")
	route53profiles_listProfilesCmd.Flags().String("next-token", "", "For the first call to this list request, omit this value.")
	route53profilesCmd.AddCommand(route53profiles_listProfilesCmd)
}
