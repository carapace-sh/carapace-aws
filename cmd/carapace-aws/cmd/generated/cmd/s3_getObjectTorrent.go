package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_getObjectTorrentCmd = &cobra.Command{
	Use:   "get-object-torrent",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_getObjectTorrentCmd).Standalone()

	s3_getObjectTorrentCmd.Flags().String("bucket", "", "The name of the bucket containing the object for which to get the torrent files.")
	s3_getObjectTorrentCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
	s3_getObjectTorrentCmd.Flags().String("key", "", "The object key for which to get the information.")
	s3_getObjectTorrentCmd.Flags().String("request-payer", "", "")
	s3_getObjectTorrentCmd.MarkFlagRequired("bucket")
	s3_getObjectTorrentCmd.MarkFlagRequired("key")
	s3Cmd.AddCommand(s3_getObjectTorrentCmd)
}
