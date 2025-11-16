package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_updateExperienceCmd = &cobra.Command{
	Use:   "update-experience",
	Short: "Updates your Amazon Kendra experience such as a search application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_updateExperienceCmd).Standalone()

	kendra_updateExperienceCmd.Flags().String("configuration", "", "Configuration information you want to update for your Amazon Kendra experience.")
	kendra_updateExperienceCmd.Flags().String("description", "", "A new description for your Amazon Kendra experience.")
	kendra_updateExperienceCmd.Flags().String("id", "", "The identifier of your Amazon Kendra experience you want to update.")
	kendra_updateExperienceCmd.Flags().String("index-id", "", "The identifier of the index for your Amazon Kendra experience.")
	kendra_updateExperienceCmd.Flags().String("name", "", "A new name for your Amazon Kendra experience.")
	kendra_updateExperienceCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of an IAM role with permission to access the `Query` API, `QuerySuggestions` API, `SubmitFeedback` API, and IAM Identity Center that stores your users and groups information.")
	kendra_updateExperienceCmd.MarkFlagRequired("id")
	kendra_updateExperienceCmd.MarkFlagRequired("index-id")
	kendraCmd.AddCommand(kendra_updateExperienceCmd)
}
