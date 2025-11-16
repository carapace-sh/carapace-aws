package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_disassociateEntitiesFromExperienceCmd = &cobra.Command{
	Use:   "disassociate-entities-from-experience",
	Short: "Prevents users or groups in your IAM Identity Center identity source from accessing your Amazon Kendra experience.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_disassociateEntitiesFromExperienceCmd).Standalone()

	kendra_disassociateEntitiesFromExperienceCmd.Flags().String("entity-list", "", "Lists users or groups in your IAM Identity Center identity source.")
	kendra_disassociateEntitiesFromExperienceCmd.Flags().String("id", "", "The identifier of your Amazon Kendra experience.")
	kendra_disassociateEntitiesFromExperienceCmd.Flags().String("index-id", "", "The identifier of the index for your Amazon Kendra experience.")
	kendra_disassociateEntitiesFromExperienceCmd.MarkFlagRequired("entity-list")
	kendra_disassociateEntitiesFromExperienceCmd.MarkFlagRequired("id")
	kendra_disassociateEntitiesFromExperienceCmd.MarkFlagRequired("index-id")
	kendraCmd.AddCommand(kendra_disassociateEntitiesFromExperienceCmd)
}
