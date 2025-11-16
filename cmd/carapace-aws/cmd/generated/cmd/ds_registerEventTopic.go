package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_registerEventTopicCmd = &cobra.Command{
	Use:   "register-event-topic",
	Short: "Associates a directory with an Amazon SNS topic.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_registerEventTopicCmd).Standalone()

	ds_registerEventTopicCmd.Flags().String("directory-id", "", "The Directory ID that will publish status messages to the Amazon SNS topic.")
	ds_registerEventTopicCmd.Flags().String("topic-name", "", "The Amazon SNS topic name to which the directory will publish status messages.")
	ds_registerEventTopicCmd.MarkFlagRequired("directory-id")
	ds_registerEventTopicCmd.MarkFlagRequired("topic-name")
	dsCmd.AddCommand(ds_registerEventTopicCmd)
}
