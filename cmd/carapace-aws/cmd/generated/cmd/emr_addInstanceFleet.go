package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_addInstanceFleetCmd = &cobra.Command{
	Use:   "add-instance-fleet",
	Short: "Adds an instance fleet to a running cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_addInstanceFleetCmd).Standalone()

	emr_addInstanceFleetCmd.Flags().String("cluster-id", "", "The unique identifier of the cluster.")
	emr_addInstanceFleetCmd.Flags().String("instance-fleet", "", "Specifies the configuration of the instance fleet.")
	emr_addInstanceFleetCmd.MarkFlagRequired("cluster-id")
	emr_addInstanceFleetCmd.MarkFlagRequired("instance-fleet")
	emrCmd.AddCommand(emr_addInstanceFleetCmd)
}
