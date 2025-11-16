package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_deleteProfileKeyCmd = &cobra.Command{
	Use:   "delete-profile-key",
	Short: "Removes a searchable key from a customer profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_deleteProfileKeyCmd).Standalone()

	customerProfiles_deleteProfileKeyCmd.Flags().String("domain-name", "", "The unique name of the domain.")
	customerProfiles_deleteProfileKeyCmd.Flags().String("key-name", "", "A searchable identifier of a customer profile.")
	customerProfiles_deleteProfileKeyCmd.Flags().String("profile-id", "", "The unique identifier of a customer profile.")
	customerProfiles_deleteProfileKeyCmd.Flags().String("values", "", "A list of key values.")
	customerProfiles_deleteProfileKeyCmd.MarkFlagRequired("domain-name")
	customerProfiles_deleteProfileKeyCmd.MarkFlagRequired("key-name")
	customerProfiles_deleteProfileKeyCmd.MarkFlagRequired("profile-id")
	customerProfiles_deleteProfileKeyCmd.MarkFlagRequired("values")
	customerProfilesCmd.AddCommand(customerProfiles_deleteProfileKeyCmd)
}
