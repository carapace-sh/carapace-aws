package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_createOptionGroupCmd = &cobra.Command{
	Use:   "create-option-group",
	Short: "Creates a new option group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_createOptionGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_createOptionGroupCmd).Standalone()

		rds_createOptionGroupCmd.Flags().String("engine-name", "", "The name of the engine to associate this option group with.")
		rds_createOptionGroupCmd.Flags().String("major-engine-version", "", "Specifies the major version of the engine that this option group should be associated with.")
		rds_createOptionGroupCmd.Flags().String("option-group-description", "", "The description of the option group.")
		rds_createOptionGroupCmd.Flags().String("option-group-name", "", "Specifies the name of the option group to be created.")
		rds_createOptionGroupCmd.Flags().String("tags", "", "Tags to assign to the option group.")
		rds_createOptionGroupCmd.MarkFlagRequired("engine-name")
		rds_createOptionGroupCmd.MarkFlagRequired("major-engine-version")
		rds_createOptionGroupCmd.MarkFlagRequired("option-group-description")
		rds_createOptionGroupCmd.MarkFlagRequired("option-group-name")
	})
	rdsCmd.AddCommand(rds_createOptionGroupCmd)
}
