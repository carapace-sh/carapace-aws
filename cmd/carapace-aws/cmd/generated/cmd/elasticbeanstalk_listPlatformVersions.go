package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_listPlatformVersionsCmd = &cobra.Command{
	Use:   "list-platform-versions",
	Short: "Lists the platform versions available for your account in an AWS Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_listPlatformVersionsCmd).Standalone()

	elasticbeanstalk_listPlatformVersionsCmd.Flags().String("filters", "", "Criteria for restricting the resulting list of platform versions.")
	elasticbeanstalk_listPlatformVersionsCmd.Flags().String("max-records", "", "The maximum number of platform version values returned in one call.")
	elasticbeanstalk_listPlatformVersionsCmd.Flags().String("next-token", "", "For a paginated request.")
	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_listPlatformVersionsCmd)
}
