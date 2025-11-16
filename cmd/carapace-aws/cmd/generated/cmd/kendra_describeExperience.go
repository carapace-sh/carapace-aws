package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_describeExperienceCmd = &cobra.Command{
	Use:   "describe-experience",
	Short: "Gets information about your Amazon Kendra experience such as a search application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_describeExperienceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kendra_describeExperienceCmd).Standalone()

		kendra_describeExperienceCmd.Flags().String("id", "", "The identifier of your Amazon Kendra experience you want to get information on.")
		kendra_describeExperienceCmd.Flags().String("index-id", "", "The identifier of the index for your Amazon Kendra experience.")
		kendra_describeExperienceCmd.MarkFlagRequired("id")
		kendra_describeExperienceCmd.MarkFlagRequired("index-id")
	})
	kendraCmd.AddCommand(kendra_describeExperienceCmd)
}
