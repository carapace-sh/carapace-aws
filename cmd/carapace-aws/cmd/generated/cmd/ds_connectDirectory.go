package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_connectDirectoryCmd = &cobra.Command{
	Use:   "connect-directory",
	Short: "Creates an AD Connector to connect to a self-managed directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_connectDirectoryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ds_connectDirectoryCmd).Standalone()

		ds_connectDirectoryCmd.Flags().String("connect-settings", "", "A [DirectoryConnectSettings]() object that contains additional information for the operation.")
		ds_connectDirectoryCmd.Flags().String("description", "", "A description for the directory.")
		ds_connectDirectoryCmd.Flags().String("name", "", "The fully qualified name of your self-managed directory, such as `corp.example.com`.")
		ds_connectDirectoryCmd.Flags().String("network-type", "", "The network type for your directory.")
		ds_connectDirectoryCmd.Flags().String("password", "", "The password for your self-managed user account.")
		ds_connectDirectoryCmd.Flags().String("short-name", "", "The NetBIOS name of your self-managed directory, such as `CORP`.")
		ds_connectDirectoryCmd.Flags().String("size", "", "The size of the directory.")
		ds_connectDirectoryCmd.Flags().String("tags", "", "The tags to be assigned to AD Connector.")
		ds_connectDirectoryCmd.MarkFlagRequired("connect-settings")
		ds_connectDirectoryCmd.MarkFlagRequired("name")
		ds_connectDirectoryCmd.MarkFlagRequired("password")
		ds_connectDirectoryCmd.MarkFlagRequired("size")
	})
	dsCmd.AddCommand(ds_connectDirectoryCmd)
}
