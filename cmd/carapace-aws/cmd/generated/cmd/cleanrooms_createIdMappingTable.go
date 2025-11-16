package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_createIdMappingTableCmd = &cobra.Command{
	Use:   "create-id-mapping-table",
	Short: "Creates an ID mapping table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_createIdMappingTableCmd).Standalone()

	cleanrooms_createIdMappingTableCmd.Flags().String("description", "", "A description of the ID mapping table.")
	cleanrooms_createIdMappingTableCmd.Flags().String("input-reference-config", "", "The input reference configuration needed to create the ID mapping table.")
	cleanrooms_createIdMappingTableCmd.Flags().String("kms-key-arn", "", "The Amazon Resource Name (ARN) of the Amazon Web Services KMS key.")
	cleanrooms_createIdMappingTableCmd.Flags().String("membership-identifier", "", "The unique identifier of the membership that contains the ID mapping table.")
	cleanrooms_createIdMappingTableCmd.Flags().String("name", "", "A name for the ID mapping table.")
	cleanrooms_createIdMappingTableCmd.Flags().String("tags", "", "An optional label that you can assign to a resource when you create it.")
	cleanrooms_createIdMappingTableCmd.MarkFlagRequired("input-reference-config")
	cleanrooms_createIdMappingTableCmd.MarkFlagRequired("membership-identifier")
	cleanrooms_createIdMappingTableCmd.MarkFlagRequired("name")
	cleanroomsCmd.AddCommand(cleanrooms_createIdMappingTableCmd)
}
