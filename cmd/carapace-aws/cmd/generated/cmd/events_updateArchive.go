package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_updateArchiveCmd = &cobra.Command{
	Use:   "update-archive",
	Short: "Updates the specified archive.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_updateArchiveCmd).Standalone()

	events_updateArchiveCmd.Flags().String("archive-name", "", "The name of the archive to update.")
	events_updateArchiveCmd.Flags().String("description", "", "The description for the archive.")
	events_updateArchiveCmd.Flags().String("event-pattern", "", "The event pattern to use to filter events sent to the archive.")
	events_updateArchiveCmd.Flags().String("kms-key-identifier", "", "The identifier of the KMS customer managed key for EventBridge to use, if you choose to use a customer managed key to encrypt this archive.")
	events_updateArchiveCmd.Flags().String("retention-days", "", "The number of days to retain events in the archive.")
	events_updateArchiveCmd.MarkFlagRequired("archive-name")
	eventsCmd.AddCommand(events_updateArchiveCmd)
}
