package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var arcZonalShift_updateZonalAutoshiftConfigurationCmd = &cobra.Command{
	Use:   "update-zonal-autoshift-configuration",
	Short: "The zonal autoshift configuration for a resource includes the practice run configuration and the status for running autoshifts, zonal autoshift status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(arcZonalShift_updateZonalAutoshiftConfigurationCmd).Standalone()

	arcZonalShift_updateZonalAutoshiftConfigurationCmd.Flags().String("resource-identifier", "", "The identifier for the resource that you want to update the zonal autoshift configuration for.")
	arcZonalShift_updateZonalAutoshiftConfigurationCmd.Flags().String("zonal-autoshift-status", "", "The zonal autoshift status for the resource that you want to update the zonal autoshift configuration for.")
	arcZonalShift_updateZonalAutoshiftConfigurationCmd.MarkFlagRequired("resource-identifier")
	arcZonalShift_updateZonalAutoshiftConfigurationCmd.MarkFlagRequired("zonal-autoshift-status")
	arcZonalShiftCmd.AddCommand(arcZonalShift_updateZonalAutoshiftConfigurationCmd)
}
