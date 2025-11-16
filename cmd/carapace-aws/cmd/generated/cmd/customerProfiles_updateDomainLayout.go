package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_updateDomainLayoutCmd = &cobra.Command{
	Use:   "update-domain-layout",
	Short: "Updates the layout used to view data for a specific domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_updateDomainLayoutCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(customerProfiles_updateDomainLayoutCmd).Standalone()

		customerProfiles_updateDomainLayoutCmd.Flags().String("description", "", "The description of the layout")
		customerProfiles_updateDomainLayoutCmd.Flags().String("display-name", "", "The display name of the layout")
		customerProfiles_updateDomainLayoutCmd.Flags().String("domain-name", "", "The unique name of the domain.")
		customerProfiles_updateDomainLayoutCmd.Flags().String("is-default", "", "If set to true for a layout, this layout will be used by default to view data.")
		customerProfiles_updateDomainLayoutCmd.Flags().String("layout", "", "A customizable layout that can be used to view data under a Customer Profiles domain.")
		customerProfiles_updateDomainLayoutCmd.Flags().String("layout-definition-name", "", "The unique name of the layout.")
		customerProfiles_updateDomainLayoutCmd.Flags().String("layout-type", "", "The type of layout that can be used to view data under a Customer Profiles domain.")
		customerProfiles_updateDomainLayoutCmd.MarkFlagRequired("domain-name")
		customerProfiles_updateDomainLayoutCmd.MarkFlagRequired("layout-definition-name")
	})
	customerProfilesCmd.AddCommand(customerProfiles_updateDomainLayoutCmd)
}
