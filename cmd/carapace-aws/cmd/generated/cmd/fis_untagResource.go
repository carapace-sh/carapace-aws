package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fis_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes the specified tags from the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fis_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fis_untagResourceCmd).Standalone()

		fis_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
		fis_untagResourceCmd.Flags().String("tag-keys", "", "The tag keys to remove.")
		fis_untagResourceCmd.MarkFlagRequired("resource-arn")
	})
	fisCmd.AddCommand(fis_untagResourceCmd)
}
