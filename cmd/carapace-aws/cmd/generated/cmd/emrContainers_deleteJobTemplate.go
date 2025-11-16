package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emrContainers_deleteJobTemplateCmd = &cobra.Command{
	Use:   "delete-job-template",
	Short: "Deletes a job template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emrContainers_deleteJobTemplateCmd).Standalone()

	emrContainers_deleteJobTemplateCmd.Flags().String("id", "", "The ID of the job template that will be deleted.")
	emrContainers_deleteJobTemplateCmd.MarkFlagRequired("id")
	emrContainersCmd.AddCommand(emrContainers_deleteJobTemplateCmd)
}
