package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var b2bi_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Attaches a key-value pair to a resource, as identified by its Amazon Resource Name (ARN).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(b2bi_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(b2bi_tagResourceCmd).Standalone()

		b2bi_tagResourceCmd.Flags().String("resource-arn", "", "Specifies an Amazon Resource Name (ARN) for a specific Amazon Web Services resource, such as a capability, partnership, profile, or transformer.")
		b2bi_tagResourceCmd.Flags().String("tags", "", "Specifies the key-value pairs assigned to ARNs that you can use to group and search for resources by type.")
		b2bi_tagResourceCmd.MarkFlagRequired("resource-arn")
		b2bi_tagResourceCmd.MarkFlagRequired("tags")
	})
	b2biCmd.AddCommand(b2bi_tagResourceCmd)
}
