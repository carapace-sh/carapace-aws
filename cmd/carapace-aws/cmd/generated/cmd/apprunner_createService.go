package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apprunner_createServiceCmd = &cobra.Command{
	Use:   "create-service",
	Short: "Create an App Runner service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apprunner_createServiceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apprunner_createServiceCmd).Standalone()

		apprunner_createServiceCmd.Flags().String("auto-scaling-configuration-arn", "", "The Amazon Resource Name (ARN) of an App Runner automatic scaling configuration resource that you want to associate with your service.")
		apprunner_createServiceCmd.Flags().String("encryption-configuration", "", "An optional custom encryption key that App Runner uses to encrypt the copy of your source repository that it maintains and your service logs.")
		apprunner_createServiceCmd.Flags().String("health-check-configuration", "", "The settings for the health check that App Runner performs to monitor the health of the App Runner service.")
		apprunner_createServiceCmd.Flags().String("instance-configuration", "", "The runtime configuration of instances (scaling units) of your service.")
		apprunner_createServiceCmd.Flags().String("network-configuration", "", "Configuration settings related to network traffic of the web application that the App Runner service runs.")
		apprunner_createServiceCmd.Flags().String("observability-configuration", "", "The observability configuration of your service.")
		apprunner_createServiceCmd.Flags().String("service-name", "", "A name for the App Runner service.")
		apprunner_createServiceCmd.Flags().String("source-configuration", "", "The source to deploy to the App Runner service.")
		apprunner_createServiceCmd.Flags().String("tags", "", "An optional list of metadata items that you can associate with the App Runner service resource.")
		apprunner_createServiceCmd.MarkFlagRequired("service-name")
		apprunner_createServiceCmd.MarkFlagRequired("source-configuration")
	})
	apprunnerCmd.AddCommand(apprunner_createServiceCmd)
}
