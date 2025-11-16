package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_listGrantsCmd = &cobra.Command{
	Use:   "list-grants",
	Short: "Gets a list of all grants for the specified KMS key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_listGrantsCmd).Standalone()

	kms_listGrantsCmd.Flags().String("grant-id", "", "Returns only the grant with the specified grant ID.")
	kms_listGrantsCmd.Flags().String("grantee-principal", "", "Returns only grants where the specified principal is the grantee principal for the grant.")
	kms_listGrantsCmd.Flags().String("key-id", "", "Returns only grants for the specified KMS key.")
	kms_listGrantsCmd.Flags().String("limit", "", "Use this parameter to specify the maximum number of items to return.")
	kms_listGrantsCmd.Flags().String("marker", "", "Use this parameter in a subsequent request after you receive a response with truncated results.")
	kms_listGrantsCmd.MarkFlagRequired("key-id")
	kmsCmd.AddCommand(kms_listGrantsCmd)
}
