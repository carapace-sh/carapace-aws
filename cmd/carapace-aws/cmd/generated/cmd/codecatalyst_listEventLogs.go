package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecatalyst_listEventLogsCmd = &cobra.Command{
	Use:   "list-event-logs",
	Short: "Retrieves a list of events that occurred during a specific time in a space.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecatalyst_listEventLogsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecatalyst_listEventLogsCmd).Standalone()

		codecatalyst_listEventLogsCmd.Flags().String("end-time", "", "The time after which you do not want any events retrieved, in coordinated universal time (UTC) timestamp format as specified in [RFC 3339](https://www.rfc-editor.org/rfc/rfc3339#section-5.6).")
		codecatalyst_listEventLogsCmd.Flags().String("event-name", "", "The name of the event.")
		codecatalyst_listEventLogsCmd.Flags().String("max-results", "", "The maximum number of results to show in a single call to this API.")
		codecatalyst_listEventLogsCmd.Flags().String("next-token", "", "A token returned from a call to this API to indicate the next batch of results to return, if any.")
		codecatalyst_listEventLogsCmd.Flags().String("space-name", "", "The name of the space.")
		codecatalyst_listEventLogsCmd.Flags().String("start-time", "", "The date and time when you want to start retrieving events, in coordinated universal time (UTC) timestamp format as specified in [RFC 3339](https://www.rfc-editor.org/rfc/rfc3339#section-5.6).")
		codecatalyst_listEventLogsCmd.MarkFlagRequired("end-time")
		codecatalyst_listEventLogsCmd.MarkFlagRequired("space-name")
		codecatalyst_listEventLogsCmd.MarkFlagRequired("start-time")
	})
	codecatalystCmd.AddCommand(codecatalyst_listEventLogsCmd)
}
