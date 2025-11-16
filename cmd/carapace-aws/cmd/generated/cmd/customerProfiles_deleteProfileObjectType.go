package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_deleteProfileObjectTypeCmd = &cobra.Command{
	Use:   "delete-profile-object-type",
	Short: "Removes a ProfileObjectType from a specific domain as well as removes all the ProfileObjects of that type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_deleteProfileObjectTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(customerProfiles_deleteProfileObjectTypeCmd).Standalone()

		customerProfiles_deleteProfileObjectTypeCmd.Flags().String("domain-name", "", "The unique name of the domain.")
		customerProfiles_deleteProfileObjectTypeCmd.Flags().String("object-type-name", "", "The name of the profile object type.")
		customerProfiles_deleteProfileObjectTypeCmd.MarkFlagRequired("domain-name")
		customerProfiles_deleteProfileObjectTypeCmd.MarkFlagRequired("object-type-name")
	})
	customerProfilesCmd.AddCommand(customerProfiles_deleteProfileObjectTypeCmd)
}
