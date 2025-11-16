package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_updateKxClusterCodeConfigurationCmd = &cobra.Command{
	Use:   "update-kx-cluster-code-configuration",
	Short: "Allows you to update code configuration on a running cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_updateKxClusterCodeConfigurationCmd).Standalone()

	finspace_updateKxClusterCodeConfigurationCmd.Flags().String("client-token", "", "A token that ensures idempotency.")
	finspace_updateKxClusterCodeConfigurationCmd.Flags().String("cluster-name", "", "The name of the cluster.")
	finspace_updateKxClusterCodeConfigurationCmd.Flags().String("code", "", "")
	finspace_updateKxClusterCodeConfigurationCmd.Flags().String("command-line-arguments", "", "Specifies the key-value pairs to make them available inside the cluster.")
	finspace_updateKxClusterCodeConfigurationCmd.Flags().String("deployment-configuration", "", "The configuration that allows you to choose how you want to update the code on a cluster.")
	finspace_updateKxClusterCodeConfigurationCmd.Flags().String("environment-id", "", "A unique identifier of the kdb environment.")
	finspace_updateKxClusterCodeConfigurationCmd.Flags().String("initialization-script", "", "Specifies a Q program that will be run at launch of a cluster.")
	finspace_updateKxClusterCodeConfigurationCmd.MarkFlagRequired("cluster-name")
	finspace_updateKxClusterCodeConfigurationCmd.MarkFlagRequired("code")
	finspace_updateKxClusterCodeConfigurationCmd.MarkFlagRequired("environment-id")
	finspaceCmd.AddCommand(finspace_updateKxClusterCodeConfigurationCmd)
}
