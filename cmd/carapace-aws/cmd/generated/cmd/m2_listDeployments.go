package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var m2_listDeploymentsCmd = &cobra.Command{
	Use:   "list-deployments",
	Short: "Returns a list of all deployments of a specific application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(m2_listDeploymentsCmd).Standalone()

	m2_listDeploymentsCmd.Flags().String("application-id", "", "The application identifier.")
	m2_listDeploymentsCmd.Flags().String("max-results", "", "The maximum number of objects to return.")
	m2_listDeploymentsCmd.Flags().String("next-token", "", "A pagination token returned from a previous call to this operation.")
	m2_listDeploymentsCmd.MarkFlagRequired("application-id")
	m2Cmd.AddCommand(m2_listDeploymentsCmd)
}
