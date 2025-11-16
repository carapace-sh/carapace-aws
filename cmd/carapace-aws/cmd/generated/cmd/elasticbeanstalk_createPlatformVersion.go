package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_createPlatformVersionCmd = &cobra.Command{
	Use:   "create-platform-version",
	Short: "Create a new version of your custom platform.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_createPlatformVersionCmd).Standalone()

	elasticbeanstalk_createPlatformVersionCmd.Flags().String("environment-name", "", "The name of the builder environment.")
	elasticbeanstalk_createPlatformVersionCmd.Flags().String("option-settings", "", "The configuration option settings to apply to the builder environment.")
	elasticbeanstalk_createPlatformVersionCmd.Flags().String("platform-definition-bundle", "", "The location of the platform definition archive in Amazon S3.")
	elasticbeanstalk_createPlatformVersionCmd.Flags().String("platform-name", "", "The name of your custom platform.")
	elasticbeanstalk_createPlatformVersionCmd.Flags().String("platform-version", "", "The number, such as 1.0.2, for the new platform version.")
	elasticbeanstalk_createPlatformVersionCmd.Flags().String("tags", "", "Specifies the tags applied to the new platform version.")
	elasticbeanstalk_createPlatformVersionCmd.MarkFlagRequired("platform-definition-bundle")
	elasticbeanstalk_createPlatformVersionCmd.MarkFlagRequired("platform-name")
	elasticbeanstalk_createPlatformVersionCmd.MarkFlagRequired("platform-version")
	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_createPlatformVersionCmd)
}
