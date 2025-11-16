package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evidently_updateProjectDataDeliveryCmd = &cobra.Command{
	Use:   "update-project-data-delivery",
	Short: "Updates the data storage options for this project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evidently_updateProjectDataDeliveryCmd).Standalone()

	evidently_updateProjectDataDeliveryCmd.Flags().String("cloud-watch-logs", "", "A structure containing the CloudWatch Logs log group where you want to store evaluation events.")
	evidently_updateProjectDataDeliveryCmd.Flags().String("project", "", "The name or ARN of the project that you want to modify the data storage options for.")
	evidently_updateProjectDataDeliveryCmd.Flags().String("s3-destination", "", "A structure containing the S3 bucket name and bucket prefix where you want to store evaluation events.")
	evidently_updateProjectDataDeliveryCmd.MarkFlagRequired("project")
	evidentlyCmd.AddCommand(evidently_updateProjectDataDeliveryCmd)
}
