package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var panorama_createJobForDevicesCmd = &cobra.Command{
	Use:   "create-job-for-devices",
	Short: "Creates a job to run on a device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(panorama_createJobForDevicesCmd).Standalone()

	panorama_createJobForDevicesCmd.Flags().String("device-ids", "", "ID of target device.")
	panorama_createJobForDevicesCmd.Flags().String("device-job-config", "", "Configuration settings for a software update job.")
	panorama_createJobForDevicesCmd.Flags().String("job-type", "", "The type of job to run.")
	panorama_createJobForDevicesCmd.MarkFlagRequired("device-ids")
	panorama_createJobForDevicesCmd.MarkFlagRequired("job-type")
	panoramaCmd.AddCommand(panorama_createJobForDevicesCmd)
}
