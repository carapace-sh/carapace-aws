package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_describeEnvironmentsCmd = &cobra.Command{
	Use:   "describe-environments",
	Short: "Returns descriptions for existing environments.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_describeEnvironmentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticbeanstalk_describeEnvironmentsCmd).Standalone()

		elasticbeanstalk_describeEnvironmentsCmd.Flags().String("application-name", "", "If specified, AWS Elastic Beanstalk restricts the returned descriptions to include only those that are associated with this application.")
		elasticbeanstalk_describeEnvironmentsCmd.Flags().String("environment-ids", "", "If specified, AWS Elastic Beanstalk restricts the returned descriptions to include only those that have the specified IDs.")
		elasticbeanstalk_describeEnvironmentsCmd.Flags().String("environment-names", "", "If specified, AWS Elastic Beanstalk restricts the returned descriptions to include only those that have the specified names.")
		elasticbeanstalk_describeEnvironmentsCmd.Flags().String("include-deleted", "", "Indicates whether to include deleted environments:")
		elasticbeanstalk_describeEnvironmentsCmd.Flags().String("included-deleted-back-to", "", "If specified when `IncludeDeleted` is set to `true`, then environments deleted after this date are displayed.")
		elasticbeanstalk_describeEnvironmentsCmd.Flags().String("max-records", "", "For a paginated request.")
		elasticbeanstalk_describeEnvironmentsCmd.Flags().String("next-token", "", "For a paginated request.")
		elasticbeanstalk_describeEnvironmentsCmd.Flags().String("version-label", "", "If specified, AWS Elastic Beanstalk restricts the returned descriptions to include only those that are associated with this application version.")
	})
	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_describeEnvironmentsCmd)
}
