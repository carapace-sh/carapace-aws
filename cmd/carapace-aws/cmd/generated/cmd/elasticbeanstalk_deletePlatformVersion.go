package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_deletePlatformVersionCmd = &cobra.Command{
	Use:   "delete-platform-version",
	Short: "Deletes the specified version of a custom platform.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_deletePlatformVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticbeanstalk_deletePlatformVersionCmd).Standalone()

		elasticbeanstalk_deletePlatformVersionCmd.Flags().String("platform-arn", "", "The ARN of the version of the custom platform.")
	})
	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_deletePlatformVersionCmd)
}
