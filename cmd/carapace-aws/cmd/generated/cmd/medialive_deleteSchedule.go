package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_deleteScheduleCmd = &cobra.Command{
	Use:   "delete-schedule",
	Short: "Delete all schedule actions on a channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_deleteScheduleCmd).Standalone()

	medialive_deleteScheduleCmd.Flags().String("channel-id", "", "Id of the channel whose schedule is being deleted.")
	medialive_deleteScheduleCmd.MarkFlagRequired("channel-id")
	medialiveCmd.AddCommand(medialive_deleteScheduleCmd)
}
