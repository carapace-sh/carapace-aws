package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glacier_deleteArchiveCmd = &cobra.Command{
	Use:   "delete-archive",
	Short: "This operation deletes an archive from a vault.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glacier_deleteArchiveCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glacier_deleteArchiveCmd).Standalone()

		glacier_deleteArchiveCmd.Flags().String("account-id", "", "The `AccountId` value is the AWS account ID of the account that owns the vault.")
		glacier_deleteArchiveCmd.Flags().String("archive-id", "", "The ID of the archive to delete.")
		glacier_deleteArchiveCmd.Flags().String("vault-name", "", "The name of the vault.")
		glacier_deleteArchiveCmd.MarkFlagRequired("account-id")
		glacier_deleteArchiveCmd.MarkFlagRequired("archive-id")
		glacier_deleteArchiveCmd.MarkFlagRequired("vault-name")
	})
	glacierCmd.AddCommand(glacier_deleteArchiveCmd)
}
