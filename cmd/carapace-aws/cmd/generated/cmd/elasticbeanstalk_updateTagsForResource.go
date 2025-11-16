package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_updateTagsForResourceCmd = &cobra.Command{
	Use:   "update-tags-for-resource",
	Short: "Update the list of tags applied to an AWS Elastic Beanstalk resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_updateTagsForResourceCmd).Standalone()

	elasticbeanstalk_updateTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resouce to be updated.")
	elasticbeanstalk_updateTagsForResourceCmd.Flags().String("tags-to-add", "", "A list of tags to add or update.")
	elasticbeanstalk_updateTagsForResourceCmd.Flags().String("tags-to-remove", "", "A list of tag keys to remove.")
	elasticbeanstalk_updateTagsForResourceCmd.MarkFlagRequired("resource-arn")
	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_updateTagsForResourceCmd)
}
