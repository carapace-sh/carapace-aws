package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_deletePolicyCmd = &cobra.Command{
	Use:   "delete-policy",
	Short: "Deletes the specified policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_deletePolicyCmd).Standalone()

	iot_deletePolicyCmd.Flags().String("policy-name", "", "The name of the policy to delete.")
	iot_deletePolicyCmd.MarkFlagRequired("policy-name")
	iotCmd.AddCommand(iot_deletePolicyCmd)
}
