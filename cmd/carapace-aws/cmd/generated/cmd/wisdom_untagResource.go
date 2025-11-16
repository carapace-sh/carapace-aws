package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wisdom_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes the specified tags from the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wisdom_untagResourceCmd).Standalone()

	wisdom_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	wisdom_untagResourceCmd.Flags().String("tag-keys", "", "The tag keys.")
	wisdom_untagResourceCmd.MarkFlagRequired("resource-arn")
	wisdom_untagResourceCmd.MarkFlagRequired("tag-keys")
	wisdomCmd.AddCommand(wisdom_untagResourceCmd)
}
