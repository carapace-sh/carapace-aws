package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3_renameObjectCmd = &cobra.Command{
	Use:   "rename-object",
	Short: "Renames an existing object in a directory bucket that uses the S3 Express One Zone storage class.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3_renameObjectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3_renameObjectCmd).Standalone()

		s3_renameObjectCmd.Flags().String("bucket", "", "The bucket name of the directory bucket containing the object.")
		s3_renameObjectCmd.Flags().String("client-token", "", "A unique string with a max of 64 ASCII characters in the ASCII range of 33 - 126.")
		s3_renameObjectCmd.Flags().String("destination-if-match", "", "Renames the object only if the ETag (entity tag) value provided during the operation matches the ETag of the object in S3.")
		s3_renameObjectCmd.Flags().String("destination-if-modified-since", "", "Renames the object if the destination exists and if it has been modified since the specified time.")
		s3_renameObjectCmd.Flags().String("destination-if-none-match", "", "Renames the object only if the destination does not already exist in the specified directory bucket.")
		s3_renameObjectCmd.Flags().String("destination-if-unmodified-since", "", "Renames the object if it hasn't been modified since the specified time.")
		s3_renameObjectCmd.Flags().String("key", "", "Key name of the object to rename.")
		s3_renameObjectCmd.Flags().String("rename-source", "", "Specifies the source for the rename operation.")
		s3_renameObjectCmd.Flags().String("source-if-match", "", "Renames the object if the source exists and if its entity tag (ETag) matches the specified ETag.")
		s3_renameObjectCmd.Flags().String("source-if-modified-since", "", "Renames the object if the source exists and if it has been modified since the specified time.")
		s3_renameObjectCmd.Flags().String("source-if-none-match", "", "Renames the object if the source exists and if its entity tag (ETag) is different than the specified ETag.")
		s3_renameObjectCmd.Flags().String("source-if-unmodified-since", "", "Renames the object if the source exists and hasn't been modified since the specified time.")
		s3_renameObjectCmd.MarkFlagRequired("bucket")
		s3_renameObjectCmd.MarkFlagRequired("key")
		s3_renameObjectCmd.MarkFlagRequired("rename-source")
	})
	s3Cmd.AddCommand(s3_renameObjectCmd)
}
