package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listTargetsForPolicyCmd = &cobra.Command{
	Use:   "list-targets-for-policy",
	Short: "List targets for the specified policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listTargetsForPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_listTargetsForPolicyCmd).Standalone()

		iot_listTargetsForPolicyCmd.Flags().String("marker", "", "A marker used to get the next set of results.")
		iot_listTargetsForPolicyCmd.Flags().String("page-size", "", "The maximum number of results to return at one time.")
		iot_listTargetsForPolicyCmd.Flags().String("policy-name", "", "The policy name.")
		iot_listTargetsForPolicyCmd.MarkFlagRequired("policy-name")
	})
	iotCmd.AddCommand(iot_listTargetsForPolicyCmd)
}
