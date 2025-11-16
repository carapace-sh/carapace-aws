package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listAttachedPoliciesCmd = &cobra.Command{
	Use:   "list-attached-policies",
	Short: "Lists the policies attached to the specified thing group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listAttachedPoliciesCmd).Standalone()

	iot_listAttachedPoliciesCmd.Flags().String("marker", "", "The token to retrieve the next set of results.")
	iot_listAttachedPoliciesCmd.Flags().String("page-size", "", "The maximum number of results to be returned per request.")
	iot_listAttachedPoliciesCmd.Flags().String("recursive", "", "When true, recursively list attached policies.")
	iot_listAttachedPoliciesCmd.Flags().String("target", "", "The group or principal for which the policies will be listed.")
	iot_listAttachedPoliciesCmd.MarkFlagRequired("target")
	iotCmd.AddCommand(iot_listAttachedPoliciesCmd)
}
