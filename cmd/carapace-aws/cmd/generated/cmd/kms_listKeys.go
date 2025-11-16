package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_listKeysCmd = &cobra.Command{
	Use:   "list-keys",
	Short: "Gets a list of all KMS keys in the caller's Amazon Web Services account and Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_listKeysCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kms_listKeysCmd).Standalone()

		kms_listKeysCmd.Flags().String("limit", "", "Use this parameter to specify the maximum number of items to return.")
		kms_listKeysCmd.Flags().String("marker", "", "Use this parameter in a subsequent request after you receive a response with truncated results.")
	})
	kmsCmd.AddCommand(kms_listKeysCmd)
}
