package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_createEnvironmentCmd = &cobra.Command{
	Use:   "create-environment",
	Short: "Create a new FinSpace environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_createEnvironmentCmd).Standalone()

	finspace_createEnvironmentCmd.Flags().String("data-bundles", "", "The list of Amazon Resource Names (ARN) of the data bundles to install.")
	finspace_createEnvironmentCmd.Flags().String("description", "", "The description of the FinSpace environment to be created.")
	finspace_createEnvironmentCmd.Flags().String("federation-mode", "", "Authentication mode for the environment.")
	finspace_createEnvironmentCmd.Flags().String("federation-parameters", "", "Configuration information when authentication mode is FEDERATED.")
	finspace_createEnvironmentCmd.Flags().String("kms-key-id", "", "The KMS key id to encrypt your data in the FinSpace environment.")
	finspace_createEnvironmentCmd.Flags().String("name", "", "The name of the FinSpace environment to be created.")
	finspace_createEnvironmentCmd.Flags().String("superuser-parameters", "", "Configuration information for the superuser.")
	finspace_createEnvironmentCmd.Flags().String("tags", "", "Add tags to your FinSpace environment.")
	finspace_createEnvironmentCmd.MarkFlagRequired("name")
	finspaceCmd.AddCommand(finspace_createEnvironmentCmd)
}
