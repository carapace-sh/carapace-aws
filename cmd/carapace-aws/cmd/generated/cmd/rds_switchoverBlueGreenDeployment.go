package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_switchoverBlueGreenDeploymentCmd = &cobra.Command{
	Use:   "switchover-blue-green-deployment",
	Short: "Switches over a blue/green deployment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_switchoverBlueGreenDeploymentCmd).Standalone()

	rds_switchoverBlueGreenDeploymentCmd.Flags().String("blue-green-deployment-identifier", "", "The resource ID of the blue/green deployment.")
	rds_switchoverBlueGreenDeploymentCmd.Flags().String("switchover-timeout", "", "The amount of time, in seconds, for the switchover to complete.")
	rds_switchoverBlueGreenDeploymentCmd.MarkFlagRequired("blue-green-deployment-identifier")
	rdsCmd.AddCommand(rds_switchoverBlueGreenDeploymentCmd)
}
