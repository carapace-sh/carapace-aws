package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_createDbsubnetGroupCmd = &cobra.Command{
	Use:   "create-dbsubnet-group",
	Short: "Creates a new DB subnet group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_createDbsubnetGroupCmd).Standalone()

	neptune_createDbsubnetGroupCmd.Flags().String("dbsubnet-group-description", "", "The description for the DB subnet group.")
	neptune_createDbsubnetGroupCmd.Flags().String("dbsubnet-group-name", "", "The name for the DB subnet group.")
	neptune_createDbsubnetGroupCmd.Flags().String("subnet-ids", "", "The EC2 Subnet IDs for the DB subnet group.")
	neptune_createDbsubnetGroupCmd.Flags().String("tags", "", "The tags to be assigned to the new DB subnet group.")
	neptune_createDbsubnetGroupCmd.MarkFlagRequired("dbsubnet-group-description")
	neptune_createDbsubnetGroupCmd.MarkFlagRequired("dbsubnet-group-name")
	neptune_createDbsubnetGroupCmd.MarkFlagRequired("subnet-ids")
	neptuneCmd.AddCommand(neptune_createDbsubnetGroupCmd)
}
