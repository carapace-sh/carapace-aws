package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var m2_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes one or more tags from the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(m2_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(m2_untagResourceCmd).Standalone()

		m2_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
		m2_untagResourceCmd.Flags().String("tag-keys", "", "The keys of the tags to remove.")
		m2_untagResourceCmd.MarkFlagRequired("resource-arn")
		m2_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	m2Cmd.AddCommand(m2_untagResourceCmd)
}
