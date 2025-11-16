package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Remove a customer tag from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_untagResourceCmd).Standalone()

		proton_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to remove customer tags from.")
		proton_untagResourceCmd.Flags().String("tag-keys", "", "A list of customer tag keys that indicate the customer tags to be removed from the resource.")
		proton_untagResourceCmd.MarkFlagRequired("resource-arn")
		proton_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	protonCmd.AddCommand(proton_untagResourceCmd)
}
