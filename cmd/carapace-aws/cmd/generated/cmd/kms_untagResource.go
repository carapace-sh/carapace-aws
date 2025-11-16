package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Deletes tags from a [customer managed key](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#customer-mgn-key). To delete a tag, specify the tag key and the KMS key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_untagResourceCmd).Standalone()

	kms_untagResourceCmd.Flags().String("key-id", "", "Identifies the KMS key from which you are removing tags.")
	kms_untagResourceCmd.Flags().String("tag-keys", "", "One or more tag keys.")
	kms_untagResourceCmd.MarkFlagRequired("key-id")
	kms_untagResourceCmd.MarkFlagRequired("tag-keys")
	kmsCmd.AddCommand(kms_untagResourceCmd)
}
