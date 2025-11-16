package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glacier_uploadMultipartPartCmd = &cobra.Command{
	Use:   "upload-multipart-part",
	Short: "This operation uploads a part of an archive.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glacier_uploadMultipartPartCmd).Standalone()

	glacier_uploadMultipartPartCmd.Flags().String("account-id", "", "The `AccountId` value is the AWS account ID of the account that owns the vault.")
	glacier_uploadMultipartPartCmd.Flags().String("body", "", "The data to upload.")
	glacier_uploadMultipartPartCmd.Flags().String("checksum", "", "The SHA256 tree hash of the data being uploaded.")
	glacier_uploadMultipartPartCmd.Flags().String("range", "", "Identifies the range of bytes in the assembled archive that will be uploaded in this part.")
	glacier_uploadMultipartPartCmd.Flags().String("upload-id", "", "The upload ID of the multipart upload.")
	glacier_uploadMultipartPartCmd.Flags().String("vault-name", "", "The name of the vault.")
	glacier_uploadMultipartPartCmd.MarkFlagRequired("account-id")
	glacier_uploadMultipartPartCmd.MarkFlagRequired("upload-id")
	glacier_uploadMultipartPartCmd.MarkFlagRequired("vault-name")
	glacierCmd.AddCommand(glacier_uploadMultipartPartCmd)
}
