package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Return the tags applied to an AWS Elastic Beanstalk resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_listTagsForResourceCmd).Standalone()

	elasticbeanstalk_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resouce for which a tag list is requested.")
	elasticbeanstalk_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_listTagsForResourceCmd)
}
