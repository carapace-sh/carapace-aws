package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var b2bi_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Detaches a key-value pair from the specified resource, as identified by its Amazon Resource Name (ARN).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(b2bi_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(b2bi_untagResourceCmd).Standalone()

		b2bi_untagResourceCmd.Flags().String("resource-arn", "", "Specifies an Amazon Resource Name (ARN) for a specific Amazon Web Services resource, such as a capability, partnership, profile, or transformer.")
		b2bi_untagResourceCmd.Flags().String("tag-keys", "", "Specifies the key-value pairs assigned to ARNs that you can use to group and search for resources by type.")
		b2bi_untagResourceCmd.MarkFlagRequired("resource-arn")
		b2bi_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	b2biCmd.AddCommand(b2bi_untagResourceCmd)
}
