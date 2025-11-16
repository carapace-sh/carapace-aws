package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_listSegmentDefinitionsCmd = &cobra.Command{
	Use:   "list-segment-definitions",
	Short: "Lists all segment definitions under a domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_listSegmentDefinitionsCmd).Standalone()

	customerProfiles_listSegmentDefinitionsCmd.Flags().String("domain-name", "", "The unique identifier of the domain.")
	customerProfiles_listSegmentDefinitionsCmd.Flags().String("max-results", "", "The maximum number of objects returned per page.")
	customerProfiles_listSegmentDefinitionsCmd.Flags().String("next-token", "", "The pagination token from the previous call.")
	customerProfiles_listSegmentDefinitionsCmd.MarkFlagRequired("domain-name")
	customerProfilesCmd.AddCommand(customerProfiles_listSegmentDefinitionsCmd)
}
