package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_createDirectoryCmd = &cobra.Command{
	Use:   "create-directory",
	Short: "Creates a Simple AD directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_createDirectoryCmd).Standalone()

	ds_createDirectoryCmd.Flags().String("description", "", "A description for the directory.")
	ds_createDirectoryCmd.Flags().String("name", "", "The fully qualified name for the directory, such as `corp.example.com`.")
	ds_createDirectoryCmd.Flags().String("network-type", "", "The network type for your directory.")
	ds_createDirectoryCmd.Flags().String("password", "", "The password for the directory administrator.")
	ds_createDirectoryCmd.Flags().String("short-name", "", "The NetBIOS name of the directory, such as `CORP`.")
	ds_createDirectoryCmd.Flags().String("size", "", "The size of the directory.")
	ds_createDirectoryCmd.Flags().String("tags", "", "The tags to be assigned to the Simple AD directory.")
	ds_createDirectoryCmd.Flags().String("vpc-settings", "", "A [DirectoryVpcSettings]() object that contains additional information for the operation.")
	ds_createDirectoryCmd.MarkFlagRequired("name")
	ds_createDirectoryCmd.MarkFlagRequired("password")
	ds_createDirectoryCmd.MarkFlagRequired("size")
	dsCmd.AddCommand(ds_createDirectoryCmd)
}
