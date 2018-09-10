package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Notification is a generated struct representing the Sophos Notification Endpoint
// GET /api/nodes/notification
type Notification struct {
	NotificationGroup        NotificationGroup        `json:"notification_group"`
	NotificationNotification NotificationNotification `json:"notification_notification"`
}

var _ sophos.Endpoint = &Notification{}

var defsNotification = map[string]sophos.RestObject{
	"NotificationGroup":        &NotificationGroup{},
	"NotificationNotification": &NotificationNotification{},
}

// RestObjects implements the sophos.Node interface and returns a map of Notification's Objects
func (Notification) RestObjects() map[string]sophos.RestObject { return defsNotification }

// GetPath implements sophos.RestGetter
func (*Notification) GetPath() string { return "/api/nodes/notification" }

// RefRequired implements sophos.RestGetter
func (*Notification) RefRequired() (string, bool) { return "", false }

var defNotification = &sophos.Definition{Description: "notification", Name: "notification", Link: "/api/definitions/notification"}

// Definition returns the /api/definitions struct of Notification
func (Notification) Definition() sophos.Definition { return *defNotification }

// ApiRoutes returns all known Notification Paths
func (Notification) ApiRoutes() []string {
	return []string{
		"/api/objects/notification/group/",
		"/api/objects/notification/group/{ref}",
		"/api/objects/notification/group/{ref}/usedby",
		"/api/objects/notification/notification/",
		"/api/objects/notification/notification/{ref}",
		"/api/objects/notification/notification/{ref}/usedby",
	}
}

// References returns the Notification's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Notification) References() []string {
	return []string{
		"REF_NotificationGroup",
		"REF_NotificationNotification",
	}
}

// NotificationGroups is an Sophos Endpoint subType and implements sophos.RestObject
type NotificationGroups []NotificationGroup

// NotificationGroup represents a UTM group
type NotificationGroup struct {
	Locked     string `json:"_locked"`
	Reference  string `json:"_ref"`
	ObjectType string `json:"_type"`
	Comment    string `json:"comment"`
	Name       string `json:"name"`
}

var _ sophos.RestGetter = &NotificationGroup{}

// GetPath implements sophos.RestObject and returns the NotificationGroups GET path
// Returns all available notification/group objects
func (*NotificationGroups) GetPath() string { return "/api/objects/notification/group/" }

// RefRequired implements sophos.RestObject
func (*NotificationGroups) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the NotificationGroups GET path
// Returns all available group types
func (n *NotificationGroup) GetPath() string {
	return fmt.Sprintf("/api/objects/notification/group/%s", n.Reference)
}

// RefRequired implements sophos.RestObject
func (n *NotificationGroup) RefRequired() (string, bool) { return n.Reference, true }

// DeletePath implements sophos.RestObject and returns the NotificationGroup DELETE path
// Creates or updates the complete object group
func (*NotificationGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/notification/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the NotificationGroup PATCH path
// Changes to parts of the object group types
func (*NotificationGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/notification/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the NotificationGroup POST path
// Create a new notification/group object
func (*NotificationGroup) PostPath() string {
	return "/api/objects/notification/group/"
}

// PutPath implements sophos.RestObject and returns the NotificationGroup PUT path
// Creates or updates the complete object group
func (*NotificationGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/notification/group/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*NotificationGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/notification/group/%s/usedby", ref)
}

// NotificationNotifications is an Sophos Endpoint subType and implements sophos.RestObject
type NotificationNotifications []NotificationNotification

// NotificationNotification represents a UTM notification
type NotificationNotification struct {
	Locked     string `json:"_locked"`
	Reference  string `json:"_ref"`
	ObjectType string `json:"_type"`
	Comment    string `json:"comment"`
	// Email default value is false
	Email bool   `json:"email"`
	Id    string `json:"id"`
	Name  string `json:"name"`
	// Snmp default value is false
	Snmp bool `json:"snmp"`
}

var _ sophos.RestGetter = &NotificationNotification{}

// GetPath implements sophos.RestObject and returns the NotificationNotifications GET path
// Returns all available notification/notification objects
func (*NotificationNotifications) GetPath() string { return "/api/objects/notification/notification/" }

// RefRequired implements sophos.RestObject
func (*NotificationNotifications) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the NotificationNotifications GET path
// Returns all available notification types
func (n *NotificationNotification) GetPath() string {
	return fmt.Sprintf("/api/objects/notification/notification/%s", n.Reference)
}

// RefRequired implements sophos.RestObject
func (n *NotificationNotification) RefRequired() (string, bool) { return n.Reference, true }

// DeletePath implements sophos.RestObject and returns the NotificationNotification DELETE path
// Creates or updates the complete object notification
func (*NotificationNotification) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/notification/notification/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the NotificationNotification PATCH path
// Changes to parts of the object notification types
func (*NotificationNotification) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/notification/notification/%s", ref)
}

// PostPath implements sophos.RestObject and returns the NotificationNotification POST path
// Create a new notification/notification object
func (*NotificationNotification) PostPath() string {
	return "/api/objects/notification/notification/"
}

// PutPath implements sophos.RestObject and returns the NotificationNotification PUT path
// Creates or updates the complete object notification
func (*NotificationNotification) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/notification/notification/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*NotificationNotification) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/notification/notification/%s/usedby", ref)
}
