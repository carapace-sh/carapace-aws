package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_selectObjectContentCmd = &cobra.Command{
	Use:   "select-object-content",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_selectObjectContentCmd).Standalone()

	s3_selectObjectContentCmd.Flags().String("bucket", "", "The S3 bucket.")
	s3_selectObjectContentCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
	s3_selectObjectContentCmd.Flags().String("expression", "", "The expression that is used to query the object.")
	s3_selectObjectContentCmd.Flags().String("expression-type", "", "The type of the provided expression (for example, SQL).")
	s3_selectObjectContentCmd.Flags().String("input-serialization", "", "Describes the format of the data in the object that is being queried.")
	s3_selectObjectContentCmd.Flags().String("key", "", "The object key.")
	s3_selectObjectContentCmd.Flags().String("output-serialization", "", "Describes the format of the data that you want Amazon S3 to return in response.")
	s3_selectObjectContentCmd.Flags().String("request-progress", "", "Specifies if periodic request progress information should be enabled.")
	s3_selectObjectContentCmd.Flags().String("scan-range", "", "Specifies the byte range of the object to get the records from.")
	s3_selectObjectContentCmd.Flags().String("ssecustomer-algorithm", "", "The server-side encryption (SSE) algorithm used to encrypt the object.")
	s3_selectObjectContentCmd.Flags().String("ssecustomer-key", "", "The server-side encryption (SSE) customer managed key.")
	s3_selectObjectContentCmd.Flags().String("ssecustomer-key-md5", "", "The MD5 server-side encryption (SSE) customer managed key.")
	s3_selectObjectContentCmd.MarkFlagRequired("bucket")
	s3_selectObjectContentCmd.MarkFlagRequired("expression")
	s3_selectObjectContentCmd.MarkFlagRequired("expression-type")
	s3_selectObjectContentCmd.MarkFlagRequired("input-serialization")
	s3_selectObjectContentCmd.MarkFlagRequired("key")
	s3_selectObjectContentCmd.MarkFlagRequired("output-serialization")
	s3Cmd.AddCommand(s3_selectObjectContentCmd)
}
