package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_deleteOptionGroupCmd = &cobra.Command{
	Use:   "delete-option-group",
	Short: "Deletes an existing option group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_deleteOptionGroupCmd).Standalone()

	rds_deleteOptionGroupCmd.Flags().String("option-group-name", "", "The name of the option group to be deleted.")
	rds_deleteOptionGroupCmd.MarkFlagRequired("option-group-name")
	rdsCmd.AddCommand(rds_deleteOptionGroupCmd)
}
