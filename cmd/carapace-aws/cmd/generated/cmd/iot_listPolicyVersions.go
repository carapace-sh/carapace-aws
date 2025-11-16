package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listPolicyVersionsCmd = &cobra.Command{
	Use:   "list-policy-versions",
	Short: "Lists the versions of the specified policy and identifies the default version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listPolicyVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_listPolicyVersionsCmd).Standalone()

		iot_listPolicyVersionsCmd.Flags().String("policy-name", "", "The policy name.")
		iot_listPolicyVersionsCmd.MarkFlagRequired("policy-name")
	})
	iotCmd.AddCommand(iot_listPolicyVersionsCmd)
}
