package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmDashboards_getResourcePolicyCmd = &cobra.Command{
	Use:   "get-resource-policy",
	Short: "Retrieves the resource-based policy attached to a dashboard, showing sharing configurations and permissions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmDashboards_getResourcePolicyCmd).Standalone()

	bcmDashboards_getResourcePolicyCmd.Flags().String("resource-arn", "", "The ARN of the dashboard whose resource-based policy you want to retrieve.")
	bcmDashboards_getResourcePolicyCmd.MarkFlagRequired("resource-arn")
	bcmDashboardsCmd.AddCommand(bcmDashboards_getResourcePolicyCmd)
}
