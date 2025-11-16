package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_listObjectsV2Cmd = &cobra.Command{
	Use:   "list-objects-v2",
	Short: "End of support notice: Beginning November 21, 2025, Amazon S3 will stop returning `DisplayName`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_listObjectsV2Cmd).Standalone()

	s3_listObjectsV2Cmd.Flags().String("bucket", "", "**Directory buckets** - When you use this operation with a directory bucket, you must use virtual-hosted-style requests in the format `Bucket-name.s3express-zone-id.region-code.amazonaws.com`. Path-style requests are not supported.")
	s3_listObjectsV2Cmd.Flags().String("continuation-token", "", "`ContinuationToken` indicates to Amazon S3 that the list is being continued on this bucket with a token.")
	s3_listObjectsV2Cmd.Flags().String("delimiter", "", "A delimiter is a character that you use to group keys.")
	s3_listObjectsV2Cmd.Flags().String("encoding-type", "", "Encoding type used by Amazon S3 to encode the [object keys](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-keys.html) in the response.")
	s3_listObjectsV2Cmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
	s3_listObjectsV2Cmd.Flags().String("fetch-owner", "", "The owner field is not present in `ListObjectsV2` by default.")
	s3_listObjectsV2Cmd.Flags().String("max-keys", "", "Sets the maximum number of keys returned in the response.")
	s3_listObjectsV2Cmd.Flags().String("optional-object-attributes", "", "Specifies the optional fields that you want returned in the response.")
	s3_listObjectsV2Cmd.Flags().String("prefix", "", "Limits the response to keys that begin with the specified prefix.")
	s3_listObjectsV2Cmd.Flags().String("request-payer", "", "Confirms that the requester knows that she or he will be charged for the list objects request in V2 style.")
	s3_listObjectsV2Cmd.Flags().String("start-after", "", "StartAfter is where you want Amazon S3 to start listing from.")
	s3_listObjectsV2Cmd.MarkFlagRequired("bucket")
	s3Cmd.AddCommand(s3_listObjectsV2Cmd)
}
