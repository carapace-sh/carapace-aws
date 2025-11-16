package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rum_updateAppMonitorCmd = &cobra.Command{
	Use:   "update-app-monitor",
	Short: "Updates the configuration of an existing app monitor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rum_updateAppMonitorCmd).Standalone()

	rum_updateAppMonitorCmd.Flags().String("app-monitor-configuration", "", "A structure that contains much of the configuration data for the app monitor.")
	rum_updateAppMonitorCmd.Flags().String("custom-events", "", "Specifies whether this app monitor allows the web client to define and send custom events.")
	rum_updateAppMonitorCmd.Flags().Bool("cw-log-enabled", false, "Data collected by RUM is kept by RUM for 30 days and then deleted.")
	rum_updateAppMonitorCmd.Flags().String("deobfuscation-configuration", "", "A structure that contains the configuration for how an app monitor can deobfuscate stack traces.")
	rum_updateAppMonitorCmd.Flags().String("domain", "", "The top-level internet domain name for which your application has administrative authority.")
	rum_updateAppMonitorCmd.Flags().String("domain-list", "", "List the domain names for which your application has administrative authority.")
	rum_updateAppMonitorCmd.Flags().String("name", "", "The name of the app monitor to update.")
	rum_updateAppMonitorCmd.Flags().Bool("no-cw-log-enabled", false, "Data collected by RUM is kept by RUM for 30 days and then deleted.")
	rum_updateAppMonitorCmd.MarkFlagRequired("name")
	rum_updateAppMonitorCmd.Flag("no-cw-log-enabled").Hidden = true
	rumCmd.AddCommand(rum_updateAppMonitorCmd)
}
