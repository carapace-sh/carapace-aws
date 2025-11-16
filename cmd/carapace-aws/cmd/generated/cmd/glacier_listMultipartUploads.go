package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glacier_listMultipartUploadsCmd = &cobra.Command{
	Use:   "list-multipart-uploads",
	Short: "This operation lists in-progress multipart uploads for the specified vault.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glacier_listMultipartUploadsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glacier_listMultipartUploadsCmd).Standalone()

		glacier_listMultipartUploadsCmd.Flags().String("account-id", "", "The `AccountId` value is the AWS account ID of the account that owns the vault.")
		glacier_listMultipartUploadsCmd.Flags().String("limit", "", "Specifies the maximum number of uploads returned in the response body.")
		glacier_listMultipartUploadsCmd.Flags().String("marker", "", "An opaque string used for pagination.")
		glacier_listMultipartUploadsCmd.Flags().String("vault-name", "", "The name of the vault.")
		glacier_listMultipartUploadsCmd.MarkFlagRequired("account-id")
		glacier_listMultipartUploadsCmd.MarkFlagRequired("vault-name")
	})
	glacierCmd.AddCommand(glacier_listMultipartUploadsCmd)
}
