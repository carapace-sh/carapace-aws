package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_createKxDataviewCmd = &cobra.Command{
	Use:   "create-kx-dataview",
	Short: "Creates a snapshot of kdb database with tiered storage capabilities and a pre-warmed cache, ready for mounting on kdb clusters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_createKxDataviewCmd).Standalone()

	finspace_createKxDataviewCmd.Flags().String("auto-update", "", "The option to specify whether you want to apply all the future additions and corrections automatically to the dataview, when you ingest new changesets.")
	finspace_createKxDataviewCmd.Flags().String("availability-zone-id", "", "The identifier of the availability zones.")
	finspace_createKxDataviewCmd.Flags().String("az-mode", "", "The number of availability zones you want to assign per volume.")
	finspace_createKxDataviewCmd.Flags().String("changeset-id", "", "A unique identifier of the changeset that you want to use to ingest data.")
	finspace_createKxDataviewCmd.Flags().String("client-token", "", "A token that ensures idempotency.")
	finspace_createKxDataviewCmd.Flags().String("database-name", "", "The name of the database where you want to create a dataview.")
	finspace_createKxDataviewCmd.Flags().String("dataview-name", "", "A unique identifier for the dataview.")
	finspace_createKxDataviewCmd.Flags().String("description", "", "A description of the dataview.")
	finspace_createKxDataviewCmd.Flags().String("environment-id", "", "A unique identifier for the kdb environment, where you want to create the dataview.")
	finspace_createKxDataviewCmd.Flags().String("read-write", "", "The option to specify whether you want to make the dataview writable to perform database maintenance.")
	finspace_createKxDataviewCmd.Flags().String("segment-configurations", "", "The configuration that contains the database path of the data that you want to place on each selected volume.")
	finspace_createKxDataviewCmd.Flags().String("tags", "", "A list of key-value pairs to label the dataview.")
	finspace_createKxDataviewCmd.MarkFlagRequired("az-mode")
	finspace_createKxDataviewCmd.MarkFlagRequired("client-token")
	finspace_createKxDataviewCmd.MarkFlagRequired("database-name")
	finspace_createKxDataviewCmd.MarkFlagRequired("dataview-name")
	finspace_createKxDataviewCmd.MarkFlagRequired("environment-id")
	finspaceCmd.AddCommand(finspace_createKxDataviewCmd)
}
