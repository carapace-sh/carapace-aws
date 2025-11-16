package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workdocs_createLabelsCmd = &cobra.Command{
	Use:   "create-labels",
	Short: "Adds the specified list of labels to the given resource (a document or folder)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workdocs_createLabelsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workdocs_createLabelsCmd).Standalone()

		workdocs_createLabelsCmd.Flags().String("authentication-token", "", "Amazon WorkDocs authentication token.")
		workdocs_createLabelsCmd.Flags().String("labels", "", "List of labels to add to the resource.")
		workdocs_createLabelsCmd.Flags().String("resource-id", "", "The ID of the resource.")
		workdocs_createLabelsCmd.MarkFlagRequired("labels")
		workdocs_createLabelsCmd.MarkFlagRequired("resource-id")
	})
	workdocsCmd.AddCommand(workdocs_createLabelsCmd)
}
