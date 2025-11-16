package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcases_createRelatedItemCmd = &cobra.Command{
	Use:   "create-related-item",
	Short: "Creates a related item (comments, tasks, and contacts) and associates it with a case.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcases_createRelatedItemCmd).Standalone()

	connectcases_createRelatedItemCmd.Flags().String("case-id", "", "A unique identifier of the case.")
	connectcases_createRelatedItemCmd.Flags().String("content", "", "The content of a related item to be created.")
	connectcases_createRelatedItemCmd.Flags().String("domain-id", "", "The unique identifier of the Cases domain.")
	connectcases_createRelatedItemCmd.Flags().String("performed-by", "", "Represents the creator of the related item.")
	connectcases_createRelatedItemCmd.Flags().String("type", "", "The type of a related item.")
	connectcases_createRelatedItemCmd.MarkFlagRequired("case-id")
	connectcases_createRelatedItemCmd.MarkFlagRequired("content")
	connectcases_createRelatedItemCmd.MarkFlagRequired("domain-id")
	connectcases_createRelatedItemCmd.MarkFlagRequired("type")
	connectcasesCmd.AddCommand(connectcases_createRelatedItemCmd)
}
