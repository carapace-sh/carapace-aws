package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var efs_describeFileSystemsCmd = &cobra.Command{
	Use:   "describe-file-systems",
	Short: "Returns the description of a specific Amazon EFS file system if either the file system `CreationToken` or the `FileSystemId` is provided.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(efs_describeFileSystemsCmd).Standalone()

	efs_describeFileSystemsCmd.Flags().String("creation-token", "", "(Optional) Restricts the list to the file system with this creation token (String).")
	efs_describeFileSystemsCmd.Flags().String("file-system-id", "", "(Optional) ID of the file system whose description you want to retrieve (String).")
	efs_describeFileSystemsCmd.Flags().String("marker", "", "(Optional) Opaque pagination token returned from a previous `DescribeFileSystems` operation (String).")
	efs_describeFileSystemsCmd.Flags().String("max-items", "", "(Optional) Specifies the maximum number of file systems to return in the response (integer).")
	efsCmd.AddCommand(efs_describeFileSystemsCmd)
}
