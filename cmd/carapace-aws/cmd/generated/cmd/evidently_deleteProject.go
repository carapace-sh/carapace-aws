package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evidently_deleteProjectCmd = &cobra.Command{
	Use:   "delete-project",
	Short: "Deletes an Evidently project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evidently_deleteProjectCmd).Standalone()

	evidently_deleteProjectCmd.Flags().String("project", "", "The name or ARN of the project to delete.")
	evidently_deleteProjectCmd.MarkFlagRequired("project")
	evidentlyCmd.AddCommand(evidently_deleteProjectCmd)
}
