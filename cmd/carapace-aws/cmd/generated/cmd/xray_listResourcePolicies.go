package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var xray_listResourcePoliciesCmd = &cobra.Command{
	Use:   "list-resource-policies",
	Short: "Returns the list of resource policies in the target Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(xray_listResourcePoliciesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(xray_listResourcePoliciesCmd).Standalone()

		xray_listResourcePoliciesCmd.Flags().String("next-token", "", "Not currently supported.")
	})
	xrayCmd.AddCommand(xray_listResourcePoliciesCmd)
}
