package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transcribe_createCallAnalyticsCategoryCmd = &cobra.Command{
	Use:   "create-call-analytics-category",
	Short: "Creates a new Call Analytics category.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transcribe_createCallAnalyticsCategoryCmd).Standalone()

	transcribe_createCallAnalyticsCategoryCmd.Flags().String("category-name", "", "A unique name, chosen by you, for your Call Analytics category.")
	transcribe_createCallAnalyticsCategoryCmd.Flags().String("input-type", "", "Choose whether you want to create a real-time or a post-call category for your Call Analytics transcription.")
	transcribe_createCallAnalyticsCategoryCmd.Flags().String("rules", "", "Rules define a Call Analytics category.")
	transcribe_createCallAnalyticsCategoryCmd.Flags().String("tags", "", "Adds one or more custom tags, each in the form of a key:value pair, to a new call analytics category at the time you start this new job.")
	transcribe_createCallAnalyticsCategoryCmd.MarkFlagRequired("category-name")
	transcribe_createCallAnalyticsCategoryCmd.MarkFlagRequired("rules")
	transcribeCmd.AddCommand(transcribe_createCallAnalyticsCategoryCmd)
}
