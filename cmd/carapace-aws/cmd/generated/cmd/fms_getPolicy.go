package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fms_getPolicyCmd = &cobra.Command{
	Use:   "get-policy",
	Short: "Returns information about the specified Firewall Manager policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fms_getPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fms_getPolicyCmd).Standalone()

		fms_getPolicyCmd.Flags().String("policy-id", "", "The ID of the Firewall Manager policy that you want the details for.")
		fms_getPolicyCmd.MarkFlagRequired("policy-id")
	})
	fmsCmd.AddCommand(fms_getPolicyCmd)
}
