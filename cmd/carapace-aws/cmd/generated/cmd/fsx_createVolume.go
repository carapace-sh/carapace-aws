package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_createVolumeCmd = &cobra.Command{
	Use:   "create-volume",
	Short: "Creates an FSx for ONTAP or Amazon FSx for OpenZFS storage volume.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_createVolumeCmd).Standalone()

	fsx_createVolumeCmd.Flags().String("client-request-token", "", "")
	fsx_createVolumeCmd.Flags().String("name", "", "Specifies the name of the volume that you're creating.")
	fsx_createVolumeCmd.Flags().String("ontap-configuration", "", "Specifies the configuration to use when creating the ONTAP volume.")
	fsx_createVolumeCmd.Flags().String("open-zfsconfiguration", "", "Specifies the configuration to use when creating the OpenZFS volume.")
	fsx_createVolumeCmd.Flags().String("tags", "", "")
	fsx_createVolumeCmd.Flags().String("volume-type", "", "Specifies the type of volume to create; `ONTAP` and `OPENZFS` are the only valid volume types.")
	fsx_createVolumeCmd.MarkFlagRequired("name")
	fsx_createVolumeCmd.MarkFlagRequired("volume-type")
	fsxCmd.AddCommand(fsx_createVolumeCmd)
}
