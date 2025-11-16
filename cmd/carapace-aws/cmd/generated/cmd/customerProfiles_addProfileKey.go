package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_addProfileKeyCmd = &cobra.Command{
	Use:   "add-profile-key",
	Short: "Associates a new key value with a specific profile, such as a Contact Record ContactId.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_addProfileKeyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(customerProfiles_addProfileKeyCmd).Standalone()

		customerProfiles_addProfileKeyCmd.Flags().String("domain-name", "", "The unique name of the domain.")
		customerProfiles_addProfileKeyCmd.Flags().String("key-name", "", "A searchable identifier of a customer profile.")
		customerProfiles_addProfileKeyCmd.Flags().String("profile-id", "", "The unique identifier of a customer profile.")
		customerProfiles_addProfileKeyCmd.Flags().String("values", "", "A list of key values.")
		customerProfiles_addProfileKeyCmd.MarkFlagRequired("domain-name")
		customerProfiles_addProfileKeyCmd.MarkFlagRequired("key-name")
		customerProfiles_addProfileKeyCmd.MarkFlagRequired("profile-id")
		customerProfiles_addProfileKeyCmd.MarkFlagRequired("values")
	})
	customerProfilesCmd.AddCommand(customerProfiles_addProfileKeyCmd)
}
