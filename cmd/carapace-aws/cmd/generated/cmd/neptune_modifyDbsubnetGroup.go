package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_modifyDbsubnetGroupCmd = &cobra.Command{
	Use:   "modify-dbsubnet-group",
	Short: "Modifies an existing DB subnet group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_modifyDbsubnetGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptune_modifyDbsubnetGroupCmd).Standalone()

		neptune_modifyDbsubnetGroupCmd.Flags().String("dbsubnet-group-description", "", "The description for the DB subnet group.")
		neptune_modifyDbsubnetGroupCmd.Flags().String("dbsubnet-group-name", "", "The name for the DB subnet group.")
		neptune_modifyDbsubnetGroupCmd.Flags().String("subnet-ids", "", "The EC2 subnet IDs for the DB subnet group.")
		neptune_modifyDbsubnetGroupCmd.MarkFlagRequired("dbsubnet-group-name")
		neptune_modifyDbsubnetGroupCmd.MarkFlagRequired("subnet-ids")
	})
	neptuneCmd.AddCommand(neptune_modifyDbsubnetGroupCmd)
}
