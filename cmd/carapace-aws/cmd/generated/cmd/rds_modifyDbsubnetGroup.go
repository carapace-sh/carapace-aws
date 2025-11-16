package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_modifyDbsubnetGroupCmd = &cobra.Command{
	Use:   "modify-dbsubnet-group",
	Short: "Modifies an existing DB subnet group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_modifyDbsubnetGroupCmd).Standalone()

	rds_modifyDbsubnetGroupCmd.Flags().String("dbsubnet-group-description", "", "The description for the DB subnet group.")
	rds_modifyDbsubnetGroupCmd.Flags().String("dbsubnet-group-name", "", "The name for the DB subnet group.")
	rds_modifyDbsubnetGroupCmd.Flags().String("subnet-ids", "", "The EC2 subnet IDs for the DB subnet group.")
	rds_modifyDbsubnetGroupCmd.MarkFlagRequired("dbsubnet-group-name")
	rds_modifyDbsubnetGroupCmd.MarkFlagRequired("subnet-ids")
	rdsCmd.AddCommand(rds_modifyDbsubnetGroupCmd)
}
