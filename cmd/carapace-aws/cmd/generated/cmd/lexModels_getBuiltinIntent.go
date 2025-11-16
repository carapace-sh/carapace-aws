package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexModels_getBuiltinIntentCmd = &cobra.Command{
	Use:   "get-builtin-intent",
	Short: "Returns information about a built-in intent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexModels_getBuiltinIntentCmd).Standalone()

	lexModels_getBuiltinIntentCmd.Flags().String("signature", "", "The unique identifier for a built-in intent.")
	lexModels_getBuiltinIntentCmd.MarkFlagRequired("signature")
	lexModelsCmd.AddCommand(lexModels_getBuiltinIntentCmd)
}
