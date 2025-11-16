package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dlm_getLifecyclePolicyCmd = &cobra.Command{
	Use:   "get-lifecycle-policy",
	Short: "Gets detailed information about the specified lifecycle policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dlm_getLifecyclePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dlm_getLifecyclePolicyCmd).Standalone()

		dlm_getLifecyclePolicyCmd.Flags().String("policy-id", "", "The identifier of the lifecycle policy.")
		dlm_getLifecyclePolicyCmd.MarkFlagRequired("policy-id")
	})
	dlmCmd.AddCommand(dlm_getLifecyclePolicyCmd)
}
