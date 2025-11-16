package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_listApplicationRevisionsCmd = &cobra.Command{
	Use:   "list-application-revisions",
	Short: "Lists information about revisions for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_listApplicationRevisionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codedeploy_listApplicationRevisionsCmd).Standalone()

		codedeploy_listApplicationRevisionsCmd.Flags().String("application-name", "", "The name of an CodeDeploy application associated with the user or Amazon Web Services account.")
		codedeploy_listApplicationRevisionsCmd.Flags().String("deployed", "", "Whether to list revisions based on whether the revision is the target revision of a deployment group:")
		codedeploy_listApplicationRevisionsCmd.Flags().String("next-token", "", "An identifier returned from the previous `ListApplicationRevisions` call.")
		codedeploy_listApplicationRevisionsCmd.Flags().String("s3-bucket", "", "An Amazon S3 bucket name to limit the search for revisions.")
		codedeploy_listApplicationRevisionsCmd.Flags().String("s3-key-prefix", "", "A key prefix for the set of Amazon S3 objects to limit the search for revisions.")
		codedeploy_listApplicationRevisionsCmd.Flags().String("sort-by", "", "The column name to use to sort the list results:")
		codedeploy_listApplicationRevisionsCmd.Flags().String("sort-order", "", "The order in which to sort the list results:")
		codedeploy_listApplicationRevisionsCmd.MarkFlagRequired("application-name")
	})
	codedeployCmd.AddCommand(codedeploy_listApplicationRevisionsCmd)
}
