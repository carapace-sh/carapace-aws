package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var savingsplans_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes the specified tags from the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(savingsplans_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(savingsplans_untagResourceCmd).Standalone()

		savingsplans_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
		savingsplans_untagResourceCmd.Flags().String("tag-keys", "", "The tag keys.")
		savingsplans_untagResourceCmd.MarkFlagRequired("resource-arn")
		savingsplans_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	savingsplansCmd.AddCommand(savingsplans_untagResourceCmd)
}
