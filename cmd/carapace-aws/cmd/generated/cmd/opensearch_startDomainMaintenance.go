package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_startDomainMaintenanceCmd = &cobra.Command{
	Use:   "start-domain-maintenance",
	Short: "Starts the node maintenance process on the data node.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_startDomainMaintenanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearch_startDomainMaintenanceCmd).Standalone()

		opensearch_startDomainMaintenanceCmd.Flags().String("action", "", "The name of the action.")
		opensearch_startDomainMaintenanceCmd.Flags().String("domain-name", "", "The name of the domain.")
		opensearch_startDomainMaintenanceCmd.Flags().String("node-id", "", "The ID of the data node.")
		opensearch_startDomainMaintenanceCmd.MarkFlagRequired("action")
		opensearch_startDomainMaintenanceCmd.MarkFlagRequired("domain-name")
	})
	opensearchCmd.AddCommand(opensearch_startDomainMaintenanceCmd)
}
