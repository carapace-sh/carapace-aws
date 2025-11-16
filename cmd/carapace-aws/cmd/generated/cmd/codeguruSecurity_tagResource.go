package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruSecurity_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Use to add one or more tags to an existing scan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruSecurity_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeguruSecurity_tagResourceCmd).Standalone()

		codeguruSecurity_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the `ScanName` object.")
		codeguruSecurity_tagResourceCmd.Flags().String("tags", "", "An array of key-value pairs used to tag an existing scan.")
		codeguruSecurity_tagResourceCmd.MarkFlagRequired("resource-arn")
		codeguruSecurity_tagResourceCmd.MarkFlagRequired("tags")
	})
	codeguruSecurityCmd.AddCommand(codeguruSecurity_tagResourceCmd)
}
