package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationSignals_listServiceLevelObjectivesCmd = &cobra.Command{
	Use:   "list-service-level-objectives",
	Short: "Returns a list of SLOs created in this account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationSignals_listServiceLevelObjectivesCmd).Standalone()

	applicationSignals_listServiceLevelObjectivesCmd.Flags().String("dependency-config", "", "Identifies the dependency using the `DependencyKeyAttributes` and `DependencyOperationName`.")
	applicationSignals_listServiceLevelObjectivesCmd.Flags().Bool("include-linked-accounts", false, "If you are using this operation in a monitoring account, specify `true` to include SLO from source accounts in the returned data.")
	applicationSignals_listServiceLevelObjectivesCmd.Flags().String("key-attributes", "", "You can use this optional field to specify which services you want to retrieve SLO information for.")
	applicationSignals_listServiceLevelObjectivesCmd.Flags().String("max-results", "", "The maximum number of results to return in one operation.")
	applicationSignals_listServiceLevelObjectivesCmd.Flags().String("metric-source-types", "", "Use this optional field to only include SLOs with the specified metric source types in the output.")
	applicationSignals_listServiceLevelObjectivesCmd.Flags().String("next-token", "", "Include this value, if it was returned by the previous operation, to get the next set of service level objectives.")
	applicationSignals_listServiceLevelObjectivesCmd.Flags().Bool("no-include-linked-accounts", false, "If you are using this operation in a monitoring account, specify `true` to include SLO from source accounts in the returned data.")
	applicationSignals_listServiceLevelObjectivesCmd.Flags().String("operation-name", "", "The name of the operation that this SLO is associated with.")
	applicationSignals_listServiceLevelObjectivesCmd.Flags().String("slo-owner-aws-account-id", "", "SLO's Amazon Web Services account ID.")
	applicationSignals_listServiceLevelObjectivesCmd.Flag("no-include-linked-accounts").Hidden = true
	applicationSignalsCmd.AddCommand(applicationSignals_listServiceLevelObjectivesCmd)
}
