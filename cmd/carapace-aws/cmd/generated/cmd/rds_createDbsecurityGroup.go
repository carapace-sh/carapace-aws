package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_createDbsecurityGroupCmd = &cobra.Command{
	Use:   "create-dbsecurity-group",
	Short: "Creates a new DB security group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_createDbsecurityGroupCmd).Standalone()

	rds_createDbsecurityGroupCmd.Flags().String("dbsecurity-group-description", "", "The description for the DB security group.")
	rds_createDbsecurityGroupCmd.Flags().String("dbsecurity-group-name", "", "The name for the DB security group.")
	rds_createDbsecurityGroupCmd.Flags().String("tags", "", "Tags to assign to the DB security group.")
	rds_createDbsecurityGroupCmd.MarkFlagRequired("dbsecurity-group-description")
	rds_createDbsecurityGroupCmd.MarkFlagRequired("dbsecurity-group-name")
	rdsCmd.AddCommand(rds_createDbsecurityGroupCmd)
}
