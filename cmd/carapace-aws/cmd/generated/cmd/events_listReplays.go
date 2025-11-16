package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_listReplaysCmd = &cobra.Command{
	Use:   "list-replays",
	Short: "Lists your replays.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_listReplaysCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(events_listReplaysCmd).Standalone()

		events_listReplaysCmd.Flags().String("event-source-arn", "", "The ARN of the archive from which the events are replayed.")
		events_listReplaysCmd.Flags().String("limit", "", "The maximum number of replays to retrieve.")
		events_listReplaysCmd.Flags().String("name-prefix", "", "A name prefix to filter the replays returned.")
		events_listReplaysCmd.Flags().String("next-token", "", "The token returned by a previous call, which you can use to retrieve the next set of results.")
		events_listReplaysCmd.Flags().String("state", "", "The state of the replay.")
	})
	eventsCmd.AddCommand(events_listReplaysCmd)
}
