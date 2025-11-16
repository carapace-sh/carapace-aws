package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_deleteProfileCmd = &cobra.Command{
	Use:   "delete-profile",
	Short: "Deletes the standard customer profile and all data pertaining to the profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_deleteProfileCmd).Standalone()

	customerProfiles_deleteProfileCmd.Flags().String("domain-name", "", "The unique name of the domain.")
	customerProfiles_deleteProfileCmd.Flags().String("profile-id", "", "The unique identifier of a customer profile.")
	customerProfiles_deleteProfileCmd.MarkFlagRequired("domain-name")
	customerProfiles_deleteProfileCmd.MarkFlagRequired("profile-id")
	customerProfilesCmd.AddCommand(customerProfiles_deleteProfileCmd)
}
