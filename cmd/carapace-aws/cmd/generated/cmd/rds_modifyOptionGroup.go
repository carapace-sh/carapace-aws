package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_modifyOptionGroupCmd = &cobra.Command{
	Use:   "modify-option-group",
	Short: "Modifies an existing option group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_modifyOptionGroupCmd).Standalone()

	rds_modifyOptionGroupCmd.Flags().Bool("apply-immediately", false, "Specifies whether to apply the change immediately or during the next maintenance window for each instance associated with the option group.")
	rds_modifyOptionGroupCmd.Flags().Bool("no-apply-immediately", false, "Specifies whether to apply the change immediately or during the next maintenance window for each instance associated with the option group.")
	rds_modifyOptionGroupCmd.Flags().String("option-group-name", "", "The name of the option group to be modified.")
	rds_modifyOptionGroupCmd.Flags().String("options-to-include", "", "Options in this list are added to the option group or, if already present, the specified configuration is used to update the existing configuration.")
	rds_modifyOptionGroupCmd.Flags().String("options-to-remove", "", "Options in this list are removed from the option group.")
	rds_modifyOptionGroupCmd.Flag("no-apply-immediately").Hidden = true
	rds_modifyOptionGroupCmd.MarkFlagRequired("option-group-name")
	rdsCmd.AddCommand(rds_modifyOptionGroupCmd)
}
