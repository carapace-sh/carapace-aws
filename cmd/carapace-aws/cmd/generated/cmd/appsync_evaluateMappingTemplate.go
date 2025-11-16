package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_evaluateMappingTemplateCmd = &cobra.Command{
	Use:   "evaluate-mapping-template",
	Short: "Evaluates a given template and returns the response.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_evaluateMappingTemplateCmd).Standalone()

	appsync_evaluateMappingTemplateCmd.Flags().String("context", "", "The map that holds all of the contextual information for your resolver invocation.")
	appsync_evaluateMappingTemplateCmd.Flags().String("template", "", "The mapping template; this can be a request or response template.")
	appsync_evaluateMappingTemplateCmd.MarkFlagRequired("context")
	appsync_evaluateMappingTemplateCmd.MarkFlagRequired("template")
	appsyncCmd.AddCommand(appsync_evaluateMappingTemplateCmd)
}
