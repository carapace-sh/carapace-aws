package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_updateWorkloadCmd = &cobra.Command{
	Use:   "update-workload",
	Short: "Update an existing workload.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_updateWorkloadCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wellarchitected_updateWorkloadCmd).Standalone()

		wellarchitected_updateWorkloadCmd.Flags().String("account-ids", "", "")
		wellarchitected_updateWorkloadCmd.Flags().String("applications", "", "List of AppRegistry application ARNs to associate to the workload.")
		wellarchitected_updateWorkloadCmd.Flags().String("architectural-design", "", "")
		wellarchitected_updateWorkloadCmd.Flags().String("aws-regions", "", "")
		wellarchitected_updateWorkloadCmd.Flags().String("description", "", "")
		wellarchitected_updateWorkloadCmd.Flags().String("discovery-config", "", "Well-Architected discovery configuration settings to associate to the workload.")
		wellarchitected_updateWorkloadCmd.Flags().String("environment", "", "")
		wellarchitected_updateWorkloadCmd.Flags().String("improvement-status", "", "")
		wellarchitected_updateWorkloadCmd.Flags().String("industry", "", "")
		wellarchitected_updateWorkloadCmd.Flags().String("industry-type", "", "")
		wellarchitected_updateWorkloadCmd.Flags().String("is-review-owner-update-acknowledged", "", "Flag indicating whether the workload owner has acknowledged that the *Review owner* field is required.")
		wellarchitected_updateWorkloadCmd.Flags().String("jira-configuration", "", "Configuration of the Jira integration.")
		wellarchitected_updateWorkloadCmd.Flags().String("non-aws-regions", "", "")
		wellarchitected_updateWorkloadCmd.Flags().String("notes", "", "")
		wellarchitected_updateWorkloadCmd.Flags().String("pillar-priorities", "", "")
		wellarchitected_updateWorkloadCmd.Flags().String("review-owner", "", "")
		wellarchitected_updateWorkloadCmd.Flags().String("workload-id", "", "")
		wellarchitected_updateWorkloadCmd.Flags().String("workload-name", "", "")
		wellarchitected_updateWorkloadCmd.MarkFlagRequired("workload-id")
	})
	wellarchitectedCmd.AddCommand(wellarchitected_updateWorkloadCmd)
}
