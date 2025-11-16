package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sns_createTopicCmd = &cobra.Command{
	Use:   "create-topic",
	Short: "Creates a topic to which notifications can be published.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sns_createTopicCmd).Standalone()

	sns_createTopicCmd.Flags().String("attributes", "", "A map of attributes with their corresponding values.")
	sns_createTopicCmd.Flags().String("data-protection-policy", "", "The body of the policy document you want to use for this topic.")
	sns_createTopicCmd.Flags().String("name", "", "The name of the topic you want to create.")
	sns_createTopicCmd.Flags().String("tags", "", "The list of tags to add to a new topic.")
	sns_createTopicCmd.MarkFlagRequired("name")
	snsCmd.AddCommand(sns_createTopicCmd)
}
