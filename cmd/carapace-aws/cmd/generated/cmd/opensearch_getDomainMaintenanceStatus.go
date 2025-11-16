package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_getDomainMaintenanceStatusCmd = &cobra.Command{
	Use:   "get-domain-maintenance-status",
	Short: "The status of the maintenance action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_getDomainMaintenanceStatusCmd).Standalone()

	opensearch_getDomainMaintenanceStatusCmd.Flags().String("domain-name", "", "The name of the domain.")
	opensearch_getDomainMaintenanceStatusCmd.Flags().String("maintenance-id", "", "The request ID of the maintenance action.")
	opensearch_getDomainMaintenanceStatusCmd.MarkFlagRequired("domain-name")
	opensearch_getDomainMaintenanceStatusCmd.MarkFlagRequired("maintenance-id")
	opensearchCmd.AddCommand(opensearch_getDomainMaintenanceStatusCmd)
}
