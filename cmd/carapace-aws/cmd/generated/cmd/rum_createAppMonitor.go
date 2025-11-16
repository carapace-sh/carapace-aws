package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rum_createAppMonitorCmd = &cobra.Command{
	Use:   "create-app-monitor",
	Short: "Creates a Amazon CloudWatch RUM app monitor, which collects telemetry data from your application and sends that data to RUM.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rum_createAppMonitorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rum_createAppMonitorCmd).Standalone()

		rum_createAppMonitorCmd.Flags().String("app-monitor-configuration", "", "A structure that contains much of the configuration data for the app monitor.")
		rum_createAppMonitorCmd.Flags().String("custom-events", "", "Specifies whether this app monitor allows the web client to define and send custom events.")
		rum_createAppMonitorCmd.Flags().Bool("cw-log-enabled", false, "Data collected by RUM is kept by RUM for 30 days and then deleted.")
		rum_createAppMonitorCmd.Flags().String("deobfuscation-configuration", "", "A structure that contains the configuration for how an app monitor can deobfuscate stack traces.")
		rum_createAppMonitorCmd.Flags().String("domain", "", "The top-level internet domain name for which your application has administrative authority.")
		rum_createAppMonitorCmd.Flags().String("domain-list", "", "List the domain names for which your application has administrative authority.")
		rum_createAppMonitorCmd.Flags().String("name", "", "A name for the app monitor.")
		rum_createAppMonitorCmd.Flags().Bool("no-cw-log-enabled", false, "Data collected by RUM is kept by RUM for 30 days and then deleted.")
		rum_createAppMonitorCmd.Flags().String("tags", "", "Assigns one or more tags (key-value pairs) to the app monitor.")
		rum_createAppMonitorCmd.MarkFlagRequired("name")
		rum_createAppMonitorCmd.Flag("no-cw-log-enabled").Hidden = true
	})
	rumCmd.AddCommand(rum_createAppMonitorCmd)
}
