package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dlm_deleteLifecyclePolicyCmd = &cobra.Command{
	Use:   "delete-lifecycle-policy",
	Short: "Deletes the specified lifecycle policy and halts the automated operations that the policy specified.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dlm_deleteLifecyclePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dlm_deleteLifecyclePolicyCmd).Standalone()

		dlm_deleteLifecyclePolicyCmd.Flags().String("policy-id", "", "The identifier of the lifecycle policy.")
		dlm_deleteLifecyclePolicyCmd.MarkFlagRequired("policy-id")
	})
	dlmCmd.AddCommand(dlm_deleteLifecyclePolicyCmd)
}
