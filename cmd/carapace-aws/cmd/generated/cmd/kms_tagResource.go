package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds or edits tags on a [customer managed key](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#customer-mgn-key).\n\nTagging or untagging a KMS key can allow or deny permission to the KMS key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_tagResourceCmd).Standalone()

	kms_tagResourceCmd.Flags().String("key-id", "", "Identifies a customer managed key in the account and Region.")
	kms_tagResourceCmd.Flags().String("tags", "", "One or more tags.")
	kms_tagResourceCmd.MarkFlagRequired("key-id")
	kms_tagResourceCmd.MarkFlagRequired("tags")
	kmsCmd.AddCommand(kms_tagResourceCmd)
}
