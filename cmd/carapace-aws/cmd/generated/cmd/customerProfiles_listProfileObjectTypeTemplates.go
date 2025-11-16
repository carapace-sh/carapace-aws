package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_listProfileObjectTypeTemplatesCmd = &cobra.Command{
	Use:   "list-profile-object-type-templates",
	Short: "Lists all of the template information for object types.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_listProfileObjectTypeTemplatesCmd).Standalone()

	customerProfiles_listProfileObjectTypeTemplatesCmd.Flags().String("max-results", "", "The maximum number of objects returned per page.")
	customerProfiles_listProfileObjectTypeTemplatesCmd.Flags().String("next-token", "", "The pagination token from the previous ListObjectTypeTemplates API call.")
	customerProfilesCmd.AddCommand(customerProfiles_listProfileObjectTypeTemplatesCmd)
}
