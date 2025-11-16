package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmSap_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists all tags on an SAP HANA application and/or database registered with AWS Systems Manager for SAP.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmSap_listTagsForResourceCmd).Standalone()

	ssmSap_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	ssmSap_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	ssmSapCmd.AddCommand(ssmSap_listTagsForResourceCmd)
}
