package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_createArchiveCmd = &cobra.Command{
	Use:   "create-archive",
	Short: "Creates an archive of events with the specified settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_createArchiveCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(events_createArchiveCmd).Standalone()

		events_createArchiveCmd.Flags().String("archive-name", "", "The name for the archive to create.")
		events_createArchiveCmd.Flags().String("description", "", "A description for the archive.")
		events_createArchiveCmd.Flags().String("event-pattern", "", "An event pattern to use to filter events sent to the archive.")
		events_createArchiveCmd.Flags().String("event-source-arn", "", "The ARN of the event bus that sends events to the archive.")
		events_createArchiveCmd.Flags().String("kms-key-identifier", "", "The identifier of the KMS customer managed key for EventBridge to use, if you choose to use a customer managed key to encrypt this archive.")
		events_createArchiveCmd.Flags().String("retention-days", "", "The number of days to retain events for.")
		events_createArchiveCmd.MarkFlagRequired("archive-name")
		events_createArchiveCmd.MarkFlagRequired("event-source-arn")
	})
	eventsCmd.AddCommand(events_createArchiveCmd)
}
