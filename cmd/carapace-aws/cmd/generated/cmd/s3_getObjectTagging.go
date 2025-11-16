package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_getObjectTaggingCmd = &cobra.Command{
	Use:   "get-object-tagging",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_getObjectTaggingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_getObjectTaggingCmd).Standalone()

		s3_getObjectTaggingCmd.Flags().String("bucket", "", "The bucket name containing the object for which to get the tagging information.")
		s3_getObjectTaggingCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_getObjectTaggingCmd.Flags().String("key", "", "Object key for which to get the tagging information.")
		s3_getObjectTaggingCmd.Flags().String("request-payer", "", "")
		s3_getObjectTaggingCmd.Flags().String("version-id", "", "The versionId of the object for which to get the tagging information.")
		s3_getObjectTaggingCmd.MarkFlagRequired("bucket")
		s3_getObjectTaggingCmd.MarkFlagRequired("key")
	})
	s3Cmd.AddCommand(s3_getObjectTaggingCmd)
}
