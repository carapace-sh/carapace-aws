package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Remove tags from a Global Accelerator resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(globalaccelerator_untagResourceCmd).Standalone()

		globalaccelerator_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the Global Accelerator resource to remove tags from.")
		globalaccelerator_untagResourceCmd.Flags().String("tag-keys", "", "The tag key pairs that you want to remove from the specified resources.")
		globalaccelerator_untagResourceCmd.MarkFlagRequired("resource-arn")
		globalaccelerator_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	globalacceleratorCmd.AddCommand(globalaccelerator_untagResourceCmd)
}
