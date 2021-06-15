package crud

import (
	"context"

	"github.com/gigamono/gigamono/pkg/configs"
	controller "github.com/gigamono/gigamono/pkg/database/controllers/resource"
	model "github.com/gigamono/gigamono/pkg/database/models/resource"
	"github.com/gigamono/gigamono/pkg/errs"
	"github.com/gigamono/gigamono/pkg/files"
	"github.com/gigamono/gigamono/pkg/inits"
	"github.com/gigamono/gigamono/pkg/services/graphql/response"
	"github.com/gigamono/gigamono/pkg/services/rest/middleware"
	"github.com/gofrs/uuid"
)

// CreateWorkflow creates a new workflow in the database.
func CreateWorkflow(ctx context.Context, app *inits.App, specification string) (*model.Workflow, error) {
	// TODO: Sec: Validation, Permission.
	userID := ctx.Value(middleware.SessionDataKey).(middleware.SessionData).UserID

	// TODO: Validate workflow config.
	workflowConfig, err := configs.NewWorkflowConfig(specification, configs.JSON)
	if err != nil {
		return nil, response.Error(ctx, err.Error())
	}

	// TODO: figure workspace id // Generate obfuscated filepath.
	filePath, err := files.GenerateObfuscatedFilePath("json", uuid.Nil, "workflow", nil)
	if err != nil {
		panic(errs.NewSystemError("", "generating workflow spec obfuscated filename", err))
	}

	// Save integration to a file.
	if _, err := app.Filestore.Project.WriteToFile(filePath, []byte(specification)); err != nil {
		panic(errs.NewSystemError("", "writing workflow spec to file", err))
	}

	// TODO: Compile workflow config.

	// Create the workflow in db.
	workflow, err := controller.CreateWorkflow(&app.DB, &userID, workflowConfig.Metadata.Name, filePath)
	if err != nil {
		panic(errs.NewSystemError("", "creating workflow", err))
	}

	return workflow, nil
}

// GetWorkflow gets an existing workflow from the database.
func GetWorkflow(ctx context.Context, app *inits.App, workflowID string) (*model.Workflow, error) {
	// TODO: Sec: Validation, Permission.
	userID := ctx.Value(middleware.SessionDataKey).(middleware.SessionData).UserID

	workflowUUID, err := uuid.FromString(workflowID)
	if err != nil {
		panic(err)
	}

	// Get the workflow from db.
	workflow, err := controller.GetWorkflow(&app.DB, &userID, &workflowUUID)
	if err != nil {
		panic(errs.NewSystemError("", "getting workflow", err))
	}

	return workflow, nil
}
