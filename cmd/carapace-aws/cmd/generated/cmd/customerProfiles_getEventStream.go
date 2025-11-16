package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_getEventStreamCmd = &cobra.Command{
	Use:   "get-event-stream",
	Short: "Returns information about the specified event stream in a specific domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_getEventStreamCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(customerProfiles_getEventStreamCmd).Standalone()

		customerProfiles_getEventStreamCmd.Flags().String("domain-name", "", "The unique name of the domain.")
		customerProfiles_getEventStreamCmd.Flags().String("event-stream-name", "", "The name of the event stream provided during create operations.")
		customerProfiles_getEventStreamCmd.MarkFlagRequired("domain-name")
		customerProfiles_getEventStreamCmd.MarkFlagRequired("event-stream-name")
	})
	customerProfilesCmd.AddCommand(customerProfiles_getEventStreamCmd)
}
