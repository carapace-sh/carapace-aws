package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sns_setTopicAttributesCmd = &cobra.Command{
	Use:   "set-topic-attributes",
	Short: "Allows a topic owner to set an attribute of the topic to a new value.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sns_setTopicAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sns_setTopicAttributesCmd).Standalone()

		sns_setTopicAttributesCmd.Flags().String("attribute-name", "", "A map of attributes with their corresponding values.")
		sns_setTopicAttributesCmd.Flags().String("attribute-value", "", "The new value for the attribute.")
		sns_setTopicAttributesCmd.Flags().String("topic-arn", "", "The ARN of the topic to modify.")
		sns_setTopicAttributesCmd.MarkFlagRequired("attribute-name")
		sns_setTopicAttributesCmd.MarkFlagRequired("topic-arn")
	})
	snsCmd.AddCommand(sns_setTopicAttributesCmd)
}
