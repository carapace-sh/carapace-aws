package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_createKxEnvironmentCmd = &cobra.Command{
	Use:   "create-kx-environment",
	Short: "Creates a managed kdb environment for the account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_createKxEnvironmentCmd).Standalone()

	finspace_createKxEnvironmentCmd.Flags().String("client-token", "", "A token that ensures idempotency.")
	finspace_createKxEnvironmentCmd.Flags().String("description", "", "A description for the kdb environment.")
	finspace_createKxEnvironmentCmd.Flags().String("kms-key-id", "", "The KMS key ID to encrypt your data in the FinSpace environment.")
	finspace_createKxEnvironmentCmd.Flags().String("name", "", "The name of the kdb environment that you want to create.")
	finspace_createKxEnvironmentCmd.Flags().String("tags", "", "A list of key-value pairs to label the kdb environment.")
	finspace_createKxEnvironmentCmd.MarkFlagRequired("kms-key-id")
	finspace_createKxEnvironmentCmd.MarkFlagRequired("name")
	finspaceCmd.AddCommand(finspace_createKxEnvironmentCmd)
}
