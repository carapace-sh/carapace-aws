package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amp_describeResourcePolicyCmd = &cobra.Command{
	Use:   "describe-resource-policy",
	Short: "Returns information about the resource-based policy attached to an Amazon Managed Service for Prometheus workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amp_describeResourcePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amp_describeResourcePolicyCmd).Standalone()

		amp_describeResourcePolicyCmd.Flags().String("workspace-id", "", "The ID of the workspace to describe the resource-based policy for.")
		amp_describeResourcePolicyCmd.MarkFlagRequired("workspace-id")
	})
	ampCmd.AddCommand(amp_describeResourcePolicyCmd)
}
