package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_describeArchiveCmd = &cobra.Command{
	Use:   "describe-archive",
	Short: "Retrieves details about an archive.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_describeArchiveCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(events_describeArchiveCmd).Standalone()

		events_describeArchiveCmd.Flags().String("archive-name", "", "The name of the archive to retrieve.")
		events_describeArchiveCmd.MarkFlagRequired("archive-name")
	})
	eventsCmd.AddCommand(events_describeArchiveCmd)
}
