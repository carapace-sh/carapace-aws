package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_deleteTrafficPolicyCmd = &cobra.Command{
	Use:   "delete-traffic-policy",
	Short: "Delete a traffic policy resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_deleteTrafficPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mailmanager_deleteTrafficPolicyCmd).Standalone()

		mailmanager_deleteTrafficPolicyCmd.Flags().String("traffic-policy-id", "", "The identifier of the traffic policy that you want to delete.")
		mailmanager_deleteTrafficPolicyCmd.MarkFlagRequired("traffic-policy-id")
	})
	mailmanagerCmd.AddCommand(mailmanager_deleteTrafficPolicyCmd)
}
