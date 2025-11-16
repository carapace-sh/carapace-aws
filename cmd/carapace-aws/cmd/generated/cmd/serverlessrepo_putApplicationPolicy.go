package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serverlessrepo_putApplicationPolicyCmd = &cobra.Command{
	Use:   "put-application-policy",
	Short: "Sets the permission policy for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serverlessrepo_putApplicationPolicyCmd).Standalone()

	serverlessrepo_putApplicationPolicyCmd.Flags().String("application-id", "", "The Amazon Resource Name (ARN) of the application.")
	serverlessrepo_putApplicationPolicyCmd.Flags().String("statements", "", "An array of policy statements applied to the application.")
	serverlessrepo_putApplicationPolicyCmd.MarkFlagRequired("application-id")
	serverlessrepo_putApplicationPolicyCmd.MarkFlagRequired("statements")
	serverlessrepoCmd.AddCommand(serverlessrepo_putApplicationPolicyCmd)
}
