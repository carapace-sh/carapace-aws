package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_describeFileSystemsCmd = &cobra.Command{
	Use:   "describe-file-systems",
	Short: "Returns the description of specific Amazon FSx file systems, if a `FileSystemIds` value is provided for that file system.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_describeFileSystemsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fsx_describeFileSystemsCmd).Standalone()

		fsx_describeFileSystemsCmd.Flags().String("file-system-ids", "", "IDs of the file systems whose descriptions you want to retrieve (String).")
		fsx_describeFileSystemsCmd.Flags().String("max-results", "", "Maximum number of file systems to return in the response (integer).")
		fsx_describeFileSystemsCmd.Flags().String("next-token", "", "Opaque pagination token returned from a previous `DescribeFileSystems` operation (String).")
	})
	fsxCmd.AddCommand(fsx_describeFileSystemsCmd)
}
