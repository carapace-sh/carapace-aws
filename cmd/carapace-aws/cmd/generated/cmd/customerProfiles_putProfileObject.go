package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_putProfileObjectCmd = &cobra.Command{
	Use:   "put-profile-object",
	Short: "Adds additional objects to customer profiles of a given ObjectType.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_putProfileObjectCmd).Standalone()

	customerProfiles_putProfileObjectCmd.Flags().String("domain-name", "", "The unique name of the domain.")
	customerProfiles_putProfileObjectCmd.Flags().String("object", "", "A string that is serialized from a JSON object.")
	customerProfiles_putProfileObjectCmd.Flags().String("object-type-name", "", "The name of the profile object type.")
	customerProfiles_putProfileObjectCmd.MarkFlagRequired("domain-name")
	customerProfiles_putProfileObjectCmd.MarkFlagRequired("object")
	customerProfiles_putProfileObjectCmd.MarkFlagRequired("object-type-name")
	customerProfilesCmd.AddCommand(customerProfiles_putProfileObjectCmd)
}
