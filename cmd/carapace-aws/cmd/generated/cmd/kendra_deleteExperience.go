package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_deleteExperienceCmd = &cobra.Command{
	Use:   "delete-experience",
	Short: "Deletes your Amazon Kendra experience such as a search application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_deleteExperienceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kendra_deleteExperienceCmd).Standalone()

		kendra_deleteExperienceCmd.Flags().String("id", "", "The identifier of your Amazon Kendra experience you want to delete.")
		kendra_deleteExperienceCmd.Flags().String("index-id", "", "The identifier of the index for your Amazon Kendra experience.")
		kendra_deleteExperienceCmd.MarkFlagRequired("id")
		kendra_deleteExperienceCmd.MarkFlagRequired("index-id")
	})
	kendraCmd.AddCommand(kendra_deleteExperienceCmd)
}
