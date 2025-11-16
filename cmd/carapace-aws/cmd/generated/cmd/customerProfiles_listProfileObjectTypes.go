package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_listProfileObjectTypesCmd = &cobra.Command{
	Use:   "list-profile-object-types",
	Short: "Lists all of the templates available within the service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_listProfileObjectTypesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(customerProfiles_listProfileObjectTypesCmd).Standalone()

		customerProfiles_listProfileObjectTypesCmd.Flags().String("domain-name", "", "The unique name of the domain.")
		customerProfiles_listProfileObjectTypesCmd.Flags().String("max-results", "", "The maximum number of objects returned per page.")
		customerProfiles_listProfileObjectTypesCmd.Flags().String("next-token", "", "Identifies the next page of results to return.")
		customerProfiles_listProfileObjectTypesCmd.MarkFlagRequired("domain-name")
	})
	customerProfilesCmd.AddCommand(customerProfiles_listProfileObjectTypesCmd)
}
