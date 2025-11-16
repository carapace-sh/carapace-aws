package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeElasticGpusCmd = &cobra.Command{
	Use:   "describe-elastic-gpus",
	Short: "Amazon Elastic Graphics reached end of life on January 8, 2024.\n\nDescribes the Elastic Graphics accelerator associated with your instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeElasticGpusCmd).Standalone()

	ec2_describeElasticGpusCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeElasticGpusCmd.Flags().String("elastic-gpu-ids", "", "The Elastic Graphics accelerator IDs.")
	ec2_describeElasticGpusCmd.Flags().String("filters", "", "The filters.")
	ec2_describeElasticGpusCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
	ec2_describeElasticGpusCmd.Flags().String("next-token", "", "The token to request the next page of results.")
	ec2_describeElasticGpusCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeElasticGpusCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeElasticGpusCmd)
}
