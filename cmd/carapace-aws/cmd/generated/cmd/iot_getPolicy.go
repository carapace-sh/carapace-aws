package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_getPolicyCmd = &cobra.Command{
	Use:   "get-policy",
	Short: "Gets information about the specified policy with the policy document of the default version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_getPolicyCmd).Standalone()

	iot_getPolicyCmd.Flags().String("policy-name", "", "The name of the policy.")
	iot_getPolicyCmd.MarkFlagRequired("policy-name")
	iotCmd.AddCommand(iot_getPolicyCmd)
}
