package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalyticsv2_updateApplicationCmd = &cobra.Command{
	Use:   "update-application",
	Short: "Updates an existing Managed Service for Apache Flink application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalyticsv2_updateApplicationCmd).Standalone()

	kinesisanalyticsv2_updateApplicationCmd.Flags().String("application-configuration-update", "", "Describes application configuration updates.")
	kinesisanalyticsv2_updateApplicationCmd.Flags().String("application-name", "", "The name of the application to update.")
	kinesisanalyticsv2_updateApplicationCmd.Flags().String("cloud-watch-logging-option-updates", "", "Describes application Amazon CloudWatch logging option updates.")
	kinesisanalyticsv2_updateApplicationCmd.Flags().String("conditional-token", "", "A value you use to implement strong concurrency for application updates.")
	kinesisanalyticsv2_updateApplicationCmd.Flags().String("current-application-version-id", "", "The current application version ID.")
	kinesisanalyticsv2_updateApplicationCmd.Flags().String("run-configuration-update", "", "Describes updates to the application's starting parameters.")
	kinesisanalyticsv2_updateApplicationCmd.Flags().String("runtime-environment-update", "", "Updates the Managed Service for Apache Flink runtime environment used to run your code.")
	kinesisanalyticsv2_updateApplicationCmd.Flags().String("service-execution-role-update", "", "Describes updates to the service execution role.")
	kinesisanalyticsv2_updateApplicationCmd.MarkFlagRequired("application-name")
	kinesisanalyticsv2Cmd.AddCommand(kinesisanalyticsv2_updateApplicationCmd)
}
