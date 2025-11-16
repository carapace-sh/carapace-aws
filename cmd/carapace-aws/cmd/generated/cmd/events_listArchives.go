package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_listArchivesCmd = &cobra.Command{
	Use:   "list-archives",
	Short: "Lists your archives.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_listArchivesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(events_listArchivesCmd).Standalone()

		events_listArchivesCmd.Flags().String("event-source-arn", "", "The ARN of the event source associated with the archive.")
		events_listArchivesCmd.Flags().String("limit", "", "The maximum number of results to return.")
		events_listArchivesCmd.Flags().String("name-prefix", "", "A name prefix to filter the archives returned.")
		events_listArchivesCmd.Flags().String("next-token", "", "The token returned by a previous call, which you can use to retrieve the next set of results.")
		events_listArchivesCmd.Flags().String("state", "", "The state of the archive.")
	})
	eventsCmd.AddCommand(events_listArchivesCmd)
}
