package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serverlessrepo_getApplicationPolicyCmd = &cobra.Command{
	Use:   "get-application-policy",
	Short: "Retrieves the policy for the application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serverlessrepo_getApplicationPolicyCmd).Standalone()

	serverlessrepo_getApplicationPolicyCmd.Flags().String("application-id", "", "The Amazon Resource Name (ARN) of the application.")
	serverlessrepo_getApplicationPolicyCmd.MarkFlagRequired("application-id")
	serverlessrepoCmd.AddCommand(serverlessrepo_getApplicationPolicyCmd)
}
