package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emrContainers_describeJobTemplateCmd = &cobra.Command{
	Use:   "describe-job-template",
	Short: "Displays detailed information about a specified job template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emrContainers_describeJobTemplateCmd).Standalone()

	emrContainers_describeJobTemplateCmd.Flags().String("id", "", "The ID of the job template that will be described.")
	emrContainers_describeJobTemplateCmd.MarkFlagRequired("id")
	emrContainersCmd.AddCommand(emrContainers_describeJobTemplateCmd)
}
