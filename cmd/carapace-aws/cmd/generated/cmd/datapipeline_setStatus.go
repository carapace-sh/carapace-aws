package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datapipeline_setStatusCmd = &cobra.Command{
	Use:   "set-status",
	Short: "Requests that the status of the specified physical or logical pipeline objects be updated in the specified pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datapipeline_setStatusCmd).Standalone()

	datapipeline_setStatusCmd.Flags().String("object-ids", "", "The IDs of the objects.")
	datapipeline_setStatusCmd.Flags().String("pipeline-id", "", "The ID of the pipeline that contains the objects.")
	datapipeline_setStatusCmd.Flags().String("status", "", "The status to be set on all the objects specified in `objectIds`.")
	datapipeline_setStatusCmd.MarkFlagRequired("object-ids")
	datapipeline_setStatusCmd.MarkFlagRequired("pipeline-id")
	datapipeline_setStatusCmd.MarkFlagRequired("status")
	datapipelineCmd.AddCommand(datapipeline_setStatusCmd)
}
