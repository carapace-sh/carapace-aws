package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ram_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes the specified tag key and value pairs from the specified resource share or managed permission.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ram_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ram_untagResourceCmd).Standalone()

		ram_untagResourceCmd.Flags().String("resource-arn", "", "Specifies the [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the managed permission that you want to remove tags from.")
		ram_untagResourceCmd.Flags().String("resource-share-arn", "", "Specifies the [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the resource share that you want to remove tags from.")
		ram_untagResourceCmd.Flags().String("tag-keys", "", "Specifies a list of one or more tag keys that you want to remove.")
		ram_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	ramCmd.AddCommand(ram_untagResourceCmd)
}
