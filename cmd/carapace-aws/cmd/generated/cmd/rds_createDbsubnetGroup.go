package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_createDbsubnetGroupCmd = &cobra.Command{
	Use:   "create-dbsubnet-group",
	Short: "Creates a new DB subnet group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_createDbsubnetGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_createDbsubnetGroupCmd).Standalone()

		rds_createDbsubnetGroupCmd.Flags().String("dbsubnet-group-description", "", "The description for the DB subnet group.")
		rds_createDbsubnetGroupCmd.Flags().String("dbsubnet-group-name", "", "The name for the DB subnet group.")
		rds_createDbsubnetGroupCmd.Flags().String("subnet-ids", "", "The EC2 Subnet IDs for the DB subnet group.")
		rds_createDbsubnetGroupCmd.Flags().String("tags", "", "Tags to assign to the DB subnet group.")
		rds_createDbsubnetGroupCmd.MarkFlagRequired("dbsubnet-group-description")
		rds_createDbsubnetGroupCmd.MarkFlagRequired("dbsubnet-group-name")
		rds_createDbsubnetGroupCmd.MarkFlagRequired("subnet-ids")
	})
	rdsCmd.AddCommand(rds_createDbsubnetGroupCmd)
}
