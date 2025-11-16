package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_detectProfileObjectTypeCmd = &cobra.Command{
	Use:   "detect-profile-object-type",
	Short: "The process of detecting profile object type mapping by using given objects.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_detectProfileObjectTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(customerProfiles_detectProfileObjectTypeCmd).Standalone()

		customerProfiles_detectProfileObjectTypeCmd.Flags().String("domain-name", "", "The unique name of the domain.")
		customerProfiles_detectProfileObjectTypeCmd.Flags().String("objects", "", "A string that is serialized from a JSON object.")
		customerProfiles_detectProfileObjectTypeCmd.MarkFlagRequired("domain-name")
		customerProfiles_detectProfileObjectTypeCmd.MarkFlagRequired("objects")
	})
	customerProfilesCmd.AddCommand(customerProfiles_detectProfileObjectTypeCmd)
}
