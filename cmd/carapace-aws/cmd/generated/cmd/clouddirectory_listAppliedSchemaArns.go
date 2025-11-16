package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_listAppliedSchemaArnsCmd = &cobra.Command{
	Use:   "list-applied-schema-arns",
	Short: "Lists schema major versions applied to a directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_listAppliedSchemaArnsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(clouddirectory_listAppliedSchemaArnsCmd).Standalone()

		clouddirectory_listAppliedSchemaArnsCmd.Flags().String("directory-arn", "", "The ARN of the directory you are listing.")
		clouddirectory_listAppliedSchemaArnsCmd.Flags().String("max-results", "", "The maximum number of results to retrieve.")
		clouddirectory_listAppliedSchemaArnsCmd.Flags().String("next-token", "", "The pagination token.")
		clouddirectory_listAppliedSchemaArnsCmd.Flags().String("schema-arn", "", "The response for `ListAppliedSchemaArns` when this parameter is used will list all minor version ARNs for a major version.")
		clouddirectory_listAppliedSchemaArnsCmd.MarkFlagRequired("directory-arn")
	})
	clouddirectoryCmd.AddCommand(clouddirectory_listAppliedSchemaArnsCmd)
}
