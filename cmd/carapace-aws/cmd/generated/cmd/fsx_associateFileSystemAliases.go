package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_associateFileSystemAliasesCmd = &cobra.Command{
	Use:   "associate-file-system-aliases",
	Short: "Use this action to associate one or more Domain Name Server (DNS) aliases with an existing Amazon FSx for Windows File Server file system.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_associateFileSystemAliasesCmd).Standalone()

	fsx_associateFileSystemAliasesCmd.Flags().String("aliases", "", "An array of one or more DNS alias names to associate with the file system.")
	fsx_associateFileSystemAliasesCmd.Flags().String("client-request-token", "", "")
	fsx_associateFileSystemAliasesCmd.Flags().String("file-system-id", "", "Specifies the file system with which you want to associate one or more DNS aliases.")
	fsx_associateFileSystemAliasesCmd.MarkFlagRequired("aliases")
	fsx_associateFileSystemAliasesCmd.MarkFlagRequired("file-system-id")
	fsxCmd.AddCommand(fsx_associateFileSystemAliasesCmd)
}
