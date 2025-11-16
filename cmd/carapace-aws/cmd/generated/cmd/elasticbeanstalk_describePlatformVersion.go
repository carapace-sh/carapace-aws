package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_describePlatformVersionCmd = &cobra.Command{
	Use:   "describe-platform-version",
	Short: "Describes a platform version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_describePlatformVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticbeanstalk_describePlatformVersionCmd).Standalone()

		elasticbeanstalk_describePlatformVersionCmd.Flags().String("platform-arn", "", "The ARN of the platform version.")
	})
	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_describePlatformVersionCmd)
}
