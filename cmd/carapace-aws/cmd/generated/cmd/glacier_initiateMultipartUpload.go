package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glacier_initiateMultipartUploadCmd = &cobra.Command{
	Use:   "initiate-multipart-upload",
	Short: "This operation initiates a multipart upload.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glacier_initiateMultipartUploadCmd).Standalone()

	glacier_initiateMultipartUploadCmd.Flags().String("account-id", "", "The `AccountId` value is the AWS account ID of the account that owns the vault.")
	glacier_initiateMultipartUploadCmd.Flags().String("archive-description", "", "The archive description that you are uploading in parts.")
	glacier_initiateMultipartUploadCmd.Flags().String("part-size", "", "The size of each part except the last, in bytes.")
	glacier_initiateMultipartUploadCmd.Flags().String("vault-name", "", "The name of the vault.")
	glacier_initiateMultipartUploadCmd.MarkFlagRequired("account-id")
	glacier_initiateMultipartUploadCmd.MarkFlagRequired("vault-name")
	glacierCmd.AddCommand(glacier_initiateMultipartUploadCmd)
}
