package constant

import "time"

const (
	// LockerTTL defines the time-to-live for lockers.
	LockerTTL = 2 * time.Minute

	// CacheTTL defines the time-to-live for cache entries.
	CacheTTL = 2 * time.Hour

	// SessionTTL defines the time-to-live for user sessions.
	SessionTTL = 24 * time.Hour
)
