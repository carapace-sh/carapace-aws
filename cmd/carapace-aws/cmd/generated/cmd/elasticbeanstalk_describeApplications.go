package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_describeApplicationsCmd = &cobra.Command{
	Use:   "describe-applications",
	Short: "Returns the descriptions of existing applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_describeApplicationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticbeanstalk_describeApplicationsCmd).Standalone()

		elasticbeanstalk_describeApplicationsCmd.Flags().String("application-names", "", "If specified, AWS Elastic Beanstalk restricts the returned descriptions to only include those with the specified names.")
	})
	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_describeApplicationsCmd)
}
