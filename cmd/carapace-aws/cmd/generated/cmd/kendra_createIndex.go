package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_createIndexCmd = &cobra.Command{
	Use:   "create-index",
	Short: "Creates an Amazon Kendra index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_createIndexCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kendra_createIndexCmd).Standalone()

		kendra_createIndexCmd.Flags().String("client-token", "", "A token that you provide to identify the request to create an index.")
		kendra_createIndexCmd.Flags().String("description", "", "A description for the index.")
		kendra_createIndexCmd.Flags().String("edition", "", "The Amazon Kendra edition to use for the index.")
		kendra_createIndexCmd.Flags().String("name", "", "A name for the index.")
		kendra_createIndexCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of an IAM role with permission to access your Amazon CloudWatch logs and metrics.")
		kendra_createIndexCmd.Flags().String("server-side-encryption-configuration", "", "The identifier of the KMS customer managed key (CMK) that's used to encrypt data indexed by Amazon Kendra.")
		kendra_createIndexCmd.Flags().String("tags", "", "A list of key-value pairs that identify or categorize the index.")
		kendra_createIndexCmd.Flags().String("user-context-policy", "", "The user context policy.")
		kendra_createIndexCmd.Flags().String("user-group-resolution-configuration", "", "Gets users and groups from IAM Identity Center identity source.")
		kendra_createIndexCmd.Flags().String("user-token-configurations", "", "The user token configuration.")
		kendra_createIndexCmd.MarkFlagRequired("name")
		kendra_createIndexCmd.MarkFlagRequired("role-arn")
	})
	kendraCmd.AddCommand(kendra_createIndexCmd)
}
