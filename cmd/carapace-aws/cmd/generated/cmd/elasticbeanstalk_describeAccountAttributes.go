package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_describeAccountAttributesCmd = &cobra.Command{
	Use:   "describe-account-attributes",
	Short: "Returns attributes related to AWS Elastic Beanstalk that are associated with the calling AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_describeAccountAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticbeanstalk_describeAccountAttributesCmd).Standalone()

	})
	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_describeAccountAttributesCmd)
}
