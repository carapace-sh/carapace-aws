package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_modifyClusterMaintenanceCmd = &cobra.Command{
	Use:   "modify-cluster-maintenance",
	Short: "Modifies the maintenance settings of a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_modifyClusterMaintenanceCmd).Standalone()

	redshift_modifyClusterMaintenanceCmd.Flags().String("cluster-identifier", "", "A unique identifier for the cluster.")
	redshift_modifyClusterMaintenanceCmd.Flags().String("defer-maintenance", "", "A boolean indicating whether to enable the deferred maintenance window.")
	redshift_modifyClusterMaintenanceCmd.Flags().String("defer-maintenance-duration", "", "An integer indicating the duration of the maintenance window in days.")
	redshift_modifyClusterMaintenanceCmd.Flags().String("defer-maintenance-end-time", "", "A timestamp indicating end time for the deferred maintenance window.")
	redshift_modifyClusterMaintenanceCmd.Flags().String("defer-maintenance-identifier", "", "A unique identifier for the deferred maintenance window.")
	redshift_modifyClusterMaintenanceCmd.Flags().String("defer-maintenance-start-time", "", "A timestamp indicating the start time for the deferred maintenance window.")
	redshift_modifyClusterMaintenanceCmd.MarkFlagRequired("cluster-identifier")
	redshiftCmd.AddCommand(redshift_modifyClusterMaintenanceCmd)
}
