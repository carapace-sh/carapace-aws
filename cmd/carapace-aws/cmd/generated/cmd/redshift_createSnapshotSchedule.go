package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_createSnapshotScheduleCmd = &cobra.Command{
	Use:   "create-snapshot-schedule",
	Short: "Create a snapshot schedule that can be associated to a cluster and which overrides the default system backup schedule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_createSnapshotScheduleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_createSnapshotScheduleCmd).Standalone()

		redshift_createSnapshotScheduleCmd.Flags().String("dry-run", "", "")
		redshift_createSnapshotScheduleCmd.Flags().String("next-invocations", "", "")
		redshift_createSnapshotScheduleCmd.Flags().String("schedule-definitions", "", "The definition of the snapshot schedule.")
		redshift_createSnapshotScheduleCmd.Flags().String("schedule-description", "", "The description of the snapshot schedule.")
		redshift_createSnapshotScheduleCmd.Flags().String("schedule-identifier", "", "A unique identifier for a snapshot schedule.")
		redshift_createSnapshotScheduleCmd.Flags().String("tags", "", "An optional set of tags you can use to search for the schedule.")
	})
	redshiftCmd.AddCommand(redshift_createSnapshotScheduleCmd)
}
