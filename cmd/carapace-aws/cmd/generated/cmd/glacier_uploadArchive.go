package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glacier_uploadArchiveCmd = &cobra.Command{
	Use:   "upload-archive",
	Short: "This operation adds an archive to a vault.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glacier_uploadArchiveCmd).Standalone()

	glacier_uploadArchiveCmd.Flags().String("account-id", "", "The `AccountId` value is the AWS account ID of the account that owns the vault.")
	glacier_uploadArchiveCmd.Flags().String("archive-description", "", "The optional description of the archive you are uploading.")
	glacier_uploadArchiveCmd.Flags().String("body", "", "The data to upload.")
	glacier_uploadArchiveCmd.Flags().String("checksum", "", "The SHA256 tree hash of the data being uploaded.")
	glacier_uploadArchiveCmd.Flags().String("vault-name", "", "The name of the vault.")
	glacier_uploadArchiveCmd.MarkFlagRequired("account-id")
	glacier_uploadArchiveCmd.MarkFlagRequired("vault-name")
	glacierCmd.AddCommand(glacier_uploadArchiveCmd)
}
