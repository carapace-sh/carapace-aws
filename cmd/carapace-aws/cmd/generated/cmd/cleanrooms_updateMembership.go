package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_updateMembershipCmd = &cobra.Command{
	Use:   "update-membership",
	Short: "Updates a membership.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_updateMembershipCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_updateMembershipCmd).Standalone()

		cleanrooms_updateMembershipCmd.Flags().String("default-job-result-configuration", "", "The default job result configuration.")
		cleanrooms_updateMembershipCmd.Flags().String("default-result-configuration", "", "The default protected query result configuration as specified by the member who can receive results.")
		cleanrooms_updateMembershipCmd.Flags().String("job-log-status", "", "An indicator as to whether job logging has been enabled or disabled for the collaboration.")
		cleanrooms_updateMembershipCmd.Flags().String("membership-identifier", "", "The unique identifier of the membership.")
		cleanrooms_updateMembershipCmd.Flags().String("query-log-status", "", "An indicator as to whether query logging has been enabled or disabled for the membership.")
		cleanrooms_updateMembershipCmd.MarkFlagRequired("membership-identifier")
	})
	cleanroomsCmd.AddCommand(cleanrooms_updateMembershipCmd)
}
