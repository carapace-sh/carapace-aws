package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var databrew_deleteScheduleCmd = &cobra.Command{
	Use:   "delete-schedule",
	Short: "Deletes the specified DataBrew schedule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databrew_deleteScheduleCmd).Standalone()

	databrew_deleteScheduleCmd.Flags().String("name", "", "The name of the schedule to be deleted.")
	databrew_deleteScheduleCmd.MarkFlagRequired("name")
	databrewCmd.AddCommand(databrew_deleteScheduleCmd)
}
