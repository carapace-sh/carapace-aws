package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_createVolumeFromBackupCmd = &cobra.Command{
	Use:   "create-volume-from-backup",
	Short: "Creates a new Amazon FSx for NetApp ONTAP volume from an existing Amazon FSx volume backup.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_createVolumeFromBackupCmd).Standalone()

	fsx_createVolumeFromBackupCmd.Flags().String("backup-id", "", "")
	fsx_createVolumeFromBackupCmd.Flags().String("client-request-token", "", "")
	fsx_createVolumeFromBackupCmd.Flags().String("name", "", "The name of the new volume you're creating.")
	fsx_createVolumeFromBackupCmd.Flags().String("ontap-configuration", "", "Specifies the configuration of the ONTAP volume that you are creating.")
	fsx_createVolumeFromBackupCmd.Flags().String("tags", "", "")
	fsx_createVolumeFromBackupCmd.MarkFlagRequired("backup-id")
	fsx_createVolumeFromBackupCmd.MarkFlagRequired("name")
	fsxCmd.AddCommand(fsx_createVolumeFromBackupCmd)
}
