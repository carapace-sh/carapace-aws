package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruSecurity_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Use to remove one or more tags from an existing scan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruSecurity_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeguruSecurity_untagResourceCmd).Standalone()

		codeguruSecurity_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the `ScanName` object.")
		codeguruSecurity_untagResourceCmd.Flags().String("tag-keys", "", "A list of keys for each tag you want to remove from a scan.")
		codeguruSecurity_untagResourceCmd.MarkFlagRequired("resource-arn")
		codeguruSecurity_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	codeguruSecurityCmd.AddCommand(codeguruSecurity_untagResourceCmd)
}
