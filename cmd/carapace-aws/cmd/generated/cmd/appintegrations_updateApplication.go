package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appintegrations_updateApplicationCmd = &cobra.Command{
	Use:   "update-application",
	Short: "Updates and persists an Application resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appintegrations_updateApplicationCmd).Standalone()

	appintegrations_updateApplicationCmd.Flags().String("application-config", "", "The configuration settings for the application.")
	appintegrations_updateApplicationCmd.Flags().String("application-source-config", "", "The configuration for where the application should be loaded from.")
	appintegrations_updateApplicationCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the Application.")
	appintegrations_updateApplicationCmd.Flags().String("description", "", "The description of the application.")
	appintegrations_updateApplicationCmd.Flags().String("iframe-config", "", "The iframe configuration for the application.")
	appintegrations_updateApplicationCmd.Flags().String("initialization-timeout", "", "The maximum time in milliseconds allowed to establish a connection with the workspace.")
	appintegrations_updateApplicationCmd.Flags().Bool("is-service", false, "Indicates whether the application is a service.")
	appintegrations_updateApplicationCmd.Flags().String("name", "", "The name of the application.")
	appintegrations_updateApplicationCmd.Flags().Bool("no-is-service", false, "Indicates whether the application is a service.")
	appintegrations_updateApplicationCmd.Flags().String("permissions", "", "The configuration of events or requests that the application has access to.")
	appintegrations_updateApplicationCmd.Flags().String("publications", "", "The events that the application publishes.")
	appintegrations_updateApplicationCmd.Flags().String("subscriptions", "", "The events that the application subscribes.")
	appintegrations_updateApplicationCmd.MarkFlagRequired("arn")
	appintegrations_updateApplicationCmd.Flag("no-is-service").Hidden = true
	appintegrationsCmd.AddCommand(appintegrations_updateApplicationCmd)
}
