package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_createGrantCmd = &cobra.Command{
	Use:   "create-grant",
	Short: "Adds a grant to a KMS key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_createGrantCmd).Standalone()

	kms_createGrantCmd.Flags().String("constraints", "", "Specifies a grant constraint.")
	kms_createGrantCmd.Flags().String("dry-run", "", "Checks if your request will succeed.")
	kms_createGrantCmd.Flags().String("grant-tokens", "", "A list of grant tokens.")
	kms_createGrantCmd.Flags().String("grantee-principal", "", "The identity that gets the permissions specified in the grant.")
	kms_createGrantCmd.Flags().String("key-id", "", "Identifies the KMS key for the grant.")
	kms_createGrantCmd.Flags().String("name", "", "A friendly name for the grant.")
	kms_createGrantCmd.Flags().String("operations", "", "A list of operations that the grant permits.")
	kms_createGrantCmd.Flags().String("retiring-principal", "", "The principal that has permission to use the [RetireGrant]() operation to retire the grant.")
	kms_createGrantCmd.MarkFlagRequired("grantee-principal")
	kms_createGrantCmd.MarkFlagRequired("key-id")
	kms_createGrantCmd.MarkFlagRequired("operations")
	kmsCmd.AddCommand(kms_createGrantCmd)
}
