package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_describeInstancesHealthCmd = &cobra.Command{
	Use:   "describe-instances-health",
	Short: "Retrieves detailed information about the health of instances in your AWS Elastic Beanstalk.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_describeInstancesHealthCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticbeanstalk_describeInstancesHealthCmd).Standalone()

		elasticbeanstalk_describeInstancesHealthCmd.Flags().String("attribute-names", "", "Specifies the response elements you wish to receive.")
		elasticbeanstalk_describeInstancesHealthCmd.Flags().String("environment-id", "", "Specify the AWS Elastic Beanstalk environment by ID.")
		elasticbeanstalk_describeInstancesHealthCmd.Flags().String("environment-name", "", "Specify the AWS Elastic Beanstalk environment by name.")
		elasticbeanstalk_describeInstancesHealthCmd.Flags().String("next-token", "", "Specify the pagination token returned by a previous call.")
	})
	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_describeInstancesHealthCmd)
}
