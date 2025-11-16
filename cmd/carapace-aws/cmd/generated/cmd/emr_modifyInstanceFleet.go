package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_modifyInstanceFleetCmd = &cobra.Command{
	Use:   "modify-instance-fleet",
	Short: "Modifies the target On-Demand and target Spot capacities for the instance fleet with the specified InstanceFleetID within the cluster specified using ClusterID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_modifyInstanceFleetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emr_modifyInstanceFleetCmd).Standalone()

		emr_modifyInstanceFleetCmd.Flags().String("cluster-id", "", "The unique identifier of the cluster.")
		emr_modifyInstanceFleetCmd.Flags().String("instance-fleet", "", "The configuration parameters of the instance fleet.")
		emr_modifyInstanceFleetCmd.MarkFlagRequired("cluster-id")
		emr_modifyInstanceFleetCmd.MarkFlagRequired("instance-fleet")
	})
	emrCmd.AddCommand(emr_modifyInstanceFleetCmd)
}
