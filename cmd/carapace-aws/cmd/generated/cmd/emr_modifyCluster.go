package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_modifyClusterCmd = &cobra.Command{
	Use:   "modify-cluster",
	Short: "Modifies the number of steps that can be executed concurrently for the cluster specified using ClusterID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_modifyClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emr_modifyClusterCmd).Standalone()

		emr_modifyClusterCmd.Flags().String("cluster-id", "", "The unique identifier of the cluster.")
		emr_modifyClusterCmd.Flags().String("extended-support", "", "Reserved.")
		emr_modifyClusterCmd.Flags().String("step-concurrency-level", "", "The number of steps that can be executed concurrently.")
		emr_modifyClusterCmd.MarkFlagRequired("cluster-id")
	})
	emrCmd.AddCommand(emr_modifyClusterCmd)
}
