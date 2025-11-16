package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_putObjectTaggingCmd = &cobra.Command{
	Use:   "put-object-tagging",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_putObjectTaggingCmd).Standalone()

	s3_putObjectTaggingCmd.Flags().String("bucket", "", "The bucket name containing the object.")
	s3_putObjectTaggingCmd.Flags().String("checksum-algorithm", "", "Indicates the algorithm used to create the checksum for the object when you use the SDK.")
	s3_putObjectTaggingCmd.Flags().String("content-md5", "", "The MD5 hash for the request body.")
	s3_putObjectTaggingCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
	s3_putObjectTaggingCmd.Flags().String("key", "", "Name of the object key.")
	s3_putObjectTaggingCmd.Flags().String("request-payer", "", "Confirms that the requester knows that she or he will be charged for the tagging object request.")
	s3_putObjectTaggingCmd.Flags().String("tagging", "", "Container for the `TagSet` and `Tag` elements")
	s3_putObjectTaggingCmd.Flags().String("version-id", "", "The versionId of the object that the tag-set will be added to.")
	s3_putObjectTaggingCmd.MarkFlagRequired("bucket")
	s3_putObjectTaggingCmd.MarkFlagRequired("key")
	s3_putObjectTaggingCmd.MarkFlagRequired("tagging")
	s3Cmd.AddCommand(s3_putObjectTaggingCmd)
}
