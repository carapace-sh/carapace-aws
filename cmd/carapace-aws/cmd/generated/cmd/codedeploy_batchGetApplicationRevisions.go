package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_batchGetApplicationRevisionsCmd = &cobra.Command{
	Use:   "batch-get-application-revisions",
	Short: "Gets information about one or more application revisions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_batchGetApplicationRevisionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codedeploy_batchGetApplicationRevisionsCmd).Standalone()

		codedeploy_batchGetApplicationRevisionsCmd.Flags().String("application-name", "", "The name of an CodeDeploy application about which to get revision information.")
		codedeploy_batchGetApplicationRevisionsCmd.Flags().String("revisions", "", "An array of `RevisionLocation` objects that specify information to get about the application revisions, including type and location.")
		codedeploy_batchGetApplicationRevisionsCmd.MarkFlagRequired("application-name")
		codedeploy_batchGetApplicationRevisionsCmd.MarkFlagRequired("revisions")
	})
	codedeployCmd.AddCommand(codedeploy_batchGetApplicationRevisionsCmd)
}
