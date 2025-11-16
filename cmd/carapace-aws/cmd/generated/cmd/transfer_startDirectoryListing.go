package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_startDirectoryListingCmd = &cobra.Command{
	Use:   "start-directory-listing",
	Short: "Retrieves a list of the contents of a directory from a remote SFTP server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_startDirectoryListingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transfer_startDirectoryListingCmd).Standalone()

		transfer_startDirectoryListingCmd.Flags().String("connector-id", "", "The unique identifier for the connector.")
		transfer_startDirectoryListingCmd.Flags().String("max-items", "", "An optional parameter where you can specify the maximum number of file/directory names to retrieve.")
		transfer_startDirectoryListingCmd.Flags().String("output-directory-path", "", "Specifies the path (bucket and prefix) in Amazon S3 storage to store the results of the directory listing.")
		transfer_startDirectoryListingCmd.Flags().String("remote-directory-path", "", "Specifies the directory on the remote SFTP server for which you want to list its contents.")
		transfer_startDirectoryListingCmd.MarkFlagRequired("connector-id")
		transfer_startDirectoryListingCmd.MarkFlagRequired("output-directory-path")
		transfer_startDirectoryListingCmd.MarkFlagRequired("remote-directory-path")
	})
	transferCmd.AddCommand(transfer_startDirectoryListingCmd)
}
