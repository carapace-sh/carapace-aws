package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_putComplianceItemsCmd = &cobra.Command{
	Use:   "put-compliance-items",
	Short: "Registers a compliance type and other compliance details on a designated resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_putComplianceItemsCmd).Standalone()

	ssm_putComplianceItemsCmd.Flags().String("compliance-type", "", "Specify the compliance type.")
	ssm_putComplianceItemsCmd.Flags().String("execution-summary", "", "A summary of the call execution that includes an execution ID, the type of execution (for example, `Command`), and the date/time of the execution using a datetime object that is saved in the following format: `yyyy-MM-dd'T'HH:mm:ss'Z'`")
	ssm_putComplianceItemsCmd.Flags().String("item-content-hash", "", "MD5 or SHA-256 content hash.")
	ssm_putComplianceItemsCmd.Flags().String("items", "", "Information about the compliance as defined by the resource type.")
	ssm_putComplianceItemsCmd.Flags().String("resource-id", "", "Specify an ID for this resource.")
	ssm_putComplianceItemsCmd.Flags().String("resource-type", "", "Specify the type of resource.")
	ssm_putComplianceItemsCmd.Flags().String("upload-type", "", "The mode for uploading compliance items.")
	ssm_putComplianceItemsCmd.MarkFlagRequired("compliance-type")
	ssm_putComplianceItemsCmd.MarkFlagRequired("execution-summary")
	ssm_putComplianceItemsCmd.MarkFlagRequired("items")
	ssm_putComplianceItemsCmd.MarkFlagRequired("resource-id")
	ssm_putComplianceItemsCmd.MarkFlagRequired("resource-type")
	ssmCmd.AddCommand(ssm_putComplianceItemsCmd)
}
