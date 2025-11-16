package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_updateKxVolumeCmd = &cobra.Command{
	Use:   "update-kx-volume",
	Short: "Updates the throughput or capacity of a volume.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_updateKxVolumeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(finspace_updateKxVolumeCmd).Standalone()

		finspace_updateKxVolumeCmd.Flags().String("client-token", "", "A token that ensures idempotency.")
		finspace_updateKxVolumeCmd.Flags().String("description", "", "A description of the volume.")
		finspace_updateKxVolumeCmd.Flags().String("environment-id", "", "A unique identifier for the kdb environment where you created the storage volume.")
		finspace_updateKxVolumeCmd.Flags().String("nas1-configuration", "", "Specifies the configuration for the Network attached storage (NAS\\_1) file system volume.")
		finspace_updateKxVolumeCmd.Flags().String("volume-name", "", "A unique identifier for the volume.")
		finspace_updateKxVolumeCmd.MarkFlagRequired("environment-id")
		finspace_updateKxVolumeCmd.MarkFlagRequired("volume-name")
	})
	finspaceCmd.AddCommand(finspace_updateKxVolumeCmd)
}
