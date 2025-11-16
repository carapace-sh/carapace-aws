package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_createConfiguredTableAssociationCmd = &cobra.Command{
	Use:   "create-configured-table-association",
	Short: "Creates a configured table association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_createConfiguredTableAssociationCmd).Standalone()

	cleanrooms_createConfiguredTableAssociationCmd.Flags().String("configured-table-identifier", "", "A unique identifier for the configured table to be associated to.")
	cleanrooms_createConfiguredTableAssociationCmd.Flags().String("description", "", "A description for the configured table association.")
	cleanrooms_createConfiguredTableAssociationCmd.Flags().String("membership-identifier", "", "A unique identifier for one of your memberships for a collaboration.")
	cleanrooms_createConfiguredTableAssociationCmd.Flags().String("name", "", "The name of the configured table association.")
	cleanrooms_createConfiguredTableAssociationCmd.Flags().String("role-arn", "", "The service will assume this role to access catalog metadata and query the table.")
	cleanrooms_createConfiguredTableAssociationCmd.Flags().String("tags", "", "An optional label that you can assign to a resource when you create it.")
	cleanrooms_createConfiguredTableAssociationCmd.MarkFlagRequired("configured-table-identifier")
	cleanrooms_createConfiguredTableAssociationCmd.MarkFlagRequired("membership-identifier")
	cleanrooms_createConfiguredTableAssociationCmd.MarkFlagRequired("name")
	cleanrooms_createConfiguredTableAssociationCmd.MarkFlagRequired("role-arn")
	cleanroomsCmd.AddCommand(cleanrooms_createConfiguredTableAssociationCmd)
}
