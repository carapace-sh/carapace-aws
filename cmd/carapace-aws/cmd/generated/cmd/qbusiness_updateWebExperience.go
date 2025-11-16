package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_updateWebExperienceCmd = &cobra.Command{
	Use:   "update-web-experience",
	Short: "Updates an Amazon Q Business web experience.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_updateWebExperienceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qbusiness_updateWebExperienceCmd).Standalone()

		qbusiness_updateWebExperienceCmd.Flags().String("application-id", "", "The identifier of the Amazon Q Business application attached to the web experience.")
		qbusiness_updateWebExperienceCmd.Flags().String("authentication-configuration", "", "The authentication configuration of the Amazon Q Business web experience.")
		qbusiness_updateWebExperienceCmd.Flags().String("browser-extension-configuration", "", "The browser extension configuration for an Amazon Q Business web experience.")
		qbusiness_updateWebExperienceCmd.Flags().String("customization-configuration", "", "Updates the custom logo, favicon, font, and color used in the Amazon Q web experience.")
		qbusiness_updateWebExperienceCmd.Flags().String("identity-provider-configuration", "", "Information about the identity provider (IdP) used to authenticate end users of an Amazon Q Business web experience.")
		qbusiness_updateWebExperienceCmd.Flags().String("origins", "", "Updates the website domain origins that are allowed to embed the Amazon Q Business web experience.")
		qbusiness_updateWebExperienceCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the role with permission to access the Amazon Q Business web experience and required resources.")
		qbusiness_updateWebExperienceCmd.Flags().String("sample-prompts-control-mode", "", "Determines whether sample prompts are enabled in the web experience for an end user.")
		qbusiness_updateWebExperienceCmd.Flags().String("subtitle", "", "The subtitle of the Amazon Q Business web experience.")
		qbusiness_updateWebExperienceCmd.Flags().String("title", "", "The title of the Amazon Q Business web experience.")
		qbusiness_updateWebExperienceCmd.Flags().String("web-experience-id", "", "The identifier of the Amazon Q Business web experience.")
		qbusiness_updateWebExperienceCmd.Flags().String("welcome-message", "", "A customized welcome message for an end user in an Amazon Q Business web experience.")
		qbusiness_updateWebExperienceCmd.MarkFlagRequired("application-id")
		qbusiness_updateWebExperienceCmd.MarkFlagRequired("web-experience-id")
	})
	qbusinessCmd.AddCommand(qbusiness_updateWebExperienceCmd)
}
