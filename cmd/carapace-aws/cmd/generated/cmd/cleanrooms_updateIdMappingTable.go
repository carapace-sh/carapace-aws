package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_updateIdMappingTableCmd = &cobra.Command{
	Use:   "update-id-mapping-table",
	Short: "Provides the details that are necessary to update an ID mapping table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_updateIdMappingTableCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_updateIdMappingTableCmd).Standalone()

		cleanrooms_updateIdMappingTableCmd.Flags().String("description", "", "A new description for the ID mapping table.")
		cleanrooms_updateIdMappingTableCmd.Flags().String("id-mapping-table-identifier", "", "The unique identifier of the ID mapping table that you want to update.")
		cleanrooms_updateIdMappingTableCmd.Flags().String("kms-key-arn", "", "The Amazon Resource Name (ARN) of the Amazon Web Services KMS key.")
		cleanrooms_updateIdMappingTableCmd.Flags().String("membership-identifier", "", "The unique identifier of the membership that contains the ID mapping table that you want to update.")
		cleanrooms_updateIdMappingTableCmd.MarkFlagRequired("id-mapping-table-identifier")
		cleanrooms_updateIdMappingTableCmd.MarkFlagRequired("membership-identifier")
	})
	cleanroomsCmd.AddCommand(cleanrooms_updateIdMappingTableCmd)
}
