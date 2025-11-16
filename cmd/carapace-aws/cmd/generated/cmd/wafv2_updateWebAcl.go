package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_updateWebAclCmd = &cobra.Command{
	Use:   "update-web-acl",
	Short: "Updates the specified [WebACL]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_updateWebAclCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafv2_updateWebAclCmd).Standalone()

		wafv2_updateWebAclCmd.Flags().String("application-config", "", "Configures the ability for the WAF console to store and retrieve application attributes.")
		wafv2_updateWebAclCmd.Flags().String("association-config", "", "Specifies custom configurations for the associations between the web ACL and protected resources.")
		wafv2_updateWebAclCmd.Flags().String("captcha-config", "", "Specifies how WAF should handle `CAPTCHA` evaluations for rules that don't have their own `CaptchaConfig` settings.")
		wafv2_updateWebAclCmd.Flags().String("challenge-config", "", "Specifies how WAF should handle challenge evaluations for rules that don't have their own `ChallengeConfig` settings.")
		wafv2_updateWebAclCmd.Flags().String("custom-response-bodies", "", "A map of custom response keys and content bodies.")
		wafv2_updateWebAclCmd.Flags().String("data-protection-config", "", "Specifies data protection to apply to the web request data for the web ACL.")
		wafv2_updateWebAclCmd.Flags().String("default-action", "", "The action to perform if none of the `Rules` contained in the `WebACL` match.")
		wafv2_updateWebAclCmd.Flags().String("description", "", "A description of the web ACL that helps with identification.")
		wafv2_updateWebAclCmd.Flags().String("id", "", "The unique identifier for the web ACL.")
		wafv2_updateWebAclCmd.Flags().String("lock-token", "", "A token used for optimistic locking.")
		wafv2_updateWebAclCmd.Flags().String("name", "", "The name of the web ACL.")
		wafv2_updateWebAclCmd.Flags().String("on-source-ddo-sprotection-config", "", "Specifies the type of DDoS protection to apply to web request data for a web ACL.")
		wafv2_updateWebAclCmd.Flags().String("rules", "", "The [Rule]() statements used to identify the web requests that you want to manage.")
		wafv2_updateWebAclCmd.Flags().String("scope", "", "Specifies whether this is for a global resource type, such as a Amazon CloudFront distribution.")
		wafv2_updateWebAclCmd.Flags().String("token-domains", "", "Specifies the domains that WAF should accept in a web request token.")
		wafv2_updateWebAclCmd.Flags().String("visibility-config", "", "Defines and enables Amazon CloudWatch metrics and web request sample collection.")
		wafv2_updateWebAclCmd.MarkFlagRequired("default-action")
		wafv2_updateWebAclCmd.MarkFlagRequired("id")
		wafv2_updateWebAclCmd.MarkFlagRequired("lock-token")
		wafv2_updateWebAclCmd.MarkFlagRequired("name")
		wafv2_updateWebAclCmd.MarkFlagRequired("scope")
		wafv2_updateWebAclCmd.MarkFlagRequired("visibility-config")
	})
	wafv2Cmd.AddCommand(wafv2_updateWebAclCmd)
}
