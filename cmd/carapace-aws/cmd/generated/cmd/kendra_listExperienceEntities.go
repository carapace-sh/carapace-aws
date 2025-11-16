package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_listExperienceEntitiesCmd = &cobra.Command{
	Use:   "list-experience-entities",
	Short: "Lists users or groups in your IAM Identity Center identity source that are granted access to your Amazon Kendra experience.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_listExperienceEntitiesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kendra_listExperienceEntitiesCmd).Standalone()

		kendra_listExperienceEntitiesCmd.Flags().String("id", "", "The identifier of your Amazon Kendra experience.")
		kendra_listExperienceEntitiesCmd.Flags().String("index-id", "", "The identifier of the index for your Amazon Kendra experience.")
		kendra_listExperienceEntitiesCmd.Flags().String("next-token", "", "If the previous response was incomplete (because there is more data to retrieve), Amazon Kendra returns a pagination token in the response.")
		kendra_listExperienceEntitiesCmd.MarkFlagRequired("id")
		kendra_listExperienceEntitiesCmd.MarkFlagRequired("index-id")
	})
	kendraCmd.AddCommand(kendra_listExperienceEntitiesCmd)
}
