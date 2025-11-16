package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_modifyClusterSnapshotCmd = &cobra.Command{
	Use:   "modify-cluster-snapshot",
	Short: "Modifies the settings for a snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_modifyClusterSnapshotCmd).Standalone()

	redshift_modifyClusterSnapshotCmd.Flags().Bool("force", false, "A Boolean option to override an exception if the retention period has already passed.")
	redshift_modifyClusterSnapshotCmd.Flags().String("manual-snapshot-retention-period", "", "The number of days that a manual snapshot is retained.")
	redshift_modifyClusterSnapshotCmd.Flags().Bool("no-force", false, "A Boolean option to override an exception if the retention period has already passed.")
	redshift_modifyClusterSnapshotCmd.Flags().String("snapshot-identifier", "", "The identifier of the snapshot whose setting you want to modify.")
	redshift_modifyClusterSnapshotCmd.Flag("no-force").Hidden = true
	redshift_modifyClusterSnapshotCmd.MarkFlagRequired("snapshot-identifier")
	redshiftCmd.AddCommand(redshift_modifyClusterSnapshotCmd)
}
