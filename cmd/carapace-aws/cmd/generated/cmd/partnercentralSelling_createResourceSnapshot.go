package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var partnercentralSelling_createResourceSnapshotCmd = &cobra.Command{
	Use:   "create-resource-snapshot",
	Short: "This action allows you to create an immutable snapshot of a specific resource, such as an opportunity, within the context of an engagement.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(partnercentralSelling_createResourceSnapshotCmd).Standalone()

	partnercentralSelling_createResourceSnapshotCmd.Flags().String("catalog", "", "Specifies the catalog where the snapshot is created.")
	partnercentralSelling_createResourceSnapshotCmd.Flags().String("client-token", "", "Specifies a unique, client-generated UUID to ensure that the request is handled exactly once.")
	partnercentralSelling_createResourceSnapshotCmd.Flags().String("engagement-identifier", "", "The unique identifier of the engagement associated with this snapshot.")
	partnercentralSelling_createResourceSnapshotCmd.Flags().String("resource-identifier", "", "The unique identifier of the specific resource to be snapshotted.")
	partnercentralSelling_createResourceSnapshotCmd.Flags().String("resource-snapshot-template-identifier", "", "The name of the template that defines the schema for the snapshot.")
	partnercentralSelling_createResourceSnapshotCmd.Flags().String("resource-type", "", "Specifies the type of resource for which the snapshot is being created.")
	partnercentralSelling_createResourceSnapshotCmd.MarkFlagRequired("catalog")
	partnercentralSelling_createResourceSnapshotCmd.MarkFlagRequired("client-token")
	partnercentralSelling_createResourceSnapshotCmd.MarkFlagRequired("engagement-identifier")
	partnercentralSelling_createResourceSnapshotCmd.MarkFlagRequired("resource-identifier")
	partnercentralSelling_createResourceSnapshotCmd.MarkFlagRequired("resource-snapshot-template-identifier")
	partnercentralSelling_createResourceSnapshotCmd.MarkFlagRequired("resource-type")
	partnercentralSellingCmd.AddCommand(partnercentralSelling_createResourceSnapshotCmd)
}
