package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_putObjectRetentionCmd = &cobra.Command{
	Use:   "put-object-retention",
	Short: "This operation is not supported for directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_putObjectRetentionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_putObjectRetentionCmd).Standalone()

		s3_putObjectRetentionCmd.Flags().String("bucket", "", "The bucket name that contains the object you want to apply this Object Retention configuration to.")
		s3_putObjectRetentionCmd.Flags().String("bypass-governance-retention", "", "Indicates whether this action should bypass Governance-mode restrictions.")
		s3_putObjectRetentionCmd.Flags().String("checksum-algorithm", "", "Indicates the algorithm used to create the checksum for the object when you use the SDK.")
		s3_putObjectRetentionCmd.Flags().String("content-md5", "", "The MD5 hash for the request body.")
		s3_putObjectRetentionCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_putObjectRetentionCmd.Flags().String("key", "", "The key name for the object that you want to apply this Object Retention configuration to.")
		s3_putObjectRetentionCmd.Flags().String("request-payer", "", "")
		s3_putObjectRetentionCmd.Flags().String("retention", "", "The container element for the Object Retention configuration.")
		s3_putObjectRetentionCmd.Flags().String("version-id", "", "The version ID for the object that you want to apply this Object Retention configuration to.")
		s3_putObjectRetentionCmd.MarkFlagRequired("bucket")
		s3_putObjectRetentionCmd.MarkFlagRequired("key")
	})
	s3Cmd.AddCommand(s3_putObjectRetentionCmd)
}
