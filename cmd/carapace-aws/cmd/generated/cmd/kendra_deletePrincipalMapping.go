package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_deletePrincipalMappingCmd = &cobra.Command{
	Use:   "delete-principal-mapping",
	Short: "Deletes a group so that all users that belong to the group can no longer access documents only available to that group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_deletePrincipalMappingCmd).Standalone()

	kendra_deletePrincipalMappingCmd.Flags().String("data-source-id", "", "The identifier of the data source you want to delete a group from.")
	kendra_deletePrincipalMappingCmd.Flags().String("group-id", "", "The identifier of the group you want to delete.")
	kendra_deletePrincipalMappingCmd.Flags().String("index-id", "", "The identifier of the index you want to delete a group from.")
	kendra_deletePrincipalMappingCmd.Flags().String("ordering-id", "", "The timestamp identifier you specify to ensure Amazon Kendra does not override the latest `DELETE` action with previous actions.")
	kendra_deletePrincipalMappingCmd.MarkFlagRequired("group-id")
	kendra_deletePrincipalMappingCmd.MarkFlagRequired("index-id")
	kendraCmd.AddCommand(kendra_deletePrincipalMappingCmd)
}
