package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_updateDirectorySetupCmd = &cobra.Command{
	Use:   "update-directory-setup",
	Short: "Updates directory configuration for the specified update type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_updateDirectorySetupCmd).Standalone()

	ds_updateDirectorySetupCmd.Flags().String("create-snapshot-before-update", "", "Specifies whether to create a directory snapshot before performing the update.")
	ds_updateDirectorySetupCmd.Flags().String("directory-id", "", "The identifier of the directory to update.")
	ds_updateDirectorySetupCmd.Flags().String("directory-size-update-settings", "", "Directory size configuration to apply during the update operation.")
	ds_updateDirectorySetupCmd.Flags().String("network-update-settings", "", "Network configuration to apply during the directory update operation.")
	ds_updateDirectorySetupCmd.Flags().String("osupdate-settings", "", "Operating system configuration to apply during the directory update operation.")
	ds_updateDirectorySetupCmd.Flags().String("update-type", "", "The type of update to perform on the directory.")
	ds_updateDirectorySetupCmd.MarkFlagRequired("directory-id")
	ds_updateDirectorySetupCmd.MarkFlagRequired("update-type")
	dsCmd.AddCommand(ds_updateDirectorySetupCmd)
}
