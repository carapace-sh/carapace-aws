package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_createWebAclCmd = &cobra.Command{
	Use:   "create-web-acl",
	Short: "Creates a [WebACL]() per the specifications provided.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_createWebAclCmd).Standalone()

	wafv2_createWebAclCmd.Flags().String("application-config", "", "Configures the ability for the WAF console to store and retrieve application attributes during the web ACL creation process.")
	wafv2_createWebAclCmd.Flags().String("association-config", "", "Specifies custom configurations for the associations between the web ACL and protected resources.")
	wafv2_createWebAclCmd.Flags().String("captcha-config", "", "Specifies how WAF should handle `CAPTCHA` evaluations for rules that don't have their own `CaptchaConfig` settings.")
	wafv2_createWebAclCmd.Flags().String("challenge-config", "", "Specifies how WAF should handle challenge evaluations for rules that don't have their own `ChallengeConfig` settings.")
	wafv2_createWebAclCmd.Flags().String("custom-response-bodies", "", "A map of custom response keys and content bodies.")
	wafv2_createWebAclCmd.Flags().String("data-protection-config", "", "Specifies data protection to apply to the web request data for the web ACL.")
	wafv2_createWebAclCmd.Flags().String("default-action", "", "The action to perform if none of the `Rules` contained in the `WebACL` match.")
	wafv2_createWebAclCmd.Flags().String("description", "", "A description of the web ACL that helps with identification.")
	wafv2_createWebAclCmd.Flags().String("name", "", "The name of the web ACL.")
	wafv2_createWebAclCmd.Flags().String("on-source-ddo-sprotection-config", "", "Specifies the type of DDoS protection to apply to web request data for a web ACL.")
	wafv2_createWebAclCmd.Flags().String("rules", "", "The [Rule]() statements used to identify the web requests that you want to manage.")
	wafv2_createWebAclCmd.Flags().String("scope", "", "Specifies whether this is for a global resource type, such as a Amazon CloudFront distribution.")
	wafv2_createWebAclCmd.Flags().String("tags", "", "An array of key:value pairs to associate with the resource.")
	wafv2_createWebAclCmd.Flags().String("token-domains", "", "Specifies the domains that WAF should accept in a web request token.")
	wafv2_createWebAclCmd.Flags().String("visibility-config", "", "Defines and enables Amazon CloudWatch metrics and web request sample collection.")
	wafv2_createWebAclCmd.MarkFlagRequired("default-action")
	wafv2_createWebAclCmd.MarkFlagRequired("name")
	wafv2_createWebAclCmd.MarkFlagRequired("scope")
	wafv2_createWebAclCmd.MarkFlagRequired("visibility-config")
	wafv2Cmd.AddCommand(wafv2_createWebAclCmd)
}
