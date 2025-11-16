package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serverlessrepo_unshareApplicationCmd = &cobra.Command{
	Use:   "unshare-application",
	Short: "Unshares an application from an AWS Organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serverlessrepo_unshareApplicationCmd).Standalone()

	serverlessrepo_unshareApplicationCmd.Flags().String("application-id", "", "The Amazon Resource Name (ARN) of the application.")
	serverlessrepo_unshareApplicationCmd.Flags().String("organization-id", "", "The AWS Organization ID to unshare the application from.")
	serverlessrepo_unshareApplicationCmd.MarkFlagRequired("application-id")
	serverlessrepo_unshareApplicationCmd.MarkFlagRequired("organization-id")
	serverlessrepoCmd.AddCommand(serverlessrepo_unshareApplicationCmd)
}
