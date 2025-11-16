package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glacier_listPartsCmd = &cobra.Command{
	Use:   "list-parts",
	Short: "This operation lists the parts of an archive that have been uploaded in a specific multipart upload.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glacier_listPartsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glacier_listPartsCmd).Standalone()

		glacier_listPartsCmd.Flags().String("account-id", "", "The `AccountId` value is the AWS account ID of the account that owns the vault.")
		glacier_listPartsCmd.Flags().String("limit", "", "The maximum number of parts to be returned.")
		glacier_listPartsCmd.Flags().String("marker", "", "An opaque string used for pagination.")
		glacier_listPartsCmd.Flags().String("upload-id", "", "The upload ID of the multipart upload.")
		glacier_listPartsCmd.Flags().String("vault-name", "", "The name of the vault.")
		glacier_listPartsCmd.MarkFlagRequired("account-id")
		glacier_listPartsCmd.MarkFlagRequired("upload-id")
		glacier_listPartsCmd.MarkFlagRequired("vault-name")
	})
	glacierCmd.AddCommand(glacier_listPartsCmd)
}
