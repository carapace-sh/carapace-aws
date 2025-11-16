package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_createSolutionCmd = &cobra.Command{
	Use:   "create-solution",
	Short: "By default, all new solutions use automatic training.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_createSolutionCmd).Standalone()

	personalize_createSolutionCmd.Flags().String("dataset-group-arn", "", "The Amazon Resource Name (ARN) of the dataset group that provides the training data.")
	personalize_createSolutionCmd.Flags().String("event-type", "", "When your have multiple event types (using an `EVENT_TYPE` schema field), this parameter specifies which event type (for example, 'click' or 'like') is used for training the model.")
	personalize_createSolutionCmd.Flags().String("name", "", "The name for the solution.")
	personalize_createSolutionCmd.Flags().Bool("no-perform-hpo", false, "Whether to perform hyperparameter optimization (HPO) on the specified or selected recipe.")
	personalize_createSolutionCmd.Flags().String("perform-auto-ml", "", "We don't recommend enabling automated machine learning.")
	personalize_createSolutionCmd.Flags().String("perform-auto-training", "", "Whether the solution uses automatic training to create new solution versions (trained models).")
	personalize_createSolutionCmd.Flags().Bool("perform-hpo", false, "Whether to perform hyperparameter optimization (HPO) on the specified or selected recipe.")
	personalize_createSolutionCmd.Flags().String("recipe-arn", "", "The Amazon Resource Name (ARN) of the recipe to use for model training.")
	personalize_createSolutionCmd.Flags().String("solution-config", "", "The configuration properties for the solution.")
	personalize_createSolutionCmd.Flags().String("tags", "", "A list of [tags](https://docs.aws.amazon.com/personalize/latest/dg/tagging-resources.html) to apply to the solution.")
	personalize_createSolutionCmd.MarkFlagRequired("dataset-group-arn")
	personalize_createSolutionCmd.MarkFlagRequired("name")
	personalize_createSolutionCmd.Flag("no-perform-hpo").Hidden = true
	personalizeCmd.AddCommand(personalize_createSolutionCmd)
}
