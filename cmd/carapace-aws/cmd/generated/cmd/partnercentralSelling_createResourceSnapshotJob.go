package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var partnercentralSelling_createResourceSnapshotJobCmd = &cobra.Command{
	Use:   "create-resource-snapshot-job",
	Short: "Use this action to create a job to generate a snapshot of the specified resource within an engagement.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(partnercentralSelling_createResourceSnapshotJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(partnercentralSelling_createResourceSnapshotJobCmd).Standalone()

		partnercentralSelling_createResourceSnapshotJobCmd.Flags().String("catalog", "", "Specifies the catalog in which to create the snapshot job.")
		partnercentralSelling_createResourceSnapshotJobCmd.Flags().String("client-token", "", "A client-generated UUID used for idempotency check.")
		partnercentralSelling_createResourceSnapshotJobCmd.Flags().String("engagement-identifier", "", "Specifies the identifier of the engagement associated with the resource to be snapshotted.")
		partnercentralSelling_createResourceSnapshotJobCmd.Flags().String("resource-identifier", "", "Specifies the identifier of the specific resource to be snapshotted.")
		partnercentralSelling_createResourceSnapshotJobCmd.Flags().String("resource-snapshot-template-identifier", "", "Specifies the name of the template that defines the schema for the snapshot.")
		partnercentralSelling_createResourceSnapshotJobCmd.Flags().String("resource-type", "", "The type of resource for which the snapshot job is being created.")
		partnercentralSelling_createResourceSnapshotJobCmd.Flags().String("tags", "", "A map of the key-value pairs of the tag or tags to assign.")
		partnercentralSelling_createResourceSnapshotJobCmd.MarkFlagRequired("catalog")
		partnercentralSelling_createResourceSnapshotJobCmd.MarkFlagRequired("client-token")
		partnercentralSelling_createResourceSnapshotJobCmd.MarkFlagRequired("engagement-identifier")
		partnercentralSelling_createResourceSnapshotJobCmd.MarkFlagRequired("resource-identifier")
		partnercentralSelling_createResourceSnapshotJobCmd.MarkFlagRequired("resource-snapshot-template-identifier")
		partnercentralSelling_createResourceSnapshotJobCmd.MarkFlagRequired("resource-type")
	})
	partnercentralSellingCmd.AddCommand(partnercentralSelling_createResourceSnapshotJobCmd)
}
