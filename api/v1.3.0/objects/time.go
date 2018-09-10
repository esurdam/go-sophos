package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Time is a generated struct representing the Sophos Time Endpoint
// GET /api/nodes/time
type Time struct {
	TimeGroup     TimeGroup     `json:"time_group"`
	TimeRecurring TimeRecurring `json:"time_recurring"`
	TimeSingle    TimeSingle    `json:"time_single"`
}

var _ sophos.Endpoint = &Time{}

var defsTime = map[string]sophos.RestObject{
	"TimeGroup":     &TimeGroup{},
	"TimeRecurring": &TimeRecurring{},
	"TimeSingle":    &TimeSingle{},
}

// RestObjects implements the sophos.Node interface and returns a map of Time's Objects
func (Time) RestObjects() map[string]sophos.RestObject { return defsTime }

// GetPath implements sophos.RestGetter
func (*Time) GetPath() string { return "/api/nodes/time" }

// RefRequired implements sophos.RestGetter
func (*Time) RefRequired() (string, bool) { return "", false }

var defTime = &sophos.Definition{Description: "time", Name: "time", Link: "/api/definitions/time"}

// Definition returns the /api/definitions struct of Time
func (Time) Definition() sophos.Definition { return *defTime }

// ApiRoutes returns all known Time Paths
func (Time) ApiRoutes() []string {
	return []string{
		"/api/objects/time/group/",
		"/api/objects/time/group/{ref}",
		"/api/objects/time/group/{ref}/usedby",
		"/api/objects/time/recurring/",
		"/api/objects/time/recurring/{ref}",
		"/api/objects/time/recurring/{ref}/usedby",
		"/api/objects/time/single/",
		"/api/objects/time/single/{ref}",
		"/api/objects/time/single/{ref}/usedby",
	}
}

// References returns the Time's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Time) References() []string {
	return []string{
		"REF_TimeGroup",
		"REF_TimeRecurring",
		"REF_TimeSingle",
	}
}

// TimeGroups is an Sophos Endpoint subType and implements sophos.RestObject
type TimeGroups []TimeGroup

// TimeGroup represents a UTM group
type TimeGroup struct {
	Locked    string `json:"_locked"`
	Reference string `json:"_ref"`
	_type     string `json:"_type"`
	Comment   string `json:"comment"`
	Name      string `json:"name"`
}

var _ sophos.RestGetter = &TimeGroup{}

// GetPath implements sophos.RestObject and returns the TimeGroups GET path
// Returns all available time/group objects
func (*TimeGroups) GetPath() string { return "/api/objects/time/group/" }

// RefRequired implements sophos.RestObject
func (*TimeGroups) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the TimeGroups GET path
// Returns all available group types
func (t *TimeGroup) GetPath() string { return fmt.Sprintf("/api/objects/time/group/%s", t.Reference) }

// RefRequired implements sophos.RestObject
func (t *TimeGroup) RefRequired() (string, bool) { return t.Reference, true }

// DeletePath implements sophos.RestObject and returns the TimeGroup DELETE path
// Creates or updates the complete object group
func (*TimeGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/time/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the TimeGroup PATCH path
// Changes to parts of the object group types
func (*TimeGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/time/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the TimeGroup POST path
// Create a new time/group object
func (*TimeGroup) PostPath() string {
	return "/api/objects/time/group/"
}

// PutPath implements sophos.RestObject and returns the TimeGroup PUT path
// Creates or updates the complete object group
func (*TimeGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/time/group/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*TimeGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/time/group/%s/usedby", ref)
}

// TimeRecurrings is an Sophos Endpoint subType and implements sophos.RestObject
type TimeRecurrings []TimeRecurring

// TimeRecurring is a generated Sophos object
type TimeRecurring struct {
	Locked    string   `json:"_locked"`
	Reference string   `json:"_ref"`
	_type     string   `json:"_type"`
	Comment   string   `json:"comment"`
	EndTime   string   `json:"end_time"`
	Name      string   `json:"name"`
	StartTime string   `json:"start_time"`
	Weekdays  []string `json:"weekdays"`
}

var _ sophos.RestGetter = &TimeRecurring{}

// GetPath implements sophos.RestObject and returns the TimeRecurrings GET path
// Returns all available time/recurring objects
func (*TimeRecurrings) GetPath() string { return "/api/objects/time/recurring/" }

// RefRequired implements sophos.RestObject
func (*TimeRecurrings) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the TimeRecurrings GET path
// Returns all available recurring types
func (t *TimeRecurring) GetPath() string {
	return fmt.Sprintf("/api/objects/time/recurring/%s", t.Reference)
}

// RefRequired implements sophos.RestObject
func (t *TimeRecurring) RefRequired() (string, bool) { return t.Reference, true }

// DeletePath implements sophos.RestObject and returns the TimeRecurring DELETE path
// Creates or updates the complete object recurring
func (*TimeRecurring) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/time/recurring/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the TimeRecurring PATCH path
// Changes to parts of the object recurring types
func (*TimeRecurring) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/time/recurring/%s", ref)
}

// PostPath implements sophos.RestObject and returns the TimeRecurring POST path
// Create a new time/recurring object
func (*TimeRecurring) PostPath() string {
	return "/api/objects/time/recurring/"
}

// PutPath implements sophos.RestObject and returns the TimeRecurring PUT path
// Creates or updates the complete object recurring
func (*TimeRecurring) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/time/recurring/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*TimeRecurring) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/time/recurring/%s/usedby", ref)
}

// GetType implements sophos.Object
func (t *TimeRecurring) GetType() string { return t._type }

// TimeSingles is an Sophos Endpoint subType and implements sophos.RestObject
type TimeSingles []TimeSingle

// TimeSingle represents a UTM single time period
type TimeSingle struct {
	Locked    string `json:"_locked"`
	Reference string `json:"_ref"`
	_type     string `json:"_type"`
	// StartDate description: (DATE)
	StartDate string `json:"start_date"`
	// StartTime description: (TIME)
	StartTime string `json:"start_time"`
	Comment   string `json:"comment"`
	// EndDate description: (DATE)
	EndDate string `json:"end_date"`
	// EndTime description: (TIME)
	EndTime string `json:"end_time"`
	Name    string `json:"name"`
}

var _ sophos.RestGetter = &TimeSingle{}

// GetPath implements sophos.RestObject and returns the TimeSingles GET path
// Returns all available time/single objects
func (*TimeSingles) GetPath() string { return "/api/objects/time/single/" }

// RefRequired implements sophos.RestObject
func (*TimeSingles) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the TimeSingles GET path
// Returns all available single types
func (t *TimeSingle) GetPath() string { return fmt.Sprintf("/api/objects/time/single/%s", t.Reference) }

// RefRequired implements sophos.RestObject
func (t *TimeSingle) RefRequired() (string, bool) { return t.Reference, true }

// DeletePath implements sophos.RestObject and returns the TimeSingle DELETE path
// Creates or updates the complete object single
func (*TimeSingle) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/time/single/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the TimeSingle PATCH path
// Changes to parts of the object single types
func (*TimeSingle) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/time/single/%s", ref)
}

// PostPath implements sophos.RestObject and returns the TimeSingle POST path
// Create a new time/single object
func (*TimeSingle) PostPath() string {
	return "/api/objects/time/single/"
}

// PutPath implements sophos.RestObject and returns the TimeSingle PUT path
// Creates or updates the complete object single
func (*TimeSingle) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/time/single/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*TimeSingle) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/time/single/%s/usedby", ref)
}
