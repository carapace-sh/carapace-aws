package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcs_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Deletes tags from an PCS resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcs_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pcs_untagResourceCmd).Standalone()

		pcs_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
		pcs_untagResourceCmd.Flags().String("tag-keys", "", "1 or more tag keys to remove from the resource.")
		pcs_untagResourceCmd.MarkFlagRequired("resource-arn")
		pcs_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	pcsCmd.AddCommand(pcs_untagResourceCmd)
}
