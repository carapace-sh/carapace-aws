package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var m2_createDeploymentCmd = &cobra.Command{
	Use:   "create-deployment",
	Short: "Creates and starts a deployment to deploy an application into a runtime environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(m2_createDeploymentCmd).Standalone()

	m2_createDeploymentCmd.Flags().String("application-id", "", "The application identifier.")
	m2_createDeploymentCmd.Flags().String("application-version", "", "The version of the application to deploy.")
	m2_createDeploymentCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier you provide to ensure the idempotency of the request to create a deployment.")
	m2_createDeploymentCmd.Flags().String("environment-id", "", "The identifier of the runtime environment where you want to deploy this application.")
	m2_createDeploymentCmd.MarkFlagRequired("application-id")
	m2_createDeploymentCmd.MarkFlagRequired("application-version")
	m2_createDeploymentCmd.MarkFlagRequired("environment-id")
	m2Cmd.AddCommand(m2_createDeploymentCmd)
}
