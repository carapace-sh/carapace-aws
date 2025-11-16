package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appintegrations_createApplicationCmd = &cobra.Command{
	Use:   "create-application",
	Short: "Creates and persists an Application resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appintegrations_createApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appintegrations_createApplicationCmd).Standalone()

		appintegrations_createApplicationCmd.Flags().String("application-config", "", "The configuration settings for the application.")
		appintegrations_createApplicationCmd.Flags().String("application-source-config", "", "The configuration for where the application should be loaded from.")
		appintegrations_createApplicationCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		appintegrations_createApplicationCmd.Flags().String("description", "", "The description of the application.")
		appintegrations_createApplicationCmd.Flags().String("iframe-config", "", "The iframe configuration for the application.")
		appintegrations_createApplicationCmd.Flags().String("initialization-timeout", "", "The maximum time in milliseconds allowed to establish a connection with the workspace.")
		appintegrations_createApplicationCmd.Flags().Bool("is-service", false, "Indicates whether the application is a service.")
		appintegrations_createApplicationCmd.Flags().String("name", "", "The name of the application.")
		appintegrations_createApplicationCmd.Flags().String("namespace", "", "The namespace of the application.")
		appintegrations_createApplicationCmd.Flags().Bool("no-is-service", false, "Indicates whether the application is a service.")
		appintegrations_createApplicationCmd.Flags().String("permissions", "", "The configuration of events or requests that the application has access to.")
		appintegrations_createApplicationCmd.Flags().String("publications", "", "The events that the application publishes.")
		appintegrations_createApplicationCmd.Flags().String("subscriptions", "", "The events that the application subscribes.")
		appintegrations_createApplicationCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
		appintegrations_createApplicationCmd.MarkFlagRequired("application-source-config")
		appintegrations_createApplicationCmd.MarkFlagRequired("name")
		appintegrations_createApplicationCmd.MarkFlagRequired("namespace")
		appintegrations_createApplicationCmd.Flag("no-is-service").Hidden = true
	})
	appintegrationsCmd.AddCommand(appintegrations_createApplicationCmd)
}
