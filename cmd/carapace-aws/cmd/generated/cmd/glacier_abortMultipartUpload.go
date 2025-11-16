package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glacier_abortMultipartUploadCmd = &cobra.Command{
	Use:   "abort-multipart-upload",
	Short: "This operation aborts a multipart upload identified by the upload ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glacier_abortMultipartUploadCmd).Standalone()

	glacier_abortMultipartUploadCmd.Flags().String("account-id", "", "The `AccountId` value is the AWS account ID of the account that owns the vault.")
	glacier_abortMultipartUploadCmd.Flags().String("upload-id", "", "The upload ID of the multipart upload to delete.")
	glacier_abortMultipartUploadCmd.Flags().String("vault-name", "", "The name of the vault.")
	glacier_abortMultipartUploadCmd.MarkFlagRequired("account-id")
	glacier_abortMultipartUploadCmd.MarkFlagRequired("upload-id")
	glacier_abortMultipartUploadCmd.MarkFlagRequired("vault-name")
	glacierCmd.AddCommand(glacier_abortMultipartUploadCmd)
}
