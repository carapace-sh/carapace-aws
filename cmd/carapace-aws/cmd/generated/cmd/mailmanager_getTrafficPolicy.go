package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_getTrafficPolicyCmd = &cobra.Command{
	Use:   "get-traffic-policy",
	Short: "Fetch attributes of a traffic policy resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_getTrafficPolicyCmd).Standalone()

	mailmanager_getTrafficPolicyCmd.Flags().String("traffic-policy-id", "", "The identifier of the traffic policy resource.")
	mailmanager_getTrafficPolicyCmd.MarkFlagRequired("traffic-policy-id")
	mailmanagerCmd.AddCommand(mailmanager_getTrafficPolicyCmd)
}
