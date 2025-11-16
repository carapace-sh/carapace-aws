package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var grafana_associateLicenseCmd = &cobra.Command{
	Use:   "associate-license",
	Short: "Assigns a Grafana Enterprise license to a workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(grafana_associateLicenseCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(grafana_associateLicenseCmd).Standalone()

		grafana_associateLicenseCmd.Flags().String("grafana-token", "", "A token from Grafana Labs that ties your Amazon Web Services account with a Grafana Labs account.")
		grafana_associateLicenseCmd.Flags().String("license-type", "", "The type of license to associate with the workspace.")
		grafana_associateLicenseCmd.Flags().String("workspace-id", "", "The ID of the workspace to associate the license with.")
		grafana_associateLicenseCmd.MarkFlagRequired("license-type")
		grafana_associateLicenseCmd.MarkFlagRequired("workspace-id")
	})
	grafanaCmd.AddCommand(grafana_associateLicenseCmd)
}
