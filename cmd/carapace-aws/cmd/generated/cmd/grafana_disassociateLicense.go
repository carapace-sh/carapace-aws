package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var grafana_disassociateLicenseCmd = &cobra.Command{
	Use:   "disassociate-license",
	Short: "Removes the Grafana Enterprise license from a workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(grafana_disassociateLicenseCmd).Standalone()

	grafana_disassociateLicenseCmd.Flags().String("license-type", "", "The type of license to remove from the workspace.")
	grafana_disassociateLicenseCmd.Flags().String("workspace-id", "", "The ID of the workspace to remove the Grafana Enterprise license from.")
	grafana_disassociateLicenseCmd.MarkFlagRequired("license-type")
	grafana_disassociateLicenseCmd.MarkFlagRequired("workspace-id")
	grafanaCmd.AddCommand(grafana_disassociateLicenseCmd)
}
