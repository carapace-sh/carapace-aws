package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_getDeviceFleetReportCmd = &cobra.Command{
	Use:   "get-device-fleet-report",
	Short: "Describes a fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_getDeviceFleetReportCmd).Standalone()

	sagemaker_getDeviceFleetReportCmd.Flags().String("device-fleet-name", "", "The name of the fleet.")
	sagemaker_getDeviceFleetReportCmd.MarkFlagRequired("device-fleet-name")
	sagemakerCmd.AddCommand(sagemaker_getDeviceFleetReportCmd)
}
