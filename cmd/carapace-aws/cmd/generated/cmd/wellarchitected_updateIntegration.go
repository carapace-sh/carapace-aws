package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_updateIntegrationCmd = &cobra.Command{
	Use:   "update-integration",
	Short: "Update integration features.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_updateIntegrationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wellarchitected_updateIntegrationCmd).Standalone()

		wellarchitected_updateIntegrationCmd.Flags().String("client-request-token", "", "")
		wellarchitected_updateIntegrationCmd.Flags().String("integrating-service", "", "Which integrated service to update.")
		wellarchitected_updateIntegrationCmd.Flags().String("workload-id", "", "")
		wellarchitected_updateIntegrationCmd.MarkFlagRequired("client-request-token")
		wellarchitected_updateIntegrationCmd.MarkFlagRequired("integrating-service")
		wellarchitected_updateIntegrationCmd.MarkFlagRequired("workload-id")
	})
	wellarchitectedCmd.AddCommand(wellarchitected_updateIntegrationCmd)
}
