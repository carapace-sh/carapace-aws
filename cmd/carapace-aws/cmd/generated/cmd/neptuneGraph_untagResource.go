package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptuneGraph_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes the specified tags from the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptuneGraph_untagResourceCmd).Standalone()

	neptuneGraph_untagResourceCmd.Flags().String("resource-arn", "", "ARN of the resource whose tag needs to be removed.")
	neptuneGraph_untagResourceCmd.Flags().String("tag-keys", "", "Tag keys for the tags to be removed.")
	neptuneGraph_untagResourceCmd.MarkFlagRequired("resource-arn")
	neptuneGraph_untagResourceCmd.MarkFlagRequired("tag-keys")
	neptuneGraphCmd.AddCommand(neptuneGraph_untagResourceCmd)
}
