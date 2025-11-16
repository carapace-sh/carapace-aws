package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_modifyCustomDbengineVersionCmd = &cobra.Command{
	Use:   "modify-custom-dbengine-version",
	Short: "Modifies the status of a custom engine version (CEV).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_modifyCustomDbengineVersionCmd).Standalone()

	rds_modifyCustomDbengineVersionCmd.Flags().String("description", "", "An optional description of your CEV.")
	rds_modifyCustomDbengineVersionCmd.Flags().String("engine", "", "The database engine.")
	rds_modifyCustomDbengineVersionCmd.Flags().String("engine-version", "", "The custom engine version (CEV) that you want to modify.")
	rds_modifyCustomDbengineVersionCmd.Flags().String("status", "", "The availability status to be assigned to the CEV.")
	rds_modifyCustomDbengineVersionCmd.MarkFlagRequired("engine")
	rds_modifyCustomDbengineVersionCmd.MarkFlagRequired("engine-version")
	rdsCmd.AddCommand(rds_modifyCustomDbengineVersionCmd)
}
