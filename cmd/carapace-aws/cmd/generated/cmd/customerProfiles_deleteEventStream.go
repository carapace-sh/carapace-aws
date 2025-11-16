package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_deleteEventStreamCmd = &cobra.Command{
	Use:   "delete-event-stream",
	Short: "Disables and deletes the specified event stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_deleteEventStreamCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(customerProfiles_deleteEventStreamCmd).Standalone()

		customerProfiles_deleteEventStreamCmd.Flags().String("domain-name", "", "The unique name of the domain.")
		customerProfiles_deleteEventStreamCmd.Flags().String("event-stream-name", "", "The name of the event stream")
		customerProfiles_deleteEventStreamCmd.MarkFlagRequired("domain-name")
		customerProfiles_deleteEventStreamCmd.MarkFlagRequired("event-stream-name")
	})
	customerProfilesCmd.AddCommand(customerProfiles_deleteEventStreamCmd)
}
