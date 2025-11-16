package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_listPlatformBranchesCmd = &cobra.Command{
	Use:   "list-platform-branches",
	Short: "Lists the platform branches available for your account in an AWS Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_listPlatformBranchesCmd).Standalone()

	elasticbeanstalk_listPlatformBranchesCmd.Flags().String("filters", "", "Criteria for restricting the resulting list of platform branches.")
	elasticbeanstalk_listPlatformBranchesCmd.Flags().String("max-records", "", "The maximum number of platform branch values returned in one call.")
	elasticbeanstalk_listPlatformBranchesCmd.Flags().String("next-token", "", "For a paginated request.")
	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_listPlatformBranchesCmd)
}
