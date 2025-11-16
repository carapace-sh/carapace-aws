package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serverlessrepo_updateApplicationCmd = &cobra.Command{
	Use:   "update-application",
	Short: "Updates the specified application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serverlessrepo_updateApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(serverlessrepo_updateApplicationCmd).Standalone()

		serverlessrepo_updateApplicationCmd.Flags().String("application-id", "", "The Amazon Resource Name (ARN) of the application.")
		serverlessrepo_updateApplicationCmd.Flags().String("author", "", "The name of the author publishing the app.")
		serverlessrepo_updateApplicationCmd.Flags().String("description", "", "The description of the application.")
		serverlessrepo_updateApplicationCmd.Flags().String("home-page-url", "", "A URL with more information about the application, for example the location of your GitHub repository for the application.")
		serverlessrepo_updateApplicationCmd.Flags().String("labels", "", "Labels to improve discovery of apps in search results.")
		serverlessrepo_updateApplicationCmd.Flags().String("readme-body", "", "A text readme file in Markdown language that contains a more detailed description of the application and how it works.")
		serverlessrepo_updateApplicationCmd.Flags().String("readme-url", "", "A link to the readme file in Markdown language that contains a more detailed description of the application and how it works.")
		serverlessrepo_updateApplicationCmd.MarkFlagRequired("application-id")
	})
	serverlessrepoCmd.AddCommand(serverlessrepo_updateApplicationCmd)
}
