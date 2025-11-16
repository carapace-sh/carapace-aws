package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_setUnhealthyNodeReplacementCmd = &cobra.Command{
	Use:   "set-unhealthy-node-replacement",
	Short: "Specify whether to enable unhealthy node replacement, which lets Amazon EMR gracefully replace core nodes on a cluster if any nodes become unhealthy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_setUnhealthyNodeReplacementCmd).Standalone()

	emr_setUnhealthyNodeReplacementCmd.Flags().String("job-flow-ids", "", "The list of strings that uniquely identify the clusters for which to turn on unhealthy node replacement.")
	emr_setUnhealthyNodeReplacementCmd.Flags().String("unhealthy-node-replacement", "", "Indicates whether to turn on or turn off graceful unhealthy node replacement.")
	emr_setUnhealthyNodeReplacementCmd.MarkFlagRequired("job-flow-ids")
	emr_setUnhealthyNodeReplacementCmd.MarkFlagRequired("unhealthy-node-replacement")
	emrCmd.AddCommand(emr_setUnhealthyNodeReplacementCmd)
}
