package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_listAliasesCmd = &cobra.Command{
	Use:   "list-aliases",
	Short: "Gets a list of aliases in the caller's Amazon Web Services account and region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_listAliasesCmd).Standalone()

	kms_listAliasesCmd.Flags().String("key-id", "", "Lists only aliases that are associated with the specified KMS key.")
	kms_listAliasesCmd.Flags().String("limit", "", "Use this parameter to specify the maximum number of items to return.")
	kms_listAliasesCmd.Flags().String("marker", "", "Use this parameter in a subsequent request after you receive a response with truncated results.")
	kmsCmd.AddCommand(kms_listAliasesCmd)
}
