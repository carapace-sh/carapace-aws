package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var partnercentralSelling_getResourceSnapshotCmd = &cobra.Command{
	Use:   "get-resource-snapshot",
	Short: "Use this action to retrieve a specific snapshot record.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(partnercentralSelling_getResourceSnapshotCmd).Standalone()

	partnercentralSelling_getResourceSnapshotCmd.Flags().String("catalog", "", "Specifies the catalog related to the request.")
	partnercentralSelling_getResourceSnapshotCmd.Flags().String("engagement-identifier", "", "The unique identifier of the engagement associated with the snapshot.")
	partnercentralSelling_getResourceSnapshotCmd.Flags().String("resource-identifier", "", "The unique identifier of the specific resource that was snapshotted.")
	partnercentralSelling_getResourceSnapshotCmd.Flags().String("resource-snapshot-template-identifier", "", "he name of the template that defines the schema for the snapshot.")
	partnercentralSelling_getResourceSnapshotCmd.Flags().String("resource-type", "", "Specifies the type of resource that was snapshotted.")
	partnercentralSelling_getResourceSnapshotCmd.Flags().String("revision", "", "Specifies which revision of the snapshot to retrieve.")
	partnercentralSelling_getResourceSnapshotCmd.MarkFlagRequired("catalog")
	partnercentralSelling_getResourceSnapshotCmd.MarkFlagRequired("engagement-identifier")
	partnercentralSelling_getResourceSnapshotCmd.MarkFlagRequired("resource-identifier")
	partnercentralSelling_getResourceSnapshotCmd.MarkFlagRequired("resource-snapshot-template-identifier")
	partnercentralSelling_getResourceSnapshotCmd.MarkFlagRequired("resource-type")
	partnercentralSellingCmd.AddCommand(partnercentralSelling_getResourceSnapshotCmd)
}
