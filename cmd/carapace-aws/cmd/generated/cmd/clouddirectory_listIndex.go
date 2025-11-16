package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_listIndexCmd = &cobra.Command{
	Use:   "list-index",
	Short: "Lists objects attached to the specified index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_listIndexCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(clouddirectory_listIndexCmd).Standalone()

		clouddirectory_listIndexCmd.Flags().String("consistency-level", "", "The consistency level to execute the request at.")
		clouddirectory_listIndexCmd.Flags().String("directory-arn", "", "The ARN of the directory that the index exists in.")
		clouddirectory_listIndexCmd.Flags().String("index-reference", "", "The reference to the index to list.")
		clouddirectory_listIndexCmd.Flags().String("max-results", "", "The maximum number of objects in a single page to retrieve from the index during a request.")
		clouddirectory_listIndexCmd.Flags().String("next-token", "", "The pagination token.")
		clouddirectory_listIndexCmd.Flags().String("ranges-on-indexed-values", "", "Specifies the ranges of indexed values that you want to query.")
		clouddirectory_listIndexCmd.MarkFlagRequired("directory-arn")
		clouddirectory_listIndexCmd.MarkFlagRequired("index-reference")
	})
	clouddirectoryCmd.AddCommand(clouddirectory_listIndexCmd)
}
