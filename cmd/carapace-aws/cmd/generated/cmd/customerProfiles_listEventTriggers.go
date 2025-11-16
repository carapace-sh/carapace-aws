package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_listEventTriggersCmd = &cobra.Command{
	Use:   "list-event-triggers",
	Short: "List all Event Triggers under a domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_listEventTriggersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(customerProfiles_listEventTriggersCmd).Standalone()

		customerProfiles_listEventTriggersCmd.Flags().String("domain-name", "", "The unique name of the domain.")
		customerProfiles_listEventTriggersCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		customerProfiles_listEventTriggersCmd.Flags().String("next-token", "", "The pagination token to use with ListEventTriggers.")
		customerProfiles_listEventTriggersCmd.MarkFlagRequired("domain-name")
	})
	customerProfilesCmd.AddCommand(customerProfiles_listEventTriggersCmd)
}
