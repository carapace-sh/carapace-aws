package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_describeFileSystemAliasesCmd = &cobra.Command{
	Use:   "describe-file-system-aliases",
	Short: "Returns the DNS aliases that are associated with the specified Amazon FSx for Windows File Server file system.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_describeFileSystemAliasesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fsx_describeFileSystemAliasesCmd).Standalone()

		fsx_describeFileSystemAliasesCmd.Flags().String("client-request-token", "", "")
		fsx_describeFileSystemAliasesCmd.Flags().String("file-system-id", "", "The ID of the file system to return the associated DNS aliases for (String).")
		fsx_describeFileSystemAliasesCmd.Flags().String("max-results", "", "Maximum number of DNS aliases to return in the response (integer).")
		fsx_describeFileSystemAliasesCmd.Flags().String("next-token", "", "Opaque pagination token returned from a previous `DescribeFileSystemAliases` operation (String).")
		fsx_describeFileSystemAliasesCmd.MarkFlagRequired("file-system-id")
	})
	fsxCmd.AddCommand(fsx_describeFileSystemAliasesCmd)
}
