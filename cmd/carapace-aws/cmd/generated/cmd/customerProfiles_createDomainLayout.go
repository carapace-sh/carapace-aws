package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_createDomainLayoutCmd = &cobra.Command{
	Use:   "create-domain-layout",
	Short: "Creates the layout to view data for a specific domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_createDomainLayoutCmd).Standalone()

	customerProfiles_createDomainLayoutCmd.Flags().String("description", "", "The description of the layout")
	customerProfiles_createDomainLayoutCmd.Flags().String("display-name", "", "The display name of the layout")
	customerProfiles_createDomainLayoutCmd.Flags().String("domain-name", "", "The unique name of the domain.")
	customerProfiles_createDomainLayoutCmd.Flags().String("is-default", "", "If set to true for a layout, this layout will be used by default to view data.")
	customerProfiles_createDomainLayoutCmd.Flags().String("layout", "", "A customizable layout that can be used to view data under a Customer Profiles domain.")
	customerProfiles_createDomainLayoutCmd.Flags().String("layout-definition-name", "", "The unique name of the layout.")
	customerProfiles_createDomainLayoutCmd.Flags().String("layout-type", "", "The type of layout that can be used to view data under a Customer Profiles domain.")
	customerProfiles_createDomainLayoutCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
	customerProfiles_createDomainLayoutCmd.MarkFlagRequired("description")
	customerProfiles_createDomainLayoutCmd.MarkFlagRequired("display-name")
	customerProfiles_createDomainLayoutCmd.MarkFlagRequired("domain-name")
	customerProfiles_createDomainLayoutCmd.MarkFlagRequired("layout")
	customerProfiles_createDomainLayoutCmd.MarkFlagRequired("layout-definition-name")
	customerProfiles_createDomainLayoutCmd.MarkFlagRequired("layout-type")
	customerProfilesCmd.AddCommand(customerProfiles_createDomainLayoutCmd)
}
