package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_deleteUsageProfileCmd = &cobra.Command{
	Use:   "delete-usage-profile",
	Short: "Deletes the Glue specified usage profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_deleteUsageProfileCmd).Standalone()

	glue_deleteUsageProfileCmd.Flags().String("name", "", "The name of the usage profile to delete.")
	glue_deleteUsageProfileCmd.MarkFlagRequired("name")
	glueCmd.AddCommand(glue_deleteUsageProfileCmd)
}
