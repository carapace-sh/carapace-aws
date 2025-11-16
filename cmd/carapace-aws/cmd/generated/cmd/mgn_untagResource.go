package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Deletes the specified set of tags from the specified set of Application Migration Service resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_untagResourceCmd).Standalone()

	mgn_untagResourceCmd.Flags().String("resource-arn", "", "Untag resource by ARN.")
	mgn_untagResourceCmd.Flags().String("tag-keys", "", "Untag resource by Keys.")
	mgn_untagResourceCmd.MarkFlagRequired("resource-arn")
	mgn_untagResourceCmd.MarkFlagRequired("tag-keys")
	mgnCmd.AddCommand(mgn_untagResourceCmd)
}
