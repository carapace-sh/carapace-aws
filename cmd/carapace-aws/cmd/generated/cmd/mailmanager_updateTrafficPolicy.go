package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_updateTrafficPolicyCmd = &cobra.Command{
	Use:   "update-traffic-policy",
	Short: "Update attributes of an already provisioned traffic policy resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_updateTrafficPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mailmanager_updateTrafficPolicyCmd).Standalone()

		mailmanager_updateTrafficPolicyCmd.Flags().String("default-action", "", "Default action instructs the traﬃc policy to either Allow or Deny (block) messages that fall outside of (or not addressed by) the conditions of your policy statements")
		mailmanager_updateTrafficPolicyCmd.Flags().String("max-message-size-bytes", "", "The maximum message size in bytes of email which is allowed in by this traffic policy—anything larger will be blocked.")
		mailmanager_updateTrafficPolicyCmd.Flags().String("policy-statements", "", "The list of conditions to be updated for filtering email traffic.")
		mailmanager_updateTrafficPolicyCmd.Flags().String("traffic-policy-id", "", "The identifier of the traffic policy that you want to update.")
		mailmanager_updateTrafficPolicyCmd.Flags().String("traffic-policy-name", "", "A user-friendly name for the traffic policy resource.")
		mailmanager_updateTrafficPolicyCmd.MarkFlagRequired("traffic-policy-id")
	})
	mailmanagerCmd.AddCommand(mailmanager_updateTrafficPolicyCmd)
}
