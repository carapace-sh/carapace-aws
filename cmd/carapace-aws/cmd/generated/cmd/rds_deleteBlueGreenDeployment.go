package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_deleteBlueGreenDeploymentCmd = &cobra.Command{
	Use:   "delete-blue-green-deployment",
	Short: "Deletes a blue/green deployment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_deleteBlueGreenDeploymentCmd).Standalone()

	rds_deleteBlueGreenDeploymentCmd.Flags().String("blue-green-deployment-identifier", "", "The unique identifier of the blue/green deployment to delete.")
	rds_deleteBlueGreenDeploymentCmd.Flags().String("delete-target", "", "Specifies whether to delete the resources in the green environment.")
	rds_deleteBlueGreenDeploymentCmd.MarkFlagRequired("blue-green-deployment-identifier")
	rdsCmd.AddCommand(rds_deleteBlueGreenDeploymentCmd)
}
