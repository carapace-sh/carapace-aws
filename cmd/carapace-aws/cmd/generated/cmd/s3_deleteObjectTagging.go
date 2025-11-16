package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_deleteObjectTaggingCmd = &cobra.Command{
	Use:   "delete-object-tagging",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_deleteObjectTaggingCmd).Standalone()

	s3_deleteObjectTaggingCmd.Flags().String("bucket", "", "The bucket name containing the objects from which to remove the tags.")
	s3_deleteObjectTaggingCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
	s3_deleteObjectTaggingCmd.Flags().String("key", "", "The key that identifies the object in the bucket from which to remove all tags.")
	s3_deleteObjectTaggingCmd.Flags().String("version-id", "", "The versionId of the object that the tag-set will be removed from.")
	s3_deleteObjectTaggingCmd.MarkFlagRequired("bucket")
	s3_deleteObjectTaggingCmd.MarkFlagRequired("key")
	s3Cmd.AddCommand(s3_deleteObjectTaggingCmd)
}
