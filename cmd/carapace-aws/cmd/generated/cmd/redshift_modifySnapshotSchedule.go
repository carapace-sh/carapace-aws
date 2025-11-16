package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_modifySnapshotScheduleCmd = &cobra.Command{
	Use:   "modify-snapshot-schedule",
	Short: "Modifies a snapshot schedule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_modifySnapshotScheduleCmd).Standalone()

	redshift_modifySnapshotScheduleCmd.Flags().String("schedule-definitions", "", "An updated list of schedule definitions.")
	redshift_modifySnapshotScheduleCmd.Flags().String("schedule-identifier", "", "A unique alphanumeric identifier of the schedule to modify.")
	redshift_modifySnapshotScheduleCmd.MarkFlagRequired("schedule-definitions")
	redshift_modifySnapshotScheduleCmd.MarkFlagRequired("schedule-identifier")
	redshiftCmd.AddCommand(redshift_modifySnapshotScheduleCmd)
}
