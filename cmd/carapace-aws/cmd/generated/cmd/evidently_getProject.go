package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evidently_getProjectCmd = &cobra.Command{
	Use:   "get-project",
	Short: "Returns the details about one launch.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evidently_getProjectCmd).Standalone()

	evidently_getProjectCmd.Flags().String("project", "", "The name or ARN of the project that you want to see the details of.")
	evidently_getProjectCmd.MarkFlagRequired("project")
	evidentlyCmd.AddCommand(evidently_getProjectCmd)
}
