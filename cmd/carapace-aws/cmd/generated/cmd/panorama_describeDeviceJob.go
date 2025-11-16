package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var panorama_describeDeviceJobCmd = &cobra.Command{
	Use:   "describe-device-job",
	Short: "Returns information about a device job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(panorama_describeDeviceJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(panorama_describeDeviceJobCmd).Standalone()

		panorama_describeDeviceJobCmd.Flags().String("job-id", "", "The job's ID.")
		panorama_describeDeviceJobCmd.MarkFlagRequired("job-id")
	})
	panoramaCmd.AddCommand(panorama_describeDeviceJobCmd)
}
