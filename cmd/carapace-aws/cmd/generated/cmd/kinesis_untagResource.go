package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesis_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes tags from the specified Kinesis resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesis_untagResourceCmd).Standalone()

	kinesis_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the Kinesis resource from which to remove tags.")
	kinesis_untagResourceCmd.Flags().String("tag-keys", "", "A list of tag key-value pairs.")
	kinesis_untagResourceCmd.MarkFlagRequired("resource-arn")
	kinesis_untagResourceCmd.MarkFlagRequired("tag-keys")
	kinesisCmd.AddCommand(kinesis_untagResourceCmd)
}
