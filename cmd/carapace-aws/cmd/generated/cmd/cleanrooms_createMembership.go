package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_createMembershipCmd = &cobra.Command{
	Use:   "create-membership",
	Short: "Creates a membership for a specific collaboration identifier and joins the collaboration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_createMembershipCmd).Standalone()

	cleanrooms_createMembershipCmd.Flags().String("collaboration-identifier", "", "The unique ID for the associated collaboration.")
	cleanrooms_createMembershipCmd.Flags().String("default-job-result-configuration", "", "The default job result configuration that determines how job results are protected and managed within this membership.")
	cleanrooms_createMembershipCmd.Flags().String("default-result-configuration", "", "The default protected query result configuration as specified by the member who can receive results.")
	cleanrooms_createMembershipCmd.Flags().String("job-log-status", "", "An indicator as to whether job logging has been enabled or disabled for the collaboration.")
	cleanrooms_createMembershipCmd.Flags().String("payment-configuration", "", "The payment responsibilities accepted by the collaboration member.")
	cleanrooms_createMembershipCmd.Flags().String("query-log-status", "", "An indicator as to whether query logging has been enabled or disabled for the membership.")
	cleanrooms_createMembershipCmd.Flags().String("tags", "", "An optional label that you can assign to a resource when you create it.")
	cleanrooms_createMembershipCmd.MarkFlagRequired("collaboration-identifier")
	cleanrooms_createMembershipCmd.MarkFlagRequired("query-log-status")
	cleanroomsCmd.AddCommand(cleanrooms_createMembershipCmd)
}
