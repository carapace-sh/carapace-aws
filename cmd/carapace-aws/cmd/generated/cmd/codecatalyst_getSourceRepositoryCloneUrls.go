package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecatalyst_getSourceRepositoryCloneUrlsCmd = &cobra.Command{
	Use:   "get-source-repository-clone-urls",
	Short: "Returns information about the URLs that can be used with a Git client to clone a source repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecatalyst_getSourceRepositoryCloneUrlsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecatalyst_getSourceRepositoryCloneUrlsCmd).Standalone()

		codecatalyst_getSourceRepositoryCloneUrlsCmd.Flags().String("project-name", "", "The name of the project in the space.")
		codecatalyst_getSourceRepositoryCloneUrlsCmd.Flags().String("source-repository-name", "", "The name of the source repository.")
		codecatalyst_getSourceRepositoryCloneUrlsCmd.Flags().String("space-name", "", "The name of the space.")
		codecatalyst_getSourceRepositoryCloneUrlsCmd.MarkFlagRequired("project-name")
		codecatalyst_getSourceRepositoryCloneUrlsCmd.MarkFlagRequired("source-repository-name")
		codecatalyst_getSourceRepositoryCloneUrlsCmd.MarkFlagRequired("space-name")
	})
	codecatalystCmd.AddCommand(codecatalyst_getSourceRepositoryCloneUrlsCmd)
}
