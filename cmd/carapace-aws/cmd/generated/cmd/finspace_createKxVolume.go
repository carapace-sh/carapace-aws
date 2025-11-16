package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_createKxVolumeCmd = &cobra.Command{
	Use:   "create-kx-volume",
	Short: "Creates a new volume with a specific amount of throughput and storage capacity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_createKxVolumeCmd).Standalone()

	finspace_createKxVolumeCmd.Flags().String("availability-zone-ids", "", "The identifier of the availability zones.")
	finspace_createKxVolumeCmd.Flags().String("az-mode", "", "The number of availability zones you want to assign per volume.")
	finspace_createKxVolumeCmd.Flags().String("client-token", "", "A token that ensures idempotency.")
	finspace_createKxVolumeCmd.Flags().String("description", "", "A description of the volume.")
	finspace_createKxVolumeCmd.Flags().String("environment-id", "", "A unique identifier for the kdb environment, whose clusters can attach to the volume.")
	finspace_createKxVolumeCmd.Flags().String("nas1-configuration", "", "Specifies the configuration for the Network attached storage (NAS\\_1) file system volume.")
	finspace_createKxVolumeCmd.Flags().String("tags", "", "A list of key-value pairs to label the volume.")
	finspace_createKxVolumeCmd.Flags().String("volume-name", "", "A unique identifier for the volume.")
	finspace_createKxVolumeCmd.Flags().String("volume-type", "", "The type of file system volume.")
	finspace_createKxVolumeCmd.MarkFlagRequired("availability-zone-ids")
	finspace_createKxVolumeCmd.MarkFlagRequired("az-mode")
	finspace_createKxVolumeCmd.MarkFlagRequired("environment-id")
	finspace_createKxVolumeCmd.MarkFlagRequired("volume-name")
	finspace_createKxVolumeCmd.MarkFlagRequired("volume-type")
	finspaceCmd.AddCommand(finspace_createKxVolumeCmd)
}
