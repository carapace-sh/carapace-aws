package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_createWebExperienceCmd = &cobra.Command{
	Use:   "create-web-experience",
	Short: "Creates an Amazon Q Business web experience.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_createWebExperienceCmd).Standalone()

	qbusiness_createWebExperienceCmd.Flags().String("application-id", "", "The identifier of the Amazon Q Business web experience.")
	qbusiness_createWebExperienceCmd.Flags().String("browser-extension-configuration", "", "The browser extension configuration for an Amazon Q Business web experience.")
	qbusiness_createWebExperienceCmd.Flags().String("client-token", "", "A token you provide to identify a request to create an Amazon Q Business web experience.")
	qbusiness_createWebExperienceCmd.Flags().String("customization-configuration", "", "Sets the custom logo, favicon, font, and color used in the Amazon Q web experience.")
	qbusiness_createWebExperienceCmd.Flags().String("identity-provider-configuration", "", "Information about the identity provider (IdP) used to authenticate end users of an Amazon Q Business web experience.")
	qbusiness_createWebExperienceCmd.Flags().String("origins", "", "Sets the website domain origins that are allowed to embed the Amazon Q Business web experience.")
	qbusiness_createWebExperienceCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the service role attached to your web experience.")
	qbusiness_createWebExperienceCmd.Flags().String("sample-prompts-control-mode", "", "Determines whether sample prompts are enabled in the web experience for an end user.")
	qbusiness_createWebExperienceCmd.Flags().String("subtitle", "", "A subtitle to personalize your Amazon Q Business web experience.")
	qbusiness_createWebExperienceCmd.Flags().String("tags", "", "A list of key-value pairs that identify or categorize your Amazon Q Business web experience.")
	qbusiness_createWebExperienceCmd.Flags().String("title", "", "The title for your Amazon Q Business web experience.")
	qbusiness_createWebExperienceCmd.Flags().String("welcome-message", "", "The customized welcome message for end users of an Amazon Q Business web experience.")
	qbusiness_createWebExperienceCmd.MarkFlagRequired("application-id")
	qbusinessCmd.AddCommand(qbusiness_createWebExperienceCmd)
}
