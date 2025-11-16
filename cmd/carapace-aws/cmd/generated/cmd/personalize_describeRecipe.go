package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_describeRecipeCmd = &cobra.Command{
	Use:   "describe-recipe",
	Short: "Describes a recipe.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_describeRecipeCmd).Standalone()

	personalize_describeRecipeCmd.Flags().String("recipe-arn", "", "The Amazon Resource Name (ARN) of the recipe to describe.")
	personalize_describeRecipeCmd.MarkFlagRequired("recipe-arn")
	personalizeCmd.AddCommand(personalize_describeRecipeCmd)
}
