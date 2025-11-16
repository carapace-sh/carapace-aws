package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_getProfileObjectTypeCmd = &cobra.Command{
	Use:   "get-profile-object-type",
	Short: "Returns the object types for a specific domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_getProfileObjectTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(customerProfiles_getProfileObjectTypeCmd).Standalone()

		customerProfiles_getProfileObjectTypeCmd.Flags().String("domain-name", "", "The unique name of the domain.")
		customerProfiles_getProfileObjectTypeCmd.Flags().String("object-type-name", "", "The name of the profile object type.")
		customerProfiles_getProfileObjectTypeCmd.MarkFlagRequired("domain-name")
		customerProfiles_getProfileObjectTypeCmd.MarkFlagRequired("object-type-name")
	})
	customerProfilesCmd.AddCommand(customerProfiles_getProfileObjectTypeCmd)
}
