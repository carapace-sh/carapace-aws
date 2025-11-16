package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_modifyClusterSnapshotScheduleCmd = &cobra.Command{
	Use:   "modify-cluster-snapshot-schedule",
	Short: "Modifies a snapshot schedule for a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_modifyClusterSnapshotScheduleCmd).Standalone()

	redshift_modifyClusterSnapshotScheduleCmd.Flags().String("cluster-identifier", "", "A unique identifier for the cluster whose snapshot schedule you want to modify.")
	redshift_modifyClusterSnapshotScheduleCmd.Flags().String("disassociate-schedule", "", "A boolean to indicate whether to remove the assoiciation between the cluster and the schedule.")
	redshift_modifyClusterSnapshotScheduleCmd.Flags().String("schedule-identifier", "", "A unique alphanumeric identifier for the schedule that you want to associate with the cluster.")
	redshift_modifyClusterSnapshotScheduleCmd.MarkFlagRequired("cluster-identifier")
	redshiftCmd.AddCommand(redshift_modifyClusterSnapshotScheduleCmd)
}
