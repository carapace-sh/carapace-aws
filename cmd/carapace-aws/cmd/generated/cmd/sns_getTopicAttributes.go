package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sns_getTopicAttributesCmd = &cobra.Command{
	Use:   "get-topic-attributes",
	Short: "Returns all of the properties of a topic.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sns_getTopicAttributesCmd).Standalone()

	sns_getTopicAttributesCmd.Flags().String("topic-arn", "", "The ARN of the topic whose properties you want to get.")
	sns_getTopicAttributesCmd.MarkFlagRequired("topic-arn")
	snsCmd.AddCommand(sns_getTopicAttributesCmd)
}
