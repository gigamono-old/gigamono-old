package resource

import (
	"encoding/json"

	"github.com/gigamono/gigamono/pkg/database/models"
	"github.com/gofrs/uuid"
)

// Preferences stores session preferences of a user.
type Preferences struct {
	models.Base
	Details Details   `json:"preferences"`
	UserID  uuid.UUID `pg:"type:uuid,notnull" json:"user_id"`
}

// Details contains details of the preferences.
type Details struct {
	FocusWorkspaceIndex uint64 `json:"focus_workspace_index"`
	Workspaces          []struct {
		ID              string `json:"id"`
		FocusSpaceIndex uint64 `json:"focus_space_index"`
		Spaces          []struct {
			ID                string `json:"id"`
			FocusProjectIndex uint64 `json:"focus_project_index"`
			Projects          []struct {
				ID                 string `json:"id"`
				FocusBoardIndex    uint64 `json:"focus_board_index"`
				FocusScreenIndex   uint64 `json:"focus_screen_index"`
				FocusWorkflowIndex uint64 `json:"focus_workflow_index"`
				FocusTableIndex    uint64 `json:"focus_table_index"`
				Boards             []struct {
					ID string `json:"id"`
				} `json:"boards"`
				Screens []struct {
					ID string `json:"id"`
				} `json:"screens"`
				Workflows []struct {
					ID string `json:"id"`
				} `json:"workflows"`
				Tables []struct {
					ID string `json:"id"`
				} `json:"tables"`
			} `json:"projects"`
		} `json:"spaces"`
		Layout struct {
			MainShortcuts      Shortcut `json:"main_shortcuts"`
			WorkspaceShortcuts Shortcut `json:"workspace_shortcuts"`
			OtherShortcuts     Shortcut `json:"other_shortcuts"`
		} `json:"layout"`
	} `json:"workspaces"`
}

// Shortcut represents a shortcut icon or button in the UI. Usually situated in sidebar.
type Shortcut struct {
	IconName   string `json:"icon_name"`
	EntityName string `json:"entity_name"`
	Route      string `json:"route"`
}

// JSON converts details to json.
func (details *Details) JSON() (string, error) {
	// TODO: Sec: Validation
	bytes, err := json.Marshal(details)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
