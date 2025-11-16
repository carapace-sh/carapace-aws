package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rolesanywhere_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Attaches tags to a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rolesanywhere_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rolesanywhere_tagResourceCmd).Standalone()

		rolesanywhere_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource.")
		rolesanywhere_tagResourceCmd.Flags().String("tags", "", "The tags to attach to the resource.")
		rolesanywhere_tagResourceCmd.MarkFlagRequired("resource-arn")
		rolesanywhere_tagResourceCmd.MarkFlagRequired("tags")
	})
	rolesanywhereCmd.AddCommand(rolesanywhere_tagResourceCmd)
}
