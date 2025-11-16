package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_deleteObjectCmd = &cobra.Command{
	Use:   "delete-object",
	Short: "Removes an object from a bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_deleteObjectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_deleteObjectCmd).Standalone()

		s3_deleteObjectCmd.Flags().String("bucket", "", "The bucket name of the bucket containing the object.")
		s3_deleteObjectCmd.Flags().String("bypass-governance-retention", "", "Indicates whether S3 Object Lock should bypass Governance-mode restrictions to process this operation.")
		s3_deleteObjectCmd.Flags().String("expected-bucket-owner", "", "The account ID of the expected bucket owner.")
		s3_deleteObjectCmd.Flags().String("if-match", "", "Deletes the object if the ETag (entity tag) value provided during the delete operation matches the ETag of the object in S3.")
		s3_deleteObjectCmd.Flags().String("if-match-last-modified-time", "", "If present, the object is deleted only if its modification times matches the provided `Timestamp`.")
		s3_deleteObjectCmd.Flags().String("if-match-size", "", "If present, the object is deleted only if its size matches the provided size in bytes.")
		s3_deleteObjectCmd.Flags().String("key", "", "Key name of the object to delete.")
		s3_deleteObjectCmd.Flags().String("mfa", "", "The concatenation of the authentication device's serial number, a space, and the value that is displayed on your authentication device.")
		s3_deleteObjectCmd.Flags().String("request-payer", "", "")
		s3_deleteObjectCmd.Flags().String("version-id", "", "Version ID used to reference a specific version of the object.")
		s3_deleteObjectCmd.MarkFlagRequired("bucket")
		s3_deleteObjectCmd.MarkFlagRequired("key")
	})
	s3Cmd.AddCommand(s3_deleteObjectCmd)
}
