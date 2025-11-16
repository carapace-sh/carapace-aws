package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_startBulkDeploymentCmd = &cobra.Command{
	Use:   "start-bulk-deployment",
	Short: "Deploys multiple groups in one operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_startBulkDeploymentCmd).Standalone()

	greengrass_startBulkDeploymentCmd.Flags().String("amzn-client-token", "", "A client token used to correlate requests and responses.")
	greengrass_startBulkDeploymentCmd.Flags().String("execution-role-arn", "", "The ARN of the execution role to associate with the bulk deployment operation.")
	greengrass_startBulkDeploymentCmd.Flags().String("input-file-uri", "", "The URI of the input file contained in the S3 bucket.")
	greengrass_startBulkDeploymentCmd.Flags().String("tags", "", "Tag(s) to add to the new resource.")
	greengrass_startBulkDeploymentCmd.MarkFlagRequired("execution-role-arn")
	greengrass_startBulkDeploymentCmd.MarkFlagRequired("input-file-uri")
	greengrassCmd.AddCommand(greengrass_startBulkDeploymentCmd)
}
