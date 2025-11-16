package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "List all tags for an accelerator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_listTagsForResourceCmd).Standalone()

	globalaccelerator_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the accelerator to list tags for.")
	globalaccelerator_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	globalacceleratorCmd.AddCommand(globalaccelerator_listTagsForResourceCmd)
}
