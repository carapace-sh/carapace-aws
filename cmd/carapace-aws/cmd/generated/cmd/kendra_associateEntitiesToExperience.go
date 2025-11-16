package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_associateEntitiesToExperienceCmd = &cobra.Command{
	Use:   "associate-entities-to-experience",
	Short: "Grants users or groups in your IAM Identity Center identity source access to your Amazon Kendra experience.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_associateEntitiesToExperienceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kendra_associateEntitiesToExperienceCmd).Standalone()

		kendra_associateEntitiesToExperienceCmd.Flags().String("entity-list", "", "Lists users or groups in your IAM Identity Center identity source.")
		kendra_associateEntitiesToExperienceCmd.Flags().String("id", "", "The identifier of your Amazon Kendra experience.")
		kendra_associateEntitiesToExperienceCmd.Flags().String("index-id", "", "The identifier of the index for your Amazon Kendra experience.")
		kendra_associateEntitiesToExperienceCmd.MarkFlagRequired("entity-list")
		kendra_associateEntitiesToExperienceCmd.MarkFlagRequired("id")
		kendra_associateEntitiesToExperienceCmd.MarkFlagRequired("index-id")
	})
	kendraCmd.AddCommand(kendra_associateEntitiesToExperienceCmd)
}
