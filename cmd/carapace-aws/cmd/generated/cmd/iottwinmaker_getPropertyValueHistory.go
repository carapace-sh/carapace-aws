package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iottwinmaker_getPropertyValueHistoryCmd = &cobra.Command{
	Use:   "get-property-value-history",
	Short: "Retrieves information about the history of a time series property value for a component, component type, entity, or workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iottwinmaker_getPropertyValueHistoryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iottwinmaker_getPropertyValueHistoryCmd).Standalone()

		iottwinmaker_getPropertyValueHistoryCmd.Flags().String("component-name", "", "The name of the component.")
		iottwinmaker_getPropertyValueHistoryCmd.Flags().String("component-path", "", "This string specifies the path to the composite component, starting from the top-level component.")
		iottwinmaker_getPropertyValueHistoryCmd.Flags().String("component-type-id", "", "The ID of the component type.")
		iottwinmaker_getPropertyValueHistoryCmd.Flags().String("end-date-time", "", "The date and time of the latest property value to return.")
		iottwinmaker_getPropertyValueHistoryCmd.Flags().String("end-time", "", "The ISO8601 DateTime of the latest property value to return.")
		iottwinmaker_getPropertyValueHistoryCmd.Flags().String("entity-id", "", "The ID of the entity.")
		iottwinmaker_getPropertyValueHistoryCmd.Flags().String("interpolation", "", "An object that specifies the interpolation type and the interval over which to interpolate data.")
		iottwinmaker_getPropertyValueHistoryCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
		iottwinmaker_getPropertyValueHistoryCmd.Flags().String("next-token", "", "The string that specifies the next page of results.")
		iottwinmaker_getPropertyValueHistoryCmd.Flags().String("order-by-time", "", "The time direction to use in the result order.")
		iottwinmaker_getPropertyValueHistoryCmd.Flags().String("property-filters", "", "A list of objects that filter the property value history request.")
		iottwinmaker_getPropertyValueHistoryCmd.Flags().String("selected-properties", "", "A list of properties whose value histories the request retrieves.")
		iottwinmaker_getPropertyValueHistoryCmd.Flags().String("start-date-time", "", "The date and time of the earliest property value to return.")
		iottwinmaker_getPropertyValueHistoryCmd.Flags().String("start-time", "", "The ISO8601 DateTime of the earliest property value to return.")
		iottwinmaker_getPropertyValueHistoryCmd.Flags().String("workspace-id", "", "The ID of the workspace.")
		iottwinmaker_getPropertyValueHistoryCmd.MarkFlagRequired("selected-properties")
		iottwinmaker_getPropertyValueHistoryCmd.MarkFlagRequired("workspace-id")
	})
	iottwinmakerCmd.AddCommand(iottwinmaker_getPropertyValueHistoryCmd)
}
