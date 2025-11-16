package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_createDashboardCmd = &cobra.Command{
	Use:   "create-dashboard",
	Short: "Creates a custom dashboard or the Highlights dashboard.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_createDashboardCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudtrail_createDashboardCmd).Standalone()

		cloudtrail_createDashboardCmd.Flags().String("name", "", "The name of the dashboard.")
		cloudtrail_createDashboardCmd.Flags().String("refresh-schedule", "", "The refresh schedule configuration for the dashboard.")
		cloudtrail_createDashboardCmd.Flags().String("tags-list", "", "")
		cloudtrail_createDashboardCmd.Flags().String("termination-protection-enabled", "", "Specifies whether termination protection is enabled for the dashboard.")
		cloudtrail_createDashboardCmd.Flags().String("widgets", "", "An array of widgets for a custom dashboard.")
		cloudtrail_createDashboardCmd.MarkFlagRequired("name")
	})
	cloudtrailCmd.AddCommand(cloudtrail_createDashboardCmd)
}
