package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var savingsplans_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds the specified tags to the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(savingsplans_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(savingsplans_tagResourceCmd).Standalone()

		savingsplans_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
		savingsplans_tagResourceCmd.Flags().String("tags", "", "One or more tags.")
		savingsplans_tagResourceCmd.MarkFlagRequired("resource-arn")
		savingsplans_tagResourceCmd.MarkFlagRequired("tags")
	})
	savingsplansCmd.AddCommand(savingsplans_tagResourceCmd)
}
