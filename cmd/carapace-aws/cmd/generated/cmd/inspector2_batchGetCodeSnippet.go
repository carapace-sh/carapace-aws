package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_batchGetCodeSnippetCmd = &cobra.Command{
	Use:   "batch-get-code-snippet",
	Short: "Retrieves code snippets from findings that Amazon Inspector detected code vulnerabilities in.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_batchGetCodeSnippetCmd).Standalone()

	inspector2_batchGetCodeSnippetCmd.Flags().String("finding-arns", "", "An array of finding ARNs for the findings you want to retrieve code snippets from.")
	inspector2_batchGetCodeSnippetCmd.MarkFlagRequired("finding-arns")
	inspector2Cmd.AddCommand(inspector2_batchGetCodeSnippetCmd)
}
