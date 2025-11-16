package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glacier_completeMultipartUploadCmd = &cobra.Command{
	Use:   "complete-multipart-upload",
	Short: "You call this operation to inform Amazon S3 Glacier (Glacier) that all the archive parts have been uploaded and that Glacier can now assemble the archive from the uploaded parts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glacier_completeMultipartUploadCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glacier_completeMultipartUploadCmd).Standalone()

		glacier_completeMultipartUploadCmd.Flags().String("account-id", "", "The `AccountId` value is the AWS account ID of the account that owns the vault.")
		glacier_completeMultipartUploadCmd.Flags().String("archive-size", "", "The total size, in bytes, of the entire archive.")
		glacier_completeMultipartUploadCmd.Flags().String("checksum", "", "The SHA256 tree hash of the entire archive.")
		glacier_completeMultipartUploadCmd.Flags().String("upload-id", "", "The upload ID of the multipart upload.")
		glacier_completeMultipartUploadCmd.Flags().String("vault-name", "", "The name of the vault.")
		glacier_completeMultipartUploadCmd.MarkFlagRequired("account-id")
		glacier_completeMultipartUploadCmd.MarkFlagRequired("upload-id")
		glacier_completeMultipartUploadCmd.MarkFlagRequired("vault-name")
	})
	glacierCmd.AddCommand(glacier_completeMultipartUploadCmd)
}
