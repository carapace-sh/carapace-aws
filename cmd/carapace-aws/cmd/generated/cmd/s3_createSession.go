package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_createSessionCmd = &cobra.Command{
	Use:   "create-session",
	Short: "Creates a session that establishes temporary security credentials to support fast authentication and authorization for the Zonal endpoint API operations on directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_createSessionCmd).Standalone()

	s3_createSessionCmd.Flags().String("bucket", "", "The name of the bucket that you create a session for.")
	s3_createSessionCmd.Flags().String("bucket-key-enabled", "", "Specifies whether Amazon S3 should use an S3 Bucket Key for object encryption with server-side encryption using KMS keys (SSE-KMS).")
	s3_createSessionCmd.Flags().String("server-side-encryption", "", "The server-side encryption algorithm to use when you store objects in the directory bucket.")
	s3_createSessionCmd.Flags().String("session-mode", "", "Specifies the mode of the session that will be created, either `ReadWrite` or `ReadOnly`.")
	s3_createSessionCmd.Flags().String("ssekmsencryption-context", "", "Specifies the Amazon Web Services KMS Encryption Context as an additional encryption context to use for object encryption.")
	s3_createSessionCmd.Flags().String("ssekmskey-id", "", "If you specify `x-amz-server-side-encryption` with `aws:kms`, you must specify the `x-amz-server-side-encryption-aws-kms-key-id` header with the ID (Key ID or Key ARN) of the KMS symmetric encryption customer managed key to use.")
	s3_createSessionCmd.MarkFlagRequired("bucket")
	s3Cmd.AddCommand(s3_createSessionCmd)
}
