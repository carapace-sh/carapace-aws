package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_describePrincipalMappingCmd = &cobra.Command{
	Use:   "describe-principal-mapping",
	Short: "Describes the processing of `PUT` and `DELETE` actions for mapping users to their groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_describePrincipalMappingCmd).Standalone()

	kendra_describePrincipalMappingCmd.Flags().String("data-source-id", "", "The identifier of the data source to check the processing of `PUT` and `DELETE` actions for mapping users to their groups.")
	kendra_describePrincipalMappingCmd.Flags().String("group-id", "", "The identifier of the group required to check the processing of `PUT` and `DELETE` actions for mapping users to their groups.")
	kendra_describePrincipalMappingCmd.Flags().String("index-id", "", "The identifier of the index required to check the processing of `PUT` and `DELETE` actions for mapping users to their groups.")
	kendra_describePrincipalMappingCmd.MarkFlagRequired("group-id")
	kendra_describePrincipalMappingCmd.MarkFlagRequired("index-id")
	kendraCmd.AddCommand(kendra_describePrincipalMappingCmd)
}
