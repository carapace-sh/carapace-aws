package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_describeApplicationVersionsCmd = &cobra.Command{
	Use:   "describe-application-versions",
	Short: "Retrieve a list of application versions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_describeApplicationVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticbeanstalk_describeApplicationVersionsCmd).Standalone()

		elasticbeanstalk_describeApplicationVersionsCmd.Flags().String("application-name", "", "Specify an application name to show only application versions for that application.")
		elasticbeanstalk_describeApplicationVersionsCmd.Flags().String("max-records", "", "For a paginated request.")
		elasticbeanstalk_describeApplicationVersionsCmd.Flags().String("next-token", "", "For a paginated request.")
		elasticbeanstalk_describeApplicationVersionsCmd.Flags().String("version-labels", "", "Specify a version label to show a specific application version.")
	})
	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_describeApplicationVersionsCmd)
}
