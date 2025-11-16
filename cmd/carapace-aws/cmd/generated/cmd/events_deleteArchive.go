package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_deleteArchiveCmd = &cobra.Command{
	Use:   "delete-archive",
	Short: "Deletes the specified archive.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_deleteArchiveCmd).Standalone()

	events_deleteArchiveCmd.Flags().String("archive-name", "", "The name of the archive to delete.")
	events_deleteArchiveCmd.MarkFlagRequired("archive-name")
	eventsCmd.AddCommand(events_deleteArchiveCmd)
}
