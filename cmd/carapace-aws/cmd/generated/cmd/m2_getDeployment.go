package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var m2_getDeploymentCmd = &cobra.Command{
	Use:   "get-deployment",
	Short: "Gets details of a specific deployment with a given deployment identifier.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(m2_getDeploymentCmd).Standalone()

	m2_getDeploymentCmd.Flags().String("application-id", "", "The unique identifier of the application.")
	m2_getDeploymentCmd.Flags().String("deployment-id", "", "The unique identifier for the deployment.")
	m2_getDeploymentCmd.MarkFlagRequired("application-id")
	m2_getDeploymentCmd.MarkFlagRequired("deployment-id")
	m2Cmd.AddCommand(m2_getDeploymentCmd)
}
