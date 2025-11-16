package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Tag a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_tagResourceCmd).Standalone()

		proton_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the Proton resource to apply customer tags to.")
		proton_tagResourceCmd.Flags().String("tags", "", "A list of customer tags to apply to the Proton resource.")
		proton_tagResourceCmd.MarkFlagRequired("resource-arn")
		proton_tagResourceCmd.MarkFlagRequired("tags")
	})
	protonCmd.AddCommand(proton_tagResourceCmd)
}
