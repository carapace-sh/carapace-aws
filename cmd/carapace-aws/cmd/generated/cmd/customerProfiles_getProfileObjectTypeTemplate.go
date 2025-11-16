package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_getProfileObjectTypeTemplateCmd = &cobra.Command{
	Use:   "get-profile-object-type-template",
	Short: "Returns the template information for a specific object type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_getProfileObjectTypeTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(customerProfiles_getProfileObjectTypeTemplateCmd).Standalone()

		customerProfiles_getProfileObjectTypeTemplateCmd.Flags().String("template-id", "", "A unique identifier for the object template.")
		customerProfiles_getProfileObjectTypeTemplateCmd.MarkFlagRequired("template-id")
	})
	customerProfilesCmd.AddCommand(customerProfiles_getProfileObjectTypeTemplateCmd)
}
