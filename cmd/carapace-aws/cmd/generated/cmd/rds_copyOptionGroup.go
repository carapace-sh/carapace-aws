package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_copyOptionGroupCmd = &cobra.Command{
	Use:   "copy-option-group",
	Short: "Copies the specified option group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_copyOptionGroupCmd).Standalone()

	rds_copyOptionGroupCmd.Flags().String("source-option-group-identifier", "", "The identifier for the source option group.")
	rds_copyOptionGroupCmd.Flags().String("tags", "", "")
	rds_copyOptionGroupCmd.Flags().String("target-option-group-description", "", "The description for the copied option group.")
	rds_copyOptionGroupCmd.Flags().String("target-option-group-identifier", "", "The identifier for the copied option group.")
	rds_copyOptionGroupCmd.MarkFlagRequired("source-option-group-identifier")
	rds_copyOptionGroupCmd.MarkFlagRequired("target-option-group-description")
	rds_copyOptionGroupCmd.MarkFlagRequired("target-option-group-identifier")
	rdsCmd.AddCommand(rds_copyOptionGroupCmd)
}
