package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_shareDirectoryCmd = &cobra.Command{
	Use:   "share-directory",
	Short: "Shares a specified directory (`DirectoryId`) in your Amazon Web Services account (directory owner) with another Amazon Web Services account (directory consumer).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_shareDirectoryCmd).Standalone()

	ds_shareDirectoryCmd.Flags().String("directory-id", "", "Identifier of the Managed Microsoft AD directory that you want to share with other Amazon Web Services accounts.")
	ds_shareDirectoryCmd.Flags().String("share-method", "", "The method used when sharing a directory to determine whether the directory should be shared within your Amazon Web Services organization (`ORGANIZATIONS`) or with any Amazon Web Services account by sending a directory sharing request (`HANDSHAKE`).")
	ds_shareDirectoryCmd.Flags().String("share-notes", "", "A directory share request that is sent by the directory owner to the directory consumer.")
	ds_shareDirectoryCmd.Flags().String("share-target", "", "Identifier for the directory consumer account with whom the directory is to be shared.")
	ds_shareDirectoryCmd.MarkFlagRequired("directory-id")
	ds_shareDirectoryCmd.MarkFlagRequired("share-method")
	ds_shareDirectoryCmd.MarkFlagRequired("share-target")
	dsCmd.AddCommand(ds_shareDirectoryCmd)
}
