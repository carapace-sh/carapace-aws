package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_disassociateFileSystemAliasesCmd = &cobra.Command{
	Use:   "disassociate-file-system-aliases",
	Short: "Use this action to disassociate, or remove, one or more Domain Name Service (DNS) aliases from an Amazon FSx for Windows File Server file system.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_disassociateFileSystemAliasesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fsx_disassociateFileSystemAliasesCmd).Standalone()

		fsx_disassociateFileSystemAliasesCmd.Flags().String("aliases", "", "An array of one or more DNS alias names to disassociate, or remove, from the file system.")
		fsx_disassociateFileSystemAliasesCmd.Flags().String("client-request-token", "", "")
		fsx_disassociateFileSystemAliasesCmd.Flags().String("file-system-id", "", "Specifies the file system from which to disassociate the DNS aliases.")
		fsx_disassociateFileSystemAliasesCmd.MarkFlagRequired("aliases")
		fsx_disassociateFileSystemAliasesCmd.MarkFlagRequired("file-system-id")
	})
	fsxCmd.AddCommand(fsx_disassociateFileSystemAliasesCmd)
}
