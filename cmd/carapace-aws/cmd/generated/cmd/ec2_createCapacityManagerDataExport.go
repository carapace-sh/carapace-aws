package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createCapacityManagerDataExportCmd = &cobra.Command{
	Use:   "create-capacity-manager-data-export",
	Short: "Creates a new data export configuration for EC2 Capacity Manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createCapacityManagerDataExportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_createCapacityManagerDataExportCmd).Standalone()

		ec2_createCapacityManagerDataExportCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		ec2_createCapacityManagerDataExportCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createCapacityManagerDataExportCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createCapacityManagerDataExportCmd.Flags().String("output-format", "", "The file format for the exported data.")
		ec2_createCapacityManagerDataExportCmd.Flags().String("s3-bucket-name", "", "The name of the S3 bucket where the capacity data export files will be delivered.")
		ec2_createCapacityManagerDataExportCmd.Flags().String("s3-bucket-prefix", "", "The S3 key prefix for the exported data files.")
		ec2_createCapacityManagerDataExportCmd.Flags().String("schedule", "", "The frequency at which data exports are generated.")
		ec2_createCapacityManagerDataExportCmd.Flags().String("tag-specifications", "", "The tags to apply to the data export configuration.")
		ec2_createCapacityManagerDataExportCmd.Flag("no-dry-run").Hidden = true
		ec2_createCapacityManagerDataExportCmd.MarkFlagRequired("output-format")
		ec2_createCapacityManagerDataExportCmd.MarkFlagRequired("s3-bucket-name")
		ec2_createCapacityManagerDataExportCmd.MarkFlagRequired("schedule")
	})
	ec2Cmd.AddCommand(ec2_createCapacityManagerDataExportCmd)
}
