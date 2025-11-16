package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_createAnonymousWebExperienceUrlCmd = &cobra.Command{
	Use:   "create-anonymous-web-experience-url",
	Short: "Creates a unique URL for anonymous Amazon Q Business web experience.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_createAnonymousWebExperienceUrlCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qbusiness_createAnonymousWebExperienceUrlCmd).Standalone()

		qbusiness_createAnonymousWebExperienceUrlCmd.Flags().String("application-id", "", "The identifier of the Amazon Q Business application environment attached to the web experience.")
		qbusiness_createAnonymousWebExperienceUrlCmd.Flags().String("session-duration-in-minutes", "", "The duration of the session associated with the unique URL for the web experience.")
		qbusiness_createAnonymousWebExperienceUrlCmd.Flags().String("web-experience-id", "", "The identifier of the web experience.")
		qbusiness_createAnonymousWebExperienceUrlCmd.MarkFlagRequired("application-id")
		qbusiness_createAnonymousWebExperienceUrlCmd.MarkFlagRequired("web-experience-id")
	})
	qbusinessCmd.AddCommand(qbusiness_createAnonymousWebExperienceUrlCmd)
}
