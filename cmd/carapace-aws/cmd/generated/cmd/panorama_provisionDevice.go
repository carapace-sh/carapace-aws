package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var panorama_provisionDeviceCmd = &cobra.Command{
	Use:   "provision-device",
	Short: "Creates a device and returns a configuration archive.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(panorama_provisionDeviceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(panorama_provisionDeviceCmd).Standalone()

		panorama_provisionDeviceCmd.Flags().String("description", "", "A description for the device.")
		panorama_provisionDeviceCmd.Flags().String("name", "", "A name for the device.")
		panorama_provisionDeviceCmd.Flags().String("networking-configuration", "", "A networking configuration for the device.")
		panorama_provisionDeviceCmd.Flags().String("tags", "", "Tags for the device.")
		panorama_provisionDeviceCmd.MarkFlagRequired("name")
	})
	panoramaCmd.AddCommand(panorama_provisionDeviceCmd)
}
