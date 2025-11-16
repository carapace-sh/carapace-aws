package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_putProfileObjectTypeCmd = &cobra.Command{
	Use:   "put-profile-object-type",
	Short: "Defines a ProfileObjectType.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_putProfileObjectTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(customerProfiles_putProfileObjectTypeCmd).Standalone()

		customerProfiles_putProfileObjectTypeCmd.Flags().String("allow-profile-creation", "", "Indicates whether a profile should be created when data is received if one doesn’t exist for an object of this type.")
		customerProfiles_putProfileObjectTypeCmd.Flags().String("description", "", "Description of the profile object type.")
		customerProfiles_putProfileObjectTypeCmd.Flags().String("domain-name", "", "The unique name of the domain.")
		customerProfiles_putProfileObjectTypeCmd.Flags().String("encryption-key", "", "The customer-provided key to encrypt the profile object that will be created in this profile object type.")
		customerProfiles_putProfileObjectTypeCmd.Flags().String("expiration-days", "", "The number of days until the data in the object expires.")
		customerProfiles_putProfileObjectTypeCmd.Flags().String("fields", "", "A map of the name and ObjectType field.")
		customerProfiles_putProfileObjectTypeCmd.Flags().String("keys", "", "A list of unique keys that can be used to map data to the profile.")
		customerProfiles_putProfileObjectTypeCmd.Flags().String("max-profile-object-count", "", "The amount of profile object max count assigned to the object type")
		customerProfiles_putProfileObjectTypeCmd.Flags().String("object-type-name", "", "The name of the profile object type.")
		customerProfiles_putProfileObjectTypeCmd.Flags().String("source-last-updated-timestamp-format", "", "The format of your `sourceLastUpdatedTimestamp` that was previously set up.")
		customerProfiles_putProfileObjectTypeCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
		customerProfiles_putProfileObjectTypeCmd.Flags().String("template-id", "", "A unique identifier for the object template.")
		customerProfiles_putProfileObjectTypeCmd.MarkFlagRequired("description")
		customerProfiles_putProfileObjectTypeCmd.MarkFlagRequired("domain-name")
		customerProfiles_putProfileObjectTypeCmd.MarkFlagRequired("object-type-name")
	})
	customerProfilesCmd.AddCommand(customerProfiles_putProfileObjectTypeCmd)
}
