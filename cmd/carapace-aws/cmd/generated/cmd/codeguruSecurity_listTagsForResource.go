package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruSecurity_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns a list of all tags associated with a scan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruSecurity_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeguruSecurity_listTagsForResourceCmd).Standalone()

		codeguruSecurity_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the `ScanName` object.")
		codeguruSecurity_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	codeguruSecurityCmd.AddCommand(codeguruSecurity_listTagsForResourceCmd)
}
