package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_getKxVolumeCmd = &cobra.Command{
	Use:   "get-kx-volume",
	Short: "Retrieves the information about the volume.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_getKxVolumeCmd).Standalone()

	finspace_getKxVolumeCmd.Flags().String("environment-id", "", "A unique identifier for the kdb environment, whose clusters can attach to the volume.")
	finspace_getKxVolumeCmd.Flags().String("volume-name", "", "A unique identifier for the volume.")
	finspace_getKxVolumeCmd.MarkFlagRequired("environment-id")
	finspace_getKxVolumeCmd.MarkFlagRequired("volume-name")
	finspaceCmd.AddCommand(finspace_getKxVolumeCmd)
}
