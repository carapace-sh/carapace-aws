package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workdocs_deleteLabelsCmd = &cobra.Command{
	Use:   "delete-labels",
	Short: "Deletes the specified list of labels from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workdocs_deleteLabelsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workdocs_deleteLabelsCmd).Standalone()

		workdocs_deleteLabelsCmd.Flags().String("authentication-token", "", "Amazon WorkDocs authentication token.")
		workdocs_deleteLabelsCmd.Flags().String("delete-all", "", "Flag to request removal of all labels from the specified resource.")
		workdocs_deleteLabelsCmd.Flags().String("labels", "", "List of labels to delete from the resource.")
		workdocs_deleteLabelsCmd.Flags().String("resource-id", "", "The ID of the resource.")
		workdocs_deleteLabelsCmd.MarkFlagRequired("resource-id")
	})
	workdocsCmd.AddCommand(workdocs_deleteLabelsCmd)
}
