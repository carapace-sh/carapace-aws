package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_deleteEmailMonitoringConfigurationCmd = &cobra.Command{
	Use:   "delete-email-monitoring-configuration",
	Short: "Deletes the email monitoring configuration for a specified organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_deleteEmailMonitoringConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workmail_deleteEmailMonitoringConfigurationCmd).Standalone()

		workmail_deleteEmailMonitoringConfigurationCmd.Flags().String("organization-id", "", "The ID of the organization from which the email monitoring configuration is deleted.")
		workmail_deleteEmailMonitoringConfigurationCmd.MarkFlagRequired("organization-id")
	})
	workmailCmd.AddCommand(workmail_deleteEmailMonitoringConfigurationCmd)
}
