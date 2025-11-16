package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_listBootstrapActionsCmd = &cobra.Command{
	Use:   "list-bootstrap-actions",
	Short: "Provides information about the bootstrap actions associated with a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_listBootstrapActionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emr_listBootstrapActionsCmd).Standalone()

		emr_listBootstrapActionsCmd.Flags().String("cluster-id", "", "The cluster identifier for the bootstrap actions to list.")
		emr_listBootstrapActionsCmd.Flags().String("marker", "", "The pagination token that indicates the next set of results to retrieve.")
		emr_listBootstrapActionsCmd.MarkFlagRequired("cluster-id")
	})
	emrCmd.AddCommand(emr_listBootstrapActionsCmd)
}
