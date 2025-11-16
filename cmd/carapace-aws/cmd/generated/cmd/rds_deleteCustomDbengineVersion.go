package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_deleteCustomDbengineVersionCmd = &cobra.Command{
	Use:   "delete-custom-dbengine-version",
	Short: "Deletes a custom engine version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_deleteCustomDbengineVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_deleteCustomDbengineVersionCmd).Standalone()

		rds_deleteCustomDbengineVersionCmd.Flags().String("engine", "", "The database engine.")
		rds_deleteCustomDbengineVersionCmd.Flags().String("engine-version", "", "The custom engine version (CEV) for your DB instance.")
		rds_deleteCustomDbengineVersionCmd.MarkFlagRequired("engine")
		rds_deleteCustomDbengineVersionCmd.MarkFlagRequired("engine-version")
	})
	rdsCmd.AddCommand(rds_deleteCustomDbengineVersionCmd)
}
