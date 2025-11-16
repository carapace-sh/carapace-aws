package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var supplychain_createInstanceCmd = &cobra.Command{
	Use:   "create-instance",
	Short: "Enables you to programmatically create an Amazon Web Services Supply Chain instance by applying KMS keys and relevant information associated with the API without using the Amazon Web Services console.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(supplychain_createInstanceCmd).Standalone()

	supplychain_createInstanceCmd.Flags().String("client-token", "", "The client token for idempotency.")
	supplychain_createInstanceCmd.Flags().String("instance-description", "", "The AWS Supply Chain instance description.")
	supplychain_createInstanceCmd.Flags().String("instance-name", "", "The AWS Supply Chain instance name.")
	supplychain_createInstanceCmd.Flags().String("kms-key-arn", "", "The ARN (Amazon Resource Name) of the Key Management Service (KMS) key you provide for encryption.")
	supplychain_createInstanceCmd.Flags().String("tags", "", "The Amazon Web Services tags of an instance to be created.")
	supplychain_createInstanceCmd.Flags().String("web-app-dns-domain", "", "The DNS subdomain of the web app.")
	supplychainCmd.AddCommand(supplychain_createInstanceCmd)
}
