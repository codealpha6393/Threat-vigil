package domain

import "time"

// Incident represents a security alert or threat detected in the system.
// This replaces the 'Task' struct.
type Incident struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`       // e.g., "SQL Injection", "Brute Force"
	Severity  string    `json:"severity"`   // e.g., "CRITICAL", "LOW"
	RawLog    string    `json:"raw_log"`    // The log line that triggered the alert
	Status    string    `json:"status"`     // e.g., "OPEN", "RESOLVED"
	CreatedAt time.Time `json:"created_at"`
}