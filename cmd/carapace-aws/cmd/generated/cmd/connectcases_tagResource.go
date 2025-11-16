package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcases_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds tags to a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcases_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcases_tagResourceCmd).Standalone()

		connectcases_tagResourceCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN)")
		connectcases_tagResourceCmd.Flags().String("tags", "", "A map of of key-value pairs that represent tags on a resource.")
		connectcases_tagResourceCmd.MarkFlagRequired("arn")
		connectcases_tagResourceCmd.MarkFlagRequired("tags")
	})
	connectcasesCmd.AddCommand(connectcases_tagResourceCmd)
}
