package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_getEventTriggerCmd = &cobra.Command{
	Use:   "get-event-trigger",
	Short: "Get a specific Event Trigger from the domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_getEventTriggerCmd).Standalone()

	customerProfiles_getEventTriggerCmd.Flags().String("domain-name", "", "The unique name of the domain.")
	customerProfiles_getEventTriggerCmd.Flags().String("event-trigger-name", "", "The unique name of the event trigger.")
	customerProfiles_getEventTriggerCmd.MarkFlagRequired("domain-name")
	customerProfiles_getEventTriggerCmd.MarkFlagRequired("event-trigger-name")
	customerProfilesCmd.AddCommand(customerProfiles_getEventTriggerCmd)
}
