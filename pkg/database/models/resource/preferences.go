package resource

import (
	"encoding/json"
	"fmt"

	"github.com/gigamono/gigamono/pkg/database"
	"github.com/gigamono/gigamono/pkg/database/models"
	"github.com/gofrs/uuid"
)

// Preferences stores session preferences of a user.
type Preferences struct {
	models.Base
	Details *Details  `json:"details"`
	UserID  uuid.UUID `pg:"type:uuid,notnull" json:"user_id"`
}

// Details contains details of the preferences.
type Details struct {
	FocusWorkspaceIndex int64            `json:"focus_workspace_index"`
	Workspaces          []PrefsWorkspace `json:"workspaces"`
}

// PrefsWorkspace contains details about a workspace.
type PrefsWorkspace struct {
	ID              string       `json:"id"`
	FocusSpaceIndex int64        `json:"focus_space_index"`
	Spaces          []PrefsSpace `json:"spaces"`
	Layout          PrefsLayout  `json:"layout"`
}

// PrefsSpace contains details about a space.
type PrefsSpace struct {
	ID                   string            `json:"id"`
	FocusDeckIndex       int64             `json:"focus_deck_index"`
	FocusAppIndex        int64             `json:"focus_app_index"`
	FocusAutomationIndex int64             `json:"focus_automation_index"`
	FocusBaseIndex       int64             `json:"focus_base_index"`
	Decks                []PrefsDeck       `json:"decks"`
	Automations          []PrefsAutomation `json:"automations"`
	Bases                []PrefsBase       `json:"bases"`
}

// PrefsDeck contains details about a deck.
type PrefsDeck struct {
	ID              string `json:"id"`
	FocusBoardIndex int64  `json:"focus_board_index"`
	Boards          []struct {
		ID string `json:"id"`
	} `json:"boards"`
}

// PrefsAutomation contains details about an automation.
type PrefsAutomation struct {
	ID                 string `json:"id"`
	FocusWorkflowIndex int64  `json:"focus_workflow_index"`
	Workflows          []struct {
		ID string `json:"id"`
	} `json:"workflows"`
}

// PrefsBase contains details about a base.
type PrefsBase struct {
	ID string `json:"id"`

	FocusTableIndex int64 `json:"focus_table_index"`
	Tables          []struct {
		ID string `json:"id"`
	} `json:"tables"`
}

// PrefsLayout contains details about a layout.
type PrefsLayout struct {
	MainShortcuts      Shortcut `json:"main_shortcuts"`
	WorkspaceShortcuts Shortcut `json:"workspace_shortcuts"`
	OtherShortcuts     Shortcut `json:"other_shortcuts"`
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

// Create creates a preferences.
func (preferences *Preferences) Create(db *database.DB) error {
	// TODO: Sec: Permission.
	// Insert preferences in db.
	if _, err := db.Model(preferences).Insert(); err != nil {
		return fmt.Errorf("creating preferences in db: %v", err)
	}

	return nil
}

// UpdateByID updates all the fields of a preferences by id.
func (preferences *Preferences) UpdateByID(db *database.DB) error {
	// TODO: Sec: Permission.
	// Update preferences in db.
	if _, err := db.
		Model(preferences).
		Column( // Sec: columns allowed to be updated.
			"details",
		).
		WherePK().
		Returning("*").
		Update(); err != nil {
		return fmt.Errorf("updating preferences in db: %v", err)
	}

	return nil
}

// GetByID gets an preferences by id.
func (preferences *Preferences) GetByID(db *database.DB) error {
	// TODO: Sec: Permission.s
	// Select the preferences with the specified preferences ID.
	if err := db.Model(preferences).WherePK().Select(); err != nil {
		return fmt.Errorf("fetching preferences from db: %v", err)
	}

	return nil
}
