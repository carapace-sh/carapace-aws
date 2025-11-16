package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_deregisterEventTopicCmd = &cobra.Command{
	Use:   "deregister-event-topic",
	Short: "Removes the specified directory as a publisher to the specified Amazon SNS topic.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_deregisterEventTopicCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ds_deregisterEventTopicCmd).Standalone()

		ds_deregisterEventTopicCmd.Flags().String("directory-id", "", "The Directory ID to remove as a publisher.")
		ds_deregisterEventTopicCmd.Flags().String("topic-name", "", "The name of the Amazon SNS topic from which to remove the directory as a publisher.")
		ds_deregisterEventTopicCmd.MarkFlagRequired("directory-id")
		ds_deregisterEventTopicCmd.MarkFlagRequired("topic-name")
	})
	dsCmd.AddCommand(ds_deregisterEventTopicCmd)
}
