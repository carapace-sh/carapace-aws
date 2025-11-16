package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_batchGetApplicationsCmd = &cobra.Command{
	Use:   "batch-get-applications",
	Short: "Gets information about one or more applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_batchGetApplicationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codedeploy_batchGetApplicationsCmd).Standalone()

		codedeploy_batchGetApplicationsCmd.Flags().String("application-names", "", "A list of application names separated by spaces.")
		codedeploy_batchGetApplicationsCmd.MarkFlagRequired("application-names")
	})
	codedeployCmd.AddCommand(codedeploy_batchGetApplicationsCmd)
}
