package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesis_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds or updates tags for the specified Kinesis resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesis_tagResourceCmd).Standalone()

	kinesis_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the Kinesis resource to which to add tags.")
	kinesis_tagResourceCmd.Flags().String("tags", "", "An array of tags to be added to the Kinesis resource.")
	kinesis_tagResourceCmd.MarkFlagRequired("resource-arn")
	kinesis_tagResourceCmd.MarkFlagRequired("tags")
	kinesisCmd.AddCommand(kinesis_tagResourceCmd)
}
