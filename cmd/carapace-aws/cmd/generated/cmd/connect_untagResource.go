package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes the specified tags from the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_untagResourceCmd).Standalone()

	connect_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	connect_untagResourceCmd.Flags().String("tag-keys", "", "The tag keys.")
	connect_untagResourceCmd.MarkFlagRequired("resource-arn")
	connect_untagResourceCmd.MarkFlagRequired("tag-keys")
	connectCmd.AddCommand(connect_untagResourceCmd)
}
