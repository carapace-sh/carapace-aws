package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_createWorkloadCmd = &cobra.Command{
	Use:   "create-workload",
	Short: "Create a new workload.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_createWorkloadCmd).Standalone()

	wellarchitected_createWorkloadCmd.Flags().String("account-ids", "", "")
	wellarchitected_createWorkloadCmd.Flags().String("applications", "", "List of AppRegistry application ARNs associated to the workload.")
	wellarchitected_createWorkloadCmd.Flags().String("architectural-design", "", "")
	wellarchitected_createWorkloadCmd.Flags().String("aws-regions", "", "")
	wellarchitected_createWorkloadCmd.Flags().String("client-request-token", "", "")
	wellarchitected_createWorkloadCmd.Flags().String("description", "", "")
	wellarchitected_createWorkloadCmd.Flags().String("discovery-config", "", "Well-Architected discovery configuration settings associated to the workload.")
	wellarchitected_createWorkloadCmd.Flags().String("environment", "", "")
	wellarchitected_createWorkloadCmd.Flags().String("industry", "", "")
	wellarchitected_createWorkloadCmd.Flags().String("industry-type", "", "")
	wellarchitected_createWorkloadCmd.Flags().String("jira-configuration", "", "Jira configuration settings when creating a workload.")
	wellarchitected_createWorkloadCmd.Flags().String("lenses", "", "")
	wellarchitected_createWorkloadCmd.Flags().String("non-aws-regions", "", "")
	wellarchitected_createWorkloadCmd.Flags().String("notes", "", "")
	wellarchitected_createWorkloadCmd.Flags().String("pillar-priorities", "", "")
	wellarchitected_createWorkloadCmd.Flags().String("profile-arns", "", "The list of profile ARNs associated with the workload.")
	wellarchitected_createWorkloadCmd.Flags().String("review-owner", "", "")
	wellarchitected_createWorkloadCmd.Flags().String("review-template-arns", "", "The list of review template ARNs to associate with the workload.")
	wellarchitected_createWorkloadCmd.Flags().String("tags", "", "The tags to be associated with the workload.")
	wellarchitected_createWorkloadCmd.Flags().String("workload-name", "", "")
	wellarchitected_createWorkloadCmd.MarkFlagRequired("client-request-token")
	wellarchitected_createWorkloadCmd.MarkFlagRequired("description")
	wellarchitected_createWorkloadCmd.MarkFlagRequired("environment")
	wellarchitected_createWorkloadCmd.MarkFlagRequired("lenses")
	wellarchitected_createWorkloadCmd.MarkFlagRequired("workload-name")
	wellarchitectedCmd.AddCommand(wellarchitected_createWorkloadCmd)
}
