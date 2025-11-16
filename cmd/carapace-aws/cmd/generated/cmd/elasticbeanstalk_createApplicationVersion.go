package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_createApplicationVersionCmd = &cobra.Command{
	Use:   "create-application-version",
	Short: "Creates an application version for the specified application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_createApplicationVersionCmd).Standalone()

	elasticbeanstalk_createApplicationVersionCmd.Flags().String("application-name", "", "The name of the application.")
	elasticbeanstalk_createApplicationVersionCmd.Flags().String("auto-create-application", "", "Set to `true` to create an application with the specified name if it doesn't already exist.")
	elasticbeanstalk_createApplicationVersionCmd.Flags().String("build-configuration", "", "Settings for an AWS CodeBuild build.")
	elasticbeanstalk_createApplicationVersionCmd.Flags().String("description", "", "A description of this application version.")
	elasticbeanstalk_createApplicationVersionCmd.Flags().String("process", "", "Pre-processes and validates the environment manifest (`env.yaml`) and configuration files (`*.config` files in the `.ebextensions` folder) in the source bundle.")
	elasticbeanstalk_createApplicationVersionCmd.Flags().String("source-build-information", "", "Specify a commit in an AWS CodeCommit Git repository to use as the source code for the application version.")
	elasticbeanstalk_createApplicationVersionCmd.Flags().String("source-bundle", "", "The Amazon S3 bucket and key that identify the location of the source bundle for this version.")
	elasticbeanstalk_createApplicationVersionCmd.Flags().String("tags", "", "Specifies the tags applied to the application version.")
	elasticbeanstalk_createApplicationVersionCmd.Flags().String("version-label", "", "A label identifying this version.")
	elasticbeanstalk_createApplicationVersionCmd.MarkFlagRequired("application-name")
	elasticbeanstalk_createApplicationVersionCmd.MarkFlagRequired("version-label")
	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_createApplicationVersionCmd)
}
