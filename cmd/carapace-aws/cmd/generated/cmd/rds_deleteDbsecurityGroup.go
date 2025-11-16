package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_deleteDbsecurityGroupCmd = &cobra.Command{
	Use:   "delete-dbsecurity-group",
	Short: "Deletes a DB security group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_deleteDbsecurityGroupCmd).Standalone()

	rds_deleteDbsecurityGroupCmd.Flags().String("dbsecurity-group-name", "", "The name of the DB security group to delete.")
	rds_deleteDbsecurityGroupCmd.MarkFlagRequired("dbsecurity-group-name")
	rdsCmd.AddCommand(rds_deleteDbsecurityGroupCmd)
}
