package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_createTrafficPolicyCmd = &cobra.Command{
	Use:   "create-traffic-policy",
	Short: "Provision a new traffic policy resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_createTrafficPolicyCmd).Standalone()

	mailmanager_createTrafficPolicyCmd.Flags().String("client-token", "", "A unique token that Amazon SES uses to recognize subsequent retries of the same request.")
	mailmanager_createTrafficPolicyCmd.Flags().String("default-action", "", "Default action instructs the traﬃc policy to either Allow or Deny (block) messages that fall outside of (or not addressed by) the conditions of your policy statements")
	mailmanager_createTrafficPolicyCmd.Flags().String("max-message-size-bytes", "", "The maximum message size in bytes of email which is allowed in by this traffic policy—anything larger will be blocked.")
	mailmanager_createTrafficPolicyCmd.Flags().String("policy-statements", "", "Conditional statements for filtering email traffic.")
	mailmanager_createTrafficPolicyCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for the resource.")
	mailmanager_createTrafficPolicyCmd.Flags().String("traffic-policy-name", "", "A user-friendly name for the traffic policy resource.")
	mailmanager_createTrafficPolicyCmd.MarkFlagRequired("default-action")
	mailmanager_createTrafficPolicyCmd.MarkFlagRequired("policy-statements")
	mailmanager_createTrafficPolicyCmd.MarkFlagRequired("traffic-policy-name")
	mailmanagerCmd.AddCommand(mailmanager_createTrafficPolicyCmd)
}
