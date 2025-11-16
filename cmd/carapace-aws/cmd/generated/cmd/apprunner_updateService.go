package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apprunner_updateServiceCmd = &cobra.Command{
	Use:   "update-service",
	Short: "Update an App Runner service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apprunner_updateServiceCmd).Standalone()

	apprunner_updateServiceCmd.Flags().String("auto-scaling-configuration-arn", "", "The Amazon Resource Name (ARN) of an App Runner automatic scaling configuration resource that you want to associate with the App Runner service.")
	apprunner_updateServiceCmd.Flags().String("health-check-configuration", "", "The settings for the health check that App Runner performs to monitor the health of the App Runner service.")
	apprunner_updateServiceCmd.Flags().String("instance-configuration", "", "The runtime configuration to apply to instances (scaling units) of your service.")
	apprunner_updateServiceCmd.Flags().String("network-configuration", "", "Configuration settings related to network traffic of the web application that the App Runner service runs.")
	apprunner_updateServiceCmd.Flags().String("observability-configuration", "", "The observability configuration of your service.")
	apprunner_updateServiceCmd.Flags().String("service-arn", "", "The Amazon Resource Name (ARN) of the App Runner service that you want to update.")
	apprunner_updateServiceCmd.Flags().String("source-configuration", "", "The source configuration to apply to the App Runner service.")
	apprunner_updateServiceCmd.MarkFlagRequired("service-arn")
	apprunnerCmd.AddCommand(apprunner_updateServiceCmd)
}
