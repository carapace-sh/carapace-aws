package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_stopSolutionVersionCreationCmd = &cobra.Command{
	Use:   "stop-solution-version-creation",
	Short: "Stops creating a solution version that is in a state of CREATE\\_PENDING or CREATE IN\\_PROGRESS.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_stopSolutionVersionCreationCmd).Standalone()

	personalize_stopSolutionVersionCreationCmd.Flags().String("solution-version-arn", "", "The Amazon Resource Name (ARN) of the solution version you want to stop creating.")
	personalize_stopSolutionVersionCreationCmd.MarkFlagRequired("solution-version-arn")
	personalizeCmd.AddCommand(personalize_stopSolutionVersionCreationCmd)
}
