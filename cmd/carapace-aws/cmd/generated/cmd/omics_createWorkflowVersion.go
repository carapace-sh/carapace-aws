package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_createWorkflowVersionCmd = &cobra.Command{
	Use:   "create-workflow-version",
	Short: "Creates a new workflow version for the workflow that you specify with the `workflowId` parameter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_createWorkflowVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(omics_createWorkflowVersionCmd).Standalone()

		omics_createWorkflowVersionCmd.Flags().String("accelerators", "", "The computational accelerator for this workflow version.")
		omics_createWorkflowVersionCmd.Flags().String("container-registry-map", "", "(Optional) Use a container registry map to specify mappings between the ECR private repository and one or more upstream registries.")
		omics_createWorkflowVersionCmd.Flags().String("container-registry-map-uri", "", "(Optional) URI of the S3 location for the registry mapping file.")
		omics_createWorkflowVersionCmd.Flags().String("definition-repository", "", "The repository information for the workflow version definition.")
		omics_createWorkflowVersionCmd.Flags().String("definition-uri", "", "The S3 URI of a definition for this workflow version.")
		omics_createWorkflowVersionCmd.Flags().String("definition-zip", "", "A ZIP archive containing the main workflow definition file and dependencies that it imports for this workflow version.")
		omics_createWorkflowVersionCmd.Flags().String("description", "", "A description for this workflow version.")
		omics_createWorkflowVersionCmd.Flags().String("engine", "", "The workflow engine for this workflow version.")
		omics_createWorkflowVersionCmd.Flags().String("main", "", "The path of the main definition file for this workflow version.")
		omics_createWorkflowVersionCmd.Flags().String("parameter-template", "", "A parameter template for this workflow version.")
		omics_createWorkflowVersionCmd.Flags().String("parameter-template-path", "", "The path to the workflow version parameter template JSON file within the repository.")
		omics_createWorkflowVersionCmd.Flags().String("readme-markdown", "", "The markdown content for the workflow version's README file.")
		omics_createWorkflowVersionCmd.Flags().String("readme-path", "", "The path to the workflow version README markdown file within the repository.")
		omics_createWorkflowVersionCmd.Flags().String("readme-uri", "", "The S3 URI of the README file for the workflow version.")
		omics_createWorkflowVersionCmd.Flags().String("request-id", "", "An idempotency token to ensure that duplicate workflows are not created when Amazon Web Services HealthOmics submits retry requests.")
		omics_createWorkflowVersionCmd.Flags().String("storage-capacity", "", "The default static storage capacity (in gibibytes) for runs that use this workflow version.")
		omics_createWorkflowVersionCmd.Flags().String("storage-type", "", "The default storage type for runs that use this workflow version.")
		omics_createWorkflowVersionCmd.Flags().String("tags", "", "Tags for this workflow version.")
		omics_createWorkflowVersionCmd.Flags().String("version-name", "", "A name for the workflow version.")
		omics_createWorkflowVersionCmd.Flags().String("workflow-bucket-owner-id", "", "Amazon Web Services Id of the owner of the S3 bucket that contains the workflow definition.")
		omics_createWorkflowVersionCmd.Flags().String("workflow-id", "", "The ID of the workflow where you are creating the new version.")
		omics_createWorkflowVersionCmd.MarkFlagRequired("request-id")
		omics_createWorkflowVersionCmd.MarkFlagRequired("version-name")
		omics_createWorkflowVersionCmd.MarkFlagRequired("workflow-id")
	})
	omicsCmd.AddCommand(omics_createWorkflowVersionCmd)
}
