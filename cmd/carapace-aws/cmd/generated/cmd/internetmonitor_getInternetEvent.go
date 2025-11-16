package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var internetmonitor_getInternetEventCmd = &cobra.Command{
	Use:   "get-internet-event",
	Short: "Gets information that Amazon CloudWatch Internet Monitor has generated about an internet event.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(internetmonitor_getInternetEventCmd).Standalone()

	internetmonitor_getInternetEventCmd.Flags().String("event-id", "", "The `EventId` of the internet event to return information for.")
	internetmonitor_getInternetEventCmd.MarkFlagRequired("event-id")
	internetmonitorCmd.AddCommand(internetmonitor_getInternetEventCmd)
}
