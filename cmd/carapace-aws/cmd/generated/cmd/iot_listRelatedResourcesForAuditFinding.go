package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listRelatedResourcesForAuditFindingCmd = &cobra.Command{
	Use:   "list-related-resources-for-audit-finding",
	Short: "The related resources of an Audit finding.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listRelatedResourcesForAuditFindingCmd).Standalone()

	iot_listRelatedResourcesForAuditFindingCmd.Flags().String("finding-id", "", "The finding Id.")
	iot_listRelatedResourcesForAuditFindingCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
	iot_listRelatedResourcesForAuditFindingCmd.Flags().String("next-token", "", "A token that can be used to retrieve the next set of results, or `null` if there are no additional results.")
	iot_listRelatedResourcesForAuditFindingCmd.MarkFlagRequired("finding-id")
	iotCmd.AddCommand(iot_listRelatedResourcesForAuditFindingCmd)
}
