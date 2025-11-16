package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var verifiedpermissions_batchGetPolicyCmd = &cobra.Command{
	Use:   "batch-get-policy",
	Short: "Retrieves information about a group (batch) of policies.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifiedpermissions_batchGetPolicyCmd).Standalone()

	verifiedpermissions_batchGetPolicyCmd.Flags().String("requests", "", "An array of up to 100 policies you want information about.")
	verifiedpermissions_batchGetPolicyCmd.MarkFlagRequired("requests")
	verifiedpermissionsCmd.AddCommand(verifiedpermissions_batchGetPolicyCmd)
}
