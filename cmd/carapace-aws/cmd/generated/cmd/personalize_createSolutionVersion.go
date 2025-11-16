package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_createSolutionVersionCmd = &cobra.Command{
	Use:   "create-solution-version",
	Short: "Trains or retrains an active solution in a Custom dataset group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_createSolutionVersionCmd).Standalone()

	personalize_createSolutionVersionCmd.Flags().String("name", "", "The name of the solution version.")
	personalize_createSolutionVersionCmd.Flags().String("solution-arn", "", "The Amazon Resource Name (ARN) of the solution containing the training configuration information.")
	personalize_createSolutionVersionCmd.Flags().String("tags", "", "A list of [tags](https://docs.aws.amazon.com/personalize/latest/dg/tagging-resources.html) to apply to the solution version.")
	personalize_createSolutionVersionCmd.Flags().String("training-mode", "", "The scope of training to be performed when creating the solution version.")
	personalize_createSolutionVersionCmd.MarkFlagRequired("solution-arn")
	personalizeCmd.AddCommand(personalize_createSolutionVersionCmd)
}
