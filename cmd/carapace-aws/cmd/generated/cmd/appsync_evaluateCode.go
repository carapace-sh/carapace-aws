package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_evaluateCodeCmd = &cobra.Command{
	Use:   "evaluate-code",
	Short: "Evaluates the given code and returns the response.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_evaluateCodeCmd).Standalone()

	appsync_evaluateCodeCmd.Flags().String("code", "", "The code definition to be evaluated.")
	appsync_evaluateCodeCmd.Flags().String("context", "", "The map that holds all of the contextual information for your resolver invocation.")
	appsync_evaluateCodeCmd.Flags().String("function", "", "The function within the code to be evaluated.")
	appsync_evaluateCodeCmd.Flags().String("runtime", "", "The runtime to be used when evaluating the code.")
	appsync_evaluateCodeCmd.MarkFlagRequired("code")
	appsync_evaluateCodeCmd.MarkFlagRequired("context")
	appsync_evaluateCodeCmd.MarkFlagRequired("runtime")
	appsyncCmd.AddCommand(appsync_evaluateCodeCmd)
}
