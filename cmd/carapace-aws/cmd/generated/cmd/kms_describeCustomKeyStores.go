package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_describeCustomKeyStoresCmd = &cobra.Command{
	Use:   "describe-custom-key-stores",
	Short: "Gets information about [custom key stores](https://docs.aws.amazon.com/kms/latest/developerguide/key-store-overview.html) in the account and Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_describeCustomKeyStoresCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kms_describeCustomKeyStoresCmd).Standalone()

		kms_describeCustomKeyStoresCmd.Flags().String("custom-key-store-id", "", "Gets only information about the specified custom key store.")
		kms_describeCustomKeyStoresCmd.Flags().String("custom-key-store-name", "", "Gets only information about the specified custom key store.")
		kms_describeCustomKeyStoresCmd.Flags().String("limit", "", "Use this parameter to specify the maximum number of items to return.")
		kms_describeCustomKeyStoresCmd.Flags().String("marker", "", "Use this parameter in a subsequent request after you receive a response with truncated results.")
	})
	kmsCmd.AddCommand(kms_describeCustomKeyStoresCmd)
}
